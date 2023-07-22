package gosupport

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
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

func (h H) ToJson() string {
	b, err := json.Marshal(h)
	if err != nil {
		return ""
	}
	return string(b)
}

func (h H) ToJsonByte() ([]byte, error) {
	return json.Marshal(h)
}

func (h H) WriteXml(w http.ResponseWriter) H {
	h.WriteContentType(w, []string{"application/xml; charset=utf-8"})
	return h
}

func (h H) WriteJson(w http.ResponseWriter) H {
	h.WriteContentType(w, []string{"application/json; charset=utf-8"})
	return h
}

func (h H) WriteContentType(w http.ResponseWriter, value []string) H {
	if len(value) == 0 {
		value = []string{"application/json; charset=utf-8"}
	}
	WriteContentType(w, value)
	return h
}

func WriteContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

// const RedisKey01 StringFormat = "user:%s:%d"; s := RedisKey01.Format("info", 123);
type StringFormat string

func (m StringFormat) Format(args ...interface{}) string {
	return fmt.Sprintf(string(m), args...)
}
