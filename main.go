package main

import (
	"fmt"
	"github.com/rjturek/go-phrase-util-oldschool/rjtphraseoldschool"
)

func GetPhrase() string {
	return phrase1
}

func main() {
	fmt.Println(GetPhrase())
	fmt.Println(rjtphraseoldschool.GetPhrase())
}
