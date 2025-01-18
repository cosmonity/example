#!/usr/bin/env bash

# This script is used to rename the example repo to the module name of your choice.
# Usage MODULE_NAME=username/repo ./scripts/rename.sh
# Disclaimer: this script do not work with multi-slash module names (e.g. my/awesome/repo) or module names with dashes (e.g. my-awesome-repo)
# If you want to use a multi-slash module name, you will need to manually rename the module name in go.mod and the proto files.
# A typical module name is 'username/proto' omitting the 'github.com/' prefix.

if [[ -z $MODULE_NAME ]]; then
    echo "MODULE_NAME must be set."
    exit 1
fi

IFS='/' read -ra PARTS <<< "$MODULE_NAME"
USERNAME="${PARTS[0]}"
REPO="${PARTS[1]}"

# remove all generated proto files
find . -type f -name "*.pb.go" -delete
find . -type f -name "*.pb.gw.go" -delete
find . -type f -name "*.pulsar.go" -delete

# rename module and imports
go mod edit -module github.com/$MODULE_NAME
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    find . -not -path './.*' -type f -exec sed -i -e "s,cosmonity/example,$MODULE_NAME,g" {} \;
    find . -name '*.proto' -type f -exec sed -i -e "s,cosmonity.example,$(echo "$MODULE_NAME" | tr '/' '.'),g" {} \;
    find . -name 'protocgen.sh' -type f -exec sed -i -e "s,rm -rf github.com cosmonity,rm -rf github.com $USERNAME,g" {} \;
    find . -not -path './.*' -type f -exec sed -i -e "s,example,$REPO,g" {} \;
else
    find . -not -path './.*' -type f -exec sed -i '' -e "s,cosmonity/example,$MODULE_NAME,g" {} \;
    find . -name '*.proto' -type f -exec sed -i '' -e "s,cosmonity.example,$(echo "$MODULE_NAME" | tr '/' '.'),g" {} \;
    find . -name 'protocgen.sh' -type f -exec sed -i '' -e "s,rm -rf github.com cosmonity,rm -rf github.com $USERNAME,g" {} \;
    find . -not -path './.*' -type f -exec sed -i '' -e "s,example,$REPO,g" {} \;
fi

# rename directory
mkdir -p proto/$MODULE_NAME
mv proto/cosmonity/example/* proto/$MODULE_NAME
rm -rf proto/cosmonity

# re-generate protos
make proto-gen

# credits
echo "# This Cosmos SDK module was generated using <https://github.com/cosmonity/example>" > THANKS.md

# removes itself
rm scripts/rename.sh
