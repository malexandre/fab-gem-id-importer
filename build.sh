#!/usr/bin/bash

targets=("darwin/amd64" "darwin/arm64" "linux/amd64" "linux/arm64" "windows/amd64" "windows/arm64")

for target in "${targets[@]}"
do
    IFS=/ read -a splits <<< "$target"
    ext=""
    if [[ ${splits[0]} == "windows" ]]; then
        ext=".exe"
    elif [[ ${splits[0]} == "darwin" ]]; then
        ext=".app"
    fi
    env GOOS=${splits[0]} GOARCH=${splits[1]} go build -o dist/fab-gem-id-importer_${splits[0]}-${splits[1]}${ext}
done
