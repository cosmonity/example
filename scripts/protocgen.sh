#!/usr/bin/env bash

set -e

echo "Generating gogo proto code"
cd proto
proto_dirs=$(find . -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
      buf generate --template buf.gen.gogo.yaml $file
  done
done

cd ..

cp -r go.cosmonity.xyz/example/* ./
rm -rf go.cosmonity.xyz github.com cosmonity