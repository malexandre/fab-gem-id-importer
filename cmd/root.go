/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func getStringFromFlag(cmd *cobra.Command, flag string) string {
	value, err := cmd.Flags().GetString(flag)

	if err != nil {
		log.Fatalf("%s flag has an err: %v\n", flag, err)
	}

	return value
}

func getStringFromFlagThenEnv(cmd *cobra.Command, flag string) string {
	value := getStringFromFlag(cmd, flag)

	if value == "" {
		value = os.Getenv(flag)

		if value == "" {
			log.Fatalf("No %s defined\n", flag)
		}
	}

	return value
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fab-gem-id-importer [flags] path_to_csv.csv",
	Short: "fab-gem-id-importer is a tool to import users from a CSV to a GEM event.",
	Long:  "fab-gem-id-importer is a tool to import users from a CSV to a GEM event. User login and password can be set through .env or with flags.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		godotenv.Load()

		username := getStringFromFlagThenEnv(cmd, "username")
		password := getStringFromFlagThenEnv(cmd, "password")
		eventId := getStringFromFlag(cmd, "event-id")

		// TODO: Replace fmt by the business logic
		fmt.Println("Login as", username, "with password", password, "for event", eventId)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("username", "u", "", "Name of the GEM account to log in as")
	rootCmd.Flags().StringP("password", "p", "", "Password of the GEM account to log in as")
	rootCmd.Flags().StringP("event-id", "e", "", "GEM event ID to update with the users. Found in the GEM URL of the event.")
	rootCmd.MarkFlagRequired("event-id")
}
