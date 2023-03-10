```sh
❯ go test -bench=.        
goos: linux
goarch: amd64
pkg: benchrand
cpu: 11th Gen Intel(R) Core(TM) i7-1185G7 @ 3.00GHz
BenchmarkMathRand-8     	 470460	     2374 ns/op
BenchmarkCryptoRand-8   	 168100	     7357 ns/op
PASS
ok  	benchrand	2.456s
```
