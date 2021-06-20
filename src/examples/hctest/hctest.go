package main

import (
   "io/ioutil"
   "log"
   "net/http"
   "os"
   "fmt"
)

func main() {
   ticker := os.Args[1]
   url := "http://www.luckyticker.com/quote?q="+ticker
   fmt.Println(url)
   resp, err := http.Get("http://www.luckyticker.com/quote?q="+ticker)
   if err != nil {
      log.Fatalln(err)
   }
//We Read the response body on the line below.
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }
//Convert the body to type string
   sb := string(body)
   //log.Printf(sb)
   fmt.Println(sb)
}
