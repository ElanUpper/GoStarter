package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logg *log.Logger

//------------------------------ cancel context ------------------------------
func someHandler() {
	ctx, cancelF := context.WithCancel(context.Background())
	go doStuff(ctx)

	//10秒后取消doStuff
	time.Sleep(10 * time.Second)
	cancelF()

}

//------------------------------ deadline/timeout context ------------------------------
func timeoutHandler() {

	// 5秒后timeout 结束doStuff
	//ctx, cancelF := context.WithTimeout(context.Background(), 5 * time.Second)
	ctx, cancelF := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))

	//doStuff(ctx)
	doTimeoutStuff(ctx)

	// 整体等待10秒
	time.Sleep(10 * time.Second)
	cancelF()
}

func doTimeoutStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)

		if deadline, ok := ctx.Deadline(); ok {
			logg.Printf("deadline set %v\n", deadline)
			// 设置error
			if time.Now().After(deadline) {
				log.Printf(ctx.Err().Error())
				return
			}
		}

		select {
		case <-ctx.Done():
			log.Printf("done")
			return
		default:
			log.Printf("work")
		}
	}
}

//------------------------------ withvalue context ------------------------------

//// NewContext returns a new Context carrying userIP.
//func NewContext(ctx context.Context, userIP net.IP) context.Context {
//	return context.WithValue(ctx, userIPKey, userIP)
//}
//
//// FromContext extracts the user IP address from ctx, if present.
//func FromContext(ctx context.Context) (net.IP, bool) {
//	// ctx.Value returns nil if ctx has no value for the key;
//	// the net.IP type assertion returns ok=false for nil.
//	userIP, ok := ctx.Value(userIPKey).(net.IP)
//	return userIP, ok
//}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logg.Printf("done")
			return
		default:
			logg.Printf("work")
		}
	}
}

func main() {
	logg = log.New(os.Stdout, "", log.Ltime)
	timeoutHandler()
	logg.Printf("down")
}
