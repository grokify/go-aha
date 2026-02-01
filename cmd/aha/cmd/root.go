package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	domain string
	apiKey string
)

var rootCmd = &cobra.Command{
	Use:   "aha",
	Short: "Aha! API CLI client",
	Long:  `A command-line interface for the Aha! Roadmap Service API.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&domain, "domain", "d", "", "Aha subdomain (or set AHA_DOMAIN env var)")
	rootCmd.PersistentFlags().StringVarP(&apiKey, "api-key", "k", "", "Aha API key (or set AHA_API_KEY env var)")
}

func getDomain() string {
	if domain != "" {
		return domain
	}
	if env := os.Getenv("AHA_DOMAIN"); env != "" {
		return env
	}
	return ""
}

func getAPIKey() string {
	if apiKey != "" {
		return apiKey
	}
	if env := os.Getenv("AHA_API_KEY"); env != "" {
		return env
	}
	return ""
}

func validateCredentials() error {
	if getDomain() == "" {
		return fmt.Errorf("domain is required: use --domain flag or set AHA_DOMAIN env var")
	}
	if getAPIKey() == "" {
		return fmt.Errorf("API key is required: use --api-key flag or set AHA_API_KEY env var")
	}
	return nil
}
