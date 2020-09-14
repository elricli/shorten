## Intro

Short url genenrate and redirect.


## Key Features

- Using [xorshift](https://en.wikipedia.org/wiki/Xorshift) genenrate pseudorandom number. Benchmark output:

  ```
  goos: windows
  goarch: amd64
  pkg: github.com/drrrMikado/shorten/fastrand
  BenchmarkCryptoRandStr-2   	  333349	      3228 ns/op	     456 B/op	      33 allocs/op
  BenchmarkMathRandStr-2     	   15608	     82073 ns/op	       8 B/op	       1 allocs/op
  BenchmarkFastrandStr-2     	 4958679	       265 ns/op	       8 B/op	       1 allocs/op
  PASS
  ok  	github.com/drrrMikado/shorten/fastrand	6.121s
  ```
- Using [bloom filter](https://en.wikipedia.org/wiki/Bloom_filter) determine whether the generated string already exists
- Bloom filter hash function use MurmurHash3 of [MurmurHash](https://en.wikipedia.org/wiki/MurmurHash)

## TODO

- [ ] Redis storage
- [ ] Generating strings multiple times for a link should be the same as before
- [ ] Reverse lookup
- [ ] Visit count