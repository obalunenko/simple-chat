#!/usr/bin/env bash

function get_dependencies() {
      declare -a packages=(
"golang.org/x/tools/cmd/cover/..."
"github.com/mattn/goveralls/..."
"github.com/Bubblyworld/gogroup/..."
"go get github.com/axw/gocov/gocov/..."
)

## now loop through the above array
for pkg in "${packages[@]}"
do
   echo "$pkg"
   go get -u -v "$pkg"
done
}




echo Gonna to update go tools and packages...
get_dependencies
echo All is done!