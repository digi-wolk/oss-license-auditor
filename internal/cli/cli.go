package cli

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetCliArguments() Arguments {
	path := flag.String("path", "", "Path to root of source code to scan")
	onlyRiskyLicenses := flag.Bool("only-risky-licenses", false, "Only show packages with risky licenses")
	failOnRisky := flag.Bool("fail-on-risky", false, "Fail if risky licenses are found")
	ci := flag.Bool("ci", false, "CI mode")
	verbose := flag.Bool("vvv", false, "Verbose output")
	commentOnPr := flag.Bool("comment-on-pr", false, "Comment on Pull Request")
	ciType := flag.String("ci-type", "", "CI type")

	flag.Parse()

	if strings.ToLower(os.Getenv("INPUT_PATH")) != "" && *path == "" {
		*path = os.Getenv("INPUT_PATH")
	}
	if strings.ToLower(os.Getenv("INPUT_ONLY-RISKY")) == "true" {
		*onlyRiskyLicenses = true
	}
	if strings.ToLower(os.Getenv("INPUT_FAIL-ON-RISKY-FAIL")) == "true" {
		*failOnRisky = true
	}
	if strings.ToLower(os.Getenv("CI")) == "true" {
		*ci = true
	}
	if strings.ToLower(os.Getenv("INPUT_VERBOSE")) == "true" {
		*verbose = true
	}
	if strings.ToLower(os.Getenv("INPUT_COMMENT-ON-PR")) == "true" {
		*commentOnPr = true
	}
	if strings.ToLower(os.Getenv("GITHUB_ACTIONS")) == "true" {
		*ciType = "github"
	}

	pathValue := *path
	// Remove extra / from the end of the path if it exists
	if (pathValue)[len(pathValue)-1:] == "/" {
		pathValue = (pathValue)[:len(pathValue)-1]
	}

	if *verbose {
		log.Println("Path: " + pathValue)
		log.Println("CiType: " + *ciType)
		log.Println("Ci: " + strconv.FormatBool(*ci))
		log.Println("OnlyRiskyLicenses: " + strconv.FormatBool(*onlyRiskyLicenses))
		log.Println("FailOnRisky: " + strconv.FormatBool(*failOnRisky))
		log.Println("Verbose: " + strconv.FormatBool(*verbose))
		log.Println("CommentOnPr: " + strconv.FormatBool(*commentOnPr))
	}

	return Arguments{
		Path:              pathValue,
		OnlyRiskyLicenses: *onlyRiskyLicenses,
		FailOnRisky:       *failOnRisky,
		Ci:                *ci,
		CiType:            *ciType,
		Verbose:           *verbose,
		CommentOnPr:       *commentOnPr,
	}
}
