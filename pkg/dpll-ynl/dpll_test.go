package dpll

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func Test_PinGet(t *testing.T) {
	conn, err := Dial(nil)
	if err != nil {
		log.Fatal(err)
	}
	pinReply, err := conn.DoPinGet(DoPinGetRequest{Id: 0})
	if err != nil {
		log.Fatal(err)
	}
	timestamp := time.Now().UTC()
	pinInfo, err := GetPinInfoHR(pinReply, timestamp)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(pinInfo))
}
