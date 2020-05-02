package fmp

import (
    "log"
    "io/ioutil"
    "net/http"
)

func GetRequest(url string) string {
    response, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    result := string(body)
    return result
}

