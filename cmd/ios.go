/*
Copyright © 2024 Mohammed Aman Khan <mohammed.aman@apptile.io>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// iosCmd represents the ios command
var iosCmd = &cobra.Command{
	Use:   "ios",
	Short: "This command will build the iOS applications for Apptile's react-native applications",
	// Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Building iOS application...")
	},
}

func init() {
	buildCmd.AddCommand(iosCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// iosCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// iosCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
