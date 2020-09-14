介绍：

短链接生成/跳转

Feature：
- 使用xorshift生成伪随机数，该方案生成随机数快。

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
- 使用 [bloom filter](https://en.wikipedia.org/wiki/Bloom_filter) 判断是否生成的字符串已存在
- bloom filter使用 [MurmurHash](https://en.wikipedia.org/wiki/MurmurHash) 中的 MurmurHash3 生成hash

TODO：
- [ ] Redis存储
- [ ] 对一个链接多次生成随机字符串应当与之前一致
- [ ] 反向查找
- [ ] 访问次数统计