package cli

import "flag"

func GetCliArguments() Arguments {
	path := flag.String("path", "", "Path to root of source code to scan")
	onlyRiskyLicenses := flag.Bool("only-risky-licenses", false, "Only show packages with risky licenses")
	failOnRisky := flag.Bool("fail-on-risky", false, "Fail if risky licenses are found")
	ci := flag.Bool("ci", false, "CI mode")
	verbose := flag.Bool("vvv", false, "Verbose output")
	commentOnGithubPr := flag.Bool("comment-on-github-pr", false, "Comment on GitHub Pull Request")
	flag.Parse()

	pathValue := *path
	// Remove extra / from the end of the path if it exists
	if (pathValue)[len(pathValue)-1:] == "/" {
		pathValue = (pathValue)[:len(pathValue)-1]
	}

	return Arguments{
		Path:              pathValue,
		OnlyRiskyLicenses: *onlyRiskyLicenses,
		FailOnRisky:       *failOnRisky,
		Ci:                *ci,
		Verbose:           *verbose,
		CommentOnGithubPr: *commentOnGithubPr,
	}
}
