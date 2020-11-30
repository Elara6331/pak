package main

import (
	"fmt"
	"strings"
)

// Print help screen
func printHelpMessage(packageManagerCommand string, useRoot bool, rootCommand string, commands []string, shortcuts []string, shortcutMappings []string, isOverridden bool) {
	fmt.Println("Arsen Musayelyan's Package Manager Wrapper")
	fmt.Print("Current package manager is: ", packageManagerCommand)
	if isOverridden { fmt.Println(" (overridden)") } else { fmt.Print("\n") }
	if useRoot { fmt.Println("Using root with command:", rootCommand) } else { fmt.Println("Not using root") }
	fmt.Println()
	fmt.Println("Usage: pak <command> [package]")
	fmt.Println("Example: pak in hello")
	fmt.Println()
	fmt.Println("The available commands are:")
	fmt.Println(strings.Join(commands, "\n"))
	fmt.Println()
	fmt.Println("The available shortcuts are:")
	for index, element := range shortcuts { fmt.Println(element + ":", shortcutMappings[index]) }
	fmt.Println()
	fmt.Println("The available flags are:")
	fmt.Println("--help, -h: Shows this help screen")
	fmt.Println("--root, -r: Bypasses root user check")
	fmt.Println()
	fmt.Println("Pak uses a string distance algorithm, so `pak in` is valid as is `pak inst` or `pak install`")
}