package clients

import (
  "log"
  //"strconv"
  "io/ioutil"
  "net/http"

  "github.com/ychen11/frontdesk/utils"
)


func HttpGet(url string) int {
  resp, err := http.Get(url)
  if err != nil {
    log.Printf("%s", err)
  }
  defer resp.Body.Close()

  return resp.StatusCode
}