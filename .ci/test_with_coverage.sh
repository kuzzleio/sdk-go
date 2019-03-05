#!/usr/bin/env bash
set -e

cover_dir="$(pwd)/.cover"
profile="$cover_dir/cover.out"
mode=count
timeout=${TIMEOUT:=1m}

generate_cover_data() {
    rm -rf "$cover_dir"
    mkdir "$cover_dir"
    go test -timeout "$timeout" -covermode="$mode" -coverprofile="$profile" ./...
}

show_cover_report() {
    go tool cover -${1}="$profile"
}

linter_check() {
    invalid_files=$(gofmt -l "$1")

    if [ -n "${invalid_files}" ]; then
        echo "Lint errors on the following files:"
        echo ${invalid_files}
        exit 1
    fi
}

linter_check .
generate_cover_data
show_cover_report func

case "$1" in
"")
    ;;
--html)
    show_cover_report html ;;
*)
    echo >&2 "error: invalid option: $1"; exit 1 ;;
esac
