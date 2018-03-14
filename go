#!/bin/bash

set -eu

function is_ci_release_build {
  [[ -n "${TRAVIS_TAG:-}" ]]
}

function ensure_goreleaser {
  if [[ ! -f bin/goreleaser ]]; then
    mkdir -p bin
    (
      cd bin
      wget https://github.com/goreleaser/goreleaser/releases/download/v0.62.3/goreleaser_Linux_i386.tar.gz
      echo "7033817d80c1318aebd9acd4a559ffeaa0985bd8016b108ad3aa6fd006259ce3  goreleaser_Linux_i386.tar.gz" |sha256sum -c -

      tar xf goreleaser_Linux_i386.tar.gz
      rm goreleaser_Linux_i386.tar.gz
    )
  fi
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
    ./bin/goreleaser --rm-dist
  else
    ./bin/goreleaser --rm-dist --snapshot
  fi
}

args=${1:-}
shift || true
case "$args" in
  test) task_test ;;
  release) task_release ;;
  *) task_usage ;;
esac
