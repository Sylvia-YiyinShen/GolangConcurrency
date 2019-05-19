package main

import (
    "fmt"
    "time"
)

func checkQuitChannel() {
	quit := make(chan bool)
	c := make(chan string)
	go func() {
		for i := 3; i >= 0; i-- { fmt.Println(<-c) }
		quit <- true
	}()
	
	for {
        select {
			case c <- "message":
			
			case <-quit:
				return 
		}
    }
}

func checkTimeout() {
	c := boring2("Joe")
    for {
        select {
        case s := <-c:
            fmt.Printf("case s: %v\n", s)
        case <-time.After(1 * time.Second):
            fmt.Println("You're too slow.")
            return
        }
    }
}

func checkBoringFanIn() {
	// c := fanIn(boring2("Joe"), boring2("Ann"))
	c := improvedFanIn(boring2("Joe"), boring2("Ann"))
    for i := 0; i < 10; i++ {
        fmt.Println(<-c)
    }
    fmt.Println("You're both boring; I'm leaving.")
}

func improvedFanIn(input1, input2 <-chan string) <-chan string {
    c := make(chan string)
    go func() {
        for {
            select {
            case s := <-input1:  c <- s
            case s := <-input2:  c <- s
            }
        }
    }()
    return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
    c := make(chan string)
    go func() { for { c <- <-input1 } }()
    go func() { for { c <- <-input2 } }()
    return c
}

func checkBoring2() {
	joe := boring2("Joe: boring")
    ann := boring2("Ann: boring")
    for i := 0; i < 5; i++ {
        fmt.Println(<-joe)
        fmt.Println(<-ann)
    }
    fmt.Println("You're both boring; I'm leaving.")
}

func boring2(msg string) <-chan string { // Returns receive-only channel of strings.
    c := make(chan string)
    go func() { // We launch the goroutine from inside the function.
        for i := 0; ; i++ {
            fullMessage := fmt.Sprintf("%s %d", msg, i)
			fmt.Println(fullMessage)
        	c <- fullMessage
            time.Sleep(time.Second)
        }
    }()
    return c // Return the channel to the caller.
}



func checkBoring1() {
	c := make(chan string)
    go boring1("boring!", c)
    for i := 0; i < 5; i++ {
        fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value.
    }
    fmt.Println("You're boring; I'm leaving.")
}

func boring1(msg string, c chan string) {
    for i := 0; ; i++ {
		fullMessage := fmt.Sprintf("%s %d", msg, i)
		fmt.Println(fullMessage)
        c <- fullMessage
        time.Sleep(time.Second)
    }
}