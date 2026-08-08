package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/zavocc/youtube-watcher-cli/internal/dataapi"
	"github.com/zavocc/youtube-watcher-cli/internal/dataapi/structs"
	"github.com/zavocc/youtube-watcher-cli/internal/shared"
)

func showHelpSearch() {
	helpString := "Search YouTube videos by query." +
		"\n\nUsage: " + os.Args[0] + " search [options] [search query]\n" +
		"\nSearch options:\n" +
		" --filter              " + helpSearchFilterString + "\n" +
		" --max-results         " + helpMaxResultsString + "\n" +
		" --next-page-token     " + helpNextPageTokenString + "\n" +
		" --compact             " + helpCompactString + "\n" +
		" query                	Search query [REQUIRED]" +
		"\n\n" +
		"Supplemental options:\n" +
		" --help     " + helpShowHelpString

	fmt.Println(helpString)
}

func runSearch(ctx context.Context, args []string) {
	flagSet := flag.NewFlagSet("search", flag.ExitOnError)
	flagSet.Usage = showHelpSearch

	// args
	filter := flagSet.String("filter", "mixed", helpSearchFilterString)
	maxResults := flagSet.Int64("max-results", 10, helpMaxResultsString)
	nextPageToken := flagSet.String("next-page-token", "", helpNextPageTokenString)
	compact := flagSet.Bool("compact", false, helpCompactString)
	showHelp := flagSet.Bool("help", false, helpShowHelpString)

	flagSet.Parse(args)

	if *showHelp {
		showHelpSearch()
		os.Exit(1)
	}

	// Prefer the positional query, then fall back to redirected stdin.
	query, err := shared.ReadTextInput(flagSet.Args(), os.Stdin)
	if err != nil {
		if errors.Is(err, shared.ErrNoTextInput) {
			fmt.Fprintln(os.Stderr, "A search query is required as an argument or through stdin")
		} else {
			fmt.Fprintln(os.Stderr, "Unable to read the search query:", err)
		}
		showHelpSearch()
		os.Exit(1)
	}

	service, err := newYouTubeService(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	searchResponse, err := dataapi.Search(service, query, *filter, *maxResults, *nextPageToken)
	if err != nil {
		fmt.Fprintln(os.Stderr, "An error has occurred while searching YouTube:", err)
		os.Exit(1)
	}

	// Serialize and print the search results as JSON, check  if we need to compact
	if *compact {
		if err := structs.GenerateCompactJsonSearch(os.Stdout, searchResponse); err != nil {
			fmt.Fprintln(os.Stderr, "An error has occurred while serializing the search results:", err)
			os.Exit(1)
		}
	} else {
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(searchResponse); err != nil {
			fmt.Fprintln(os.Stderr, "An error has occurred while serializing the search results:", err)
			os.Exit(1)
		}
	}

	os.Exit(0)
}
