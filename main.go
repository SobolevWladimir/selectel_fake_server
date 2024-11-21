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
	Repository repository.RepositoryInterface
}

const settinsRmsPage = "/settings/rms.json"
const settinsRoutePage = "/settings/route.json"

func (ct *CounterHandler) HandlerSaveFile(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, err.Error())
		return
	}

	err = ct.Repository.SaveFile(requestBody, r.URL.Path)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, err.Error())
		return
	}
	w.WriteHeader(201)
	fmt.Fprintln(w, string("ok"))
}

func (ct *CounterHandler) HandlerGetFile(w http.ResponseWriter, r *http.Request) {
	bytes, err := ct.Repository.GetFile(r.URL.Path)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, err.Error())
		return
	}
	w.WriteHeader(201)
	w.Write(bytes)
}

func (ct *CounterHandler) HandlerDeleteFile(w http.ResponseWriter, r *http.Request) {
	err := ct.Repository.DeleteFile(r.URL.Path)
	if err != nil {
		w.WriteHeader(400)
		fmt.Println("delete", err.Error())
		fmt.Fprintln(w, err.Error())
		return
	}
	w.WriteHeader(204)
	fmt.Fprintln(w, string("ok"))
}

func (ct *CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\n----------------")
	fmt.Println(r.Method, r.URL.String())
	if r.Method == "PUT" {
		ct.HandlerSaveFile(w, r)
		return
	}
	if r.Method == "GET" {
		ct.HandlerGetFile(w, r)
		return
	}

	if r.Method == "DELETE" {
		ct.HandlerDeleteFile(w, r)
		return
	}
	w.WriteHeader(400)
	fmt.Fprintln(w, "method not exist")
	fmt.Println("---------------- end ---------------------")
}
func (ct *CounterHandler) setHeader(w http.ResponseWriter, resp *http.Response) {
	for key, value := range resp.Header {
		w.Header().Add(key, value[0])
	}
}
func main() {
	repository := &repository.RepositoryLocal{RootDir: "./files"}
	th := &CounterHandler{Repository: repository}
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
