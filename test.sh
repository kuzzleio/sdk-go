#!/usr/bin/env bash
set -e

dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
workdir=$dir/.cover
profile="$workdir/cover.out"
mode=count
dirs=(kuzzle connection/websocket collection security ms realtime index server auth document)
timeout=${TIMEOUT:=1m}

generate_cover_data() {
    rm -rf "$workdir"
    mkdir "$workdir"

    for pkg in ${dirs[@]}; do
        go test -timeout "$timeout" -covermode="$mode" -coverprofile="$workdir/$(basename $pkg).cover" "./$pkg"
    done

    echo "mode: $mode" >"$profile"
    grep -h -v "^mode:" "$workdir"/*.cover >>"$profile"
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

cd "$dir"

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
