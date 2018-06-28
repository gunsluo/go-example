package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

const (
	clientID     = "client id"
	clientSecret = "client secret"
)

func main() {
	if err := cmd().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}
}

func cmd() *cobra.Command {
	c := cobra.Command{
		Use:   "example-app",
		Short: "An example OpenID Connect client",
		Long:  "",
		Run:   run,
	}

	return &c
}

func run(cmd *cobra.Command, args []string) {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/sso/callback", handleCallback)

	fmt.Println("Run http://localhost:5556")
	http.ListenAndServe(":5556", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		authURL := fmt.Sprintf("https://www.linkedin.com/oauth/v2/authorization?response_type=code&client_id=%s&redirect_uri=http://sso-dex:5556/sso/callback&state=I wish to wash my irish wristwatch", clientID)
		http.Redirect(w, r, authURL, http.StatusSeeOther)
	}
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	if code == "" {
		w.Write([]byte("code is nil"))
		return
	}
	codeURL := fmt.Sprintf("https://www.linkedin.com/oauth/v2/accessToken?grant_type=authorization_code&code=%s&redirect_uri=http://sso-dex:5556/sso/callback&client_id=%s&client_secret=%s", code, clientID, clientSecret)
	resp, err := http.Post(codeURL, "application/x-www-form-urlencoded", nil)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println(string(body))
}
