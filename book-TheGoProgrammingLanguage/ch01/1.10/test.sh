#!/bin/bash
time ./fetchall http://introcs.cs.princeton.edu/data/pi-10million.txt > fetchall.1.out
time ./fetchall http://introcs.cs.princeton.edu/data/pi-10million.txt > fetchall.2.out
diff fetchall.1.out fetchall.2.out