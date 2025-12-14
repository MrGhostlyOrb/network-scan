package main

import (
	"fmt"
	"html/template"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("[ERROR] Unable to load environment variables: %s\n", err.Error())
	}
}

func main() {
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	if len(os.Args) == 1 {
		fmt.Println("Running in Web mode...")
		http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			template, err := template.ParseFiles("index.html")
			if err != nil {
				fmt.Println(err)
			}
			err = template.Execute(w, nil)
			if err != nil {
				fmt.Println(err)
			}
		})
		http.HandleFunc("/scan", func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()

			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err)
			}
			values, err := url.ParseQuery(string(body))
			if err != nil {
				fmt.Println(err)
			}
			port := values.Get("port")
			subnet := values.Get("subnet")
			portNumber, err := strconv.Atoi(port)
			if err != nil {
				fmt.Println(err)
			}
			subnetNumber, err := strconv.Atoi(subnet)
			if err != nil {
				fmt.Println(err)
			}
			validIps := ScanNetwork(portNumber, subnetNumber)
			fmt.Println("Valid addresses:")
			for _, ip := range validIps {
				fmt.Printf("http://%s/\n", ip)
			}
			template, err := template.ParseFiles("index.html")
			if err != nil {
				fmt.Println(err)
			}
			err = template.Execute(w, validIps)
			if err != nil {
				fmt.Println(err)
			}
		})
		fmt.Printf("Started on http://%s:%s/\n", host, port)
		err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
		if err != nil {
			fmt.Println(err)
		}
	}
	if len(os.Args) == 3 {
		fmt.Println("Running in CLI mode...")
		port := os.Args[1]
		subnet := os.Args[2]
		portNumber, err := strconv.Atoi(port)
		if err != nil {
			fmt.Println(err)
		}
		subnetNumber, err := strconv.Atoi(subnet)
		if err != nil {
			fmt.Println(err)
		}
		validIps := ScanNetwork(portNumber, subnetNumber)
		fmt.Println("Valid addresses:")
		for _, ip := range validIps {
			fmt.Printf("http://%s/\n", ip)
		}
	}

	if len(os.Args) != 1 && len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <port> <subnet>")
		fmt.Println("Example: go run main.go 8080 50")
		os.Exit(0)
	}
}

func ScanNetwork(port int, subnet int) []string {
	validIps := []string{}
	var waitGroup sync.WaitGroup
	var mutex sync.Mutex
	var timeout int64

	before := time.Now().UnixMilli()
	connection, err := net.DialTimeout("tcp", "google.com:80", 2*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	defer connection.Close()
	after := time.Now().UnixMilli()
	timeout = (after - before) * 3

	for i := 1; i < 255; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			connection, err := net.DialTimeout("tcp", fmt.Sprintf("192.168.%d.%d:%d", subnet, i, port), time.Duration(timeout)*time.Millisecond)
			if err != nil {

			} else {
				defer connection.Close()
				mutex.Lock()
				validIps = append(validIps, fmt.Sprintf("192.168.%d.%d:%d", subnet, i, port))
				mutex.Unlock()
			}
		}(i)
	}
	waitGroup.Wait()
	return validIps
}
