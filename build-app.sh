#!/bin/bash

program_name="todo_app"

platforms=("windows" "darwin" "linux")

for os in "${platforms[@]}"; do
  for arch in "amd64"; do
    export GOOS="$os"
    export GOARCH="$arch"

    output_name="${program_name}_${os}_${arch}"

    go build -o "$output_name" .  # Gantilah "main.go" dengan nama berkas sumber Anda

    unset GOOS
    unset GOARCH

    echo "Berhasil membangun untuk $os $arch"
  done
done
