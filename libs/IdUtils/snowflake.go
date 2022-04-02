package IdUtils

import (
	"github.com/bwmarrin/snowflake"
	"strconv"
	"sync"
	"teacupapi/utils"
	"time"
)

var (
	nodePool = sync.Pool{New: func() interface{} {
		nodeNum := utils.GetLocalIpToInt() % 1023
		node, err := snowflake.NewNode(nodeNum)
		if err != nil {
			panic(err)
		}
		return node
	}}
)

func GetId() int64 {
	node := nodePool.Get()
	if node != nil {
		node1 := node.(*snowflake.Node)
		return node1.Generate().Int64()
	}
	return time.Now().Unix()
}

func GetStringId() string {
	node := nodePool.Get()
	if node != nil {
		node1 := node.(*snowflake.Node)
		return node1.Generate().String()
	}
	return strconv.FormatInt(time.Now().Unix(), 10)
}
