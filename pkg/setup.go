package pkg

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aman-apptile/bob/pkg/utils"
)

// setupHomebrewPackages installs the necessary Homebrew packages.
func setupHomebrewPackages(packages []string) {
	for _, pkg := range packages {
		utils.InstallPackage(pkg)
	}
}

// setupNVM installs and configures Node Version Manager (NVM).
func setupNVM(homeDir string) {
	nvmInstallScript := "https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh"

	fmt.Println("Installing NVM...")
	err := utils.RunCommand("curl", "-o-", nvmInstallScript, "|", "bash")
	utils.CheckError(err, "Failed to install NVM")

	zshrc := filepath.Join(homeDir, ".zshrc")
	nvmConfig := `
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh" # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion" # This loads nvm bash_completion
`
	utils.AppendToFile(zshrc, nvmConfig)

	err = utils.RunCommand("source", zshrc)
	utils.CheckError(err, "Failed to source .zshrc")

	fmt.Println("Installing Node.js using NVM...")
	err = utils.RunCommand("nvm", "install", "node")
	utils.CheckError(err, "Failed to install Node.js using NVM")
}

// setupRbenv installs and configures rbenv.
func setupRbenv(homeDir string) {
	fmt.Println("Installing rbenv...")
	utils.InstallPackage("rbenv")

	fmt.Println("Installing ruby-build plugin for rbenv...")
	utils.InstallPackage("ruby-build")

	zshrc := filepath.Join(homeDir, ".zshrc")
	rbenvConfig := `
export PATH="$HOME/.rbenv/bin:$PATH"
eval "$(rbenv init -)"
`
	utils.AppendToFile(zshrc, rbenvConfig)

	utils.RunCommand("source", zshrc)

	fmt.Println("Installing Ruby using rbenv...")
	utils.RunCommand("rbenv", "install", "2.7.2")

	utils.RunCommand("rbenv", "global", "2.7.2")
}

// setupAndroidSDK sets up the Android SDK.
func setupAndroidSDK(homeDir string) {
	sdkDir := filepath.Join(homeDir, "Library", "Android", "sdk", "cmdline-tools", "latest")
	err := os.MkdirAll(sdkDir, os.ModePerm)
	utils.CheckError(err, "Failed to create Android SDK directory")

	err = utils.DownloadAndExtract("https://dl.google.com/android/repository/commandlinetools-mac-7583922_latest.zip", sdkDir)
	utils.CheckError(err, "Failed to download and extract Android SDK command line tools")

	sdkRoot := filepath.Join(homeDir, "Library", "Android", "sdk")
	zshrc := filepath.Join(homeDir, ".zshrc")
	utils.AppendToFile(zshrc, fmt.Sprintf("\nexport ANDROID_SDK_ROOT=%s\nexport PATH=$PATH:%s/cmdline-tools/latest/bin:%s/platform-tools\n", sdkRoot, sdkRoot, sdkRoot))

	sdkPackages := []string{
		"platform-tools",
		"platforms;android-30",
		"build-tools;30.0.3",
		"emulator",
		"ndk-bundle",
	}
	for _, pkg := range sdkPackages {
		err := utils.RunCommand("sdkmanager", "--install", pkg)
		utils.CheckError(err, fmt.Sprintf("Failed to install SDK package: %s", pkg))
	}
}

// setupGradle installs and sets up Gradle.
func setupGradle(homeDir string) {
	err := utils.DownloadAndExtract("https://services.gradle.org/distributions/gradle-7.2-bin.zip", "/usr/local")
	utils.CheckError(err, "Failed to download and extract Gradle")

	zshrc := filepath.Join(homeDir, ".zshrc")
	utils.AppendToFile(zshrc, "export PATH=$PATH:/usr/local/gradle-7.2/bin\n")
}

// setupXcode installs Xcode command line tools and accepts the license.
func setupXcode() {
	err := utils.RunCommand("xcode-select", "--install")
	if err != nil {
		log.Println("Xcode command line tools already installed")
	} else {
		utils.CheckError(err, "Failed to install Xcode command line tools")
	}

	err = utils.RunCommand("sudo", "xcodebuild", "-license", "accept")
	utils.CheckError(err, "Failed to accept Xcode license")
}

// setupCocoapods installs CocoaPods using Ruby gem.
func setupCocoapods() {
	err := utils.RunCommand("sudo", "gem", "install", "cocoapods")
	utils.CheckError(err, "Failed to install CocoaPods")
}

func Init() {
	homeDir, err := os.UserHomeDir()
	utils.CheckError(err, "Failed to get home directory")

	fmt.Println("Setting up development environment...")
	setupHomebrewPackages([]string{"openjdk@11", "ruby-build"})
	setupNVM(homeDir)
	setupRbenv(homeDir)
	setupAndroidSDK(homeDir)
	setupGradle(homeDir)
	setupXcode()
	setupCocoapods()
	fmt.Println("Development environment setup complete!")
}
