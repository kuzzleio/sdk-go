#!/usr/bin/env bash
set -e

workdir=$(pwd)/.cover
profile="$workdir/cover.out"
mode=count
dirs=(kuzzle connection collection security ms)

generate_cover_data() {
    rm -rf "$workdir"
    mkdir "$workdir"

    for pkg in ${dirs[@]}; do
        go test -covermode="$mode" -coverprofile="$workdir/$(basename $pkg).cover" "./$pkg"
    done

    echo "mode: $mode" >"$profile"
    grep -h -v "^mode:" "$workdir"/*.cover >>"$profile"
}

show_cover_report() {
    go tool cover -${1}="$profile"
}

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
