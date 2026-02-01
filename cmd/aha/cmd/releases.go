package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/grokify/go-aha/v3/oag7/aha"
	"github.com/grokify/go-aha/v3/oag7/client"
)

var releasesCmd = &cobra.Command{
	Use:   "releases",
	Short: "Manage releases",
	Long:  `Commands for working with Aha! releases.`,
}

var (
	releasesProductID string
	releasesPage      int
	releasesPerPage   int
)

var listReleasesCmd = &cobra.Command{
	Use:   "list",
	Short: "List releases for a product",
	Long:  `List releases for a specific product.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := validateCredentials(); err != nil {
			return err
		}

		if releasesProductID == "" {
			return fmt.Errorf("product ID is required: use --product flag")
		}

		cfg, err := client.NewConfiguration(getDomain(), getAPIKey())
		if err != nil {
			return fmt.Errorf("failed to create configuration: %w", err)
		}

		apiClient := aha.NewAPIClient(cfg)
		ctx := context.Background()

		req := apiClient.ReleasesAPI.GetProductReleases(ctx, releasesProductID)

		if releasesPage > 0 {
			req = req.Page(int32(releasesPage))
		}
		if releasesPerPage > 0 {
			req = req.PerPage(int32(releasesPerPage))
		}

		releases, resp, err := req.Execute()
		if err != nil {
			return fmt.Errorf("failed to list releases: %w", err)
		}

		if resp.StatusCode >= 400 {
			return fmt.Errorf("API error: %s", resp.Status)
		}

		output, err := json.MarshalIndent(releases, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		fmt.Println(string(output))
		return nil
	},
}

var getReleaseCmd = &cobra.Command{
	Use:   "get [release-id]",
	Short: "Get a specific release",
	Long:  `Get details for a specific release by ID or key.`,
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

		release, resp, err := apiClient.ReleasesAPI.GetRelease(ctx, args[0]).Execute()
		if err != nil {
			return fmt.Errorf("failed to get release: %w", err)
		}

		if resp.StatusCode >= 400 {
			return fmt.Errorf("API error: %s", resp.Status)
		}

		output, err := json.MarshalIndent(release, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		fmt.Println(string(output))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(releasesCmd)
	releasesCmd.AddCommand(listReleasesCmd)
	releasesCmd.AddCommand(getReleaseCmd)

	// List flags
	listReleasesCmd.Flags().StringVar(&releasesProductID, "product", "", "Product ID (required)")
	listReleasesCmd.Flags().IntVarP(&releasesPage, "page", "p", 0, "Page number")
	listReleasesCmd.Flags().IntVarP(&releasesPerPage, "per-page", "n", 0, "Results per page")
}
