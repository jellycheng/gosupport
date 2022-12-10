package gosupport

import (
	"bytes"
	"encoding/xml"
	"io"
)

func XmlUnmarshal(str string, obj interface{}) error {
	return xml.Unmarshal([]byte(str), obj)
}

type xmlStringMap map[string]string

type xmlMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

func (m xmlStringMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range m {
		_ = e.Encode(xmlMapEntry{XMLName: xml.Name{Local: k}, Value: v})
	}

	return e.EncodeToken(start.End())
}

func (m *xmlStringMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = xmlStringMap{}
	for {
		var e xmlMapEntry

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		(*m)[e.XMLName.Local] = e.Value
	}
	return nil
}

func Map2XML(kvs map[string]string) (text []byte, err error) {
	text, err = xml.Marshal(xmlStringMap(kvs))
	if err != nil {
		return
	}
	text = bytes.ReplaceAll(text, []byte(`<xmlStringMap>`), []byte(`<xml>`))
	text = bytes.ReplaceAll(text, []byte(`</xmlStringMap>`), []byte(`</xml>`))

	return
}

func Map2XMLV2(params map[string]string, rootName ...string) string {
	var xmlRootName = ""
	if len(rootName) > 0 {
		xmlRootName = rootName[0]
	}
	if xmlRootName == "" {
		xmlRootName = "xml"
	}
	var buf bytes.Buffer
	buf.WriteString(`<` + xmlRootName + `>`)
	for k, v := range params {
		buf.WriteString(`<`)
		buf.WriteString(k)
		buf.WriteString(`><![CDATA[`)
		buf.WriteString(v)
		buf.WriteString(`]]></`)
		buf.WriteString(k)
		buf.WriteString(`>`)
	}
	buf.WriteString(`</` + xmlRootName + `>`)

	return buf.String()
}

func XML2Map(text []byte) (result map[string]string, err error) {
	err = xml.Unmarshal(text, (*xmlStringMap)(&result))
	if err != nil {
		return
	}
	return
}

type CDATA string

func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}
