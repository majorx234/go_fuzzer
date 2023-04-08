package main

	

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"os"
	"time"
)

func request(id int, url string) {
    resp, err := http.Get(url)
    if err != nil {
         fmt.Println("Some error: ", err)
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Some error: ", err)
    }
    status := resp.StatusCode
    if err != nil {
        fmt.Println("Some error: ", err)
    }    
    fmt.Printf("[%d] req to url %s (%d) status: %d \n", id, url, len(body), status)
}

func readFile(){
	dat, err := os.ReadFile("./wordlist.txt")
    if err != nil {

    }
    fmt.Println(string(dat))
    words := strings.Split(string(dat), "\n")

    for i, word := range words {
    	fmt.Printf("http request nr &d to word %s ", i, word)
    	url := fmt.Sprintf("https://emile.space/%s", word)
    	resp, err := http.Get(url)
        if err != nil {

        }
        body, err := io.ReadAll(resp.Body)
        if err != nil {
 
        }
        fmt.Println(string(body))
    }
}

//func worker(done chan bool) {
//    fmt.Print("working...")
//    time.Sleep(time.Second)
//    fmt.Println("done")
//    done <- true
//}
func worker(id int, wordChan chan string, doneChan chan bool) {
    out:
    for {
        select {
            case url := <-wordChan:
            url = fmt.Sprintf("https://emile.space/%s", url)
            request(id, url)
            case <-time.After(3 * time.Second):
            fmt.Printf("worker %d couldn't get a new url after 3 seconds,quitting\n", id)
            break out
        }
    }
    doneChan <- true
}

func main() {
	// read wordlist
	dat, err := os.ReadFile("./wordlist1.txt")
	if err != nil {
	    fmt.Println("Some error: ", err)
	}
	words := strings.Split(string(dat), "\n")
	// create workers
	wordChan := make(chan string, 10)
	doneChan := make(chan bool, 4)
	for i := 1; i < 4; i++ {
	    go worker(i, wordChan, doneChan)
	}
	// fill word channel with all the words we want to fuzz
	fmt.Println("Filling wordChan")
	for _, word := range words {
	    wordChan <- word
	}
	// check that all the workers are done before ending
	for i := 1; i < 4; i++ {
	    <-doneChan
	}
}


//func main() {
//    fmt.Println("hello world")
    //readFile()
    //resp, err := http.Get("http://example.com")
    //defer resp.Body.Close()
    //if err != nil {

    //}

//    body, err := io.ReadAll(resp.Body)
//    if err != nil {
//
//    }
//    fmt.Println(string(body))
//}
