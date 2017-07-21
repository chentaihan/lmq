package main

import (
	"fmt"
	"time"
	"strconv"
	"encoding/json"

	"lmq/lmq"
	"lmq/util"
	"lmq/db"
	"lmq/api"
	"net/http"
)

type Person struct{
	Name string
	Age int
}

func test(){
	var i int = 10;
	fmt.Println(i);
	for i := 1; i <= 10; i++{
		fmt.Println(i)
	}

	personList := make([]Person, 4);

	for i := 0 ; i < len(personList) ; i++  {
		personList[i] = Person{"chentaihan"+ strconv.FormatInt(int64(i), 10), i}
	}
	parseStr, err := json.Marshal(personList);
	if err != nil{
		fmt.Println("json.Marshal error");
		return;
	}
	fmt.Println(string(parseStr))
	fmt.Println("--------------------------------------")
	personList1 := make([]Person,0)
	json.Unmarshal(parseStr, &personList1)
	parseStr1, err := json.Marshal(personList1);
	if err != nil{
		fmt.Println("json.Marshal error");
		return;
	}
	fmt.Println(string(parseStr1))
	lmq.LoadPlatform()
	parseStr11, err := json.Marshal(lmq.PlatformList);
	if err != nil{
		fmt.Println("json.Marshal error");
		return;
	}
	fmt.Println(string(parseStr11))
	lmq.LoadPlatform()
	fmt.Println(lmq.AddPlatform(lmq.PlatformItem{ Platform:"ORP", Module:"datatrans" }))
	fmt.Println(lmq.AddPlatform(lmq.PlatformItem{ Platform:"ORP", Module:"dataman" }))
	fmt.Println(lmq.AddPlatform(lmq.PlatformItem{ Platform:"ORP", Module:"datatrans" }))
	lmq.OutPutPlatformList()
	time.Sleep(1000000000)
	util.FileTest1()
}

func main() {
	db.InitDB();
	fmt.Println("server start...")
	server := api.NewServer()
	server.InitRouter()
	http.ListenAndServe(":8001", server.Router)
	fmt.Println("server start OK")
}
