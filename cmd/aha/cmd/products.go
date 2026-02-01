package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/grokify/go-aha/v3/oag7/aha"
	"github.com/grokify/go-aha/v3/oag7/client"
)

var productsCmd = &cobra.Command{
	Use:   "products",
	Short: "Manage products",
	Long:  `Commands for working with Aha! products.`,
}

var (
	productsPage    int
	productsPerPage int
)

var listProductsCmd = &cobra.Command{
	Use:   "list",
	Short: "List products",
	Long:  `List all products.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := validateCredentials(); err != nil {
			return err
		}

		cfg, err := client.NewConfiguration(getDomain(), getAPIKey())
		if err != nil {
			return fmt.Errorf("failed to create configuration: %w", err)
		}

		apiClient := aha.NewAPIClient(cfg)
		ctx := context.Background()

		req := apiClient.ProductsAPI.GetProducts(ctx)

		if productsPage > 0 {
			req = req.Page(int32(productsPage))
		}
		if productsPerPage > 0 {
			req = req.PerPage(int32(productsPerPage))
		}

		products, resp, err := req.Execute()
		if err != nil {
			return fmt.Errorf("failed to list products: %w", err)
		}

		if resp.StatusCode >= 400 {
			return fmt.Errorf("API error: %s", resp.Status)
		}

		output, err := json.MarshalIndent(products, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		fmt.Println(string(output))
		return nil
	},
}

var getProductCmd = &cobra.Command{
	Use:   "get [product-id]",
	Short: "Get a specific product",
	Long:  `Get details for a specific product by ID or key.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := validateCredentials(); err != nil {
			return err
		}

		cfg, err := client.NewConfiguration(getDomain(), getAPIKey())
		if err != nil {
			return fmt.Errorf("failed to create configuration: %w", err)
		}

		apiClient := aha.NewAPIClient(cfg)
		ctx := context.Background()

		product, resp, err := apiClient.ProductsAPI.GetProduct(ctx, args[0]).Execute()
		if err != nil {
			return fmt.Errorf("failed to get product: %w", err)
		}

		if resp.StatusCode >= 400 {
			return fmt.Errorf("API error: %s", resp.Status)
		}

		output, err := json.MarshalIndent(product, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		fmt.Println(string(output))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(productsCmd)
	productsCmd.AddCommand(listProductsCmd)
	productsCmd.AddCommand(getProductCmd)

	// List flags
	listProductsCmd.Flags().IntVarP(&productsPage, "page", "p", 0, "Page number")
	listProductsCmd.Flags().IntVarP(&productsPerPage, "per-page", "n", 0, "Results per page")
}
