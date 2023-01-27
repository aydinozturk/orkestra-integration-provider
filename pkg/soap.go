package Orkestra

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func (osoap *Soap) CreateOrkestraClient(conf *SoapConfig) {

	osoap.client = &http.Client{}
	osoap.Config = conf
}

func (osoap *Soap) Call(payload io.Reader) (string, error) {
	method := "POST"
	req, err := http.NewRequest(method, osoap.Config.WSDL, payload)
	if err != nil {
		fmt.Println(err)
	}
	req.SetBasicAuth(osoap.Config.LoginUser, osoap.Config.LoginPass)
	req.Header.Add("Content-Type", "application/json")

	res, err := osoap.client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(body), err
}

func (osoap *Soap) ReadByReference(params ReadByReferencePayload) (string, error) {
	params.WS.Params.OutputType = "json2"
	params.WS.Params.ClassType = -1

	reda, _ := json.Marshal(params)
	payload := strings.NewReader(string(reda))
	return osoap.Call(payload)
}

func (osoap *Soap) GetPage(params GetPagePayload) (string, error) {
	params.WS.Params.OutputType = "json2"
	params.WS.Params.ClassType = -1
	if params.WS.Params.Filters == nil {
		params.WS.Params.Filters = []Filters{}
	}
	if params.WS.Params.OrderBy.Desc == "" {
		params.WS.Params.OrderBy.Desc = "desc"
		params.WS.Params.OrderBy.Property = "id"
	}
	reda, _ := json.Marshal(params)
	payload := strings.NewReader(string(reda))
	return osoap.Call(payload)
}

func (osoap *Soap) Delete(params DeletePayload) (string, error) {
	params.WS.Params.OutputType = "json2"
	params.WS.Params.ClassType = -1
	params.WS.Params.SingleTx = false
	reda, _ := json.Marshal(params)
	payload := strings.NewReader(string(reda))
	return osoap.Call(payload)
}

func (osoap *Soap) Create(params CreatePayload) (string, error) {
	params.WS.Params.OutputType = "json2"
	params.WS.Params.IgnoreOverflow = true
	params.WS.Params.ForceDuplicate = false
	params.WS.Params.RebuildData = true
	params.WS.Params.ReturnReferenceOnly = true
	reda, _ := json.Marshal(params)
	payload := strings.NewReader(string(reda))
	return osoap.Call(payload)
}

func (osoap *Soap) Update(params UpdatePayload) (string, error) {
	params.WS.Params.OutputType = "json2"
	params.WS.Params.IgnoreOverflow = true
	params.WS.Params.ForceDuplicate = false
	params.WS.Params.RebuildData = true
	params.WS.Params.ReturnReferenceOnly = true
	reda, _ := json.Marshal(params)
	log.Println(string(reda))
	payload := strings.NewReader(string(reda))
	return osoap.Call(payload)
}
