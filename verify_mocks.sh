#!/bin/bash

# Verify mocks in CI

# Check if mockery is installed
if ! command -v mockery &> /dev/null
then
  echo "Installing mockery"
  go install github.com/vektra/mockery/v2@v2.17.0
fi

# Verify if mocks are up to date
GENERATED_MOCK=$(mockery --all --print --quiet --with-expecter)
EXISTING_MOCK=$(cat ./mocks/*)

#diff <(echo "$GENERATED_MOCK") <(echo "$EXISTING_MOCK")

if [[ "$GENERATED_MOCK" != "$EXISTING_MOCK" ]]
then
  echo "Mocks are not up to date"
  exit 1
else
  echo "Mocks are up to date"
fi
