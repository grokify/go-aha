package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/grokify/go-aha/v3/oag7/aha"
	"github.com/grokify/go-aha/v3/oag7/client"
)

var ideasCmd = &cobra.Command{
	Use:   "ideas",
	Short: "Manage ideas",
	Long:  `Commands for working with Aha! ideas.`,
}

var (
	// list flags
	listQuery          string
	listWorkflowStatus string
	listSort           string
	listTag            string
	listUserID         string
	listCreatedSince   string
	listUpdatedSince   string
	listPage           int
	listPerPage        int
	listSpam           bool
)

var listIdeasCmd = &cobra.Command{
	Use:   "list",
	Short: "List ideas",
	Long:  `List ideas with optional filtering.`,
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

		req := apiClient.IdeasAPI.ListIdeas(ctx)

		if listQuery != "" {
			req = req.Q(listQuery)
		}
		if listWorkflowStatus != "" {
			req = req.WorkflowStatus(listWorkflowStatus)
		}
		if listSort != "" {
			req = req.Sort(listSort)
		}
		if listTag != "" {
			req = req.Tag(listTag)
		}
		if listUserID != "" {
			req = req.UserId(listUserID)
		}
		if listCreatedSince != "" {
			if t, err := time.Parse(time.RFC3339, listCreatedSince); err == nil {
				req = req.CreatedSince(t)
			}
		}
		if listUpdatedSince != "" {
			if t, err := time.Parse(time.RFC3339, listUpdatedSince); err == nil {
				req = req.UpdatedSince(t)
			}
		}
		if listPage > 0 {
			req = req.Page(int32(listPage))
		}
		if listPerPage > 0 {
			req = req.PerPage(int32(listPerPage))
		}
		if listSpam {
			req = req.Spam(true)
		}

		ideas, resp, err := req.Execute()
		if err != nil {
			return fmt.Errorf("failed to list ideas: %w", err)
		}

		if resp.StatusCode >= 400 {
			return fmt.Errorf("API error: %s", resp.Status)
		}

		output, err := json.MarshalIndent(ideas, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		fmt.Println(string(output))
		return nil
	},
}

var getIdeaCmd = &cobra.Command{
	Use:   "get [idea-id]",
	Short: "Get a specific idea",
	Long:  `Get details for a specific idea by ID or reference number.`,
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

		idea, resp, err := apiClient.IdeasAPI.GetIdea(ctx, args[0]).Execute()
		if err != nil {
			return fmt.Errorf("failed to get idea: %w", err)
		}

		if resp.StatusCode >= 400 {
			return fmt.Errorf("API error: %s", resp.Status)
		}

		output, err := json.MarshalIndent(idea, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		fmt.Println(string(output))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(ideasCmd)
	ideasCmd.AddCommand(listIdeasCmd)
	ideasCmd.AddCommand(getIdeaCmd)

	// List flags
	listIdeasCmd.Flags().StringVarP(&listQuery, "query", "q", "", "Search term to match against idea name")
	listIdeasCmd.Flags().StringVarP(&listWorkflowStatus, "status", "s", "", "Filter by workflow status ID or name")
	listIdeasCmd.Flags().StringVar(&listSort, "sort", "", "Sort by: recent, trending, or popular")
	listIdeasCmd.Flags().StringVarP(&listTag, "tag", "t", "", "Filter by tag")
	listIdeasCmd.Flags().StringVar(&listUserID, "user-id", "", "Filter by creator user ID")
	listIdeasCmd.Flags().StringVar(&listCreatedSince, "created-since", "", "Only ideas created after this timestamp (RFC3339)")
	listIdeasCmd.Flags().StringVar(&listUpdatedSince, "updated-since", "", "Only ideas updated after this timestamp (RFC3339)")
	listIdeasCmd.Flags().IntVarP(&listPage, "page", "p", 0, "Page number")
	listIdeasCmd.Flags().IntVarP(&listPerPage, "per-page", "n", 0, "Results per page")
	listIdeasCmd.Flags().BoolVar(&listSpam, "spam", false, "Include spam ideas")

	// Suppress unused variable warnings
	_ = os.Stderr
}
