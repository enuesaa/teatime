// +build ignore

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
)

func Aaa() {
	fmt.Println("Building...")
}

func Bbb() {
    mg.Deps(Aaa)
    fmt.Println("Installing...")
}