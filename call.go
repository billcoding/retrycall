package retrycall

import (
	"fmt"
	"log"
	"os"
	"time"
)

var callLogger = log.New(os.Stdout, "[RetryCall]", log.LstdFlags)

func Call(fn func() interface{}) interface{} {
	return CallWithPeriod(fn, time.Millisecond*100)
}

func CallWithPeriod(fn func() interface{}, period time.Duration) interface{} {
	rc := make(chan interface{}, 1)
	go func() {
		for {
			rs := func() interface{} {
				defer func() {
					if re := recover(); re != nil {
						callLogger.Println(fmt.Sprintf("%v", re))
					}
				}()
				return fn()
			}()
			if rs != nil {
				rc <- rs
				break
			}
			time.Sleep(period)
		}
	}()
	return <-rc
}
