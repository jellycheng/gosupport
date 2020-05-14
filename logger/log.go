package logger

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	//默认日志格式 [级别]: 时间 - 日志内容
	defaultLogFormat       = "[%level%]: %time% - %msg%"
	defaultTimestampFormat = time.RFC3339
)

type SimpleTextFormatter struct {
	//时间格式
	TimestampFormat string
	//日志格式
	LogFormat string
}

func (f *SimpleTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	output := f.LogFormat
	if output == "" {
		output = defaultLogFormat
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	output = strings.Replace(output, "%time%", entry.Time.Format(timestampFormat), 1)
	output = strings.Replace(output, "%msg%", entry.Message, 1)
	level := strings.ToUpper(entry.Level.String())
	output = strings.Replace(output, "%level%", level, 1)

	for k, val := range entry.Data {
		switch v := val.(type) {
			case string:
				output = strings.Replace(output, "%"+k+"%", v, 1)
			case int:
				s := strconv.Itoa(v)
				output = strings.Replace(output, "%"+k+"%", s, 1)
			case bool:
				s := strconv.FormatBool(v)
				output = strings.Replace(output, "%"+k+"%", s, 1)
			default:
				s := fmt.Sprintf("%+v", v)
				output = strings.Replace(output, "%"+k+"%", s, 1)
		}
	}

	return []byte(output), nil
}

