#!/usr/bin/env bash

set -e

echo "Generating proto code"
cd proto
buf generate --template buf.gen.yaml

cd ..

rm -rf api && mkdir api
mv cosmosregistry/example/* ./api
rm -rf github.com cosmosregistry