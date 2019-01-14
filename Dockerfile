FROM golang:1.11.4-alpine as builder

# Add Maintainer Info
LABEL maintainer="Oleg Balunenko <oleg.balunenko@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/oleg-balunenko/simple-chat

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

RUN \
       apk add --no-cache bash openssh git make

ENV \
       GO111MODULE=on \
       CGO_ENABLED=0

RUN go mod download

RUN make compile-for-docker


######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/src/github.com/oleg-balunenko/simple-chat/web .
RUN find . -name "*.go" -type f -delete

COPY --from=builder /go/src/github.com/oleg-balunenko/simple-chat/bin/simple-chat .