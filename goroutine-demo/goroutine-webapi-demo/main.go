package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	json.NewEncoder(w).Encode(&Response{"後続処理は非同期で処理"})

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
	json.NewEncoder(w).Encode(&Response{"すべて同期して処理"})
	SleepAndPrint("同期処理 1") // 時間がかかる処理があると想定
	SleepAndPrint("同期処理 2") // 時間がかかる処理があると想定
}

func SleepAndPrint(s string) {
	time.Sleep(1 * time.Second)
	fmt.Println(s)
}
