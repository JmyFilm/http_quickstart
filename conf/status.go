package conf

import (
	"time"
)

func init() {
	AfterConfigExec(func() {
		go execStatusRecord()
	})
}

func execStatusRecord() {
	ticker := time.NewTicker(time.Duration(Log.StatusTickerTime) * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:

			RecordZero()
		}
	}
}

func RecordZero() {

}
