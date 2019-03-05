package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/kolo/xmlrpc"
)

var rpcClient *xmlrpc.Client

func main() {
	rpcEndpoint := os.Getenv("RPC_ENDPOINT")
	if rpcEndpoint == "" {
		panic("set RPC_ENDPOINT environment variable to a URL")
	}

	client, err := xmlrpc.NewClient(rpcEndpoint, nil)
	if err != nil {
		panic(err)
	}
	rpcClient = client

	http.HandleFunc("/login", loginHandler)

	log.Fatal(http.ListenAndServe("127.0.0.1:30303", nil))
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	user := req.PostFormValue("user")
	pass := req.PostFormValue("pass")

	rpcArgs := []interface{}{user, pass}
	var result bool

	err := rpcClient.Call("checkAuthentication", rpcArgs, &result)
	if err != nil {
		panic(err)
	}

	io.WriteString(w, strconv.FormatBool(result))
}
