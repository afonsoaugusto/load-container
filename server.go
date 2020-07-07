package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/gorilla/mux"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/mem", printMemUsage)
	r.HandleFunc("/cpu", printCPUUsage)
	r.HandleFunc("/load", createLoad)

	fmt.Println("Listening on : 8080")
	http.ListenAndServe(":8080", r)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Docker test server running!")
	fmt.Fprintln(w, "GET params were:", r.URL.Query())
	// Loop over header names
	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Fprintln(w, name, value)
		}
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, b)
}

func printMemUsage(w http.ResponseWriter, r *http.Request) {

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Fprintln(w, "Alloc = MiB", bToMb(m.Alloc))
	fmt.Fprintln(w, "\tTotalAlloc =  MiB", bToMb(m.TotalAlloc))
	fmt.Fprintln(w, "\tSys =  MiB", bToMb(m.Sys))
	fmt.Fprintln(w, "\tNumGC = \n", m.NumGC)

	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	fmt.Printf("memory total: %d \n", bToMb(memory.Total))
	fmt.Printf("memory used: %d \n", bToMb(memory.Used))
	fmt.Printf("memory cached: %d \n", bToMb(memory.Cached))
	fmt.Printf("memory free: %d \n\n", bToMb(memory.Free))

	fmt.Printf("memory total: %d bytes\n", memory.Total)
	fmt.Printf("memory used: %d bytes\n", memory.Used)
	fmt.Printf("memory cached: %d bytes\n", memory.Cached)
	fmt.Printf("memory free: %d bytes\n\n", memory.Free)
}

func printCPUUsage(w http.ResponseWriter, r *http.Request) {

	before, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	fmt.Fprintln(w, "cpu user:\n", before.User)
	fmt.Fprintln(w, "cpu system:\n", before.System)
	fmt.Fprintln(w, "cpu idle:\n", before.Idle)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func createLoad(w http.ResponseWriter, r *http.Request) {

	command := "/usr/bin/loadserver.sh"
	args := " 1"

	doCmd := exec.Command(command, args)
	doCmd.Stdout = os.Stdout
	doCmd.Stderr = os.Stderr

	err := doCmd.Run()

	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		os.Exit(1)
	}
	fmt.Fprintln(w, err)
}
