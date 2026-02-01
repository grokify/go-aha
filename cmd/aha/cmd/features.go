package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/grokify/go-aha/v3/oag7/aha"
	"github.com/grokify/go-aha/v3/oag7/client"
)

var featuresCmd = &cobra.Command{
	Use:   "features",
	Short: "Manage features",
	Long:  `Commands for working with Aha! features.`,
}

var (
	featuresQuery        string
	featuresTag          string
	featuresAssignedTo   string
	featuresUpdatedSince string
	featuresPage         int
	featuresPerPage      int
)

var listFeaturesCmd = &cobra.Command{
	Use:   "list",
	Short: "List features",
	Long:  `List features with optional filtering.`,
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

		req := apiClient.FeaturesAPI.GetFeatures(ctx)

		if featuresQuery != "" {
			req = req.Q(featuresQuery)
		}
		if featuresTag != "" {
			req = req.Tag(featuresTag)
		}
		if featuresAssignedTo != "" {
			req = req.AssignedToUser(featuresAssignedTo)
		}
		if featuresUpdatedSince != "" {
			if t, err := time.Parse(time.RFC3339, featuresUpdatedSince); err == nil {
				req = req.UpdatedSince(t)
			}
		}
		if featuresPage > 0 {
			req = req.Page(int32(featuresPage))
		}
		if featuresPerPage > 0 {
			req = req.PerPage(int32(featuresPerPage))
		}

		features, resp, err := req.Execute()
		if err != nil {
			return fmt.Errorf("failed to list features: %w", err)
		}

		if resp.StatusCode >= 400 {
			return fmt.Errorf("API error: %s", resp.Status)
		}

		output, err := json.MarshalIndent(features, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		fmt.Println(string(output))
		return nil
	},
}

var getFeatureCmd = &cobra.Command{
	Use:   "get [feature-id]",
	Short: "Get a specific feature",
	Long:  `Get details for a specific feature by ID or reference number.`,
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

		feature, resp, err := apiClient.FeaturesAPI.GetFeature(ctx, args[0]).Execute()
		if err != nil {
			return fmt.Errorf("failed to get feature: %w", err)
		}

		if resp.StatusCode >= 400 {
			return fmt.Errorf("API error: %s", resp.Status)
		}

		output, err := json.MarshalIndent(feature, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		fmt.Println(string(output))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(featuresCmd)
	featuresCmd.AddCommand(listFeaturesCmd)
	featuresCmd.AddCommand(getFeatureCmd)

	// List flags
	listFeaturesCmd.Flags().StringVarP(&featuresQuery, "query", "q", "", "Search term to match against feature name or ID")
	listFeaturesCmd.Flags().StringVarP(&featuresTag, "tag", "t", "", "Filter by tag")
	listFeaturesCmd.Flags().StringVar(&featuresAssignedTo, "assigned-to", "", "Filter by assigned user ID or email")
	listFeaturesCmd.Flags().StringVar(&featuresUpdatedSince, "updated-since", "", "Only features updated after this timestamp (RFC3339)")
	listFeaturesCmd.Flags().IntVarP(&featuresPage, "page", "p", 0, "Page number")
	listFeaturesCmd.Flags().IntVarP(&featuresPerPage, "per-page", "n", 0, "Results per page")
}
