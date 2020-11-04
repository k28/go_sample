package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	/*
		contextを使ってgoroutineを停止する方法
	*/
	fetchWithCancel()
	//fetchWithTimeout()
}

// contextを使って非同期処理をハンドリングする場合
func fetchWithCancel() {
	// キャンセルするためのContextを生成
	ctx, cancel := context.WithCancel(context.Background())
	queue := make(chan string)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go fetchURL(ctx, queue)
	}

	queue <- "https://www.example.com"
	queue <- "https://www.example.net"
	queue <- "https://www.example.net/foo"
	queue <- "https://www.example.net/bar"

	cancel()  // ctxを終了させる
	wg.Wait() // すべてのgoroutineが終了するのを待つ
}

// タイムアウトを設けて処理する場合
func fetchWithTimeout() {
	// キャンセルするためのContextを生成
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	queue := make(chan string)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go fetchURL(ctx, queue)
	}

	queue <- "https://www.example.com"
	queue <- "https://www.example.net"
	queue <- "https://www.example.net/foo"
	queue <- "https://www.example.net/bar"

	cancel()  // ctxを終了させる
	wg.Wait() // すべてのgoroutineが終了するのを待つ
}

func fetchURL(ctx context.Context, queue chan string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker exit")
			wg.Done()
			return
		case url := <-queue:
			// url取得処理
			fmt.Println("fetching", url)
		}
	}
}
