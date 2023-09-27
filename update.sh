#!/bin/bash

# This script will do the following.
# - Generate proto files
# - Add to git, commit, update version, create tag, push both tag and commit.
# - Also updates the proto dependency to latest in your local repo both agent and rcc server. (Change the path to your local path)

version=52
old_version=$version

echo -e '\033[4;32;1mGenerating Proto Code\033[m'
./generate.sh

echo -e '\033[4;32;1mAdding To Git\033[m'
git add .
echo -e '\033[4;32;1mGive a message: \033[m'
read message
git commit -m "$message"
version=$((version+1))
echo "Creating a tag - v1.0.$version"
git tag "v1.0.$version"
echo -e '\033[4;32;1mPushing Code To Github\033[m'
git push origin develop
git push --tags
echo -e '\033[4;32;1mUpdating Go Repository\033[m'

cd /home/anand/go/src/github.com/anand-dotworld/robotix-agent
GOPRIVATE=github.com/TeamDotworld/robotix-proto go get github.com/TeamDotworld/robotix-proto@v1.0.$version
go mod tidy
cd /home/anand/go/src/github.com/anand-dotworld/rcc-multi-tenant
GOPRIVATE=github.com/TeamDotworld/robotix-proto go get github.com/TeamDotworld/robotix-proto@v1.0.$version
go mod tidy

sed -i "s/version=$old_version/version=$version/g" /home/anand/go/src/github.com/TeamDotworld/robotix-proto/update.sh
echo -e '\033[4;32;1mDone!\033[m'

