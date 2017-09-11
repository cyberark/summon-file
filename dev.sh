#!/usr/bin/env bash

function finish {
  echo 'Removing environment'
  echo '-----'
  docker-compose down -v
}
trap finish EXIT

function main() {
  runDevelopment
}

function runDevelopment() {
  docker-compose build --pull tester

  # Run development
  docker-compose run --rm \
    --entrypoint "bash -c './convey.sh& bash'" \
    --service-ports \
    tester
}

main
