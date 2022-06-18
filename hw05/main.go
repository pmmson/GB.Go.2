/*
1. Напишите программу, которая запускает 𝑛 потоков и дожидается завершения их всех
2. Реализуйте функцию для разблокировки мьютекса с помощью defer
3. Протестируйте производительность операций чтения и записи на множестве
	действительных чисел, безопасность которого обеспечивается sync.Mutex и sync.RWMutex
	для разных вариантов использования: 10% запись, 90% чтение; 50% запись, 50% чтение; 90% запись, 10% чтение
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("Задание 1")
	fmt.Print("Введите n (длину случайно генерируемой строки): ")
	var n int
	fmt.Scan(&n)
	genStr := randStr(n)
	fmt.Println(genStr, "| длина строки:", len(genStr))

	fmt.Println("Задание 2")
	var statRunesInStr = make(map[rune]int)
	var mux sync.Mutex
	var wg sync.WaitGroup
	for _, v := range genStr {
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
	for key, i := range statRunesInStr /*statRunesInStr*/ {
		sum += i
		fmt.Printf("%s - %d; ", string(key), i)
	}
	fmt.Println("\nВсего rune в сумме:", sum)

	fmt.Println("Задание 3")

}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(n int) string {
	b := make([]rune, n)
	var wg = sync.WaitGroup{}
	for i := range b {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rand.Seed(time.Now().UnixMicro())
			b[i] = letters[rand.Intn(len(letters))]
		}()
		wg.Wait()
	}
	return string(b)
}
