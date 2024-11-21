package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"selectel_fake_server/src/repository"
	"time"
)

type CounterHandler struct {
}

const settinsRmsPage = "/settings/rms.json"
const settinsRoutePage = "/settings/route.json"

func (ct *CounterHandler) saveFile(r *http.Request) error {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	repository := &repository.RepositoryLocal{RootDir: "./files"}
	repository.SaveFile(requestBody, r.URL.RawPath)
	return nil
}
func (ct *CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\n----------------")
	fmt.Println(r.Method, r.URL.String())
	if r.Method == "PUT" {
		err := ct.saveFile(r)
		if err == nil {
			w.WriteHeader(201)
			fmt.Fprintln(w, string("ok"))
			return
		}
		fmt.Println("error: ", err.Error())
		w.WriteHeader(400)
		fmt.Fprintln(w, err.Error())
	} else {
		w.WriteHeader(200)

	}
	fmt.Println("---------------- end ---------------------")
}
func (ct *CounterHandler) setHeader(w http.ResponseWriter, resp *http.Response) {
	for key, value := range resp.Header {
		w.Header().Add(key, value[0])
	}
}
func main() {
	th := &CounterHandler{}
	s := &http.Server{
		Addr:           ":8051",
		Handler:        th,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
	fmt.Println("server start")
}
