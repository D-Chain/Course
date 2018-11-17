package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	//"runtime"
	"strconv"
	"time"
)

var quit chan int = make(chan int)

func mainloop(icount int) {

	start := time.Now().UnixNano() / 1e6
	url := "http://localhost:8081/txpool"

	payload := strings.NewReader("{\n    \"From\": \"1KSKahQT9n69sgqn4aVmRUPpydf6AUeeZY\",\n    \"To\": \"1EFnWYm1suorEdt5XLEJ9UMTYQjGzqmiJq\",\n    \"Value\": 1,\n    \"Data\": \"message\",\n    \"Nonce\": " + strconv.Itoa(icount) + "\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "2d31fe09-2755-4c06-bc82-645c1b33c4bb")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	end := time.Now().UnixNano() / 1e6
	fmt.Printf("执行消耗的时间为:%v毫秒,执行索引为:%d\n", end-start, icount)

	quit <- 0

}
func main() {
	//runtime.GOMAXPROCS(2)
	N := 1000

	start := time.Now().UnixNano() / 1e6
	for a := 0; a <= N; a++ {
		time.Sleep(time.Nanosecond * 1e6)
		go mainloop(a)
	}
	for i := 0; i <= N; i++ {
		<-quit
	}
	end := time.Now().UnixNano() / 1e6
	fmt.Printf("总执行消耗的时间为:%v毫秒,执行次数为:%d\n", end-start, N)

}

