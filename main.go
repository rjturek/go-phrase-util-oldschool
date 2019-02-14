package main

import (
	"fmt"
	"github.com/rjturek/go-phrase-util-oldschool/rjtphraseoldschool"
)

var phrase1 = "But when night comes round, oh gosh, oh gee"

func GetPhrase() string {
	return phrase1
}

func main() {
	fmt.Println(GetPhrase())
	fmt.Println(rjtphraseoldschool.GetPhrase())
}
