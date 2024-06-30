/*
Copyright © 2024 Mohammed Aman Khan <mohammed.aman@apptile.io>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// androidCmd represents the android command
var androidCmd = &cobra.Command{
	Use:   "android",
	Short: "This command will build the Android applications for Apptile's react-native applications",
	// Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Building Android application...")
	},
}

func init() {
	buildCmd.AddCommand(androidCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// androidCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// androidCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
