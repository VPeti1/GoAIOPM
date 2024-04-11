package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func welcome() {
	clear()
	fmt.Println("Welcome to the AIOPM!")
	fmt.Println("Available commands: 'update', 'install', 'remove', 'exit'")
	fmt.Println("Type 'flatpak' for managing flatpak")
	fmt.Println("Type 'snap' for managing snap")
	fmt.Println("Type 'pip' for managing pip")
}

func flatpak() {
	clear()
	fmt.Println("You are now managing flatpak")
	fmt.Println("Available commands: 'update', 'install', 'remove', 'exit'")
	fmt.Print("> ")
	input := getInput("")
	switch input {
	case "update", "Update":
		cmd := exec.Command("flatpak", "update")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	case "install", "Install":
		input := getInput("Enter packages name(s): ")
		cmd := exec.Command("flatpak", "install", input)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	case "remove", "Remove":
		input := getInput("Enter packages name(s): ")
		cmd := exec.Command("flatpak", "remove", input)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	case "exit", "Exit":
	default:
		fmt.Println("Invalid input! Retrying")
		systemPause()
		welcome()
		flatpak()
	}
}

func pip() {
	clear()
	fmt.Println("You are now managing pip")
	fmt.Println("Available commands: 'install', 'remove', 'exit'")
	fmt.Print("> ")
	input := getInput("")
	switch input {
	case "install", "Install":
		input := getInput("Enter packages name(s): ")
		cmd := exec.Command("pip", "install", "--break-system-packages", input)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	case "remove", "Remove":
		input := getInput("Enter packages name(s): ")
		cmd := exec.Command("pip", "uninstall", "--break-system-packages", input)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	case "exit", "Exit":
	default:
		fmt.Println("Invalid input! Retrying")
		systemPause()
		pip()
	}
}

func snap() {
	clear()
	fmt.Println("You are now managing snap")
	fmt.Println("Available commands: 'install', 'remove', 'exit'")
	fmt.Print("> ")
	input := getInput("")
	switch input {
	case "install", "Install":
		input := getInput("Enter packages name(s): ")
		cmd := exec.Command("sudo", "snap", "install", input)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	case "remove", "Remove":
		input := getInput("Enter packages name(s): ")
		cmd := exec.Command("sudo", "snap", "remove", input)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	case "exit", "Exit":
	default:
		fmt.Println("Invalid input! Retrying")
		systemPause()
		snap()
	}
}

func arch() {
	fmt.Println("You are now managing Arch Linux")
	fmt.Println("Available commands: 'update', 'install', 'remove', 'aur', 'pip', 'snap', 'mremove', 'exit'")
	fmt.Println("Enter your command:")
	input := getInput("")

	switch input {
	case "update", "Update":
		cmd := exec.Command("sudo", "pacman", "-Syy")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
		cmd = exec.Command("sudo", "pacman", "-Syu")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
		arch()
	case "install", "Install":
		clear()
		fmt.Println("Enter package(s) name(s): ")
		packages := getInput("")
		cmd := exec.Command("sudo", "pacman", "-S", packages)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
		arch()
	case "remove", "Remove":
		clear()
		fmt.Println("Enter package(s) name(s): ")
		packages := getInput("")
		cmd := exec.Command("sudo", "pacman", "-R", packages)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
		arch()
	case "pip", "Pip":
		pip()
	case "snap", "Snap":
		snap()
	case "mremove", "Mremove":
		clear()
		fmt.Println("Enter package(s) name(s): ")
		packages := getInput("")
		cmd := exec.Command("sudo", "pacman", "-R", "$(pacman", "-Qq", "|", "grep", packages+")")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
		arch()
	case "exit", "Exit":
		clear()
	default:
		fmt.Println("Invalid input! Retrying")
		systemPause()
		arch()
	}
}

func deb() {
	fmt.Println("You are now managing Debian")
	fmt.Println("Available commands: 'update', 'install', 'remove', 'flatpak', 'pip', 'snap', 'exit'")
	fmt.Println("Enter your command:")
	input := getInput("")

	switch input {
	case "update", "Update":
		cmd := exec.Command("sudo", "apt", "update")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
		deb()
	case "install", "Install":
		clear()
		fmt.Println("Enter package(s) name(s): ")
		packages := getInput("")
		cmd := exec.Command("sudo", "apt", "install", packages)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
		deb()
	case "remove", "Remove":
		clear()
		fmt.Println("Enter package(s) name(s): ")
		packages := getInput("")
		cmd := exec.Command("sudo", "apt", "remove", packages)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
		deb()
	case "flatpak", "Flatpak":
		flatpak()
	case "pip", "Pip":
		pip()
	case "snap", "Snap":
		snap()
	case "exit", "Exit":
		clear()
	default:
		fmt.Println("Invalid input! Retrying")
		systemPause()
		deb()
	}
}

func fed() {
	fmt.Println("You are now managing Fedora")
	fmt.Println("Available commands: 'update', 'install', 'remove', 'flatpak', 'pip', 'snap', 'exit'")
	fmt.Println("Enter your command:")
	input := getInput("")

	switch input {
	case "update", "Update":
		cmd := exec.Command("sudo", "yum", "update")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
		fed()
	case "install", "Install":
		clear()
		fmt.Println("Enter package(s) name(s): ")
		packages := getInput("")
		cmd := exec.Command("sudo", "yum", "install", packages)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
		fed()
	case "remove", "Remove":
		clear()
		fmt.Println("Enter package(s) name(s): ")
		packages := getInput("")
		cmd := exec.Command("sudo", "yum", "remove", packages)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
		fed()
	case "flatpak", "Flatpak":
		flatpak()
	case "pip", "Pip":
		pip()
	case "snap", "Snap":
		snap()
	case "exit", "Exit":
		clear()
	default:
		fmt.Println("Invalid input! Retrying")
		systemPause()
		fed()
	}
}

