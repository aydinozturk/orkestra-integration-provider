package Orkestra

import (
	"encoding/xml"
	"net/http"
)

type Soap struct {
	client *http.Client
	Config *SoapConfig
}

type SoapConfig struct {
	LoginUser string
	LoginPass string
	WSDL      string
}

type ConverterFunc func(payload string) (string, error)

type ReadByReferencePayload struct {
	WS struct {
		Params struct {
			Period     int    `json:"Period"`
			OutputType string `json:"OutputType"`
			EntityName string `json:"EntityName"`
			ClassType  int    `json:"ClassType"`
			Reference  int    `json:"Reference"`
		} `json:"params"`
	} `json:"ws.readByReference"`
}

type GetPagePayload struct {
	WS struct {
		Params struct {
			Period     int       `json:"Period"`
			OutputType string    `json:"OutputType"`
			EntityName string    `json:"EntityName"`
			ClassType  int       `json:"ClassType"`
			Fields     []string  `json:"Fields"`
			PageSize   int       `json:"PageSize"`
			PageIndex  int       `json:"PageIndex"`
			Filters    []Filters `json:"Filters"`
			OrderBy    OrderBy   `json:"OrderBy"`
		} `json:"params"`
	} `json:"ws.getPage"`
}

type DeletePayload struct {
	WS struct {
		Params struct {
			Period           int    `json:"Period"`
			OutputType       string `json:"OutputType"`
			EntityName       string `json:"EntityName"`
			ClassType        int    `json:"ClassType"`
			References       []int  `json:"References"`
			SingleTx         bool   `json:"SingleTx"`
			LoadBeforeDelete bool   `json:"LoadBeforeDelete"`
			IgnoreOverflow   bool   `json:"IgnoreOverflow"`
		} `json:"params"`
	} `json:"ws.delete"`
}

type CreatePayload struct {
	WS struct {
		Params UpdateCreatePayloadParams `json:"params"`
	} `json:"ws.create"`
}

type UpdatePayload struct {
	WS struct {
		Params UpdateCreatePayloadParams `json:"params"`
	} `json:"ws.update"`
}

type UpdateCreatePayloadParams struct {
	Period              int    `json:"Period"`
	OutputType          string `json:"OutputType"`
	BeanXml             string `json:"BeanXml"`
	ForceDuplicate      bool   `json:"ForceDuplicate"`
	RebuildData         bool   `json:"RebuildData"`
	ReturnReferenceOnly bool   `json:"ReturnReferenceOnly"`
	IgnoreOverflow      bool   `json:"IgnoreOverflow"`
}

type CreateResponse struct {
	Envelope struct {
		Body struct {
			CreateResponse struct {
				Return string `json:"return"`
			} `json:"createResponse"`
		} `json:"Body"`
	} `json:"Envelope"`
}

type CrossItem struct {
	Reference string `xml:"reference" json:"reference"`
}

type ContainerId struct {
	XMLName xml.Name `xml:"CrossItem" json:"CrossItem"`
	CrossItem
}

type ReceiveCreatePayload struct {
	Converter string `json:"Converter"`
	Entity    string `json:"Entity"`
	Period    int    `json:"Period"`
}

type OrderBy struct {
	Property string `json:"Property"`
	Desc     string `json:"Desc"`
}

type Filters struct {
	Type        string `json:"Type"`
	NumValue    string `json:"NumValue"`
	StringValue string `json:"StringValue"`
	Property    string `json:"Property"`
	Operator    string `json:"Operator"`
}
