#!/usr/bin/env bash
set -euf -o pipefail

rm -rf build
mkdir -p build

for platform in $(go tool dist list)
do
  os="$(echo -n $platform | cut -d'/' -f1)"
  arch="$(echo -n $platform | cut -d'/' -f2)"

  if [[ "$os" == "ios" || "$os" == "android" || "$os" == "js" ]]; then
    continue
  fi

  GOOS=$os GOARCH=$arch go build -o "build/choirpack-$os-$arch" .
  echo "Built successfully for $os/$arch"
done
