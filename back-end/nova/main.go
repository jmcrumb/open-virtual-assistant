package main

import (
	"fmt"

	data "github.com/jmcrumb/nova/data"
	nlp "github.com/jmcrumb/nova/nlp"
)

func main() {
	fmt.Println("Hello NOVA backend")

	nlp.Test()
	data.Test()
}
