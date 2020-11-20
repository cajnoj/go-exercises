#!/bin/bash
benchmark() {
    echo $*
    /usr/bin/time -v ./mandelbrot $* 2>&1 | grep -E "Elapsed|Maximum resident set size"
}

go build mandelbrot.go

# big.Rat won't be benchmarked because it's extremely slow

benchmark complex64 -2 -2 4 512
benchmark complex128 -2 -2 4 512
benchmark big.Float -2 -2 4 512
#benchmark big.Rat -2 -2 4 512

benchmark complex64 0.40052917 0.3505417 4.17e-06 1024
benchmark complex128 0.40052917 0.3505417 4.17e-06 1024
benchmark big.Float 0.40052917 0.3505417 4.17e-06 1024
#benchmark big.Rat 0.40052917 0.3505417 4.17e-06 1024
