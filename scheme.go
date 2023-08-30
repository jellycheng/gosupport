package gosupport

import "fmt"

type SchemeInfo struct {
	Protocol string `json:"protocol"` //协议，如：http、https、tcp
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func (m SchemeInfo) ToString() string {
	s := ""
	if m.Protocol != "" {
		s += m.Protocol + "://"
	}
	s += m.Host
	if m.Port > 0 {
		s = fmt.Sprintf("%s:%d", s, m.Port)
	}
	return s
}

type SchemeInfoOption func(info *SchemeInfo)

func NewSchemeInfo(opts ...SchemeInfoOption) *SchemeInfo {
	schmeInfo := new(SchemeInfo)
	for _, opt := range opts {
		opt(schmeInfo)
	}
	return schmeInfo
}

func WithSchemeInfoProtocol(s string) SchemeInfoOption {
	return func(m *SchemeInfo) {
		m.Protocol = s
	}
}

func WithSchemeInfoHost(s string) SchemeInfoOption {
	return func(m *SchemeInfo) {
		m.Host = s
	}
}

func WithSchemeInfoPort(d int) SchemeInfoOption {
	return func(m *SchemeInfo) {
		m.Port = d
	}
}
