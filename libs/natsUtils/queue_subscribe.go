package natsUtils

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"teacupapi/config"
	glog "teacupapi/logs"
	"time"
)

//主题队里，一个消息只能被一个订阅消费
func QueueSubscribe(subject, qgroup string, messageHandlerFunc stan.MsgHandler, startOpts ...stan.SubscriptionOption) {
	subject = config.GetBaseConf().Env + "_" + subject
	durableName := subject + "_queue"
	//消息接收处理函数
	if messageHandlerFunc == nil {
		panic(fmt.Sprintf("订阅消息主题【%s】没有消息处理函数", subject))
	}
	//启动配置
	if len(startOpts) == 0 {
		startOpts = []stan.SubscriptionOption{
			stan.SetManualAckMode(),
			stan.AckWait(30 * time.Second),
			stan.StartWithLastReceived(),
			stan.DurableName(durableName),
			//设置1为有序处理数据
			stan.MaxInflight(100),
		}
	}

	sub, err := subStreamingConnect.QueueSubscribe(subject, qgroup, messageHandlerFunc, startOpts...)
	if err != nil {
		glog.Errorf("订阅subj：%s,err:%s", subject, err)

		err := subStreamingConnect.Close()
		if err != nil {
			glog.Errorf("关闭 nats streaming 订阅subj：%s,err:%s", subject, err)
		}
	}
	glog.Infof("Listening on [%s], qgroup=[%s] durable=[%s]\n", subject, qgroup, durableName)
	subList = append(subList, sub)
}
