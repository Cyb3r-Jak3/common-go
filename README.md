# Common

[![Test Go](https://github.com/Cyb3r-Jak3/common-go/actions/workflows/go.yml/badge.svg)](https://github.com/Cyb3r-Jak3/common-go/actions/workflows/go.yml) [![DeepSource](https://deepsource.io/gh/Cyb3r-Jak3/common-go.svg/?label=active+issues&show_trend=true&token=lDZpKPdXNU-TQiyqQQBe4r7z)](https://deepsource.io/gh/Cyb3r-Jak3/common-go/?ref=repository-badge) [![Go Report Card](https://goreportcard.com/badge/github.com/Cyb3r-Jak3/common)](https://goreportcard.com/report/github.com/Cyb3r-Jak3/common)

This repo contains common code that I use between my programs. There are currently some tests and at some point I *might* get around to writing more.

Tested with go versions 1.13 to 1.16.

### Benchmark

Benchmarks are taken from latest-ubuntu and go version 1.16.

```
BenchmarkJSONResponse-2      	 1155928	      1016 ns/op	    1024 B/op	      10 allocs/op
BenchmarkWOAllowedMethod-2   	 1000000	      1039 ns/op	    1016 B/op	      10 allocs/op
BenchmarkAllowedMethod-2     	 1000000	      1206 ns/op	    1064 B/op	      12 allocs/op
BenchmarkContentResponse-2   	 1000000	      1021 ns/op	    1013 B/op	      10 allocs/op
BenchmarkStringResponse-2    	 1000000	      1040 ns/op	    1016 B/op	      10 allocs/op
BenchmarkJSONMarshall-2      	  933313	      1343 ns/op	    1024 B/op	      10 allocs/op
BenchmarkGenerate-2          	  614492	      1982 ns/op
BenchmarkJSONParse-2         	  106275	     11315 ns/op	    1096 B/op	      12 allocs/op
BenchmarkYAMLParse-2         	   49432	     23725 ns/op	    7616 B/op	      84 allocs/op
BenchmarkStringSearch2-2     	44019916	        26.47 ns/op
BenchmarkStringSearch10-2    	25348384	        46.87 ns/op
BenchmarkFloatSearch2-2      	99105530	        11.99 ns/op
BenchmarkFloatSearch10-2     	61142461	        18.78 ns/op
BenchmarkIntSearch2-2        	80376837	        15.03 ns/op
BenchmarkIntSearch10-2       	54109246	        22.17 ns/op
BenchmarkGetEnv-2            	26952231	        46.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetEnvMissing-2     	24062250	        50.51 ns/op
BenchmarkSHA256-2            	   62689	     18296 ns/op
BenchmarkSHA384-2            	   70474	     17190 ns/op
BenchmarkSHA512-2            	   69765	     17052 ns/op
```