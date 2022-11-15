# Common

[![Test Go](https://github.com/Cyb3r-Jak3/common-go/actions/workflows/go.yml/badge.svg)](https://github.com/Cyb3r-Jak3/common-go/actions/workflows/go.yml) [![Golanglint CI](https://github.com/Cyb3r-Jak3/common-go/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/Cyb3r-Jak3/common-go/actions/workflows/golangci-lint.yml)
[![DeepSource](https://deepsource.io/gh/Cyb3r-Jak3/common-go.svg/?label=active+issues&show_trend=true&token=lDZpKPdXNU-TQiyqQQBe4r7z)](https://deepsource.io/gh/Cyb3r-Jak3/common-go/?ref=repository-badge) [![Go Report Card](https://goreportcard.com/badge/github.com/Cyb3r-Jak3/common)](https://goreportcard.com/report/github.com/Cyb3r-Jak3/common) [![codecov](https://codecov.io/gh/Cyb3r-Jak3/common-go/branch/main/graph/badge.svg?token=L471VTTRPM)](https://codecov.io/gh/Cyb3r-Jak3/common-go)

This repo contains common code that I use between my programs. There are currently some tests and at some point I *might* get around to writing more.

Tested with go versions 1.18 & 1.19.

### Benchmark

Benchmarks are taken from latest-ubuntu and go version 1.17.

```
BenchmarkJSONResponse-2           	 1513520	       767.4 ns/op	    1024 B/op	      10 allocs/op
BenchmarkWOAllowedMethod-2        	 1550850	       773.7 ns/op	    1016 B/op	      10 allocs/op
BenchmarkAllowedMethods-2         	 1271031	      1005 ns/op	    1088 B/op	      13 allocs/op
BenchmarkDeniedAllowedMethods-2   	 1000000	      1213 ns/op	    1128 B/op	      14 allocs/op
BenchmarkContentResponse-2        	 1555020	       776.5 ns/op	    1013 B/op	      10 allocs/op
BenchmarkStringResponse-2         	 1527940	       782.8 ns/op	    1016 B/op	      10 allocs/op
BenchmarkJSONMarshall-2           	 1264372	       965.7 ns/op	    1024 B/op	      10 allocs/op
BenchmarkGenerate-2               	  764608	      1643 ns/op
BenchmarkJSONParse-2              	  711868	      1589 ns/op	      80 B/op	       2 allocs/op
BenchmarkYAMLParse-2              	  764983	      1600 ns/op	      80 B/op	       2 allocs/op
BenchmarkStringSearch2-2          	16068895	        70.55 ns/op
BenchmarkStringSearch10-2         	 6211059	       190.5 ns/op
BenchmarkFloatSearch2-2           	21945752	        54.78 ns/op
BenchmarkFloatSearch10-2          	10253065	       122.2 ns/op
BenchmarkIntSearch2-2             	21588042	        56.26 ns/op
BenchmarkIntSearch10-2            	11035240	       117.7 ns/op
BenchmarkGetEnv-2                 	30184161	        38.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetEnvMissing-2          	32234304	        37.16 ns/op
BenchmarkSHA256-2                 	   81972	     14140 ns/op
BenchmarkSHA384-2                 	   89167	     13359 ns/op
BenchmarkSHA512-2                 	   90374	     14348 ns/op
BenchmarkToHex-2                  	 7411296	       151.0 ns/op
BenchmarkSkipRoot-2               	  966310	      1210 ns/op
BenchmarkEnvironMap-2             	   52266	     22935 ns/op
```
