package main

import "demo/core"

func main() {
	bc := core.NewBlockChain()
	bc.SendData("Send 1 BTC to Jacky")
	bc.SendData("Send 1 EOS to Jack")
	bc.Print()

}
