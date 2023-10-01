package main

import (
	"bufio"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

// This function will make request to the target URL with provided options
func makeRequest(preparedURL string, word string, showSuccessful bool) {
	// Scanning for GET Method
	req, errReq := http.NewRequest(http.MethodGet, preparedURL, nil)
	if errReq != nil {
		PrintLog("request error", "There was an error while making request to this URL: "+preparedURL+"\nERROR: "+errReq.Error())
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		PrintLog("respose error", "There was an error encounted while making GET request to this URL: "+preparedURL+"\nERROR: "+err.Error())
	} else {
		if resp.StatusCode == 400 && showSuccessful {
			return
		}
		PrintLog(strconv.Itoa(resp.StatusCode), "[Method: GET] "+preparedURL)
	}

	// TODO: Scanning for POST Method
	// resp, err = http.Post(requestURL, "json")
	// if err != nil {
	// 	PrintLog("response error", )
	// }
}

func Scanner(targetURL string, showSuccessful bool, wordlistFile *os.File, secure bool) {
	_, err := url.Parse(targetURL)
	if err != nil {
		PrintLog("error", "The provided target URL is not valid!")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(wordlistFile)
	for scanner.Scan() {
		line := scanner.Text()

		wg.Add(1)

		throwProtocol := func() string {
			if secure {
				return "https://"
			} else {
				return "http://"
			}
		}()
		requestURL := throwProtocol + strings.Replace(targetURL, "FUZZTHIS", line, -1)

		go func(url string, word string, showSuccessful bool) {
			defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
			makeRequest(url, word, showSuccessful)
		}(requestURL, line, showSuccessful)
	}

	wg.Wait()
}
