package kafka

import (
	"errors"
	"github.com/jellycheng/gosupport"
	"sync"
)

type Config struct {
	Topic            string `default:""`
	Group            string `default:""`
	Brokers          string `default:"localhost:9092"`
	AutoOffsetReset  string `default:"earliest"` //earliest、latest
	SessionTimeoutMs int64  `default:"6000"`
}

func NewKafkaConfig() *Config {
	c := &Config{}
	gosupport.InitStruct4DefaultTag(c)
	return c
}

func (c *Config) GetTopic() string {
	return c.Topic
}

func (c *Config) SetTopic(topic string) *Config {
	c.Topic = topic
	return c
}

func (c *Config) GetGroup() string {
	return c.Group
}

func (c *Config) SetGroup(group string) *Config {
	c.Group = group
	return c
}

func (c *Config) GetBrokers() string {
	brokerServer := c.Brokers
	return brokerServer
}

func (c *Config) SetBrokers(brokerServer string) *Config {
	c.Brokers = brokerServer
	return c
}

func (c *Config) GetAutoOffsetReset() string {
	autoOffsetReset := c.AutoOffsetReset
	return autoOffsetReset
}

func (c *Config) SetAutoOffsetReset(autoOffsetReset string) *Config {
	c.AutoOffsetReset = autoOffsetReset
	return c
}

type ConfigManage struct {
	data map[string]*Config
	l    sync.Mutex
}

func (cm *ConfigManage) SetData(group string, c *Config) {
	cm.l.Lock()
	defer cm.l.Unlock()
	cm.data[group] = c
}

func (cm *ConfigManage) GetData(group string) (*Config, error) {
	cm.l.Lock()
	defer cm.l.Unlock()
	if c, ok := cm.data[group]; ok {
		return c, nil
	}
	return nil, errors.New("配置组不存在：" + group)
}

// 单例
var kafkaCfgManageOnce sync.Once
var kafkaCfgManage *ConfigManage

func NewKafkaConfigManage() *ConfigManage {
	kafkaCfgManageOnce.Do(func() {
		kafkaCfgManage = &ConfigManage{data: make(map[string]*Config)}
	})
	return kafkaCfgManage
}
