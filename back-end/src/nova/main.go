package main

import (
	"fmt"

	"github.com/jmcrumb/nova/data"
	"github.com/jmcrumb/nova/nlp"
)

func main() {
	fmt.Println("Hello NOVA backend")

	nlp.Test()
	data.Test()
}