func opensuse() {
	fmt.Println("You are now managing openSUSE")
	fmt.Println("Available commands: 'update', 'install', 'remove', 'flatpak', 'pip', 'snap', 'exit'")
	fmt.Println("Enter your command:")
	input := getInput("")

	switch input {
	case "update", "Update":
		cmd := exec.Command("sudo", "zypper", "update")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
		opensuse()
	case "install", "Install":
		clear()
		fmt.Println("Enter package(s) name(s): ")
		packages := getInput("")
		cmd := exec.Command("sudo", "zypper", "install", packages)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
		opensuse()
	case "remove", "Remove":
		clear()
		fmt.Println("Enter package(s) name(s): ")
		packages := getInput("")
		cmd := exec.Command("sudo", "zypper", "remove", packages)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
		opensuse()
	case "flatpak", "Flatpak":
		flatpak()
	case "pip", "Pip":
		pip()
	case "snap", "Snap":
		snap()
	case "exit", "Exit":
		clear()
	default:
		fmt.Println("Invalid input! Retrying")
		systemPause()
		opensuse()
	}
}

func voidl() {
	fmt.Println("You are now managing Void Linux")
	fmt.Println("Available commands: 'update', 'install', 'remove', 'flatpak', 'pip', 'snap', 'exit'")
	fmt.Println("Enter your command:")
	input := getInput("")

	switch input {
	case "update", "Update":
		cmd := exec.Command("sudo", "xbps-install", "-Su")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
		voidl()
	case "install", "Install":
		clear()
		fmt.Println("Enter package(s) name(s): ")
		packages := getInput("")
		cmd := exec.Command("sudo", "xbps-install", packages)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
		voidl()
	case "remove", "Remove":
		clear()
		fmt.Println("Enter package(s) name(s): ")
		packages := getInput("")
		cmd := exec.Command("sudo", "xbps-remove", packages)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
		voidl()
	case "flatpak", "Flatpak":
		flatpak()
	case "pip", "Pip":
		pip()
	case "snap", "Snap":
		snap()
	case "exit", "Exit":
		clear()
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid input! Retrying")
		systemPause()
		voidl()
	}
}

func fst() {
	clear()
	fmt.Println("Welcome to AIOPM Setup!")
	fmt.Println("What distro are you using?")
	fmt.Println("(arch,debian,fedora,opensuse or void)")
	fmt.Println("(Derivatives included)")
	fmt.Print("> ")
	input := getInput("")
	switch input {
	case "arch", "Arch":
		fmt.Println("Setting configuration for Arch")
		exec.Command("sudo", "mkdir", "/usr/aiopm").Run()
		exec.Command("sudo", "touch", "/usr/aiopm/a1.cw").Run()
	case "debian", "Debian":
		fmt.Println("Setting configuration for Debian")
		exec.Command("sudo", "mkdir", "/usr/aiopm").Run()
		exec.Command("sudo", "touch", "/usr/aiopm/a2.cw").Run()
	case "fedora", "Fedora":
		fmt.Println("Setting configuration for Fedora")
		exec.Command("sudo", "mkdir", "/usr/aiopm").Run()
		exec.Command("sudo", "touch", "/usr/aiopm/a3.cw").Run()
	case "opensuse", "Opensuse":
		exec.Command("sudo", "mkdir", "/usr/aiopm").Run()
		exec.Command("sudo", "touch", "/usr/aiopm/a4.cw").Run()
	case "void", "Void":
		exec.Command("sudo", "mkdir", "/usr/aiopm").Run()
		exec.Command("sudo", "touch", "/usr/aiopm/a5.cw").Run()
	default:
		fmt.Println("Invalid input! Retrying")
		systemPause()
		fst()
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func systemPause() {
	fmt.Println("Press Enter to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	clear()
	fmt.Println(" $$$$$$\\  $$$$$$\\  $$$$$$\\  $$$$$$$\\  $$\\      $$\\ ")
	fmt.Println("$$  __$$\\ \\_$$  _|$$  __$$\\ $$  __$$\\ $$$\\    $$$ |")
	fmt.Println("$$ /  $$ |  $$ |  $$ /  $$ |$$ |  $$ |$$$$\\  $$$$ |")
	fmt.Println("$$$$$$$$ |  $$ |  $$ |  $$ |$$$$$$$  |$$\\$$\\$$ $$ |")
	fmt.Println("$$  __$$ |  $$ |  $$ |  $$ |$$  ____/ $$ \\$$$  $$ |")
	fmt.Println("$$ |  $$ |  $$ |  $$ |  $$ |$$ |      $$ |\\$  /$$ |")
	fmt.Println("$$ |  $$ |$$$$$$\\  $$$$$$  |$$ |      $$ | \\_/ $$ |")
	fmt.Println("\\__|  \\__|\\______| \\______/ \\__|      \\__|     \\__|")
	fmt.Println("By VPeti")
	time.Sleep(2 * time.Second)
	_, err := os.Open("/usr/aiopm/a1.cw")
	if err == nil {
		welcome()
		arch()
		return
	}
	_, err = os.Open("/usr/aiopm/a2.cw")
	if err == nil {
		welcome()
		deb()
		return
	}
	_, err = os.Open("/usr/aiopm/a3.cw")
	if err == nil {
		welcome()
		fed()
		return
	}
	_, err = os.Open("/usr/aiopm/a4.cw")
	if err == nil {
		welcome()
		opensuse()
		return
	}
	_, err = os.Open("/usr/aiopm/a5.cw")
	if err == nil {
		welcome()
		voidl()
		return
	}
	fst()
}
