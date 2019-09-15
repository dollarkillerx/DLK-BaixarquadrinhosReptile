package test

import (
	"fmt"
	"github.com/dollarkillerx/easyutils/compression"
	"log"
	"testing"
	"time"
)

func TestCh(t *testing.T) {
	ints := make(chan int, 1000)

	go func(ch chan int) {
		t.Log("sadsa")
		for i := 0; i <= 100; i++ {
			ints <- i
			log.Println("插入数据", i)
		}
		// 往chan 中插入数据100条数据 如果塞满就 关闭chan
		close(ch)
		fmt.Println("通道关闭")
	}(ints)

	go func(ch chan int) {
		for {
			select {
			case a, ok := <-ch:
				if ok {
					log.Println(a)
				} else {
					panic("完毕")
				}
			}
			time.Sleep(time.Millisecond * 200)
		}
	}(ints)

	time.Sleep(1000 * time.Second)
}


func TestZip(t *testing.T) {
	data := "/Users/cpx/Github/Work/DLK-BaixarquadrinhosReptile/img/A Poderosa Thor 21 – Jason Aaron/d079a9070ec73d5a07449478797538bda48645adf3fdf3f33459395625e0476e/one.zip"
	unzip := compression.Unzip(data, ".")
	if unzip!= nil {
		log.Println(unzip)
	}
}