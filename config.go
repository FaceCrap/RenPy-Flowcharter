package main

import (
	"log"
	"os"
)

func getDefaultConfig() {
	_, err := os.ReadFile("Flowcharter.config")
	if err != nil {
		log.Println("Creating default config...")
		b := []byte(defaultConfig())
		err := os.WriteFile("Flowcharter.config", b, 0o644)
		if err != nil {
			log.Fatal("Couldn't create default config file. Please create a Flowcharter.config file manually.")
		}
	}
}

func defaultConfig() string {
	return `### RENPY-FLOWCHART TOOL CONFIGURATION ###
# Select what you want to show and what you want to hide
# Hiding everything isn't always very useful
# Showing everything can make the graph pretty ugly

# Just change true/false to false/true and do not touch anything else
# If there is any problem, or if you want to restore defaults
# just delete this file and restart the program

[config]

### You can edit ↓ below ↓
# Show the nodes with no neighbors or not?
atoms = false # default: false

# Try to display menu choices on the edges
edges = true # default: true

# Open the .png created directly or not?
open = true # default: true

# Shows screens or not (might be relevant to your project)
screens = true # default: true

# Shows nested screens (keyword 'use' inside screens) or not?
nested-screens = true # default: true

# Output something to the stdout, or not?
silent = false # default: false

# Debug mode
debug = false # default: false

# Regex for skipping files
skip-files = "" # default: ""

### You can edit ↑ above ↑

### HOW TO USE THE PROGRAM? ###
# If you know how to use the command line, just type 'Flowchart'.
# You can get help with 'Flowchart -h'
# If you don't, just place the .exe on your game folder and launch the executable

### Advanced usage ###
# If This customization isn't enough and you want more precise control?
# Use TAGS
# Example:
# label my_awesome_game_chapter_1: # Flowcharter: TITLE
# This will make a pretty bubble for this label!
# You can also override the file for 1 command by using CLI flags
# Check Flowchart -h for more information
`
}
