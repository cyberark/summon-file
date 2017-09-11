#!/bin/bash -e

function finish {
  echo 'Removing environment'
  echo '-----'
  docker-compose down -v
}
trap finish EXIT

function main() {
  runTests
}

function runTests() {
  docker-compose build --pull tester

  # Execute tests
  docker-compose run --rm \
    tester
}

main
