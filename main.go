package main

import (
	"fmt"
	"time"
)

func Func(str string) {
	fmt.Print(str + " ")
}

func main() {
	// とても健全な言葉なのに、goroutineで動かすと...?
	for i := 0; i < 10; i++ {
		go Func("お")
		go Func("ちん")
		go Func("ぎん")
		go Func("ほしい")
		go Func("お")
		go Func("ちん")
		go Func("ぎん")
		go Func("ほしい")
		time.Sleep(time.Second)
		Func("\n")
	}
	fmt.Scanln()
}
