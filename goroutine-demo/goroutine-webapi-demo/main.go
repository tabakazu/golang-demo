// Usage
//   $ go run main.go
//   $ curl http://localhost:8080/goroutine
//   $ curl http://localhost:8080/not_goroutine

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	// curl http://localhost:8080/goroutine
	http.HandleFunc("/goroutine", GoroutineHandler)
	// curl http://localhost:8080/not_goroutine
	http.HandleFunc("/not_goroutine", NotGoroutineHandler)

	http.ListenAndServe(":8080", nil)
}

type Response struct {
	Name string
}

// GoroutineHandler ある処理を非同期化したハンドラー
func GoroutineHandler(w http.ResponseWriter, r *http.Request) {
	defer log.Println("finish GoroutineHandler.")
	log.Printf("%v - %v", r.Method, r.URL)

	// 時間のかかる事前作業を並行実行して同期
	res := func() *Response {
		ret := &Response{"後続処理は非同期で処理"}
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			SleepAndPrint("事前処理 1")
			ret.Name = ret.Name + "!"
		}()
		go func() {
			defer wg.Done()
			SleepAndPrint("事前処理 2")
			ret.Name = ret.Name + "!!"
		}()

		wg.Wait()
		return ret
	}()
	json.NewEncoder(w).Encode(res)

	// 遅延処理で問題ない時間のかかる後続作業を非同期処理
	go func() {
		SleepAndPrint("遅延処理 1") // 時間がかかる処理があると想定
	}()
	go func() {
		SleepAndPrint("遅延処理 2") // 時間がかかる処理があると想定
	}()
}

// NotGoroutineHandler 非同期化してないハンドラー
func NotGoroutineHandler(w http.ResponseWriter, r *http.Request) {
	defer log.Println("finish NotGoroutineHandler.")
	log.Printf("%v - %v", r.Method, r.URL)

	res := func() *Response {
		ret := &Response{"すべて同期して処理"}

		SleepAndPrint("事前処理 1")
		ret.Name = ret.Name + "!"
		SleepAndPrint("事前処理 2")
		ret.Name = ret.Name + "!!"

		return ret
	}()
	json.NewEncoder(w).Encode(res)

	SleepAndPrint("同期処理 1") // 時間がかかる処理があると想定
	SleepAndPrint("同期処理 2") // 時間がかかる処理があると想定
}

func SleepAndPrint(s string) {
	time.Sleep(1 * time.Second)
	log.Println(s)
}
