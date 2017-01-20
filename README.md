# go-echo

HTTP application which writes the request body to the response

## Server
Install Go from https://golang.org/doc/install

```
set GOPATH=C:\Work
mkdir C:\Work\src\github.com\mikeharder
cd C:\Work\src\github.com\mikeharder
git clone https://github.com/mikeharder/go-echo
cd go-echo
build.cmd
run.cmd
```

## Client
```
curl -v -H "Content-Type: application/json" -d @payload.json http://server:port
```
