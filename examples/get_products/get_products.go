package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/antihax/optional"
	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"

	"github.com/grokify/go-aha/aha"
	au "github.com/grokify/go-aha/ahautil"
)

func getProducts(apis au.ClientAPIs) {
	api := apis.APIClient.ProductsApi
	ctx := context.Background()

	params := aha.GetProductsOpts{
		Page:    optional.NewInt32(int32(1)),
		PerPage: optional.NewInt32(int32(500))}

	info, resp, err := api.GetProducts(ctx, &params)

	if err != nil {
		log.Fatal("Error retrieving features")
	}

	fmt.Println(resp.StatusCode)
	fmtutil.PrintJSON(info)
	fmt.Printf("Found %v products\n", len(info.Products))
	fmt.Println("===")

	for _, prod := range info.Products {
		fmtutil.PrintJSON(prod)

		prod, resp, err := api.GetProduct(ctx, prod.Id)
		if err != nil {
			log.Fatal("Error retrieving product")
		} else if resp.StatusCode >= 300 {
			log.Fatal(fmt.Sprintf("Error calling API: %v", resp.StatusCode))
		}

		fmtutil.PrintJSON(prod)

		break
	}
}

func main() {
	err := config.LoadDotEnvSkipEmpty(os.Getenv("ENV_PATH"))
	if err != nil {
		panic(err)
	}

	apis := au.NewClientAPIs(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))

	getProducts(apis)

	fmt.Println("DONE")
}
