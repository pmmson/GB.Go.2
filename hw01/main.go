package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	task01()
	err := task02()
	if err != nil {
		fmt.Println(err)
	}
	//task03(1000000) // закомментировано, так как паники не достиг при 1 млн файлов
	task04()
}

func task01() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("panic situation", err)
		}
	}()
	var arr = []int{1, 2, 3, 4, 5}
	arrNum := len(arr) + 1
	fmt.Println(arr[arrNum])
}

type errTask02 struct {
	errnum  int
	errmsg  string
	errtime time.Time
}

func (e *errTask02) Error() string {
	return fmt.Sprintf("%s EXCEPTION: %d: %s", e.errtime.Format(time.RFC3339), e.errnum, e.errmsg)
}

func New(msg string) error {
	return &errTask02{
		errnum:  -100,
		errmsg:  msg,
		errtime: time.Now(),
	}
}

func task02() error {

	var arr = []int{1, 2, 3, 4, 5}
	arrNum := len(arr) + 1
	if arrNum > len(arr) {
		return New("index out of range")
	}
	fmt.Println(arr[arrNum])
	return nil
}

func task03(cnt int) {
	for i := 0; i < cnt; i++ {
		f, _ := os.Create("/Users/pmmson/TEST/file" + strconv.Itoa(i))
		fmt.Println(f.Stat())
	}
} // создало в папке 1000000 файлов, паники не было. ulimit -n (1048575) ?

func task04() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered", v)
		}
	}()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recovered from panic", err)
			}
		}()
		panic("A-A-A!!!")
	}()
	time.Sleep(time.Second)
}
