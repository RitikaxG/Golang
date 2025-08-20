package main

import (
	"context"
	"fmt"
	"time"
)

/* Context : Signal/ Extra info that travels with your function calls and goroutines

1. Cancelling Work (stop goroutines if they’re no longer needed)
2. Setting timeouts or deadlines (stop waiting after some time).
3. Carrying request-scoped values (passing metadata like request ID, user info).

*/

func GoContext() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3) // This ctx fires up after 3s

	ch := make(chan string)
	go func() { // This fn sends hi there to channel after 2s
		time.Sleep(time.Second * 2)
		ch <- "hi there"
	}()

	select {
	case <-ctx.Done(): // If context fires up 1st print hello alse print what's in channel
		fmt.Println("hello")
	case x := <-ch:
		fmt.Println(x)
	}
}

func GoContextCancel() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3) // context will fire up after 3s
	ch := make(chan string)

	go func() { // Prints hi there after 2s
		time.Sleep(time.Second * 2)
		fmt.Println("Hi there")
	}()

	cancel() // Immmediately fires up context ( If u cancel a context, it is marked done )

	select {
	case <-ctx.Done():
		fmt.Println("hello")
	case x := <-ch:
		fmt.Println(x)
	}
}
