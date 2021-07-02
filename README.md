# Common

[![Test Go](https://github.com/Cyb3r-Jak3/common-go/actions/workflows/go.yml/badge.svg)](https://github.com/Cyb3r-Jak3/common-go/actions/workflows/go.yml) [![DeepSource](https://deepsource.io/gh/Cyb3r-Jak3/common-go.svg/?label=active+issues&show_trend=true&token=lDZpKPdXNU-TQiyqQQBe4r7z)](https://deepsource.io/gh/Cyb3r-Jak3/common-go/?ref=repository-badge) [![Go Report Card](https://goreportcard.com/badge/github.com/Cyb3r-Jak3/common)](https://goreportcard.com/report/github.com/Cyb3r-Jak3/common)

This repo contains common code that I use between my programs. There are currently some tests and at some point I *might* get around to writing more.

Tested with go versions 1.13 to 1.16.

### Benchmark

    BenchmarkJSONResponse-2      	 1361775	       908.0 ns/op	    1024 B/op	      10 allocs/op
    BenchmarkWOAllowedMethod-2   	 1374057	       869.0 ns/op	    1016 B/op	      10 allocs/op
    BenchmarkAllowedMethod-2     	 1000000	      1038 ns/op	    1064 B/op	      12 allocs/op
    BenchmarkContentResponse-2   	 1378454	       859.3 ns/op	    1013 B/op	      10 allocs/op
    BenchmarkStringResponse-2    	 1360341	       869.7 ns/op	    1016 B/op	      10 allocs/op
    BenchmarkJSONMarshall-2      	 1000000	      1119 ns/op	    1024 B/op	      10 allocs/op
    BenchmarkGenerate-2          	  700604	      1746 ns/op
    BenchmarkJSONParse-2         	  124810	      9510 ns/op	    1096 B/op	      12 allocs/op
    BenchmarkYAMLParse-2         	   59620	     20274 ns/op	    7616 B/op	      84 allocs/op
    BenchmarkStringSearch2-2     	51719239	        23.08 ns/op
    BenchmarkStringSearch10-2    	29695441	        40.16 ns/op
    BenchmarkFloatSearch2-2      	100000000	        10.15 ns/op
    BenchmarkFloatSearch10-2     	69445460	        16.40 ns/op
    BenchmarkIntSearch2-2        	80205716	        12.94 ns/op
    BenchmarkIntSearch10-2       	60177013	        18.84 ns/op
    BenchmarkGetEnv-2            	30803790	        39.05 ns/op	       0 B/op	       0 allocs/op
    BenchmarkGetEnvMissing-2     	32854902	        36.49 ns/op
