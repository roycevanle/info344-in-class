package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/roycevanle/info344-in-class/zipsvr/handlers"

	"github.com/roycevanle/info344-in-class/zipsvr/models"
)

const zipsPath = "/zips/"

//Pointer, small 64 bit number that gets passed. Can get back to computer mem.
// This is essentially a pass by reference (not entire structure)
//Pass by value, makes complete complete copy. Private copy.
//Pass by reference and make changes, can be seen outside.
//Interfaces are always passed by reference so don't need star
func helloHandler(w http.ResponseWriter, r *http.Request) {
	runtime.GC()
	name := r.URL.Query().Get("name")
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Hello %s!", name)
	//w.Write([]byte("Hello, World!")) //Strings can be byte slice & byte slice to string
}

func memoryHandler(w http.ResponseWriter, r *http.Request) {
	//stack is segment on ram, heap is on machine
	stats := &runtime.MemStats{} //Struct on heap & pointer on heap
	runtime.ReadMemStats(stats)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func main() {
	//Set envi variable using export ADDR=localhost:4001 on term
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443" //Will accept req from any other computer
	}

	tlskey := os.Getenv("TLSKEY")
	tlscert := os.Getenv("TLSCERT")
	if len(tlskey) == 0 || len(tlscert) == 0 {
		log.Fatal("please set TLSKEY and TLSCERT")
	}

	zips, err := models.LoadZips("zips.csv")
	if err != nil {
		log.Fatalf("error load zips: %v", err)
	}
	log.Printf("loaded %d zips", len(zips))

	cityIndex := models.ZipIndex{}
	//index, item
	for _, z := range zips {
		cityLower := strings.ToLower(z.City)
		cityIndex[cityLower] = append(cityIndex[cityLower], z)
	}

	//fmt.Println("HelloWorld!")
	mux := http.NewServeMux()
	//We don't know what it'll be ahead of time (key). If it starts w/ pattern,
	//call the function rather than /hello cause it will match exactly
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/memory", memoryHandler)

	//When declare struct, need comma after everyone of lines
	cityHandler := &handlers.CityHandler{
		Index:      cityIndex,
		PathPrefix: zipsPath,
	}
	mux.Handle(zipsPath, cityHandler)

	fmt.Printf("server is listening at https://%s\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlscert, tlskey, mux))
}
