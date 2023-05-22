package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//changeStrSlice()
	producerAndConsumer()
}

func changeStrSlice() {
	strList := []string{"I", "am", "stupid", "and", "weak"}
	for i, str := range strList {
		if strings.EqualFold(str, "stupid") {
			strList[i] = "smart"
		} else if strings.EqualFold(str, "weak") {
			strList[i] = "strong"
		}

		fmt.Printf("%s ", strList[i])
	}
	fmt.Println()
}

func producerAndConsumer() {
	data := make(chan int, 10)
	done := make(chan bool)

	// consumer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("child process interrupt......")
				return
			default:
				fmt.Printf("receive msg : %d\n", <-data)
			}
		}
	}()

	i := 0
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		data <- i
		i++
		if i == 20 {
			break
		}
	}
	close(done)
	time.Sleep(5 * time.Second)
	fmt.Println("main process interrupt ......")
	time.Sleep(time.Second)
}
