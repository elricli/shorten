#!/usr/bin/env -S bash -e

usage() {
  cat <<EOUSAGE
Usage: $0 [up|down|force|version] {#}"
EOUSAGE
}

# Redirect stderr to stdout because migrate outputs to stderr, and we want
# to be able to use ordinary output redirection.
case "$1" in
  up|down|force|version)
    migrate \
      -source file:migrations \
      -database "postgres://postgres@localhost:5432/shorten-db?sslmode=disable" \
      "$@" 2>&1
    ;;
  *)
    usage
    exit 1
    ;;
esac
