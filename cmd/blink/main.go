package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/sapirrior/blink/internal/browser"
	"github.com/sapirrior/blink/internal/resolver"
)

// Version is injected during compilation
var Version = "0.2.0"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: blink [flags] <target> [query...]\n\n")
		fmt.Fprintf(os.Stderr, "Blink is a fast, cross-platform CLI tool to open URLs, files, folders, and search shortcuts.\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}

	versionFlag := flag.Bool("version", false, "Print the version of blink and exit")
	dryRunFlag := flag.Bool("dry-run", false, "Print the resolved URL/URI instead of opening it")
	listAliasesFlag := flag.Bool("list-aliases", false, "List all available search aliases")

	flag.Parse()

	if *versionFlag {
		fmt.Printf("blink version %s\n", Version)
		return
	}

	if *listAliasesFlag {
		fmt.Println("Available aliases:")
		for _, a := range resolver.BuiltInAliases {
			fmt.Printf("  %-16s -> %s\n", strings.Join(a.Keys, ", "), a.BaseURL)
		}
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	resolved, err := resolver.Resolve(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "blink: %v\n", err)
		os.Exit(1)
	}

	if *dryRunFlag {
		fmt.Println(resolved)
		return
	}

	if err := browser.Open(resolved); err != nil {
		fmt.Fprintf(os.Stderr, "blink: failed to open: %v\n", err)
		os.Exit(1)
	}
}
