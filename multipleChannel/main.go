package main

import "fmt"

func message(msg string, c chan string) {
	c <- msg
}

func main(){
	c := make(chan string, 2)
	c <- "Message 1"
	c <- "Message 2"

	fmt.Println(len(c), cap(c))

	// Range and close
	close(c)

	//c <- "Message 3" // panic: send on closed channel
	for msg := range c {
		fmt.Println(msg)
	}

	//Select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("Email 1", email1)
	go message("Email 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email 1 received", m1)
		case m2 := <-email2:
			fmt.Println("Email 2 received", m2)
		}
	}

}