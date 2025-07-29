/**--------------------------------------------------------
 * Author: jiang5630@outlook.com 2025-07-29
 * Description:
 * This file is part of the NTPack project.
 * --------------------------------------------------------
 * 作者：jiang5630@outlook.com  2025年07月29日
 * 描述：通用类型定义
 --------------------------------------------------------*/

package NTPack

import (
	"time"
)

// TCBComponentHeader 组件头信息结构，用于标识自身和发送报文
type TCBComponentHeader struct {
	ComponentId  string    //组件名
	PublisherID  string    //发布消息的署名
	Version      string    //版本号
	Intro        string    //介绍
	Author       string    //作者名称
	StartTime    time.Time //启动时间
	SnowID       string    //雪花ID
	ServerNodeId int64     //服务器节点ID
	Pid          int       //进程ID
}
