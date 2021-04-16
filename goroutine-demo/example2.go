package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	log.Println("start...")
	starttime := time.Now()

	type Post struct {
		title string
	}
	type User struct {
		posts []Post
	}
	var users []User = []User{
		User{posts: []Post{}},
		User{posts: []Post{Post{title: "a-1"}}},
		User{posts: []Post{Post{title: "b-1"}, Post{title: "b-2"}}},
		User{posts: []Post{Post{title: "c-1"}, Post{title: "c-2"}}},
		User{posts: []Post{Post{title: "d-1"}, Post{title: "d-2"}, Post{title: "d-3"}, Post{title: "d-4"}}},
	}

	limit := 2
	ch := make(chan struct{}, limit)

	var wg sync.WaitGroup
	wg.Add(len(users))

	for idx, u := range users {
		// ゴルーチン外のなんらかの処理
		fmt.Println("Notゴルーチン", "user", idx, u)
		ch <- struct{}{}

		// ゴルーチン
		go func(idx int, u User) {
			defer wg.Done()
			for _, p := range u.posts {
				time.Sleep(1 * time.Second)
				fmt.Println("ゴルーチン", "post", idx, p)
			}
			<-ch
		}(idx, u)

	}

	wg.Wait()

	endtime := time.Now()
	fmt.Printf("%f秒\n", (endtime.Sub(starttime)).Seconds())
	log.Println("end...")
}
