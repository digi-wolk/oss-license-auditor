#!/usr/bin/env bash
# Script to run code coverage tests with configurable package exclusions

set -e

# Default configuration file path
CONFIG_FILE=".codecov.yml"
OUTPUT_FILE="coverage.out"

# Parse command line arguments
while [[ $# -gt 0 ]]; do
  case $1 in
    --config)
      CONFIG_FILE="$2"
      shift 2
      ;;
    --output)
      OUTPUT_FILE="$2"
      shift 2
      ;;
    *)
      echo "Unknown option: $1"
      exit 1
      ;;
  esac
done

# Check if configuration file exists
if [[ ! -f "$CONFIG_FILE" ]]; then
  echo "Configuration file not found: $CONFIG_FILE"
  echo "Running coverage on all packages..."
  go test -coverprofile="$OUTPUT_FILE" ./...
  exit 0
fi

# Extract excluded packages from the configuration file
# This uses yq, a YAML processor. If not available, it falls back to grep and sed
if command -v yq &> /dev/null; then
  EXCLUDED_PACKAGES=$(yq eval '.exclude_packages[]' "$CONFIG_FILE" 2>/dev/null || echo "")
else
  # Fallback to grep and sed if yq is not available
  EXCLUDED_PACKAGES=$(grep -A 100 "exclude_packages:" "$CONFIG_FILE" | grep -v "exclude_packages:" | grep "^\s*-" | sed 's/^\s*-\s*//' | tr -d ' ')
fi

# Remove any leading dashes from package names
EXCLUDED_PACKAGES=$(echo "$EXCLUDED_PACKAGES" | sed 's/^-//')

# Debug: Show the excluded packages
echo "Packages to exclude:"
for pkg in $EXCLUDED_PACKAGES; do
  echo "  - $pkg"
done

# Build the grep pattern for excluding packages
if [[ -n "$EXCLUDED_PACKAGES" ]]; then
  # Create a comma-separated list of packages to exclude for use with -skip flag
  SKIP_LIST=""
  for pkg in $EXCLUDED_PACKAGES; do
    if [[ -n "$SKIP_LIST" ]]; then
      SKIP_LIST="$SKIP_LIST,$pkg"
    else
      SKIP_LIST="$pkg"
    fi
  done

  echo "Excluding packages: $SKIP_LIST"

  # Use go list to get all packages except those in the skip list
  PACKAGES_TO_TEST=$(go list ./... | grep -v -E "$(echo $SKIP_LIST | sed 's/,/|/g')")

  # Run tests on the filtered package list
  go test -coverprofile="$OUTPUT_FILE" $PACKAGES_TO_TEST
else
  echo "No packages to exclude. Running coverage on all packages..."
  go test -coverprofile="$OUTPUT_FILE" ./...
fi

echo "Coverage profile written to $OUTPUT_FILE"
