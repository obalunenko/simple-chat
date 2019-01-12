#!/usr/bin/env bash

function vet(){
    echo "vet project..."
    go vet ./... | grep -v "vendor"
    echo ""
}

function fmt(){
go fmt $(go list ./... | grep -v /vendor/)

}
function golangci(){
    echo "golang-ci lint..."
    golangci-lint run ./...
    echo ""
}

function go-lint(){
    echo "golint..."
    declare -a lints=$(golint ./... | grep -v "vendor")
    if [[ $lints ]]; then
        echo "fix it:"
        for l in "${lints[@]}"
        do
            echo "$l"
            
        done
        exit 1
        
    else
        echo "code is ok"
        echo $lints
    fi
    echo ""
}

function go-group()
{
    echo "gogroup..."
    gogroup -order std,other,prefix=github.com/oleg-balunenko/  $(find . -type f -name "*.go" | grep -v "vendor/")
    echo ""
}

vet
golangci
go-lint
go-group