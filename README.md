# Common

[![Test Go](https://github.com/Cyb3r-Jak3/common-go/actions/workflows/go.yml/badge.svg)](https://github.com/Cyb3r-Jak3/common-go/actions/workflows/go.yml) [![Golanglint CI](https://github.com/Cyb3r-Jak3/common-go/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/Cyb3r-Jak3/common-go/actions/workflows/golangci-lint.yml)
[![DeepSource](https://deepsource.io/gh/Cyb3r-Jak3/common-go.svg/?label=active+issues&show_trend=true&token=lDZpKPdXNU-TQiyqQQBe4r7z)](https://deepsource.io/gh/Cyb3r-Jak3/common-go/?ref=repository-badge) [![Go Report Card](https://goreportcard.com/badge/github.com/Cyb3r-Jak3/common)](https://goreportcard.com/report/github.com/Cyb3r-Jak3/common) [![codecov](https://codecov.io/gh/Cyb3r-Jak3/common-go/branch/main/graph/badge.svg?token=L471VTTRPM)](https://codecov.io/gh/Cyb3r-Jak3/common-go)

This repo contains common code that I use between my programs. There are currently some tests and at some point I *might* get around to writing more.

Tested with go versions 1.18, 1.19, and 1.20.

### Benchmark

Benchmarks are taken from latest-ubuntu and go version 1.20.

```
BenchmarkJSONResponse-2           	 1480198	       874.1 ns/op	    1024 B/op	      10 allocs/op
BenchmarkWOAllowedMethod-2        	 1455447	       809.3 ns/op	    1016 B/op	      10 allocs/op
BenchmarkAllowedMethods-2         	 1000000	      1165 ns/op	    1088 B/op	      13 allocs/op
BenchmarkDeniedAllowedMethods-2   	  808725	      1269 ns/op	    1128 B/op	      14 allocs/op
BenchmarkContentResponse-2        	 1508470	       799.8 ns/op	    1013 B/op	      10 allocs/op
BenchmarkStringResponse-2         	 1480393	       816.2 ns/op	    1016 B/op	      10 allocs/op
BenchmarkJSONMarshall-2           	 1205497	      1002 ns/op	    1024 B/op	      10 allocs/op
BenchmarkGenerate-2               	  804207	      1457 ns/op
BenchmarkJSONParse-2              	  694284	      1871 ns/op	      80 B/op	       2 allocs/op
BenchmarkYAMLParse-2              	  619479	      1844 ns/op	      80 B/op	       2 allocs/op
BenchmarkStringSearch2-2          	14119131	        83.43 ns/op
BenchmarkStringSearch10-2         	 7538738	       160.3 ns/op
BenchmarkFloatSearch2-2           	17752921	        65.32 ns/op
BenchmarkFloatSearch10-2          	10868433	       109.4 ns/op
BenchmarkIntSearch2-2             	17698303	        65.83 ns/op
BenchmarkIntSearch10-2            	11625736	       103.5 ns/op
BenchmarkGetEnv-2                 	31068798	        38.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetEnvMissing-2          	32717121	        36.52 ns/op
BenchmarkSHA256-2                 	   54486	     21549 ns/op
BenchmarkSHA384-2                 	   56664	     21297 ns/op
BenchmarkSHA512-2                 	   55195	     21276 ns/op
BenchmarkSkipRoot-2               	  789014	      1295 ns/op
BenchmarkEnvironMap-2             	   45550	     28015 ns/op
BenchmarkFileExists-2             	  951858	      1225 ns/op
BenchmarkFileExistsMissing-2      	  972051	      1159 ns/op
```
