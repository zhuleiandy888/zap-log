# zap-log
zap Log module encapsulation

## Basic Usage:

```go
  // init 初始化组件
  func init() {
	  // 初始化日志
	  log.InitLogger(DEFAULT_LOG, jsonMode)
	  defer log.SugarLogger.Sync()
  }


```
