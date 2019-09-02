package test

import (
	"log"
	"strconv"
	"testing"
)

func TestUrl(t *testing.T) {
	for i := 2; i <= 83; i++ {
		url := "http://baixarquadrinhos.com/page/" + strconv.Itoa(i) + "/"
		log.Println(url)
	}
}
