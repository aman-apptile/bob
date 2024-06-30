package pkg

import (
	"os"
	"path/filepath"

	"github.com/aman-apptile/bob/pkg/utils"
)

// CheckCocoapods checks if CocoaPods is installed or not.
func CheckCocoapods() bool {
	return utils.IsGemInstalled("cocoapods")
}

// CheckHomebrew checks if Homebrew is installed or not.
func CheckHomebrew() bool {
	return utils.IsCommandAvailable("brew")
}

// CheckHomebrewPackages checks if the necessary Homebrew packages are installed or not.
func CheckHomebrewPackages(packages []string) bool {
	for _, pkg := range packages {
		if !utils.IsPackageInstalled(pkg) {
			return false
		}
	}

	return true
}

// CheckNVM checks if Node Version Manager (NVM) is installed or not.
func CheckNVM() bool {
	return utils.IsCommandAvailable("nvm")
}

// CheckNode checks if Node.js is installed or not.
func CheckNode() bool {
	return utils.IsCommandAvailable("node")
}

// CheckRbenv checks if Ruby Version Manager (Rbenv) is installed or not.
func CheckRbenv() bool {
	return utils.IsCommandAvailable("rbenv")
}

// CheckRuby checks if Ruby is installed or not.
func CheckRuby() bool {
	return utils.IsCommandAvailable("ruby")
}

// CheckAndroidEnvironment checks if Android environment is setup or not.
func CheckAndroidEnvironment(homeDir string) bool {
	sdkRoot := filepath.Join(homeDir, "Library", "Android", "sdk")

	if _, err := os.Stat(sdkRoot); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

// CheckIosEnvironment checks if iOS environment is setup or not.
func CheckIosEnvironment() bool {
	return utils.IsCommandAvailable("xcodebuild") && utils.IsCommandAvailable("pod")
}
