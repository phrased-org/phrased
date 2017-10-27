#!/bin/bash

set -eu

function ensure_goreleaser {
  go get github.com/goreleaser/goreleaser
}

function task_usage {
  echo "usage: $0 test | release"
  exit 1
}

function task_test {
  exit
}

function task_release {
  ensure_goreleaser
  goreleaser --rm-dist
}

args=${1:-}
shift || true
case "$args" in
  test) task_test ;;
  release) task_release ;;
  *) task_usage ;;
esac
