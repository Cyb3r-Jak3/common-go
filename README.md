# Common

[![Test Go](https://github.com/Cyb3r-Jak3/common-go/actions/workflows/go.yml/badge.svg)](https://github.com/Cyb3r-Jak3/common-go/actions/workflows/go.yml) [![Golanglint CI](https://github.com/Cyb3r-Jak3/common-go/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/Cyb3r-Jak3/common-go/actions/workflows/golangci-lint.yml)
[![DeepSource](https://deepsource.io/gh/Cyb3r-Jak3/common-go.svg/?label=active+issues&show_trend=true&token=lDZpKPdXNU-TQiyqQQBe4r7z)](https://deepsource.io/gh/Cyb3r-Jak3/common-go/?ref=repository-badge) [![Go Report Card](https://goreportcard.com/badge/github.com/Cyb3r-Jak3/common)](https://goreportcard.com/report/github.com/Cyb3r-Jak3/common) [![codecov](https://codecov.io/gh/Cyb3r-Jak3/common-go/branch/main/graph/badge.svg?token=L471VTTRPM)](https://codecov.io/gh/Cyb3r-Jak3/common-go)

This repo contains common code that I use between my programs. There are currently some tests and at some point I *might* get around to writing more.

Tested with go versions 1.18 & 1.19.

### Benchmark

Benchmarks are taken from latest-ubuntu and go version 1.19.

```
BenchmarkJSONResponse-2           	 1450099	       829.3 ns/op	    1024 B/op	      10 allocs/op
BenchmarkWOAllowedMethod-2        	 1439946	       836.9 ns/op	    1016 B/op	      10 allocs/op
BenchmarkAllowedMethods-2         	 1000000	      1011 ns/op	    1088 B/op	      13 allocs/op
BenchmarkDeniedAllowedMethods-2   	  849014	      1272 ns/op	    1128 B/op	      14 allocs/op
BenchmarkContentResponse-2        	 1444442	       920.3 ns/op	    1013 B/op	      10 allocs/op
BenchmarkStringResponse-2         	 1427762	       839.8 ns/op	    1016 B/op	      10 allocs/op
BenchmarkJSONMarshall-2           	  988467	      1037 ns/op	    1024 B/op	      10 allocs/op
BenchmarkGenerate-2               	  810075	      1486 ns/op
BenchmarkJSONParse-2              	  769706	      1631 ns/op	      80 B/op	       2 allocs/op
BenchmarkYAMLParse-2              	  678457	      1614 ns/op	      80 B/op	       2 allocs/op
BenchmarkStringSearch2-2          	15896445	        73.79 ns/op
BenchmarkStringSearch10-2         	 7631498	       157.9 ns/op
BenchmarkFloatSearch2-2           	20262049	        58.03 ns/op
BenchmarkFloatSearch10-2          	10148270	       117.0 ns/op
BenchmarkIntSearch2-2             	19276776	        61.46 ns/op
BenchmarkIntSearch10-2            	10009627	       118.8 ns/op
BenchmarkGetEnv-2                 	31326195	        37.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetEnvMissing-2          	32664105	        36.51 ns/op
BenchmarkSHA256-2                 	   70501	     16564 ns/op
BenchmarkSHA384-2                 	   77529	     15798 ns/op
BenchmarkSHA512-2                 	   74725	     15967 ns/op
BenchmarkSkipRoot-2               	  816931	      1310 ns/op
BenchmarkEnvironMap-2             	   44618	     26481 ns/op
```
