go-sql-benchmark
================

Benchmark among different SQL libraries in Go


Go-sql-benchmark will compare the following SQL libraries:

- raw [database/sql](http://golang.org/pkg/database/sql/driver/)
- https://github.com/gchaincl/dotsql
- https://github.com/jinzhu/gorm
- https://github.com/jmoiron/sqlx
- https://github.com/lann/squirrel


Usage:

    Run the shell script: ./run.sh


The benchmarks mentioned bellow were run on the following computer:

    Intel(R) Core(TM) i7-4500U CPU @ 1.80GHz
    8G Ram - 256GB SSD


Benchmark Output:

    PASS
    BenchmarkNative   100000         13476 ns/op         376 B/op         14 allocs/op
    BenchmarkSqlX     100000         16334 ns/op         537 B/op         17 allocs/op
    BenchmarkDotSQL   100000         13361 ns/op         376 B/op         14 allocs/op
    BenchmarkSqrl      50000         29306 ns/op        2564 B/op         53 allocs/op
    BenchmarkGorm      10000        642476 ns/op       13048 B/op        227 allocs/op
    ok      _/go-sql-benchmark  13.007s
