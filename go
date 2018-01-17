#!/bin/bash

set -eu

function is_ci_release_build {
  [[ -n "${TRAVIS_TAG:-}" ]]
}

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
  set -x

  if is_ci_release_build;
  then
    goreleaser --rm-dist
  else
    goreleaser --rm-dist --snapshot
  fi
}

args=${1:-}
shift || true
case "$args" in
  test) task_test ;;
  release) task_release ;;
  *) task_usage ;;
esac
