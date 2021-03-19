package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	res, erro := http.Get("http://www.google.com")

	if erro != nil {
		fmt.Println("Error:", erro)
		os.Exit(1)
	}

	fmt.Println(res)

}
