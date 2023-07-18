package main

import (
	"log"

	dpll "github.com/vitus133/go-dpll/pkg/dpll-ynl"
)

func main() {

	conn, err := dpll.Dial(nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Printf("%+v\n", conn.GetGenetlinkFamily().Groups)
	reply, err := conn.DumpDeviceGet()
	log.Println(reply[0])

}
