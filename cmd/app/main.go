package main

import (
	"fmt"
	"log"

	"github.com/svkorch/kvss/internal/kvss"
)

func main() {
	log.Println("KVSS App started")

	s := kvss.New()

	fmt.Println(s.Get("aa"))

	s.Add("aa", "AAAAA")
	fmt.Println(s.Get("aa"))

	s.Add("bbb", "BBBB")
	fmt.Println(s.Get("aa"))
	fmt.Println(s.Get("bbb"))
}
