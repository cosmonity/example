#!/usr/bin/env bash

if [[ -z $MODULE_NAME ]]; then
    echo "MODULE_NAME must be set."
    exit 1
fi

# remove all generated proto files
find . -type f -name "*.pb.go" -delete
find . -type f -name "*.pulsar.go" -delete

# rename module and imports
go mod edit -module github.com/$MODULE_NAME
find . -not -path './.*' -type f -exec sed -i -e "s,cosmosregistry/example,$MODULE_NAME,g" {} \;
find . -name '*.proto' -type f -exec sed -i -e "s,cosmosregistry.example,$(echo "$MODULE_NAME" | tr '/' '.'),g" {} \;

# rename directory
mkdir -p proto/$MODULE_NAME
mv proto/cosmosregistry/example/* proto/$MODULE_NAME
rm -rf proto/cosmosregistry

# re-generate protos
make proto-gen

# removes itself
# rm scripts/rename.sh