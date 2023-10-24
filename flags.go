package main

import "flag"

// Flags struct to store parsed command-line flag values
type Flags struct {
	FileName    string
	NewFileName string
	StrictMode  bool
}

// ParseFlags parses command-line flags and returns a Flags struct
func ParseFlags() Flags {
	var flags Flags

	// Define command-line flags
	flag.StringVar(&flags.FileName, "file", "", "Name of the file to parse and rebuild")
	flag.StringVar(&flags.NewFileName, "new-file", "", "Name of the new file to create after parsing")
	flag.BoolVar(&flags.StrictMode, "strict", false, "Enable strict mode")

	// Parse the command-line flags
	flag.Parse()

	return flags
}
