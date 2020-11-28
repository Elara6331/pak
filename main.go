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
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"strings"
)

func main()  {
	// Put all arguments into a variable
	args := os.Args[1:]

	// Check which currentUser is running command
	currentUser, err := user.Current()
	if err != nil { log.Fatal(err) }

	// Check to make sure root is not being used unless -r/--root specified
	if !Contains(args, "-r") && !Contains(args, "--root") {
		if strings.Contains(currentUser.Username, "root") {
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

	// Define variables for config file location, and override state boolean
	var configFileLocation string
	var isOverridden bool
	// Get PAK_MGR_OVERRIDE environment variable
	override := os.Getenv("PAK_MGR_OVERRIDE")
	// If override is set
	if override != "" {
		// Set configFileLocation to /etc/pak.d/{override}.cfg
		configFileLocation = "/etc/pak.d/" + override + ".cfg"
		// Set override state to true
		isOverridden = true
	} else {
		// Otherwise, set configFileLocation to default config
		configFileLocation = "/etc/pak.cfg"
		// Set override state to false
		isOverridden = false
	}

	// Parse config file removing all comments and empty lines
	config, err := ioutil.ReadFile(configFileLocation)
	if err != nil { log.Fatal(err) }
	commentRegex := regexp.MustCompile(`#.*`)
	emptyLineRegex := regexp.MustCompile(`(?m)^\s*\n`)
	parsedConfig := commentRegex.ReplaceAllString(string(config), "")
	parsedConfig = emptyLineRegex.ReplaceAllString(parsedConfig, "")

	cfg := strings.Split(parsedConfig, "\n")
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

	// Create similar to slice to put all matched commands into
	var similarTo []string

	// Displays help message if no arguments provided or -h/--help is passed
	if len(args) == 0 || Contains(args, "-h") || Contains(args, "--help") || Contains(args, "help") {
		printHelpMessage(packageManagerCommand, useRootBool, rootCommand, commands, shortcuts, shortcutMappings, isOverridden)
		os.Exit(0)
	}

	// Create distance slice to store JaroWinkler distance values
	var distance []float64
	// Appends JaroWinkler distance between each available command and the first argument to an array
	for _,command := range commands {
		distance = append(distance, JaroWinkler(command, args[0], 1, 0))
	}

	// Compares each distance to the max of the distance slice and appends the closest command to similarTo
	for index, element := range distance {
		// If current element is the closest to the first argument
		if element == Max(distance) {
			// Append command at same index as distance to similarTo
			similarTo = append(similarTo, commands[index])
		}
	}

	// Deals with shortcuts
	for index, shortcut := range shortcuts {
		// If the first argument is a shortcut and similarTo does not already contain its mapping, append it
		if args[0] == shortcut && !Contains(similarTo, shortcutMappings[index]) {
			similarTo = append(similarTo, shortcutMappings[index])
		}
	}

	// If similarTo is still empty, log it fatally as something is wrong with the config or the code
	if len(similarTo) == 0 { log.Fatalln("This command does not match any known commands or shortcuts") }

	// Print text showing command being run and package manager being used
	fmt.Println("Running:", strings.Title(similarTo[0]), "using", strings.Title(packageManagerCommand))
	// Run package manager with the proper arguments passed if more than one argument exists
	var cmdArr []string
	// If root is to be used, append it to cmdArr
	if useRootBool { cmdArr = append(cmdArr, rootCommand) }
	// Create slice with all commands and arguments for the package manager
	cmdArr = append(cmdArr, []string{packageManagerCommand, similarTo[0]}...)
	// If greater than 2 arguments, append them to cmdArr
	if len(args) >= 2 { cmdArr = append(cmdArr, strings.Join(args[1:], " ")) }
	// Create space separated string from cmdArr
	cmdStr := strings.Join(cmdArr, " ")
	// Instantiate exec.Command object with command sh, flag -c, and cmdStr
	command := exec.Command("sh", "-c", cmdStr)
	// Set standard outputs for command
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	// Run command
	err = command.Run()
	// If command returned an error, log fatally with explanation
	if err != nil {
		fmt.Println("Error received from child process")
		log.Fatal(err)
	}
}