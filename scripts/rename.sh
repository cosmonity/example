#!/usr/bin/env bash

# Usage MODULE_NAME=username/repo ./scripts/rename.sh

if [[ -z $MODULE_NAME ]]; then
    echo "MODULE_NAME must be set."
    exit 1
fi

IFS='/' read -ra PARTS <<< "$MODULE_NAME"
USERNAME="${PARTS[0]}"
REPO="${PARTS[-1]}"

# remove all generated proto files
find . -type f -name "*.pb.go" -delete
find . -type f -name "*.pulsar.go" -delete

# rename module and imports
go mod edit -module github.com/$MODULE_NAME
find . -not -path './.*' -type f -exec sed -i -e "s,cosmosregistry/example,$MODULE_NAME,g" {} \;
find . -name '*.proto' -type f -exec sed -i -e "s,cosmosregistry.example,$(echo "$MODULE_NAME" | tr '/' '.'),g" {} \;
find . -name 'protocgen.sh' -type f -exec sed -i -e "s,rm -rf github.com cosmosregistry,rm -rf github.com $USERNAME,g" {} \;
find . -not -path './.*' -type f -exec sed -i -e "s,example,$REPO,g" {} \;

# rename directory
mkdir -p proto/$MODULE_NAME
mv proto/cosmosregistry/example/* proto/$MODULE_NAME
rm -rf proto/cosmosregistry

# re-generate protos
make proto-gen

# credits
echo "# This project was generated using https://github.com/cosmosregistry/example" > THANKS.md

# removes itself
rm scripts/rename.sh