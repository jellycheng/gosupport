package kafka

import (
	"fmt"
	"testing"
	"time"
)

func TestNewKafkaConfig(t *testing.T) {
	kafkaConfig := NewKafkaConfig()
	fmt.Println("broker server: ", kafkaConfig.Brokers, kafkaConfig.GetBrokers())
	fmt.Println("topic:", kafkaConfig.GetTopic())
	kafkaConfig.SetTopic("testTopic")
	fmt.Println("topic:", kafkaConfig.GetTopic())
	fmt.Println("AutoOffsetReset:", kafkaConfig.AutoOffsetReset)
	fmt.Println("SessionTimeoutMs", kafkaConfig.SessionTimeoutMs)

}

func TestNewKafkaConfigManage(t *testing.T) {
	 kafkaCfgGroup := NewKafkaConfigManage()
	 go func() {
	 	//新增配置
		 for i:=0; i<10000;i++ {
			 kafkaCfgGroup.SetData("group_" + fmt.Sprint(i), NewKafkaConfig().SetTopic("topic_" + fmt.Sprint(i)))
		 }
	 }()

	 go func() {
	 	//获取配置
		for i:=10000; i>0;i-- {
			if c,err := kafkaCfgGroup.GetData("group_" + fmt.Sprint(i));err==nil {
				fmt.Println("broker server: " + c.GetBrokers() + " topic:" + c.GetTopic())
			}
		}
	 }()

	 time.Sleep(3 * time.Second)
}
