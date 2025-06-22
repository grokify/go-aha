package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/antihax/optional"
	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/fmt/fmtutil"

	"github.com/grokify/go-aha/v3/oag4/aha"
	au "github.com/grokify/go-aha/v3/oag4/ahautil"
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
	fmtutil.MustPrintJSON(info)
	fmt.Printf("Found %v products\n", len(info.Products))
	fmt.Println("===")

	for _, prod := range info.Products {
		fmtutil.MustPrintJSON(prod)

		prod, resp, err := api.GetProduct(ctx, prod.Id)
		if err != nil {
			log.Fatal("Error retrieving product")
		} else if resp.StatusCode >= 300 {
			log.Fatal(fmt.Sprintf("Error calling API: %v", resp.StatusCode))
		}

		fmtutil.MustPrintJSON(prod)

		break
	}
}

func main() {
	_, err := config.LoadDotEnv([]string{os.Getenv("ENV_PATH")}, 1)
	if err != nil {
		log.Fatal(err)
	}

	apis, err := au.NewClientAPIs(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	getProducts(apis)

	fmt.Println("DONE")
}
