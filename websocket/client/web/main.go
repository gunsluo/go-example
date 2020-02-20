package main

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command
var argPath string

func init() {
	rootCmd = &cobra.Command{
		Use:   "web-serve",
		Short: "demo for web server",
		Long:  "demo for web server",
		Run:   run,
	}

	rootCmd.Flags().StringP("address", "a", ":8000", "address to listen on")
	rootCmd.Flags().StringVarP(&argPath, "file", "f", "publish.html", "the path of html file")
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, argPath)
}

func run(cmd *cobra.Command, args []string) {
	address, err := cmd.Flags().GetString("address")
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/", home)

	log.Println("start up web server, listen on", address)
	err = http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
