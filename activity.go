package vmsparseractivity

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-vmsparseractivity")

// MyActivity is a stub for your Activity implementation
type XMLParserActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &XMLParserActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *XMLParserActivity) Metadata() *activity.Metadata {
	return a.metadata
}

//XSD
type VMSInfo struct {
	XMLName      xml.Name   `xml:"VMSInfo" json:"-"`
	VMSmsgList []VMSmsg `xml:"VMSmsg" json:"VMSmsg"`
}

type VMSmsg struct {
	XMLName xml.Name `xml:"VMSmsg" json:"-"`
	Date   string   `xml:"Date" json:"Date"`
	EquipmentID   string   `xml:"EquipmentID" json:"EquipmentID"`
	LinkID    string   `xml:"LinkID" json:"LinkID"`
	Attribute    string   `xml:"Attribute" json:"Attribute"`	
	Message string   `xml:"Message" json:"Message"`
}



/*type	VMSInfo struct {
		VMSmsg []struct {
			Date        string `json:"Date"`
			EquipmentID string `json:"EquipmentID"`
			LinkID      string `json:"LinkID"`
			Attribute   string `json:"Attribute"`
			Message     string `json:"Message"`
		} `json:"VMSmsg"`
	}
*/
// end of XSD

// Eval implements activity.Activity.Eval
func (a *XMLParserActivity) Eval(ctx activity.Context) (done bool, err error) {

	XMLString := ctx.GetInput("xmlString").(string)

	activityLog.Debugf("XML String is : [%s]", XMLString)
	//fmt.Println("XML String is : ", XMLString)

	if len(XMLString) == 0 {
		activityLog.Debugf("value in the field is empty ")
		//fmt.Println("value in  the field is empty ")

	}
	//	XMLString = (string(XMLString))

	xml_data := VMSInfo{}
	err = xml.Unmarshal([]byte(XMLString), &xml_data)

	jsondata, _ := json.Marshal(xml_data)
	if err != nil {
		activityLog.Debugf("Error ", err)
		fmt.Println("error: ", err)
		return
	}

	//fmt.Println(" JSON String ")
	//fmt.Println(string(jsondata))

	// Set the output as part of the context
	activityLog.Debugf("Activity has parsed XML Successfully")
	fmt.Println("Activity has parsed XML Successfully")

	ctx.SetOutput("output", string(jsondata))

	return true, nil
}
