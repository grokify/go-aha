package main

import (
	"flag"
	"fmt"

	"github.com/grokify/swaggman"
	"github.com/grokify/swaggman/postman2"
)

const (
	EnvAhaAccount = "AHA_ACCOUNT"
	EnvAhaApiKey  = "AHA_API_KEY"
)

func main() {
	var swaggerPath = flag.String("in", "path/to/swagger.json", "Path to swagger.json")
	var postmanPath = flag.String("out", "path/to/postman.json", "Path to postman.json")

	flag.Parse()

	fmt.Println("in has value ", *swaggerPath)

	// Instantiate a converter with default configuration
	conv := swaggman.NewConverter(swaggman.Configuration{
		PostmanURLHostname: fmt.Sprintf("{{%v}}.aha.io", EnvAhaAccount),
		PostmanHeaders: []postman2.Header{
			{
				Key:   "Authorization",
				Value: fmt.Sprintf("Bearer {{%v}}", EnvAhaApiKey),
			},
		},
	})

	// Convert a Swagger spec
	err := conv.Convert(*swaggerPath, *postmanPath)
	if err != nil {
		panic(err)
	}
}
