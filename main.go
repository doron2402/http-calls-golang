package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const numberOfSecondsDefaultValue = 1

func makeCalls(domain string, numberOfSeconds int, params []string) {

	for _, param := range params {
		tmpUrl := fmt.Sprintf("%s%s", domain, url.QueryEscape(param))
		fmt.Printf("[GET]: %s\n", tmpUrl)
		resp, err := http.Get(tmpUrl)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error making call")
		}

		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			fmt.Printf("Not Found or Bad call (%d) %s\n", resp.StatusCode, tmpUrl)
		} else if resp.StatusCode >= 500 {
			fmt.Printf("Server error (%d), %s\n", resp.StatusCode, tmpUrl)
		} else {
			fmt.Printf("Status Code: %d\n", resp.StatusCode)
		}

		sleepTime := time.Duration(numberOfSeconds) * time.Second
		time.Sleep(sleepTime)
	}
	os.Exit(1)
}

func main() {

	fmt.Println("---------------------")
	fmt.Println("---------------------")
	fmt.Println("Make Http/s calls")
	fmt.Println("Enter domain")
	fmt.Println("Enter Params")
	fmt.Println("When you done just click enter")
	fmt.Println("---------------------")
	// Get URL
	fmt.Print("Enter URL: ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	domain, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again", err)
		return
	}
	domain = strings.TrimSuffix(domain, "\n")
	u, err := url.ParseRequestURI(domain)
	if err != nil {
		fmt.Println("************ Error ************")
		fmt.Println("************ Error ************")
		fmt.Printf("%s is not a valid URL/Domain\n", domain)
		fmt.Println(err)
		fmt.Println("************ Error ************")
		os.Exit(1)
	}

	fmt.Printf("Domain: %s\n", u)

	// get interval
	fmt.Println("Enter interval (1 second by default, must be a number):")
	interval, err := reader.ReadString('\n')
	// convert CRLF to LF
	interval = strings.Replace(interval, "\n", "", -1)
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again", err)
		return
	}

	// remove the delimeter from the string
	interval = strings.TrimSuffix(interval, "\n")
	intervalInt, err := strconv.Atoi(interval)
	if err != nil {
		fmt.Println(err)
		intervalInt = numberOfSecondsDefaultValue
	}
	fmt.Printf("______________________\n")
	fmt.Printf("______________________\n")
	fmt.Printf("Domain: %s\n", domain)
	fmt.Printf("Sleeping between calls: %d seconds\n", intervalInt)
	fmt.Printf("______________________\n")
	params := []string{}

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		params = append(params, text)

		if ShouldMakeCalls(text) {
			makeCalls(domain, intervalInt, params[:len(params)-1])
		}
	}
}

func ShouldMakeCalls(word string) bool {
	switch word {
	case
		"stop",
		"done",
		"",
		"quit":
		return true
	}
	return false
}
