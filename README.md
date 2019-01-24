## Install golang 1.11.4 or later

https://golang.org/dl/

Ubuntu:
sudo snap install --classic go

## Run
navigate to directory where primes.go is

go run primes.go

output in primes.txt in same directory

## Overview

Parallelization happens on line 99. When filtering out a prime number the filtration is broken into 8 threads filtering out multiples.    


