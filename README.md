# League Backend Challenge

## Goal
Solution contains a basic web server application written in GoLang that will process given `curl` in request csv file.  

It accepts a the following requests at endpoints:

| endpoint   |      description      | 
|----------|:-------------:|
| /echo |  Returns the matrix as a string in matrix format |
| /sum |  Return the sum of the integers in the matrix |
| /multiply |  Return the product of the integers in the matrix |
| /flatten |  Return the matrix as a 1 line string, with values separated by commas |
| /invert |  Return the matrix as a string in matrix format where the columns and rows are inverted (transpose) |  


The input file to these functions is a matrix, of any dimension where the number of rows are equal to the number of columns (square). Each value is an integer, and there is no header row. matrix.csv is example valid input.  
Server accepts only `HTTP POST` method.

## Testing 
File should have at least `read` premissions for `curl` for sending

Run web server locally (Linux):
``` golang
$ go run ./api
```
or as a binary app:
``` golang
$ ./build/api
```

Send request (example csv is in `examples` dir):
``` bash
$ curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
$ curl -F 'file=@/path/matrix.csv' "localhost:8080/sum"
$ curl -F 'file=@/path/matrix.csv' "localhost:8080/multiply"
$ curl -F 'file=@/path/matrix.csv' "localhost:8080/flatten"
$ curl -F 'file=@/path/matrix.csv' "localhost:8080/invert"
```

## Assumptions and limitations

- Unit tests are only checking correctness of `utils.go` (no performance, benchmark or e2e)
- Server handles only `http POST` method, every other will be dropped with error on user side (`http 405`)
- Input csv file (matrix) is very big, as some endpoints were not changed (i.e. given `/echo`) and are loading whole csv at once (this could be potential bottleneck)
- csv contains integer variables (if not error will be thrown)
- http server listens by default at port `:8080` on `localhost` 
- no `big int` values are included in matrix (in case so - overflow error will be thrown)