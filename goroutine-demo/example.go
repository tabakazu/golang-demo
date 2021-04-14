package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("start...")

	var wg sync.WaitGroup

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
	}

	wg.Add(len(users))

	for idx, u := range users {
		fmt.Println("user", idx, u)

		// ゴルーチン
		go func(idx int, u User) {
			for _, p := range u.posts {
				time.Sleep(1 * time.Second)
				fmt.Println("post", idx, p)
			}

			wg.Done()
		}(idx, u)

	}

	wg.Wait()
	fmt.Println("end...")

}
