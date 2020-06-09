package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	t, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal("Fail get exact time ", err)
	}
	fmt.Println("current time:", time.Now().Round(time.Second))
	fmt.Println("exact time:", t.Round(time.Second))
}
