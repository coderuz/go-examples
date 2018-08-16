package main

// import "fmt"
import (
	"net/http"
	"net/url"
	"log"
	"fmt"
	"io/ioutil"
)

func main(){
	for i:=10; i<60; i++ {
		print(i)
		go fmt.Println(sendRequest("2", "0", "2018-08-14 17:18:00", "BB:0C:D0:15"))
	}
		for i:=10; i<60; i++ {
		print(i)
		go fmt.Println(sendRequest("2", "0", "2018-08-14 17:18:00", "BB:0C:D0:15"))
	}
		for i:=10; i<60; i++ {
		print(i)
		go fmt.Println(sendRequest("2", "0", "2018-08-14 17:18:00", "BB:0C:D0:15"))
	}
		for i:=10; i<60; i++ {
		print(i)
		go fmt.Println(sendRequest("2", "0", "2018-08-14 17:18:00", "BB:0C:D0:15"))
	}
}


func sendRequest(EMP_ID string, EVENT_CODE string, DT string, DEVICE_SN string) string {
	resp, err := http.PostForm("http://localhost:8080/event",
	url.Values{"EMP_ID": {EMP_ID}, "EVENT_CODE":{EVENT_CODE}, "DT":{DT}, "DEVICE_SN":{DEVICE_SN}})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	result := ""
	for _, char := range(body){
		result += string(char)
	}
	return result
}
