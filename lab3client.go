package main 

import (
    "net/http"
    "log"
    "io/ioutil"
    "strconv"
    "fmt"
    "os"
    "strings"
)


func main(){


	//client code
	var port string
	var hash_value int
	if(len(os.Args)==1){// example $go run client.go
		log.Println("Not enough arguments")
		os.Exit(1)
	}

	if(len(os.Args)==2){// example $go run client.go GET
		url := fmt.Sprintf("http://localhost:3000/keys")
		get, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(get.Body)
		get.Body.Close()
		log.Println("Key Value pairs are  : ", string(data))
	}else if os.Args[1]== "GET" {// $go run client.go GET /keys/1
		request := os.Args[2]
		keyvalue := strings.Split(request,"/")
		key,_ := strconv.Atoi(keyvalue[2])
		hash_value = key % 3
		if(hash_value == 0){
			fmt.Println("port 3000")
			port = "3000"
		}else if(hash_value == 1){
			fmt.Println("port 3001")
			port = "3001"
		}else {
			fmt.Println("port 3002")
			port = "3002"
		}
		url := fmt.Sprintf("http://localhost:%s/keys/%s",port,keyvalue[2])
		get_result, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}	
		data, err := ioutil.ReadAll(get_result.Body)
		get_result.Body.Close()
		log.Println("Key Value pairs are  : ", string(data))
	}else{//$go run client.go PUT /keys/1/a
		request := os.Args[2]
		keyvalue := strings.Split(request,"/")
		key_int,_ := strconv.Atoi(keyvalue[2])
		hash_value = key_int % 3
		if(hash_value == 0){
			fmt.Println("port 3000")
			port = "3000"
		}else if(hash_value == 1){
			fmt.Println("port 3001")
			port = "3001"
		}else {
			fmt.Println("port 3002")
			port = "3002"
		}
		put_url:="http://localhost:"+port+"/keys/"+keyvalue[2]+"/"+keyvalue[3]
		
		client := &http.Client{}
		req, _ := http.NewRequest("PUT", put_url, nil)
		resp, _ := client.Do(req)
		resp.Body.Close()
		log.Println(" Response : ", 200)
	}
	
}