/*
Copyright © 2024 Mohammed Aman Khan <mohammed.aman@apptile.io>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/aman-apptile/bob/pkg"
	"github.com/aman-apptile/bob/pkg/constants"
	"github.com/aman-apptile/bob/pkg/utils"
	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "This command will setup the environment required to make Android and iOS builds for Apptile's react-native applications",
	// Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		utils.CheckError(err, "Failed to get home directory")

		fmt.Println("Setting up development environment...")

		pkg.SetupHomebrew()
		pkg.SetupHomebrewPackages([]string{"openjdk@" + constants.REQUIRED_JDK_VERSION, "gradle"})
		pkg.SetupNVM(homeDir)
		pkg.SetupRbenv(homeDir)
		pkg.SetupCocoapods()
		pkg.SetupAndroidEnvironment(homeDir)
		pkg.SetupIosEnvironment()

		fmt.Println("Development environment setup complete!")
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
