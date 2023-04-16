package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// TODO:
// 		1) основной синтаксис языка прграммирования Golang
// 		2) минимум демонстрации
// 		3) пройдемся по https://go.dev/tour/list

func foo(i int) int {
	return 1 + i
}

// entrypoint
func main() { // <---- main func (точка входа)
	defer fmt.Println("world") // отложенная функция
	fmt.Println("Hello from Go compiler")

	// CSP - Охара Алгебра процессов

	var t int // по умолчанию 0 GLOBAL VAR
	go func() {
		t = foo(10) // происходит замыкание (closes)
	}()
	time.Sleep(time.Second * 1) // блокируемся на 1 секунду

	fmt.Printf("Что здесь выведет?: <%d>\n", t)

	req, _ := http.NewRequest(http.MethodPost, "/", nil)
	defer func() {
		if err := req.Body.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	req.Header.Add("Content-Type", "application/json")

	bb, _ := io.ReadAll(req.Body)

	tt := map[string]string{}
	if err := json.Unmarshal(bb, &tt); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//main()_____________________________|__________________
//			\foo()__________________/ (отдельная горутина aka корутина)
