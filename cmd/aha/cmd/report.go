package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/grokify/gocharts/v2/data/table"
	"github.com/spf13/cobra"

	"github.com/grokify/go-aha/v3/oag7/aha"
	"github.com/grokify/go-aha/v3/oag7/client"
	"github.com/grokify/go-aha/v3/oag7/ideas"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate reports",
	Long:  `Commands for generating various Aha! reports.`,
}

var (
	reportQuery          string
	reportWorkflowStatus string
	reportTag            string
	reportPage           int
	reportPerPage        int
	reportFormat         string
	reportOutput         string
	reportIdeaPortalURL  string
	reportFeatureBaseURL string
	reportFilterFeature  string
	reportFilterRelease  string
	reportCompact        bool
	reportFetchAll       bool
)

var ideaFeatureReportCmd = &cobra.Command{
	Use:   "idea-feature",
	Short: "Generate idea-feature-release report",
	Long: `Generate a comprehensive report of ideas with their associated
features and releases. Output can be JSON, Markdown, or XLSX.`,
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

		req := ideas.ListIdeasRequest{
			Query:          reportQuery,
			WorkflowStatus: reportWorkflowStatus,
			Tag:            reportTag,
			FetchAll:       reportFetchAll,
		}
		if reportPage > 0 {
			req.Page = int32(reportPage)
		}
		if reportPerPage > 0 {
			req.PerPage = int32(reportPerPage)
		}

		reportSet, err := ideas.GetIdeaFeatureReports(ctx, apiClient, req)
		if err != nil {
			return fmt.Errorf("failed to generate report: %w", err)
		}

		// Apply filters
		if reportFilterFeature != "" {
			hasFeature := strings.ToLower(reportFilterFeature) == "yes" || reportFilterFeature == "true"
			reportSet = reportSet.FilterByHasFeature(hasFeature)
		}
		if reportFilterRelease != "" {
			hasRelease := strings.ToLower(reportFilterRelease) == "yes" || reportFilterRelease == "true"
			reportSet = reportSet.FilterByHasRelease(hasRelease)
		}

		// Sort by votes (most popular first)
		reportSet.SortByVotes()

		// Output the report
		switch strings.ToLower(reportFormat) {
		case "json":
			return outputJSON(reportSet)
		case "markdown", "md":
			return outputMarkdown(reportSet)
		case "xlsx", "excel":
			return outputXLSX(reportSet)
		default:
			return outputJSON(reportSet)
		}
	},
}

func outputJSON(reportSet *ideas.IdeaFeatureReportSet) error {
	output, err := json.MarshalIndent(reportSet.Reports, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal report: %w", err)
	}

	if reportOutput != "" {
		return os.WriteFile(reportOutput, output, 0644)
	}
	fmt.Println(string(output))
	return nil
}

func outputMarkdown(reportSet *ideas.IdeaFeatureReportSet) error {
	var t *table.Table

	if reportIdeaPortalURL != "" || reportFeatureBaseURL != "" {
		t = reportSet.TableWithLinks(reportIdeaPortalURL, reportFeatureBaseURL)
	} else if reportCompact {
		t = reportSet.TableCompact()
	} else {
		t = reportSet.Table()
	}

	md := t.Markdown("\n", true)

	if reportOutput != "" {
		return os.WriteFile(reportOutput, []byte(md), 0644)
	}
	fmt.Println(md)
	return nil
}

func outputXLSX(reportSet *ideas.IdeaFeatureReportSet) error {
	var tbl *table.Table

	if reportCompact {
		tbl = reportSet.TableCompact()
	} else {
		tbl = reportSet.Table()
	}

	outputFile := reportOutput
	if outputFile == "" {
		outputFile = fmt.Sprintf("idea_feature_report_%s.xlsx", time.Now().Format("20060102_150405"))
	}

	if err := tbl.WriteXLSX(outputFile, "Idea Feature Report"); err != nil {
		return fmt.Errorf("failed to write XLSX: %w", err)
	}

	fmt.Printf("Report written to: %s\n", outputFile)
	return nil
}

func init() {
	rootCmd.AddCommand(reportCmd)
	reportCmd.AddCommand(ideaFeatureReportCmd)

	// Query filters
	ideaFeatureReportCmd.Flags().StringVarP(&reportQuery, "query", "q", "", "Search term to match against idea name")
	ideaFeatureReportCmd.Flags().StringVarP(&reportWorkflowStatus, "status", "s", "", "Filter by workflow status")
	ideaFeatureReportCmd.Flags().StringVarP(&reportTag, "tag", "t", "", "Filter by tag")
	ideaFeatureReportCmd.Flags().IntVarP(&reportPage, "page", "p", 0, "Page number")
	ideaFeatureReportCmd.Flags().IntVarP(&reportPerPage, "per-page", "n", 0, "Results per page (default 30)")
	ideaFeatureReportCmd.Flags().BoolVarP(&reportFetchAll, "all", "a", false, "Fetch all pages (paginate automatically)")

	// Report filters
	ideaFeatureReportCmd.Flags().StringVar(&reportFilterFeature, "has-feature", "", "Filter by has feature (yes/no)")
	ideaFeatureReportCmd.Flags().StringVar(&reportFilterRelease, "has-release", "", "Filter by has release (yes/no)")

	// Output options
	ideaFeatureReportCmd.Flags().StringVarP(&reportFormat, "format", "f", "json", "Output format: json, markdown, xlsx")
	ideaFeatureReportCmd.Flags().StringVarP(&reportOutput, "output", "o", "", "Output file path (default: stdout for json/md, auto-named for xlsx)")
	ideaFeatureReportCmd.Flags().BoolVar(&reportCompact, "compact", false, "Use compact table format")

	// URL options for linked output
	ideaFeatureReportCmd.Flags().StringVar(&reportIdeaPortalURL, "idea-portal-url", "", "Base URL for idea portal links")
	ideaFeatureReportCmd.Flags().StringVar(&reportFeatureBaseURL, "feature-base-url", "", "Base URL for feature links")
}
