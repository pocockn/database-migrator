#!/bin/bash

echo "Releasing version $1"

if git rev-parse "v$1" >/dev/null 2>&1; then
    goreleaser --rm-dist
else
    git tag -a v$1 -m "Release number: $1"
    git push origin v$1
    goreleaser --rm-dist
fi