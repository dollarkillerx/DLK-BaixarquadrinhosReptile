package ch_test

import (
	"github.com/dollarkillerx/easyutils/clog"
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestCh(t *testing.T) {
	// 生成测试数据
	a := []int{}
	for i:=0;i<100;i++ {
		a = append(a,i)
	}

	sy1 := sync.WaitGroup{}

	// go 打印测试数据

	chi := make(chan int,10)

	chi <- 1
	for _,i := range a {
		sy1.Add(1)
		go func(i int) {
			defer func() {
				<- chi

				sy1.Done()
			}()
			clog.Println(strconv.Itoa(runtime.NumGoroutine()))
			time.Sleep(time.Second)

		}(i)
	}

	sy1.Wait()

	clog.Println(strconv.Itoa(runtime.NumGoroutine()))


}
