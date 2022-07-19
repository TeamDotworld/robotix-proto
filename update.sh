#!/bin/bash
version=3
old_version=$version

echo "Generating proto code"
./generate.sh

echo "Adding to git..."
git add .
echo "Give a message: "
read message
git commit -m "$message"
version=$((version+1))
echo "Creating a tag - v1.0.$version"
git tag "v1.0.$version"
echo "Pushing code to github"
git push origin v2
git push --tags
echo "Updating go repository"



cd /home/anand/go/src/github.com/anand-dotworld/robotix-agent
go get -u github.com/TeamDotworld/robotix-proto@v1.0.$version
go mod tidy
cd /home/anand/go/src/github.com/anand-dotworld/rcc-multi-tenant
go get -u github.com/TeamDotworld/robotix-proto@v1.0.$version
go mod tidy


sed -i "s/version=$old_version/version=$version/g" /home/anand/go/src/github.com/TeamDotworld/robotix-proto/test.sh
echo "Done!"
# git tag 

