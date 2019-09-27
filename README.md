# tim
Tim (the enchanter) is a go tool to modify source code.
This could also be used in compilation process with a special Go compiler.

Tim the enchanter (from Monty Python and the Holy Grail):<br>
![Tim the enchanter](https://vignette.wikia.nocookie.net/montypython/images/f/fb/Tim.jpg/revision/latest/scale-to-width-down/247?cb=20130716232411)

## Version 1 - Hello World!
> Done on 2019-09-26.
Version 1 is an experiment. The `tim` would inject a `fmt.Println("hello world")` into `main` function in any source file.

The strategy of locating the inject point:
1. any Go source file with name given in arg[1]
2. any function named `main`
3. first line in that function

Also it would ensure the package `fmt` is imported if not. Then the tool print out the modified source to stdout.

### As a command-line tool
```
./tim my_go_file.go
```

### As a Go tool
1. Copy `tim` to `$GOROOT/pkg/tools/<GOOS>_<GOARCH>/`, which actually is the value of `go env GOTOOLDIR`
2. Execute:
```
go tool tim your_go_file.go
```
