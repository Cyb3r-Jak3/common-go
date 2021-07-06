# Common

[![Test Go](https://github.com/Cyb3r-Jak3/common-go/actions/workflows/go.yml/badge.svg)](https://github.com/Cyb3r-Jak3/common-go/actions/workflows/go.yml) [![DeepSource](https://deepsource.io/gh/Cyb3r-Jak3/common-go.svg/?label=active+issues&show_trend=true&token=lDZpKPdXNU-TQiyqQQBe4r7z)](https://deepsource.io/gh/Cyb3r-Jak3/common-go/?ref=repository-badge) [![Go Report Card](https://goreportcard.com/badge/github.com/Cyb3r-Jak3/common)](https://goreportcard.com/report/github.com/Cyb3r-Jak3/common)

This repo contains common code that I use between my programs. There are currently some tests and at some point I *might* get around to writing more.

Tested with go versions 1.13 to 1.16.

### Benchmark

Benchmarks are taken from latest-ubuntu and go version 1.16.

```
BenchmarkJSONResponse-2      	 1000000	      1040 ns/op	    1024 B/op	      10 allocs/op
BenchmarkWOAllowedMethod-2   	 1000000	      1060 ns/op	    1016 B/op	      10 allocs/op
BenchmarkAllowedMethod-2     	  973203	      1243 ns/op	    1064 B/op	      12 allocs/op
BenchmarkContentResponse-2   	 1000000	      1061 ns/op	    1013 B/op	      10 allocs/op
BenchmarkStringResponse-2    	 1000000	      1085 ns/op	    1016 B/op	      10 allocs/op
BenchmarkJSONMarshall-2      	  882159	      1332 ns/op	    1024 B/op	      10 allocs/op
BenchmarkGenerate-2          	  607874	      2062 ns/op
BenchmarkJSONParse-2         	  109234	     11144 ns/op	    1096 B/op	      12 allocs/op
BenchmarkYAMLParse-2         	   50282	     24365 ns/op	    7616 B/op	      84 allocs/op
BenchmarkStringSearch2-2     	43035110	        27.43 ns/op
BenchmarkStringSearch10-2    	25246152	        47.51 ns/op
BenchmarkFloatSearch2-2      	97814433	        12.16 ns/op
BenchmarkFloatSearch10-2     	59949708	        19.97 ns/op
BenchmarkIntSearch2-2        	76164642	        15.43 ns/op
BenchmarkIntSearch10-2       	53413640	        23.13 ns/op
BenchmarkGetEnv-2            	23637121	        48.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetEnvMissing-2     	27369729	        44.75 ns/op
BenchmarkSHA1-2              	   74755	     15850 ns/op
BenchmarkSHA256-2            	   63435	     18890 ns/op
BenchmarkSHA384-2            	   63531	     18046 ns/op
BenchmarkSHA512-2            	   64582	     17614 ns/op
```