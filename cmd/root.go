// Package cmd contains all CLI commands used by the application.
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags
	port int
)

const rootCommandLongDesc string = "dmc allows device management via a console"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:               "dmc [flags]",
	Version:           "1.0.0",
	Short:             "An app to manage core based iot devices",
	Long:              rootCommandLongDesc,
	Example:           `dmc`,
	Args:              cobra.NoArgs,
	PersistentPreRunE: ValidateFlags,
	RunE:              RunRootCmd,
}

func init() {
	// global params/flags
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 80, "Port to listen on.")
}

// ValidateFlags checks that the flags are within expected boundaries.
func ValidateFlags(cmd *cobra.Command, args []string) error {

	// place holder for validation
	return nil
}

// RunRootCmd is executed when the application is run without any subcommands.
func RunRootCmd(cmd *cobra.Command, args []string) error {
	cmd.Printf("run DMC called")
	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
