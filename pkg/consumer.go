package Orkestra

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"orkestra-integration-provider/pkg/Entities"
	"regexp"
	"time"
)

var (
	uri          = flag.String("uri", "amqp://service:123qwe123@192.168.1.48:32062/semerp", "AMQP URI")
	exchange     = flag.String("exchange", "sem-integration", "Durable, non-auto-deleted AMQP exchange name")
	exchangeType = flag.String("exchange-type", "topic", "Exchange type - direct|fanout|topic|x-custom")
	queue        = flag.String("queue", "orkestraerp.service", "Ephemeral AMQP queue name")
	bindingKey   = flag.String("key", "orkestraerp", "AMQP binding key")
	consumerTag  = flag.String("consumer-tag", "orkestra.erp.integration.service", "AMQP consumer tag (should not be blank)")
	lifetime     = flag.Duration("lifetime", 0*time.Second, "lifetime of process before shutdown (0s=infinite)")
	converter    = make(map[string]ConverterFunc)
)

func init() {
	flag.Parse()
	converter["MaterialStockItemConverter"] = Entities.MaterialStockItemConverter
}

func ConnectAMQP() {
	soapConfig := &SoapConfig{
		WSDL:      "http://192.168.1.61:9090/ws/factory?wsdl",
		LoginUser: "admin",
		LoginPass: "a123",
	}
	c.orkestra = new(Soap)
	c.orkestra.CreateOrkestraClient(soapConfig)

	c, err := NewConsumer(*uri, *exchange, *exchangeType, *queue, *bindingKey, *consumerTag)
	if err != nil {
		log.Fatalf("%s", err)
	}

	if *lifetime > 0 {
		log.Printf("running for %s", *lifetime)
		time.Sleep(*lifetime)
	} else {
		log.Printf("AMQP: running forever")
		select {}
	}

	log.Printf("shutting down")

	if err := c.Shutdown(); err != nil {
		log.Fatalf("error during shutdown: %s", err)
	}
}

type Consumer struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	tag      string
	done     chan error
	orkestra *Soap
}

var c = &Consumer{
	conn:    nil,
	channel: nil,
	tag:     *consumerTag,
	done:    make(chan error),
}

func NewConsumer(amqpURI, exchange, exchangeType, queueName, key, ctag string) (*Consumer, error) {

	var err error
	c.conn, err = amqp.Dial(amqpURI)
	if err != nil {
		return nil, fmt.Errorf("Dial: %s", err)
	}

	log.Printf("got Connection, getting Channel")
	c.channel, err = c.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("Channel: %s", err)
	}

	_, err = c.channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	err = c.channel.QueueBind(
		queueName,        // name of the queue
		*bindingKey+".*", //
		exchange,         //
		false, nil)
	if err != nil {
		return nil, err
	}

	deliveries, err := c.channel.Consume(
		queueName, // name
		c.tag,     // consumerTag,
		false,     // noAck
		false,     // exclusive
		false,     // noLocal
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("Queue Consume: %s", err)
	}

	go handle(deliveries, c.done)

	return c, nil
}

/**
 * Gracefully shutdown the consumer
 */
func (c *Consumer) Shutdown() error {
	if err := c.channel.Cancel(c.tag, true); err != nil {
		return fmt.Errorf("Consumer cancel failed: %s", err)
	}

	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}

	defer log.Printf("AMQP shutdown OK")

	// wait for handle() to exit
	return <-c.done
}

/**
 * Publishes a message to the given exchange with the given routing key.
 */
func publishReplyGrpc(body string, d amqp.Delivery) {
	err := c.channel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
		Body:        []byte(body),
		ContentType: "application/json"})

	if err != nil {
		log.Printf("error publishing response: %s", err)
	}
	err = c.channel.Ack(d.DeliveryTag, false)
	if err != nil {
		log.Printf("error acking message: %s", err)
	}
}

