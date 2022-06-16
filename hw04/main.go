package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	/*
		С помощью пула воркеров написать программу, которая запускает
		1000 горутин, каждая из которых увеличивает число на 1.
		Дождаться завершения всех горутин и убедиться, что при
		каждом запуске программы итоговое число равно 1000
	*/
	fmt.Println("Задание 1")
	for i := 0; i < 10; i++ {
		var a int
		var workers = make(chan int, 1)
		for i := 0; i <= 1000; i++ {
			workers <- 0
			go func() {
				defer func() {
					<-workers
				}()
				a++
			}()
		}
		fmt.Printf("a = %d\n", a)
	}

	/*
		Написать программу, которая при получении в канал сигнала SIGTERM
		останавливается не позднее, чем за одну секунду (установить таймаут).
	*/
	fmt.Println("Задание 2")
	ch := make(chan os.Signal, 1)
	exitCh := make(chan string)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

	pid := os.Getpid()
	log.Println("PID:", pid)

	go func() {
		for {
			val := <-ch
			switch val {
			case syscall.SIGTERM:
				log.Println("SIGTERM recieved:", val)
				time.Sleep(2 * time.Second)
				log.Println("Done!:", val)
				exitCh <- "SIGTERM"
			default:
				log.Println("Not a SIGTERM:", val)
			}
		}
	}()

	select {
	case <-exitCh:
		return
	}
}
