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
	"github.com/alessio/shellescape"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	flag "github.com/spf13/pflag"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

var Log = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

func main() {

	// Create help flag
	var helpFlagGiven bool
	flag.BoolVarP(&helpFlagGiven, "help", "h", false, "Show help screen")
	// Create package manager override flag
	var packageManagerOverride string
	flag.StringVarP(&packageManagerOverride, "package-manager", "p", os.Getenv("PAK_MGR_OVERRIDE"), "Override package manager wrapped by pak")
	// Parse arguments for flags
	flag.Parse()

	// Check which user is running command
	currentUser, err := user.Current()
	if err != nil {
		Log.Fatal().Err(err).Msg("Error getting current user")
	}
	// If running as root
	if strings.Contains(currentUser.Username, "root") {
		// Print warning message
		Log.Warn().Msg("Running as root may cause extraneous root invocation")
	}

	// Get arguments without flags
	args := flag.Args()

	// Define variables for config file location, and override state boolean
	var isOverridden bool

	// Read /etc/pak.toml into new Config{}
	config := NewConfig("/etc/pak.toml")

	// If override is set
	if packageManagerOverride != "" {
		// Set active package manager to override
		config.ActiveManager = packageManagerOverride
		// Set override state to true
		isOverridden = true
	} else {
		// Set override state to false
		isOverridden = false
	}

	// Parse list of commands in config line 2 and set to variable as array
	commands := config.Managers[config.ActiveManager].Commands
	//fmt.Println(commands) //DEBUG

	// Set the root option in config line 3 to a variable
	useRoot := config.Managers[config.ActiveManager].UseRoot
	//fmt.Println(useRoot) //DEBUG

	// Set command to use to invoke root at config line 4 to a variable
	rootCommand := config.RootCommand
	//fmt.Println(rootCommand) //DEBUG

	// Parse list of shortcuts in config and line 5 set to variable as an array
	shortcuts := config.Managers[config.ActiveManager].Shortcuts
	//fmt.Println(shortcuts) //DEBUG

	// Create similar to slice to put all matched commands into
	similarTo := []string{}

	// Displays help message if no arguments provided or -h/--help is passed
	if len(args) == 0 || helpFlagGiven || Contains(args, "help") {
		printHelpMessage(config.ActiveManager, useRoot, rootCommand, commands, shortcuts, isOverridden)
		os.Exit(0)
	}

	// Create distance slice to store JaroWinkler distance values
	distance := map[string]float64{}
	// Appends JaroWinkler distance between each available command and the first argument to an array
	for command := range commands {
		distance[command] = JaroWinkler(command, args[0], 1, 0)
	}

	// Deals with shortcuts
	for shortcut, mapping := range shortcuts {
		// If the first argument is a shortcut and similarTo does not already contain its mapping, append it
		if args[0] == shortcut && !Contains(similarTo, mapping) {
			similarTo = append(similarTo, mapping)
		}
	}

	// Compares each distance to the max of the distance slice and appends the closest command to similarTo
	for command, cmdDist := range distance {
		// If current element is the closest to the first argument
		if cmdDist == Max(GetValuesDist(distance)) {
			// Append command at same index as distance to similarTo
			similarTo = append(similarTo, commands[command])
		}
	}

	// If similarTo is still empty, log it fatally as something is wrong with the config or the code
	if len(similarTo) == 0 {
		Log.Fatal().Msg("This command does not match any known commands or shortcuts")
	}
	// Anonymous function to decide whether to print (overridden)
	printOverridden := func() string {
		if isOverridden {
			return "(overridden)"
		} else {
			return ""
		}
	}
	// Print text showing command being run and package manager being used
	fmt.Println("Running:", strings.Title(GetKey(commands, similarTo[0])), "using", strings.Title(config.ActiveManager), printOverridden())
	// Run package manager with the proper arguments passed if more than one argument exists
	var cmdArr []string
	// If root is to be used, append it to cmdArr
	if useRoot {
		cmdArr = append(cmdArr, rootCommand)
	}
	// If command to be run has a prefix of "cmd:"
	if strings.HasPrefix(similarTo[0], "cmd:") {
		// Append the command to the slice without the prefix
		cmdArr = append(cmdArr, strings.TrimPrefix(similarTo[0], "cmd:"))
	} else {
		// Otherwise, append all commands and arguments for the package manager to slice
		cmdArr = append(cmdArr, config.ActiveManager, similarTo[0])
	}
	// If greater than 2 arguments, append them to cmdArr
	if len(args) >= 2 {
		cmdArr = append(cmdArr, shellescape.QuoteCommand(args[1:]))
	}
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
		Log.Fatal().Err(err).Msg("Error received from child process")
	}
}
