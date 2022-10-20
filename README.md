# ferbench
Simple, reliable and multiplatform commandline cpu benchmark

These are here, so I don't forget how to cross compile

env GOOS=windows GOARCH=amd64 go build -ldflags="-s -w"

env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"