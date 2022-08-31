package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)


func makeRequest(request string, url string) bool {
	timeout := time.Duration(15 * time.Second)
	client := http.Client{Timeout: timeout}
	response, err := client.Get(request)
	if err != nil {
		log.Print("GET error:", err)
		file, err := os.OpenFile("list_url_timeout_error.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			log.Println(err)
			}
		file.WriteString(url + "\n")
		return false
	}
	
	defer response.Body.Close()
	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)
	file, err := os.OpenFile("cards", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
        log.Println(err)
		}	
	defer file.Close()
	message, err := json.Marshal(&result)
	if err != nil {
		log.Fatal("JSON marshaling failed:", err)
	}
	file.WriteString(string(message) + "\n")
	return true
}


func thread(domain string, urls []string) {
	for i := 0; i < 400; i++ {
		request := domain + urls[i]
		if !makeRequest(request, urls[i]) {
			continue
		}
	}
	log.Println("Goroutine finish")
}


func main_thread(domain string, urls []string) {
	for i := 0; i < 400; i++ {
		time.Sleep(500 * time.Millisecond)
		request := domain + urls[i]
		if !makeRequest(request, urls[i]) {
			continue
		}	
	}
	log.Println("Finish main func")
}

func main() {
	file, err := os.Open("list_url_timeout_error.txt")
	if err != nil{
		log.Fatalln(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var stack []string
	for scanner.Scan() {
		stack = append(stack, scanner.Text())
	}
	workers := [40]string{
		//data
		"http://51.250.106.196:8080/?url=",
		"http://51.250.102.139:8080/?url=",
		"http://51.250.96.188:8080/?url=",
		"http://51.250.23.221:8080/?url=",
		"http://51.250.28.129:8080/?url=",
		"http://51.250.102.119:8080/?url=",
		"http://51.250.97.227:8080/?url=",
		"http://51.250.28.112:8080/?url=",
		//data-0
		"http://51.250.99.229:8080/?url=",
		"http://51.250.100.176:8080/?url=",
		"http://51.250.105.202:8080/?url=",
		"http://51.250.107.208:8080/?url=",
		"http://51.250.27.86:8080/?url=",
		"http://51.250.31.146:8080/?url=",
		"http://51.250.107.225:8080/?url=",
		"http://51.250.104.145:8080/?url=",
		//data-1
		"http://51.250.110.230:8080/?url=",
		"http://84.201.153.239:8080/?url=",
		"http://84.201.139.203:8080/?url=",
		"http://62.84.121.62:8080/?url=",
		"http://51.250.109.112:8080/?url=",
		"http://84.201.163.246:8080/?url=",
		"http://84.201.137.124:8080/?url=",
		"http://84.201.138.142:8080/?url=",
		//data-2
		"http://84.252.137.115:8080/?url=",
		"http://51.250.31.238:8080/?url=",
		"http://84.201.161.49:8080/?url=",
		"http://62.84.122.220:8080/?url=",
		"http://84.201.166.40:8080/?url=",
		"http://62.84.122.201:8080/?url=",
		"http://84.201.177.247:8080/?url=",
		"http://84.201.154.186:8080/?url=",
		//data-3
		"http://84.201.179.58:8080/?url=",
		"http://84.252.142.101:8080/?url=",
		"http://84.252.138.200:8080/?url=",
		"http://84.252.136.25:8080/?url=",
		"http://51.250.108.130:8080/?url=",
		"http://84.252.136.150:8080/?url=",
		"http://84.252.142.173:8080/?url=",
		"http://51.250.109.102:8080/?url=",
		}
	
	start_index := 0
	index := 400
	for i := 0; i < 39; i++ {
		time.Sleep(20 * time.Millisecond)
		go thread(workers[i], stack[start_index: index])

		log.Println(i, " ", start_index, " ", index)
		log.Println()
		start_index = index
		index += 400
	}

	main_thread(workers[39], stack[0:1000])
}