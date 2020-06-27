package main

import (
	_ "bytes"
	"encoding/json"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"io/ioutil"
	"log"
	"net/http"
)

// Todo struct
type ActivityDto struct {
	ActivityId   int    `json:"activityId"`
	ActivityName string `json:"activityName"`
}


func main() {
	s := gocron.NewScheduler()
	s.Every(2).Hours().Do(get)
	<-s.Start()
}


func get() {
	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get("http://reminderapidev-env.eba-5pppvizn.ap-northeast-1.elasticbeanstalk.com/api/activityruntime/get")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var activityDtoStruct ActivityDto
	json.Unmarshal(bodyBytes, &activityDtoStruct)
	fmt.Printf("API Response as struct %+v\n", activityDtoStruct)
}

//func post() {
//	fmt.Println("2. Performing Http Post...")
//	todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
//	jsonReq, err := json.Marshal(todo)
//	resp, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	defer resp.Body.Close()
//	bodyBytes, _ := ioutil.ReadAll(resp.Body)
//
//	// Convert response body to string
//	bodyString := string(bodyBytes)
//	fmt.Println(bodyString)
//
//	// Convert response body to Todo struct
//	var todoStruct Todo
//	json.Unmarshal(bodyBytes, &todoStruct)
//	fmt.Printf("%+v\n", todoStruct)
//}
