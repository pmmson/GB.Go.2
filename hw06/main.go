/*
1. Написать программу, которая использует мьютекс для безопасного доступа к данным из нескольких потоков.
   Выполните трассировку программы
2. Написать многопоточную программу, в которой будет использоваться явный вызов планировщика.
   Выполните трассировку программы
3. Смоделировать ситуацию “гонки”, и проверить программу на наличии “гонки”
*/

package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

const str = "RWoQPUlisUmdENnVRgvDVYPLPpsRCGTWjnPUNctmRSXqJqvrexpNeyGPwCpzvtJYHkiuzeKMNfPwkyYoTarNPcjjBmVtcEnxZxEL"

func main() {
	// использовал решение урока 5
	for {
		fmt.Println("Для запуска демо 1 используйте: go run hw06/main.go 2>hw06/trace.out\n потом go tool trace hw06/trace.out")
		fmt.Println("Для запуска демо 2 используйте: go run hw06/main.go 2>hw06/trace2.out\n потом go tool trace hw06/trace2.out")
		fmt.Println("Для запуска демо 3 используйте: go run -race hw06/main.go")
		fmt.Println("Для выхода введите 0")

		var n int
		fmt.Print("введите номер задания для демо: ")
		fmt.Scan(&n)

		switch n {
		case 1:
			fmt.Println("Задание 1")
			// go run hw06/main.go 2>hw06/trace.out
			// go tool trace hw06/trace.out

			trace.Start(os.Stderr)
			defer trace.Stop()

			var statRunesInStr = make(map[rune]int)
			var mux sync.Mutex
			var wg sync.WaitGroup
			for _, v := range str {
				wg.Add(1)
				go func(r rune) {
					mux.Lock()
					defer mux.Unlock()
					statRunesInStr[r]++
					wg.Done()
				}(v)
			}
			wg.Wait()
			var sum int
			for key, i := range statRunesInStr {
				sum += i
				fmt.Printf("%s - %d; ", string(key), i)
			}
			fmt.Println("\nВсего rune в сумме:", sum)
			return
		case 2:
			fmt.Println("Задание 2")
			// go run hw06/main.go 2>hw06/trace.out
			// go tool trace hw06/trace.out

			trace.Start(os.Stderr)
			defer trace.Stop()

			for i := 0; i < 10; i++ {
				go func(u int) {
					fmt.Println("I want to work", u)
				}(i)
				runtime.Gosched()
				go func(u int) {
					fmt.Println("I want to work too", u)
				}(i)
			}
			return
		case 3:
			fmt.Println("Задание 3")
			// убрал sync.WaitGroup - получил состояние гонки
			// go run -race hw06/main.go

			var stRunesInStr = make(map[rune]int)
			var mx sync.Mutex

			for _, v := range str {
				go func(r rune) {
					mx.Lock()
					defer mx.Unlock()
					stRunesInStr[r]++

				}(v)
			}
			var s int
			for key, i := range stRunesInStr {
				s += i
				fmt.Printf("%s - %d; ", string(key), i)
			}
			fmt.Println("\nВсего rune в сумме:", s)
			return
		case 0:
			return
		}
	}
}
