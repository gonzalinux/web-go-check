package main

import (
	"fmt"
	"net/http"
	"time"
)

const destEmail = "yourEmailHere@x.x"
const interval = time.Minute

func checkUrl(url string) (bool, error) {

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return false, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("received status code not equal to 200: %d", resp.StatusCode) // Not successful, but not necessarily an error (informational)
	}
	return true, nil
}

func main() {
	fmt.Println("Starting web-go-check.")
	urls := []string{
		"https://url1",
		"https://url2:2999/route",
		"https://url3.com/",
	}
	for {
		fmt.Printf("%s Checking all the urls\n", time.Now().Format("2006-01-02 15:04:05"))

		for _, url := range urls {
			success, err := checkUrl(url)
			if !success {

				errorString := fmt.Sprintf("☒ %s\n", url)
				if err != nil {
					errorString += fmt.Sprintf("the exception was:\n %s \n\n", err)
				}
				fmt.Println("\033[31m", errorString, "\033[0m")
				err = SendMail(FormatErrorEmail(url, errorString), destEmail, "Error in "+url)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Printf(" \033[32m☑ %s \n\033[0m", url)
			}
		}

		time.Sleep(interval)
	}

}
