package Entities

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
)

type StockItem struct {
	XMLName   xml.Name `xml:"StockItem"`
	Text      string   `xml:",chardata"`
	Class     string   `xml:"class,attr"`
	Code      string   `xml:"code"`
	Status    string   `xml:"status"`
	Reference int      `xml:"reference"`
	Entity    struct {
		Text      string `xml:",chardata"`
		Class     string `xml:"class,attr"`
		Reference int    `xml:"reference"`
	} `xml:"entity"`
	VatRate struct {
		Text      string `xml:",chardata"`
		Class     string `xml:"class,attr"`
		Reference int    `xml:"reference"`
	} `xml:"vatRate"`
	UnitSet struct {
		Text      string `xml:",chardata"`
		Class     string `xml:"class,attr"`
		Reference int    `xml:"reference"`
	} `xml:"unitSet"`
	BaseUnit struct {
		Text      string `xml:",chardata"`
		Class     string `xml:"class,attr"`
		Reference int    `xml:"reference"`
	} `xml:"baseUnit"`
	Description string `xml:"description"`
}

type Material struct {
	XMLName          xml.Name `xml:"Material"`
	Text             string   `xml:",chardata"`
	ID               string   `xml:"id"`
	Version          string   `xml:"version"`
	CanUndoApprove   string   `xml:"canUndoApprove"`
	ChangePeriodHour string   `xml:"changePeriodHour"`
	ChangePeriodYear string   `xml:"changePeriodYear"`
	Code             string   `xml:"code"`
	CodeLocked       string   `xml:"codeLocked"`
	ConvFactor       string   `xml:"convFactor"`
	ConvFactorTep    string   `xml:"convFactorTep"`
	CreatedBy        struct {
		Text        string `xml:",chardata"`
		ID          string `xml:"id"`
		Description string `xml:"description"`
	} `xml:"createdBy"`
	Currency                    string `xml:"currency"`
	DateCreated                 string `xml:"dateCreated"`
	Detail                      string `xml:"detail"`
	EldMaintenanceLevel         string `xml:"eldMaintenanceLevel"`
	EstimatedCost               string `xml:"estimatedCost"`
	FailurePeriodHour           string `xml:"failurePeriodHour"`
	FirstMaintHour              string `xml:"firstMaintHour"`
	Gtip                        string `xml:"gtip"`
	HavingSerialLotNo           string `xml:"havingSerialLotNo"`
	Info                        string `xml:"info"`
	IsEldFollow                 string `xml:"isEldFollow"`
	IsMatNameChangeable         string `xml:"isMatNameChangeable"`
	IsPassive                   string `xml:"isPassive"`
	IsPersonnelPickingMandatory string `xml:"isPersonnelPickingMandatory"`
	LabourCost                  string `xml:"labourCost"`
	LabourCostCurrency          string `xml:"labourCostCurrency"`
	LabourManHour               string `xml:"labourManHour"`
	LastUpdated                 string `xml:"lastUpdated"`
	Leadtime                    string `xml:"leadtime"`
	LongParamName               string `xml:"longParamName"`
	MatGroup                    struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id"`
		Name      string `xml:"name"`
		OtherName string `xml:"otherName"`
	} `xml:"matGroup"`
	MatNsnCode  string `xml:"matNsnCode"`
	MatPmsCtg   string `xml:"matPmsCtg"`
	MatSpecCode string `xml:"matSpecCode"`
	MatType     struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id"`
		Name      string `xml:"name"`
		OtherName string `xml:"otherName"`
	} `xml:"matType"`
	MaterialParams []struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id"`
	} `xml:"materialParams"`
	MinLevel        string `xml:"minLevel"`
	Name            string `xml:"name"`
	OldMatCode      string `xml:"oldMatCode"`
	OldName         string `xml:"oldName"`
	OrderQty        string `xml:"orderQty"`
	OutOfCode       string `xml:"outOfCode"`
	PartNumber      string `xml:"partNumber"`
	ProcurementType struct {
		Text     string `xml:",chardata"`
		EnumType string `xml:"enumType"`
		Name     string `xml:"name"`
		Label    string `xml:"label"`
	} `xml:"procurementType"`
	ProdUnit struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id"`
		OtherName string `xml:"otherName"`
		Code      string `xml:"code"`
		Name      string `xml:"name"`
	} `xml:"prodUnit"`
	ProvideType struct {
		Text     string `xml:",chardata"`
		EnumType string `xml:"enumType"`
		Name     string `xml:"name"`
		Label    string `xml:"label"`
	} `xml:"provideType"`
	PurcUnit struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id"`
		OtherName string `xml:"otherName"`
		Code      string `xml:"code"`
		Name      string `xml:"name"`
	} `xml:"purcUnit"`
	PurchBuyer struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id"`
		Name      string `xml:"name"`
		OtherName string `xml:"otherName"`
	} `xml:"purchBuyer"`
	SafetyQty string `xml:"safetyQty"`
	ShelfLife string `xml:"shelfLife"`
	Status    struct {
		Text     string `xml:",chardata"`
		EnumType string `xml:"enumType"`
		Name     string `xml:"name"`
		Label    string `xml:"label"`
	} `xml:"status"`
	UpdatedBy struct {
		Text        string `xml:",chardata"`
		ID          string `xml:"id"`
		Description string `xml:"description"`
	} `xml:"updatedBy"`
	UseShelfLife                 string `xml:"useShelfLife"`
	Weight                       string `xml:"weight"`
	EnumeratedEvents             string `xml:"enumeratedEvents"`
	EnumeratedActiveEvents       string `xml:"enumeratedActiveEvents"`
	OnHand                       string `xml:"onHand"`
	OtherName                    string `xml:"otherName"`
	MatCategory                  string `xml:"matCategory"`
	MaterialCodeName             string `xml:"materialCodeName"`
	MaterialUnit                 string `xml:"materialUnit"`
	PurchOrderCount              string `xml:"purchOrderCount"`
	ReleasedPurchOrderCount      string `xml:"releasedPurchOrderCount"`
	PurchInquiryCount            string `xml:"purchInquiryCount"`
	PurchReqPackageCount         string `xml:"purchReqPackageCount"`
	PurchOrderPreAcceptanceCount string `xml:"purchOrderPreAcceptanceCount"`
	PurchRequestLineCount        string `xml:"purchRequestLineCount"`
	PurchPurchasePackageCount    string `xml:"purchPurchasePackageCount"`
	PurchReceiptCount            string `xml:"purchReceiptCount"`
	IssuedQtyCount               string `xml:"issuedQtyCount"`
	TotalStockCount              string `xml:"totalStockCount"`
	PurchReservedCount           string `xml:"purchReservedCount"`
	LastPreAcceptedCount         string `xml:"lastPreAcceptedCount"`
	LastPreRejectedCount         string `xml:"lastPreRejectedCount"`
}

func (stockItem *StockItem) fillClassName() {
	stockItem.Class = "com.gtech.erp.inventory.objects.StockItem"
	stockItem.Entity.Class = "com.gtech.erp.base.objects.LegalEntity"
	stockItem.VatRate.Class = "com.gtech.erp.base.objects.VatRateDefinition"
	stockItem.UnitSet.Class = "com.gtech.relax.crmbase.objects.UnitOfMeasureSet"
	stockItem.BaseUnit.Class = "com.gtech.relax.crmbase.objects.UnitOfMeasure"
}
func MaterialStockItemConverter(material string) (string, error) {
	var newMaterial = Material{}
	err := json.Unmarshal([]byte(material), &newMaterial)

	var stockItem = StockItem{}
	stockItem.Code = newMaterial.Name
	stockItem.Status = "0"
	stockItem.Reference = 758
	stockItem.Entity.Reference = 1
	stockItem.VatRate.Reference = 1
	stockItem.UnitSet.Reference = 7
	stockItem.BaseUnit.Reference = 22
	stockItem.Description = newMaterial.Detail
	stockItem.fillClassName()

	stockItemJSON, err := xml.Marshal(stockItem)
	return base64.StdEncoding.EncodeToString(stockItemJSON), err
}
