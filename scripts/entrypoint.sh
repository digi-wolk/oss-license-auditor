#!/usr/bin/env bash

# Expand the arguments into an array of strings. This is required because the GitHub action
# provides all arguments concatenated as a single string.
ARGS=("$@")

# Using -n instead of ! -z: https://github.com/koalaman/shellcheck/wiki/SC2236
if [[ -n "${GITHUB_AUTHENTICATION_TOKEN}" ]]; then
  git config --global --add url."https://x-access-token:${GITHUB_AUTHENTICATION_TOKEN}@github.com/".insteadOf "https://github.com/"
fi

/bin/gosec "${ARGS[*]}"