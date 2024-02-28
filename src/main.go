package main

import "com.tuntun.rangers/service/chaindata/src/cli"

func main() {
	client := cli.NewGX()
	client.Run()
}
