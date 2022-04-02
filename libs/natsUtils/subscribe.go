package natsUtils

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"teacupapi/config"
	glog "teacupapi/logs"
	"time"
)

//普通订阅，所有订阅都收到相同的消息
func Subscribe(subject string, messageHandlerFunc stan.MsgHandler, startOpts ...stan.SubscriptionOption) {
	subject = config.GetBaseConf().Env + "_" + subject
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
			//设置1为有序处理数据
			stan.MaxInflight(100),
			stan.DurableName(subject),
		}
	}

	sub, err := subStreamingConnect.Subscribe(subject, messageHandlerFunc, startOpts...)
	if err != nil {
		glog.Errorf("订阅subj：%s,err:%s", subject, err)

		err := subStreamingConnect.Close()
		if err != nil {
			glog.Errorf("关闭 nats streaming 订阅subj：%s,err:%s", subject, err)
		}
	}
	glog.Infof("Listening on [%s],  durable=[%s]\n", subject, subject)
	subList = append(subList, sub)
}
