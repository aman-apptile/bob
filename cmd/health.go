/*
Copyright © 2024 Mohammed Aman Khan <mohammed.aman@apptile.io>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/aman-apptile/bob/pkg"
	"github.com/aman-apptile/bob/pkg/utils"
	"github.com/spf13/cobra"
)

// healthCmd represents the health command
var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "This command will check the health of the environment required to make Android and iOS builds for Apptile's react-native applications",
	// Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		utils.CheckError(err, "Failed to get home directory")

		fmt.Println("Checking the health of the development environment...")

		s := utils.StartSpinner(" Checking Homebrew")
		result := pkg.CheckHomebrew()
		if result {
			utils.StopSpinner(s, " Homebrew is installed.", "success")
		} else {
			utils.StopSpinner(s, " Homebrew is not installed.", "failure")
		}

		s = utils.StartSpinner(" Checking required Homebrew packages")
		result = pkg.CheckHomebrewPackages([]string{"openjdk", "gradle"})
		if result {
			utils.StopSpinner(s, " Required Homebrew packages are installed.", "success")
		} else {
			utils.StopSpinner(s, " Required Homebrew packages are not installed.", "failure")
		}

		s = utils.StartSpinner(" Checking NVM")
		result = pkg.CheckNVM()
		if result {
			utils.StopSpinner(s, " NVM is installed.", "success")
		} else {
			utils.StopSpinner(s, " NVM is not installed.", "failure")
		}

		s = utils.StartSpinner(" Checking Node.js")
		result = pkg.CheckNode()
		if result {
			utils.StopSpinner(s, " Node.js is installed.", "success")
		} else {
			utils.StopSpinner(s, " Node.js is not installed.", "failure")
		}

		s = utils.StartSpinner(" Checking Rbenv")
		result = pkg.CheckRbenv()
		if result {
			utils.StopSpinner(s, " Rbenv is installed.", "success")
		} else {
			utils.StopSpinner(s, " Rbenv is not installed.", "failure")
		}

		s = utils.StartSpinner(" Checking Ruby")
		result = pkg.CheckRuby()
		if result {
			utils.StopSpinner(s, " Ruby is installed.", "success")
		} else {
			utils.StopSpinner(s, " Ruby is not installed.", "failure")
		}

		s = utils.StartSpinner(" Checking CocoaPods")
		result = pkg.CheckCocoapods()
		if result {
			utils.StopSpinner(s, " CocoaPods is installed.", "success")
		} else {
			utils.StopSpinner(s, " CocoaPods is not installed.", "failure")
		}

		s = utils.StartSpinner(" Checking Android environment")
		result = pkg.CheckAndroidEnvironment(homeDir)
		if result {
			utils.StopSpinner(s, " Android environment is setup.", "success")
		} else {
			utils.StopSpinner(s, " Android environment is not setup.", "failure")
		}

		s = utils.StartSpinner(" Checking iOS environment")
		result = pkg.CheckIosEnvironment()
		if result {
			utils.StopSpinner(s, " iOS environment is setup.", "success")
		} else {
			utils.StopSpinner(s, " iOS environment is not setup.", "failure")
		}
	},
}

func init() {
	rootCmd.AddCommand(healthCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// healthCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// healthCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
