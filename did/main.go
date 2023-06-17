package main

import (
	"fmt"
	"log"

	"github.com/ockam-network/did"
)

func main() {
	// d, err := did.Parse("did:example:q7ckgxeq1lxmra0r")
	d, err := did.Parse("did:example:q7ckgxeq1lxmra0r/abc/pqr")
	
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", d)
} 