package main

import (
	"libs/mongo"
	"push/conf"
)

func main() {
	//全量push服务标准化

	//1:读取配置文件
	cfg := conf.GetConf()

	mongo.Conn(cfg.MogodbUrl) //返回连接对象，并且全局保持唯一

	//先扫描全量的数据，并且根据对应的协程数量生成对应对量的文件

	//根据flag拿消息id

	//匹配对应的协程来进行任务处理

	//保持当前进程不关闭
}