func getPage(d amqp.Delivery) {
	var getpage = GetPagePayload{}
	err := json.Unmarshal(d.Body, &getpage.WS.Params)
	if err != nil {
		return
	}
	reference, _ := c.orkestra.GetPage(getpage)
	publishReplyGrpc(reference, d)

}

func getByReference(d amqp.Delivery) {
	var byReference = ReadByReferencePayload{}
	err := json.Unmarshal(d.Body, &byReference.WS.Params)
	if err != nil {
		return

	}
	reference, _ := c.orkestra.ReadByReference(byReference)
	publishReplyGrpc(reference, d)
}

func delete(d amqp.Delivery) {
	var delete = DeletePayload{}
	err := json.Unmarshal(d.Body, &delete.WS.Params)
	if err != nil {
		log.Fatalf(err.Error())
	}
	reference, _ := c.orkestra.Delete(delete)
	publishReplyGrpc(reference, d)
}

func create(d amqp.Delivery) {
	var receiveCreate = ReceiveCreatePayload{}
	var create = CreatePayload{}
	err := json.Unmarshal(d.Body, &receiveCreate)
	if err != nil {
		return
	}
	beanXML, err := converter[receiveCreate.Converter](receiveCreate.Entity)
	if err != nil {
		return
	}
	create.WS.Params.BeanXml = beanXML
	create.WS.Params.Period = receiveCreate.Period
	reference, _ := c.orkestra.Create(create)
	var createResponse = CreateResponse{}
	err = json.Unmarshal([]byte(reference), &createResponse)
	if createResponse.Envelope.Body.CreateResponse.Return == "" {
		publishReplyGrpc(reference, d)
	} else {
		var decodeResponseID, _ = base64.StdEncoding.DecodeString(createResponse.Envelope.Body.CreateResponse.Return)
		referenceRegex := regexp.MustCompile(`<reference>(.*?)</reference>`)
		referenceId := referenceRegex.FindStringSubmatch(string(decodeResponseID))[1]

		createResponse.Envelope.Body.CreateResponse.Return = referenceId
		var referenceOutput, _ = json.Marshal(createResponse)

		publishReplyGrpc(string(referenceOutput), d)
	}
}

func update(d amqp.Delivery) {
	var receiveCreate = ReceiveCreatePayload{}
	var create = UpdatePayload{}
	err := json.Unmarshal(d.Body, &receiveCreate)
	if err != nil {
		return
	}
	beanXML, err := converter[receiveCreate.Converter](receiveCreate.Entity)
	if err != nil {
		return
	}
	create.WS.Params.BeanXml = beanXML
	create.WS.Params.Period = receiveCreate.Period
	reference, _ := c.orkestra.Update(create)
	var createResponse = CreateResponse{}
	err = json.Unmarshal([]byte(reference), &createResponse)
	if createResponse.Envelope.Body.CreateResponse.Return == "" {
		publishReplyGrpc(reference, d)
	} else {
		var decodeResponseID, _ = base64.StdEncoding.DecodeString(createResponse.Envelope.Body.CreateResponse.Return)
		referenceRegex := regexp.MustCompile(`<reference>(.*?)</reference>`)
		referenceId := referenceRegex.FindStringSubmatch(string(decodeResponseID))[1]

		createResponse.Envelope.Body.CreateResponse.Return = referenceId
		var referenceOutput, _ = json.Marshal(createResponse)

		publishReplyGrpc(string(referenceOutput), d)
	}
}

func handle(deliveries <-chan amqp.Delivery, done chan error) {
	for d := range deliveries {
		switch d.RoutingKey {
		case *bindingKey + ".getpage":
			getPage(d)
		case *bindingKey + ".byReference":
			getByReference(d)
		case *bindingKey + ".delete":
			delete(d)
		case *bindingKey + ".create":
			create(d)
		case *bindingKey + ".update":
			update(d)
		}
	}
	log.Printf("handle: deliveries channel closed")

	done <- nil
}
