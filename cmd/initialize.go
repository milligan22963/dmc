// Package cmd contains all CLI commands used by the application.
package cmd

import (
	"github.com/spf13/cobra"
)

const initializeCommandLongDesc string = "initializes the dmc database"

// initializeCmd represents the setup command to make a system dmc ready
var initializeCmd = &cobra.Command{
	Use:               "initialize",
	Version:           "1.0.0",
	Short:             "An app to manage core based iot devices",
	Long:              initializeCommandLongDesc,
	Example:           `dmc initialize`,
	Args:              cobra.NoArgs,
	PersistentPreRunE: ValidateInitializeFlags,
	RunE:              RunInitializeCmd,
}

//func init() {
//}

// ValidateInitializeFlags checks that the flags are within expected boundaries.
func ValidateInitializeFlags(cmd *cobra.Command, args []string) error {

	// place holder for validation
	return nil
}

// RunInitializeCmd is executed when the application is run without any subcommands.
func RunInitializeCmd(cmd *cobra.Command, args []string) error {
	cmd.Printf("initialize DMC called\n")
	return nil
}
