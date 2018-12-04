package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	srv := &http.Server{Addr: ":8080"}

	http.HandleFunc("/", viewHandler)
	log.Fatal(srv.ListenAndServe())

	defer srv.Shutdown(nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	randComic := "https://xkcd.com/" + fmt.Sprintf("%d", rnd.Intn(2080)) + "/"
	fmt.Println(randComic)
	resp, err := http.Get(randComic)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Fprintf(w, string(body))
}
