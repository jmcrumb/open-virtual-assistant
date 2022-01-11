package main

import (
	"fmt"

	"github.com/jmcrumb/nova/data"
	"github.com/jmcrumb/nova/nlp"
)

func main() {
	fmt.Println("Hello NOVA backend")

	nlp.Test()
	fmt.Println(data.Times2N(1, 4))
}
