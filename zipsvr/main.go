package main

import "fmt"
import "net/http"
import "log"
import "runtime"
import "encoding/json"

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
    //fmt.Println("HelloWorld!")
    mux := http.NewServeMux()
    //We don't know what it'll be ahead of time (key). If it starts w/ pattern,
        //call the function rather than /hello cause it will match exactly
    mux.HandleFunc("/hello", helloHandler)
    mux.HandleFunc("/memory", memoryHandler)
    fmt.Printf("server is listening at http://localhost:4000\n")
    log.Fatal(http.ListenAndServe("localhost:4000", mux))
}