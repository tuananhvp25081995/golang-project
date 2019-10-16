package main

import (
	"context"
	"fmt"
	"time"
)

// func main() {
// 	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
// 	doSomeThing(ctx)
// }

// func doSomeThing(ctx context.Context) {
// 	canceledChannel := make(chan bool)
// 	go func() {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println(ctx.Err())
// 			canceledChannel <- true
// 			return
// 		}
// 	}()
// 	isCanceledChannel := <-canceledChannel
// 	if isCanceledChannel {
// 		close(canceledChannel)
// 		return
// 	}
// 	time.Sleep(time.Second * 10)
// 	fmt.Println("Done")
// }

// func main() {
// 	ctx := context.WithValue(context.Background(), "number", 10)
// 	A(ctx)
// }

// func A(ctx context.Context) {
// 	if value := ctx.Value("number"); value != nil {
// 		ctx := context.WithValue(ctx, "number1", 5)
// 		B(ctx)
// 	}
// }

// func B(ctx context.Context) {
// 	value := ctx.Value("number").(int)
// 	value1 := ctx.Value("number1").(int)
// 	fmt.Println(value1 + value)
// }

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	time.AfterFunc(time.Second*5, func() {
		cancel()
	})

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
