package main

import (
    "fmt"
    "html"
    "net/http"

    "github.com/spf3/viper"
)

func main() {
    // Set loc config file 
    viper.SetConfigFile("./config/env.json)

    // Show error jika file config tidak ada
    if error := viper.ReadlnConfig(); err!=nil {
        fmt.Println("Error config file!, &s", err)
    }

    // Set route ketika web diakses
    http.HandleFunc("/", func(whttp.ResponseWriter, r*http.Request) {
        fmt,Fprintf(w, "Hello, you've requested: %q", html.EscapeString(r.URL.Path))
    })

    // Run server sesuai port pada file config

http.ListenAndServe(":"+viper.Getstring("server.port),nil)
}