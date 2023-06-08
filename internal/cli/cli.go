package cli

import (
	"flag"
	"fmt"
	"os"
)

func GetCliArguments() Arguments {
	path := flag.String("path", "", "Path to root of source code to scan")
	onlyRiskyLicenses := flag.Bool("only-risky-licenses", false, "Only show packages with risky licenses")
	failOnRisky := flag.Bool("fail-on-risky", false, "Fail if risky licenses are found")
	ci := flag.Bool("ci", false, "CI mode")
	verbose := flag.Bool("vvv", false, "Verbose output")
	commentOnGithubPr := flag.Bool("comment-on-github-pr", false, "Comment on GitHub Pull Request")

	flag.Parse()

	fmt.Println("####################")
	// Print variables
	fmt.Println("path:", *path)
	fmt.Println("onlyRiskyLicenses:", *onlyRiskyLicenses)
	fmt.Println("failOnRisky:", *failOnRisky)
	fmt.Println("ci:", *ci)
	fmt.Println("verbose:", *verbose)
	fmt.Println("commentOnGithubPr:", *commentOnGithubPr)

	fmt.Println("####################")
	// Print environment variables
	fmt.Println("INPUT_PATH:", os.Getenv("INPUT_PATH"))
	fmt.Println("INPUT_COMMENT-ON-GITHUB-PR:" + os.Getenv("INPUT_COMMENT-ON-GITHUB-PR"))
	fmt.Println("INPUT_ONLY-RISKY:" + os.Getenv("INPUT_ONLY-RISKY"))
	fmt.Println("INPUT_FAIL-ON-RISKY-FAIL:" + os.Getenv("INPUT_FAIL-ON-RISKY-FAIL"))
	fmt.Println("HOME:" + os.Getenv("HOME"))
	fmt.Println("GITHUB_JOB:" + os.Getenv("GITHUB_JOB"))
	fmt.Println("GITHUB_REF:" + os.Getenv("GITHUB_REF"))
	fmt.Println("GITHUB_SHA:" + os.Getenv("GITHUB_SHA"))
	fmt.Println("GITHUB_REPOSITORY:" + os.Getenv("GITHUB_REPOSITORY"))
	fmt.Println("GITHUB_REPOSITORY_OWNER:" + os.Getenv("GITHUB_REPOSITORY_OWNER"))
	fmt.Println("GITHUB_REPOSITORY_OWNER_ID:" + os.Getenv("GITHUB_REPOSITORY_OWNER_ID"))
	fmt.Println("GITHUB_RUN_ID:" + os.Getenv("GITHUB_RUN_ID"))
	fmt.Println("GITHUB_RUN_NUMBER:" + os.Getenv("GITHUB_RUN_NUMBER"))
	fmt.Println("GITHUB_RETENTION_DAYS:" + os.Getenv("GITHUB_RETENTION_DAYS"))
	fmt.Println("GITHUB_RUN_ATTEMPT:" + os.Getenv("GITHUB_RUN_ATTEMPT"))
	fmt.Println("GITHUB_REPOSITORY_ID:" + os.Getenv("GITHUB_REPOSITORY_ID"))
	fmt.Println("GITHUB_ACTOR_ID:" + os.Getenv("GITHUB_ACTOR_ID"))
	fmt.Println("GITHUB_ACTOR:" + os.Getenv("GITHUB_ACTOR"))
	fmt.Println("GITHUB_TRIGGERING_ACTOR:" + os.Getenv("GITHUB_TRIGGERING_ACTOR"))
	fmt.Println("GITHUB_WORKFLOW:" + os.Getenv("GITHUB_WORKFLOW"))
	fmt.Println("GITHUB_HEAD_REF:" + os.Getenv("GITHUB_HEAD_REF"))
	fmt.Println("GITHUB_BASE_REF:" + os.Getenv("GITHUB_BASE_REF"))
	fmt.Println("GITHUB_EVENT_NAME:" + os.Getenv("GITHUB_EVENT_NAME"))
	fmt.Println("GITHUB_SERVER_URL:" + os.Getenv("GITHUB_SERVER_URL"))
	fmt.Println("GITHUB_API_URL:" + os.Getenv("GITHUB_API_URL"))
	fmt.Println("GITHUB_GRAPHQL_URL:" + os.Getenv("GITHUB_GRAPHQL_URL"))
	fmt.Println("GITHUB_REF_NAME:" + os.Getenv("GITHUB_REF_NAME"))
	fmt.Println("GITHUB_REF_PROTECTED:" + os.Getenv("GITHUB_REF_PROTECTED"))
	fmt.Println("GITHUB_REF_TYPE:" + os.Getenv("GITHUB_REF_TYPE"))
	fmt.Println("GITHUB_WORKFLOW_REF:" + os.Getenv("GITHUB_WORKFLOW_REF"))
	fmt.Println("GITHUB_WORKFLOW_SHA:" + os.Getenv("GITHUB_WORKFLOW_SHA"))
	fmt.Println("GITHUB_WORKSPACE:" + os.Getenv("GITHUB_WORKSPACE"))
	fmt.Println("GITHUB_ACTION:" + os.Getenv("GITHUB_ACTION"))
	fmt.Println("GITHUB_EVENT_PATH:" + os.Getenv("GITHUB_EVENT_PATH"))
	fmt.Println("GITHUB_ACTION_REPOSITORY:" + os.Getenv("GITHUB_ACTION_REPOSITORY"))
	fmt.Println("GITHUB_ACTION_REF:" + os.Getenv("GITHUB_ACTION_REF"))
	fmt.Println("GITHUB_PATH:" + os.Getenv("GITHUB_PATH"))
	fmt.Println("GITHUB_ENV:" + os.Getenv("GITHUB_ENV"))
	fmt.Println("GITHUB_STEP_SUMMARY:" + os.Getenv("GITHUB_STEP_SUMMARY"))
	fmt.Println("GITHUB_STATE:" + os.Getenv("GITHUB_STATE"))
	fmt.Println("GITHUB_OUTPUT:" + os.Getenv("GITHUB_OUTPUT"))
	fmt.Println("RUNNER_OS:" + os.Getenv("RUNNER_OS"))
	fmt.Println("RUNNER_ARCH:" + os.Getenv("RUNNER_ARCH"))
	fmt.Println("RUNNER_NAME:" + os.Getenv("RUNNER_NAME"))
	fmt.Println("RUNNER_TOOL_CACHE:" + os.Getenv("RUNNER_TOOL_CACHE"))
	fmt.Println("RUNNER_TEMP:" + os.Getenv("RUNNER_TEMP"))
	fmt.Println("RUNNER_WORKSPACE:" + os.Getenv("RUNNER_WORKSPACE"))
	fmt.Println("ACTIONS_RUNTIME_URL:" + os.Getenv("ACTIONS_RUNTIME_URL"))
	fmt.Println("ACTIONS_CACHE_URL:" + os.Getenv("ACTIONS_CACHE_URL"))
	fmt.Println("GITHUB_ACTIONS:" + os.Getenv("GITHUB_ACTIONS"))
	fmt.Println("CI:" + os.Getenv("CI"))
	fmt.Println("####################")

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
