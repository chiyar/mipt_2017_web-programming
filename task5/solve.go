package main

import "github.com/gorilla/mux"
import "net/http"
import "io/ioutil"
import "encoding/json"
import "strconv"

type Request struct {
  Url string `json:"url"`
}

type Response struct {
  Key string `json:"key"`
}

var shortUrls = make([]string, 0)

func getKey() int {
  return len(shortUrls)
}

func AddUrl(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var url Request
  body, _ := ioutil.ReadAll(r.Body)
  json.Unmarshal(body, &url)
  key := getKey()
  shortUrls = append(shortUrls, url.Url)
  var response Response
  response.Key = strconv.Itoa(key)
  ans, _ := json.Marshal(response)
  w.WriteHeader(http.StatusOK)
  w.Write(ans)
}

func GetUrl(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  key, _ := strconv.Atoi(vars["key"])
  w.Header().Set("Location", shortUrls[key])
  w.WriteHeader(http.StatusMovedPermanently)
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", AddUrl).Methods("POST")
  r.HandleFunc("/{key}", GetUrl).Methods("GET")
  http.ListenAndServe(":8082", r)
}
