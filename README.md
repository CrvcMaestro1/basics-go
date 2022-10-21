Learning GO
===========

To create a go module
```bash
$ go mod init example.com/hello
```

To import local module
```bash
$ go mod edit -replace example.com/greetings=../greetings
```

To sync module's dependencies
```bash
$ go mod tidy
```

To run module
````bash
$ go run .
````

To run test suite
```bash
$ go test
```

To run test suite with verbose
```bash
$ go test -v
```

To build
```bash
# go build <<module_name>>
# go build
$ go build hello.go
```

To discover install path
```bash
$ go list -f '{{.Target}}'
```

To set environment variables
```bash
$ go env -w GOBIN=path
```

To install
```bash
$ go install
```
