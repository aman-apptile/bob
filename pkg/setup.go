package pkg

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aman-apptile/bob/pkg/utils"
)

// SetupCocoapods installs CocoaPods using Ruby gem if it is not already installed.
func SetupCocoapods() {
	if !utils.IsGemInstalled("cocoapods") {
		err := utils.RunCommand("sudo", "gem", "install", "cocoapods")
		utils.CheckError(err, "Failed to install CocoaPods")
	} else {
		fmt.Println("CocoaPods is already installed.")
	}
}

// SetupHomebrew installs Homebrew if not already installed.
func SetupHomebrew() {
	if !utils.IsCommandAvailable("brew") {
		_, err := os.Stat("/opt/homebrew")
		if os.IsNotExist(err) {
			err := utils.RunCommand("/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)")
			utils.CheckError(err, "Failed to install Homebrew")
		}
	} else {
		fmt.Println("Homebrew is already installed.")
	}
}

// SetupHomebrewPackages installs the necessary Homebrew packages if they are not already installed.
func SetupHomebrewPackages(packages []string) {
	for _, pkg := range packages {
		if !utils.IsPackageInstalled(pkg) {
			utils.InstallPackage(pkg)
		} else {
			fmt.Printf("%s is already installed.\n", pkg)
		}
	}
}

// SetupNVM installs and configures Node Version Manager (NVM) if it is not already installed.
func SetupNVM(homeDir string) {
	if !utils.IsCommandAvailable("nvm") {
		REQUIRED_NODE_VERSION := os.Getenv("REQUIRED_NODE_VERSION")

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
		err = utils.RunCommand("nvm", "install", REQUIRED_NODE_VERSION)
		utils.CheckError(err, "Failed to install Node.js using NVM")

		err = utils.RunCommand("nvm", "alias", "default", REQUIRED_NODE_VERSION)
		utils.CheckError(err, "Failed to set default Node.js version")

		err = utils.RunCommand("nvm", "use", REQUIRED_NODE_VERSION)
		utils.CheckError(err, fmt.Sprintf("Failed to use Node.js version: %s", REQUIRED_NODE_VERSION))
	} else {
		fmt.Println("NVM is already installed.")
	}
}

// SetupRbenv installs and configures rbenv if it is not already installed.
func SetupRbenv(homeDir string) {
	if !utils.IsCommandAvailable("rbenv") {
		REQUIRED_RUBY_VERSION := os.Getenv("REQUIRED_RUBY_VERSION")

		fmt.Println("Installing rbenv...")
		utils.InstallPackage("rbenv")

		fmt.Println("Installing ruby-build plugin for rbenv...")
		utils.InstallPackage("ruby-build")

		lines := []string{
			"export PATH=\"$HOME/.rbenv/bin:$PATH\"",
			"eval \"$(rbenv init -)\"",
			"export PATH=\"${HOME}/.rbenv/shims:${PATH}\"",
			"export RBENV_SHELL=zsh",
			"source '/opt/homebrew/Cellar/rbenv/1.2.0/libexec/../completions/rbenv.zsh'",
			"command rbenv rehash 2>/dev/null",
			"rbenv() {",
			"local command command=\"${1:-}\"",
			"if [ \"$#\" -gt 0 ]; then",
			"shift",
			"fi",
			"case \"$command\" in rehash|shell)",
			"eval \"$(rbenv \"sh-$command\" \"$@\")\";; *)",
			"command rbenv \"$command\" \"$@\";;",
			"esac }",
		}
		utils.AppendLinesToZshrc(lines...)

		utils.RunCommand("source", filepath.Join(homeDir, ".zshrc"))

		fmt.Println("Installing Ruby using rbenv...")
		utils.RunCommand("rbenv", "install", REQUIRED_RUBY_VERSION)

		utils.RunCommand("rbenv", "global", REQUIRED_RUBY_VERSION)
	} else {
		fmt.Println("rbenv is already installed.")
	}
}

// SetupAndroidEnvironment sets up the Android SDK, NDK, Platform Tools, Emulator & Build Tools if it they are not already set up.
func SetupAndroidEnvironment(homeDir string) {
	sdkRoot := filepath.Join(homeDir, "Library", "Android", "sdk")
	if _, err := os.Stat(sdkRoot); os.IsNotExist(err) {
		sdkDir := filepath.Join(sdkRoot, "cmdline-tools", "latest")
		err := os.MkdirAll(sdkDir, os.ModePerm)
		utils.CheckError(err, "Failed to create Android SDK directory")

		err = utils.DownloadAndExtract("https://dl.google.com/android/repository/commandlinetools-mac-7583922_latest.zip", sdkDir)
		utils.CheckError(err, "Failed to download and extract Android SDK command line tools")

		lines := []string{
			"export ANDROID_SDK_ROOT=" + sdkRoot,
			"export PATH=$PATH:" + sdkDir + "/bin:" + sdkRoot + "/platform-tools",
		}
		utils.AppendLinesToZshrc(lines...)

		utils.RunCommand("source", filepath.Join(homeDir, ".zshrc"))
		// utils.AppendToFile(zshrc, fmt.Sprintf("\nexport ANDROID_SDK_ROOT=%s\nexport PATH=$PATH:%s/cmdline-tools/latest/bin:%s/platform-tools\n", sdkRoot, sdkRoot, sdkRoot))

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
	} else {
		fmt.Println("Android SDK is already set up.")
	}

	// if !utils.IsCommandAvailable("gradle") {
	// 	err := utils.DownloadAndExtract("https://services.gradle.org/distributions/gradle-7.2-bin.zip", "/usr/local")
	// 	utils.CheckError(err, "Failed to download and extract Gradle")

	// 	zshrc := filepath.Join(homeDir, ".zshrc")
	// 	utils.AppendToFile(zshrc, "export PATH=$PATH:/usr/local/gradle-7.2/bin\n")
	// } else {
	// 	fmt.Println("Gradle is already installed.")
	// }
}

// SetupIosEnvironment installs or updates Xcode command line tools and accepts the license.
func SetupIosEnvironment() {
	if !utils.IsCommandAvailable("xcode-select") {
		err := utils.RunCommand("xcode-select", "--install")
		if err != nil {
			log.Println("Xcode command line tools installation attempt failed, possibly already installed.")
		} else {
			utils.CheckError(err, "Failed to install Xcode command line tools")
		}
	} else {
		fmt.Println("Xcode command line tools are already installed. Checking for updates...")
		// Attempt to update Xcode command line tools
		err := utils.RunCommand("softwareupdate", "--install", "-a")
		utils.CheckError(err, "Failed to update Xcode command line tools")
	}

	// Accept the Xcode license
	err := utils.RunCommand("sudo", "xcodebuild", "-license", "accept")
	utils.CheckError(err, "Failed to accept Xcode license")
}
