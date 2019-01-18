package main

import (
	"fmt"
	"log"
	"os"
	"errors"
)

var destAdr string
var destPtl string



func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

func getArgs() {
	//	argsWithProg := os.Args
	//	argsWithoutProg := os.Args[1:]
	destUrl := os.Args(1)
}

func findProxy( proto string) {
    switch proto {
    case http: return os.Getenv("HTTP_PROXY")
    case https: return os.Getenv("HTTPS_PROXY")
    case ftp: return  os.Getenv("FTP_PROXY")
    case :  panic ("protocol not specified")
    }
}

func main() {
	findProxy()
	fmt.Printf("Check network settings\n")
	fmt.Println("http proxy: ",  findProxy( http))
}
