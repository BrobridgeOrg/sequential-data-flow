# sequential-data-flow

High performance sequential data processing.

## Benchmark

Here is result of benchmark with parsing JSON string:

```shell
go test -bench=. .
goos: darwin
goarch: amd64
pkg: github.com/BrobridgeOrg/sequential-data-flow
BenchmarkBaseline/Small-16                  	  743193	      1432 ns/op
BenchmarkBaseline/Large-16                  	   10000	    105195 ns/op
BenchmarkLowBufferSize/Small-16             	 1582555	       738 ns/op
BenchmarkLowBufferSize/Large-16             	   21139	     58260 ns/op
BenchmarkHighBufferSize/Small-16            	 1616371	       878 ns/op
BenchmarkHighBufferSize/Large-16            	   36951	     47298 ns/op
BenchmarkHighWorkerCount/Small-16           	 1941549	       717 ns/op
BenchmarkHighWorkerCount/Large-16           	   52912	     29032 ns/op
BenchmarkHighBufferSizeAndWorkerCount/Small-16         	 1267849	      1035 ns/op
BenchmarkHighBufferSizeAndWorkerCount/Large-16         	   41293	     38983 ns/op
PASS
ok  	github.com/BrobridgeOrg/sequential-data-flow	68.936s
```

## License
Licensed under the Apache License

## Authors
Copyright(c) 2021 Fred Chien <fred@brobridge.com>
