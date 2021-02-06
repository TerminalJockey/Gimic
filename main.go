package main

//C was imported and functions exported for easy embedding into c#

import (
	"C"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"github.com/traefik/yaegi/stdlib/unrestricted"
)

func main() {
	fmt.Println("pushing go")
	RunServer()
}

//EvalInput :
//export EvalInput
func EvalInput(incode string) {
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	i.Use(unrestricted.Symbols)
	test, err := i.Eval(incode)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(test)
}

//RunServer :
//export RunServer
func RunServer() {
	ln, err := net.Listen("tcp", ":31337")
	if err != nil {
		log.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}
		inbytes, err := ioutil.ReadAll(conn)
		EvalInput(string(inbytes))
	}
}
