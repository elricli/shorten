## Intro

Fast generate short url and redirection.

## Deployment

Set the following environment variables:

   - `SHORTEN_REDIS_ADDR` (default: localhost:6379)
   - `SHORTEN_REDIS_PASSWORD` (default: '')
   - `SHORTEN_DOMAIN` (default: localhost)
   - `SHORTEN_BLOOM_FILTER_EXPECTED_INSERTIONS` (default: 1e7)
   - `SHORTEN_BLOOM_FILTER_FPP` (default: 0.00001)
   - `SHORTEN_BLOOM_FILTER_HASH_SEED` (default: 0x1)

   See `internal/config/config.go` for details.

## Key Features

- Using [xorshift](https://en.wikipedia.org/wiki/Xorshift) fast generate pseudorandom number. Benchmark output:

  ```
  goos: windows
  goarch: amd64
  pkg: github.com/drrrMikado/shorten/fastrand
  BenchmarkCryptorandStr-12    	  500660	      2260 ns/op	     456 B/op	      33 allocs/op
  BenchmarkMathrandStr-12      	   20852	     57259 ns/op	       8 B/op	       1 allocs/op
  BenchmarkFastrandStr-12      	 6433384	       197 ns/op	       8 B/op	       1 allocs/op
  PASS
  ok  	github.com/drrrMikado/shorten/fastrand	4.538s
  ```
- Using [bloom filter](https://en.wikipedia.org/wiki/Bloom_filter) fast determine whether the generated string already exists
- Bloom filter hash function use MurmurHash3 of [MurmurHash](https://en.wikipedia.org/wiki/MurmurHash)
