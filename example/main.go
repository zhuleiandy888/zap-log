/*
 * @Notice: edit notice here
 * @Author: zhulei
 * @Date: 2022-12-26 14:19:54
 * @LastEditors: zhulei
 * @LastEditTime: 2022-12-26 16:05:52
 */
package main

import (
	log "github.com/zhuleiandy888/zap-log"
)

const (
	// 日志文件
	DEFAULT_LOG = "server.log"
)

// 日志json模式开关
var jsonMode = false

// init 初始化组件
func init() {

	// 日志保存最大时间(hours)
	log.MaxAge = 24*30
	// 日志轮转时间(hours)
	log.RotationTime = 1
	// 初始化日志
	log.InitLogger(DEFAULT_LOG, log.DEBUG, jsonMode)
	defer log.SugarLogger.Sync()
}
func main() {

	log.SugarLogger.Infof("info log...")
	log.SugarLogger.Warnf("warn log...")
	log.SugarLogger.Errorf("error log...")
}
