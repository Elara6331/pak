/*
   Pak: Wrapper designed for package managers to unify software management commands between distros
   Copyright (C) 2020 Arsen Musayelyan

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"strings"
)


// Print help screen

func printHelpMessage(packageManagerCommand string, useRootBool bool, rootCommand string, commands []string, shortcuts []string) {
	fmt.Println("Arsen Musayelyan's Package Manager Wrapper")
	fmt.Println("Current package manager is:", packageManagerCommand)
	if useRootBool { fmt.Println("Using root with:", rootCommand, "\n") } else { fmt.Println("Not using root\n") }
	fmt.Println("Usage: pak <command> [package]\nExample: pak in hello\n")
	fmt.Println("The available commands are:\n" + strings.Join(commands, "\n"),"\n")
	fmt.Println("The available shortcuts are:\n" + strings.Join(shortcuts, "\n"), "\n")
	fmt.Println("The available flags are:\n--help, -h: Shows this help screen\n--root, -r: Bypass root error")
	fmt.Println("misc: All", packageManagerCommand, "flags\n")
	fmt.Println("Writing the whole command is uneccesary, just use enough to differentiate")
	os.Exit(0)
}

// Remove an element at an index from a slice
func removeAtIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

// Check if slice contains string
func Contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func Max(array []float64) float64 {
	var max float64 = array[0]
	var min float64 = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return max
}

func Find(slice []string, val string) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}
	return -1
}

func Jaro(a, b string) float64 {
	la := float64(len(a))
	lb := float64(len(b))

	// match range = max(len(a), len(b)) / 2 - 1
	matchRange := int(math.Floor(math.Max(la, lb)/2.0)) - 1
	matchRange = int(math.Max(0, float64(matchRange-1)))
	var matches, halfs float64
	transposed := make([]bool, len(b))

	for i := 0; i < len(a); i++ {
		start := int(math.Max(0, float64(i-matchRange)))
		end := int(math.Min(lb-1, float64(i+matchRange)))

		for j := start; j <= end; j++ {
			if transposed[j] {
				continue
			}

			if a[i] == b[j] {
				if i != j {
					halfs++
				}
				matches++
				transposed[j] = true
				break
			}
		}
	}

	if matches == 0 {
		return 0
	}

	transposes := math.Floor(float64(halfs / 2))

	return ((matches / la) + (matches / lb) + (matches-transposes)/matches) / 3.0
}


func JaroWinkler(a, b string, boostThreshold float64, prefixSize int) float64 {
	j := Jaro(a, b)

	if j <= boostThreshold {
		return j
	}

	prefixSize = int(math.Min(float64(len(a)), math.Min(float64(prefixSize), float64(len(b)))))

	var prefixMatch float64
	for i := 0; i < prefixSize; i++ {
		if a[i] == b[i] {
			prefixMatch++
		}
	}

	return j + 0.1*prefixMatch*(1.0-j)
}


func main()  {
	// Put all arguments into a variable
	args := os.Args[1:]

	// Check which user is running command
	usr, err := exec.Command("whoami").Output()
	if err != nil {
		log.Fatal(err)
	}

	// Check to make sure root is not being used unless -r/--root specified
	if !Contains(args, "-r") && !Contains(args, "--root") {
		if strings.Contains(string(usr), "root") {
			fmt.Println("Do not run as root, this program will invoke root for you if selected in config.")
			fmt.Println("If you would like to bypass this, run this command with -r or --root.")
			os.Exit(1)
		}
	} else {
		if Contains(args, "-r") {
			args = removeAtIndex(args, Find(args, "-r"))
		} else if Contains(args, "--root") {
			args = removeAtIndex(args, Find(args, "--root"))
		}
	}

	// Parse config file removing all comments and empty lines
	cfgStr, err := exec.Command("sed", "-e", "s/#.*$//", "-e", "/^$/d", "/etc/pak.cfg").Output()
	if err != nil {
		log.Fatal(err)
	}
	cfg := strings.Split(string(cfgStr), "\n")
	//fmt.Println(cfg) //DEBUG

	// Set first line of config to variable
	packageManagerCommand := cfg[0]
	//fmt.Println(packageManagerCommand) //DEBUG

	// Parse list of commands in config line 2 and set to variable as array
	commands := strings.Split(cfg[1], ",")
	//fmt.Println(commands) //DEBUG

	// Set the root option in config line 3 to a variable
	useRoot := cfg[2]
	//fmt.Println(useRoot) //DEBUG

	// Set command to use to invoke root at config line 4 to a variable
	rootCommand := cfg[3]
	//fmt.Println(rootCommand) //DEBUG

	// Parse list of shortcuts in config and line 5 set to variable as an array
	shortcuts := strings.Split(cfg[4], ",")
	//fmt.Println(shortcuts) //DEBUG

	// Parse list of shortcuts in config line 6 and set to variable as array
	shortcutMappings := strings.Split(cfg[5], ",")
	//fmt.Println(shortcutMappings) //DEBUG

	// Check if config file allows root and set boolean to a variable
	var useRootBool bool
	if useRoot == "yes" {
		useRootBool = true
	} else if useRoot == "no" {
		useRootBool = false
	}
	//fmt.Println(useRootBool) //DEBUG

	// Create similar to slice to detect incomplete and misspelled commands
	var similarTo []string

	// Displays help message if no arguments provided or -h/--help is passed

	if len(args) == 0 || Contains(args, "-h") || Contains(args, "--help") || Contains(args, "help") {
		printHelpMessage(packageManagerCommand, useRootBool, rootCommand, commands, shortcuts)
	}

	// Checks for known commands in first argument
	// Appends percent similarity between command and all commands in array to an array
	var dist []float64
	for _,command := range commands {
		dist = append(dist, JaroWinkler(command, args[0], 1, 0))
	}

	// Appends the suspected command to an array
	for count, element := range dist {
		if element == Max(dist) {
			similarTo = append(similarTo, commands[count])
		}
	}

	// Deals with shortcuts
	for index, shortcut := range shortcuts {
		if args[0] == shortcut && !Contains(similarTo, shortcutMappings[index]) {
			similarTo = nil
			similarTo = append(similarTo, shortcutMappings[index])
		}
	}

	// Run package manager with the proper arguments passed`
	fmt.Println("Running:", strings.Title(similarTo[0]), "using", strings.Title(packageManagerCommand))
	if len(similarTo) == 1 && len(args) >= 2 {
		if useRootBool {
			cmdArr := []string{rootCommand, packageManagerCommand, similarTo[0], strings.Join(args[1:], " ")}
			cmdStr := strings.Join(cmdArr, " ")
			command := exec.Command("sh", "-c", cmdStr)
			command.Stdout = os.Stdout
			command.Stdin = os.Stdin
			command.Stderr = os.Stderr
			error := command.Run()
			if error != nil {
				fmt.Println("Error received from child process")
				log.Fatal(error)
			}
		} else {
			cmdArr :=[]string{packageManagerCommand, similarTo[0], strings.Join(args[1:], " ")}
			cmdStr := strings.Join(cmdArr, " ")
			command := exec.Command("sh", "-c", cmdStr)
			command.Stdout = os.Stdout
			command.Stdin = os.Stdin
			command.Stderr = os.Stderr
			error := command.Run()
			if error != nil {
				fmt.Println("Error received from child process")
				log.Fatal(error)
			}
		}
	} else if len(similarTo) != 1 && len(args) >= 1 && similarTo != nil {
		fmt.Println("Ambiguous:", similarTo)
	} else if similarTo == nil && len(args) >= 1 {
		fmt.Println("Command", args[0], "not known")
		os.Exit(1)
	} else {
		if useRootBool {
			cmdArr :=[]string{rootCommand, packageManagerCommand, similarTo[0]}
			cmdStr := strings.Join(cmdArr, " ")
			command := exec.Command("sh", "-c", cmdStr)
			command.Stdout = os.Stdout
			command.Stdin = os.Stdin
			command.Stderr = os.Stderr
			error := command.Run()
			if error != nil {
				fmt.Println("Error received from child process")
				log.Fatal(error)
			}
		} else {
			cmdArr :=[]string{packageManagerCommand, similarTo[0]}
			cmdStr := strings.Join(cmdArr, " ")
			command := exec.Command("sh", "-c", cmdStr)
			command.Stdout = os.Stdout
			command.Stdin = os.Stdin
			command.Stderr = os.Stderr
			error := command.Run()
			if error != nil {
				fmt.Println("Error received from child process")
				log.Fatal(error)
			}
		}
	}
}