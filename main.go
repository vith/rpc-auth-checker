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

	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		panic("set LISTEN_ADDR environment variable to a host:port")
	}

	client, err := xmlrpc.NewClient(rpcEndpoint, nil)
	if err != nil {
		panic(err)
	}
	rpcClient = client

	http.HandleFunc("/login", loginHandler)

	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	user := req.PostFormValue("user")
	pass := req.PostFormValue("pass")

	rpcArgs := []interface{}{user, pass}

	var rpcResp map[string]string

	err := rpcClient.Call("checkAuthentication", rpcArgs, &rpcResp)
	if err != nil {
		log.Panicf("rpc transport level error occurred: %s\n", err)
	}

	if result, ok := rpcResp["result"]; ok {
		if result == "Success" {
			io.WriteString(w, strconv.FormatBool(true))
			return
		}
		log.Panicf("unhandled result type: %s\n", result)
	}

	if errStr, ok := rpcResp["error"]; ok {
		if errStr == "Invalid password" {
			io.WriteString(w, strconv.FormatBool(false))
			return
		}
		log.Panicf("unhandled error inside xml response: %s\n", errStr)
	}

	log.Panic("couldn't interpret response")
}
