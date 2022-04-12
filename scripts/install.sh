#!/bin/bash

if ! command -v dlv 2>/dev/null; then
    go get github.com/go-delve/delve/cmd/dlv
fi

# if ! foo -v 2>/dev/null; then
#     go get -v github.com/foo/bar/...
#     go install github.com/foo/bar/baz
# fi

