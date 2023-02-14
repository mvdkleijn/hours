# Hours

To whoever reads this: the code here is quick-n-dirty.
It was written for my personal use and published here
because some people expressed their interest in it.

## What it does

It asks a couple of questions and then calculates and
spits out some info like how many workable hours in a year,
how many to meet your target and how much potential extra
earnings you can have.

The logic is fairly generic so I thought I'd publish it for
others to use if they can.

Disclaimer: YMMV, No guarantees, rough code

## Build

```
GOOS=windows GOARCH=amd64 go build -o bin/hours.exe main.go
GOOS=darwin GOARCH=amd64 go build -o bin/hours-mac-x86 main.go
GOOS=darwin GOARCH=arm64 go build -o bin/hours-mac-arm main.go
GOOS=linux GOARCH=amd64 go build -o bin/hours-linux main.go
```
