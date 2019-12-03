# learn-golang

➜  hello git:(master) ✗ go test -bench=. 
goos: darwin
goarch: amd64
pkg: github.com/martishevich/hello
BenchmarkMyStrToInt-8           196373287                7.00 ns/op
BenchmarkMyStrToInt2-8           2078641               616 ns/op
PASS
ok      github.com/martishevich/hello   4.270s


as we can see, strconv.Atoi() faster then fmt.Sscanf(s, "%d", &number)