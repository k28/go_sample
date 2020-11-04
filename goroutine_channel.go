package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	/*
		チャンネル使ったgorutineの停止
	*/
	queue := make(chan string)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go fetchURL(queue)
	}

	queue <- "https://example.com"
	queue <- "https://example.net"
	queue <- "https://example.net/foo"
	queue <- "https://example.net/bar"

	close(queue) // goroutineに終了を伝える
	wg.Wait()    // すべてのgoroutineが終了するのを待つ
}

func fetchURL(queue chan string) {
	for {
		url, more := <-queue // closeされるとmoreがfalseになる
		if more {
			// url取得処理
			fmt.Println("fetching... ", url)
		} else {
			fmt.Println("worker exit")
			wg.Done()
			return
		}
	}
}
