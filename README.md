# tim
Tim (the enchanter) is a go tool to modify source code

## Version 1 - Hello World!
Done on 2019-09-26.

### As a command-line tool
```
./tim my_go_file.go
```

It would look for the function named `main` and inject `fmt.Println("hello world")` in the first line in the main function.

Also it would ensure the package `fmt` is imported if not.

### As a Go tool
1. Copy `tim` to `$GOROOT/pkg/tools/<GOOS>_<GOARCH>/`, which actually is the value of `go env GOTOOLDIR`
2. Execute:
```
go tool tim your_go_file.go
```

This could also be used in compilation process with a special Go compiler.
