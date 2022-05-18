package gosupport

import (
	"encoding/json"
	"encoding/xml"
)

type H map[string]interface{}

// MarshalXML allows type H to be used with xml.Marshal.
func (h H) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{
		Space: "",
		Local: "xml",
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for key, value := range h {
		elem := xml.StartElement{
			Name: xml.Name{Space: "", Local: key},
			Attr: []xml.Attr{},
		}
		if err := e.EncodeElement(value, elem); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (h H)ToJson() string {
	b, err := json.Marshal(h)
	if err != nil {
		return ""
	}
	return string(b)
}

func (h H)ToJsonByte() ([]byte,error) {
	return json.Marshal(h)
}
