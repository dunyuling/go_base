package main

import (
	"net/http"
	"demo/core"
	"encoding/json"
	"io"
	"fmt"
)

var blockChain *core.BlockChain

func blockChainGetHandler(w http.ResponseWriter, req *http.Request) {
	bytes,err := 	json.Marshal(blockChain)
	fmt.Println(bytes)
	fmt.Println(len(bytes))
	if err!= nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	io.WriteString(w,string(bytes))
}

func blockChainWriteHandler(w http.ResponseWriter, req *http.Request) {
	blockData := req.URL.Query().Get("data")
	blockChain.SendData(blockData)
	blockChainGetHandler(w,req)
}

func run()  {
	http.HandleFunc("/blockchain/get", blockChainGetHandler)
	http.HandleFunc("/blockchain/write", blockChainWriteHandler)
	http.ListenAndServe(":8000",nil)
}

func main()  {
	blockChain = core.NewBlockChain()
	run()
}

