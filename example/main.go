/*
 * @Notice: edit notice here
 * @Author: zhulei
 * @Date: 2022-12-26 14:19:54
 * @LastEditors: zhulei
 * @LastEditTime: 2022-12-26 14:26:57
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
	// 初始化日志
	log.InitLogger(DEFAULT_LOG, jsonMode)
	defer log.SugarLogger.Sync()
}
func main() {
	log.SugarLogger.Infof("info log...")
	log.SugarLogger.Warnf("warn log...")
	log.SugarLogger.Errorf("error log...")
}
