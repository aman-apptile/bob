/*
Copyright © 2024 Mohammed Aman Khan <mohammed.aman@apptile.io>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "This command will setup the environment required to make Android and iOS builds for Apptile's react-native applications",
	// Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Setting up the environment for Android and iOS builds...\nInstalling required dependencies...\n\n- JDK\n- Xcode\n- Watchman\n- CocoaPods\n- Gradle\n- Android SDK\n- Android NDK\n- Android Emulator\n- Android Platform Tools\n- Android Build Tools\n- RBENV & Ruby\n- NVM & Node\n- Brew")
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
