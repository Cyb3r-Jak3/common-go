# Common

[![Test Go](https://github.com/Cyb3r-Jak3/common-go/actions/workflows/go.yml/badge.svg)](https://github.com/Cyb3r-Jak3/common-go/actions/workflows/go.yml) [![Golanglint CI](https://github.com/Cyb3r-Jak3/common-go/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/Cyb3r-Jak3/common-go/actions/workflows/golangci-lint.yml)
[![DeepSource](https://deepsource.io/gh/Cyb3r-Jak3/common-go.svg/?label=active+issues&show_trend=true&token=lDZpKPdXNU-TQiyqQQBe4r7z)](https://deepsource.io/gh/Cyb3r-Jak3/common-go/?ref=repository-badge) [![Go Report Card](https://goreportcard.com/badge/github.com/Cyb3r-Jak3/common)](https://goreportcard.com/report/github.com/Cyb3r-Jak3/common) [![codecov](https://codecov.io/gh/Cyb3r-Jak3/common-go/branch/main/graph/badge.svg?token=L471VTTRPM)](https://codecov.io/gh/Cyb3r-Jak3/common-go)

This repo contains common code that I use between my programs. There are currently some tests and at some point I *might* get around to writing more.

Tested with go versions 1.16 & 1.17.

### Benchmark

Benchmarks are taken from latest-ubuntu and go version 1.16.

```
BenchmarkJSONResponse-2      	 1392727	       859.1 ns/op	    1024 B/op	      10 allocs/op
BenchmarkWOAllowedMethod-2   	 1257962	       890.7 ns/op	    1016 B/op	      10 allocs/op
BenchmarkAllowedMethod-2     	 1000000	      1010 ns/op	    1064 B/op	      12 allocs/op
BenchmarkContentResponse-2   	 1403922	       856.2 ns/op	    1013 B/op	      10 allocs/op
BenchmarkStringResponse-2    	 1373049	       866.9 ns/op	    1016 B/op	      10 allocs/op
BenchmarkJSONMarshall-2      	 1000000	      1103 ns/op	    1024 B/op	      10 allocs/op
BenchmarkGenerate-2          	  687927	      1759 ns/op
BenchmarkJSONParse-2         	  736328	      1734 ns/op	      80 B/op	       2 allocs/op
BenchmarkYAMLParse-2         	  731004	      1747 ns/op	      80 B/op	       2 allocs/op
BenchmarkStringSearch2-2     	52436001	        22.40 ns/op
BenchmarkStringSearch10-2    	29598258	        39.59 ns/op
BenchmarkFloatSearch2-2      	100000000	        10.19 ns/op
BenchmarkFloatSearch10-2     	73469235	        16.32 ns/op
BenchmarkIntSearch2-2        	93743188	        12.67 ns/op
BenchmarkIntSearch10-2       	64075212	        18.75 ns/op
BenchmarkGetEnv-2            	30846982	        38.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetEnvMissing-2     	31355505	        37.83 ns/op
BenchmarkSHA256-2            	   72594	     16070 ns/op
BenchmarkSHA384-2            	   78955	     15485 ns/op
BenchmarkSHA512-2            	   77162	     15605 ns/op
BenchmarkToHex-2             	 7395942	       166.2 ns/op
BenchmarkSkipRoot-2          	  812551	      1494 ns/op
BenchmarkEnvironMap-2        	   45308	     26414 ns/op
```