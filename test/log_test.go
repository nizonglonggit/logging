package test

import (
	"github.com/nizonglonggit/logging/logging"
	"testing"
	"time"
)

func TestLog(t *testing.T) {

	//logging.InitLog("./info.log", "./error.log", zap.InfoLevel)

	logging.SetMultiLog(map[int]string{logging.INFO: "./info.log", logging.ERROR: "./error.log", logging.DEBUG: "./debug.log", logging.WARN: "./warn.log"})
	logging.Info(logging.GetLogLevel())

	//logging.SetLogLevel(logging.DEBUGLevel)
	count := 0

	for {
		logging.Debugf("name:%d", count)
		logging.Infof("name:%d", count)
		logging.Errorf("name:%d", count)
		logging.Warnf("name:%d", count)

		count++
		time.Sleep(time.Millisecond * 100)

		if count >= 10 {
			break
		}
	}
}
