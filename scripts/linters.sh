#!/usr/bin/env bash

function vet(){
    echo "vet project..."
    go vet $(go list ./... | grep -v /vendor/)
    echo ""
}

function fmt(){
    go fmt $(go list ./... | grep -v /vendor/)
    
}
function golangci(){
    echo "golang-ci linter running..."
    if [ -f "$(go env GOPATH)/bin/golangci-lint" ] || [ -f "/usr/local/bin/golangci-lint" ]; then
        golangci-lint run ./...
    else
        printf "Cannot check golang-ci, please run:
        curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin \n"
        exit 1
    fi
    echo ""
}

function go-lint(){
    echo "golint linter running..."
    if [ -f "$(go env GOPATH)/bin/golint" ]; then
        declare -a lints=$(golint $(go list ./... | grep -v /vendor/))
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
    else
        printf "Cannot check golint, please run:
        go get -u -v golang.org/x/lint/golint \n"
        exit 1
    fi
    
    echo ""
}

function go-group()
{
    echo "gogroup imports check is running..."
    if [ -f "$(go env GOPATH)/bin/gogroup" ]; then
        
        declare -a lints=$(gogroup -order std,other,prefix=github.com/oleg-balunenko/ $(find . -type f -name "*.go" | grep -v "vendor/"))
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
        
    else
        printf "!!!! Cannot check import order, please run:
        go get -u -v github.com/Bubblyworld/gogroup\n"
        exit 1
    fi
    echo ""
}

vet
golangci
go-lint
go-group
