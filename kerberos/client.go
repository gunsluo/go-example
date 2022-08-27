package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/jcmturner/gokrb5/v8/client"
	"github.com/jcmturner/gokrb5/v8/config"
	"github.com/jcmturner/gokrb5/v8/credentials"
	"github.com/jcmturner/gokrb5/v8/keytab"
	"github.com/jcmturner/gokrb5/v8/spnego"
)

const (
	realm = "TEST.GOKRB5"
	spn   = "sso.test.gokrb5"
	port  = ":80"
)

func main() {
	krb5ConfPath := "./docker/keytabs/krb5.conf"
	b1, err := ioutil.ReadFile(krb5ConfPath)
	if err != nil {
		panic(err)
	}

	conf, err := config.NewFromString(string(b1))
	if err != nil {
		panic(err)
	}

	// testKeytab(conf)
	testCache(conf)
}

func testKeytab(conf *config.Config) {
	// Create the client with the keytab
	cl, err := makeClientWithKeytab(conf)
	if err != nil {
		panic(err)
	}

	//Log in the client
	err = cl.Login()
	if err != nil {
		panic(err)
	}

	// Form the request
	url := "http://127.0.0.1" + port
	// url := "http://sso.test.gokrb5" + port
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	spnegoCl := spnego.NewClient(cl, nil, "HTTP/"+spn)

	// Make the request
	resp, err := spnegoCl.Do(r)
	if err != nil {
		panic(err)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func testCache(conf *config.Config) {
	cl, err := makeClientWithCache(conf)
	if err != nil {
		panic(err)
	}

	s := spnego.SPNEGOClient(cl, "HTTP/"+spn)
	if err := s.AcquireCred(); err != nil {
		panic(err)
	}

	st, err := s.InitSecContext()
	if err != nil {
		panic(err)
	}
	nb, err := st.Marshal()
	if err != nil {
		panic(err)
	}

	// Form the request
	url := "http://127.0.0.1" + port
	// url := "http://sso.test.gokrb5" + port
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	hs := "Negotiate " + base64.StdEncoding.EncodeToString(nb)
	r.Header.Set("Authorization", hs)

	c := &http.Client{}
	resp, err := c.Do(r)

	if err != nil {
		panic(err)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func makeClientWithKeytab(conf *config.Config) (*client.Client, error) {
	filename := "./docker/keytabs/luoji.keytab"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	kt := keytab.New()
	if err := kt.Unmarshal(b); err != nil {
		return nil, err
	}

	cl := client.NewWithKeytab("luoji", realm, kt, conf, client.DisablePAFXFAST(true))

	return cl, nil
}

func makeClientWithCache(conf *config.Config) (*client.Client, error) {
	ccpath := "./docker/keytabs/krb5cc_0"
	ccache, err := credentials.LoadCCache(ccpath)
	if err != nil {
		return nil, err
	}

	cl, err := client.NewFromCCache(ccache, conf, client.DisablePAFXFAST(true))
	return cl, err
}
