/*

<execution time>  <total number of
primes found>  <sum of all primes found>
< top ten maximum primes, listed in order from lowest to highest>

*/
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var initUpTo = 100


func main(){



	var numTh int = 8
	var upTo int = 100000000

	// benchmark
	/*
	var primes = []int{}

	start := time.Now()
	// simple single thread brute force approach
	for i := 3; i < upTo; i+=2{
		if SimplePrimeCheck(i,0,i){
			primes = append(primes,i)
		}
	}

	primesStr := IntSliceToString(primes)

	err:=ioutil.WriteFile("primes.txt",[]byte(primesStr),0)

	if err != nil{
		log.Fatal(err)
	}

	elapsed := time.Since(start)

	log.Printf("Single thread brute force time: %s",elapsed)
	*/
	start := time.Now()
	GoPrime(numTh, upTo)
	elapsed := time.Since(start)

	log.Printf("Multi thread brute force time: %s",elapsed)
}

/*
using both probabilistic and simple deterministic checks, quickly (hopefully) determines if a number is prime
*/
func CheckPrime(n int)bool{

	return SimplePrimeCheck(n, 0, n)
	/*
	//res := SimplePrimeCheck(n, 0, initUpTo)
	res:= true
	if res {
		// using Baillie–PSW primality test https://en.wikipedia.org/wiki/Baillie%E2%80%93PSW_primality_test
		res = ProbPrimeCheck(n)

		if res{
			// check found primes with simplePrimeCheck
			res = SimplePrimeCheck(n, 0, n)
		}
	}
	return res
	*/

}

/*nums some helper functions and merges the lists of primes*/
func GoPrime(numTh int, upTo int) {

	// create a buffer channel with numTh slots. threads will report their list of primes on completion
	lists := make(chan []int, numTh)
	done := make(chan bool, numTh)

	//var buffer bytes.Buffer

	// give each thread a start value then check <start value>+ k*numTh values
	//for each thread call ParallelPrimes
	for i:=0;i<numTh;i++{
		// start with the first numTh odds after 2. 3,5,7, ...ect

		go ParallelPrimes(done, lists, i, numTh, upTo)
	}


	// make sure each thread is done
	for i:=0;i<numTh;i++{
		<-done
	}

	/*
	// parallelized merge
	for len(lists) >1{
		buffer.WriteString( IntSliceToString(<-lists))
		//mergeLists(list)
	}
	*/

	primesInfo ,err := os.OpenFile("primes.txt", os.O_RDWR,0666)

	if err != nil{
		log.Fatal(err)
	}

	//clear file and set cursor to beginning
	primesInfo.Truncate(0)
	primesInfo.Seek(0,0)





}
/*creates the threads */
func ParallelPrimes(done chan<- bool, lists chan<- []int, start int, numTh int, upTo int){
	primes := []int{}

	// starting thread start
	fmt.Printf("Starting thread %d\n", start)

	startTime := time.Now()

	jumpAmt := numTh
	// if even jump amt is 2*numTh
	if numTh%2 == 0{
		jumpAmt = 2*numTh
	}


	for i:=3+start*2;i<upTo;i+=jumpAmt{
		if CheckPrime(i){
			primes = append(primes,i)
		}
	}


	elapsed := time.Since(startTime)
	fmt.Printf("Finished thread %d in %s\n", start, elapsed)
	fmt.Printf("\tNum Found: %d\n",len(primes))
	if len(primes)!=0{
		fmt.Printf("\tMax: %d\n",primes[len(primes)-1])
	}



	lists<-primes
	done<-true


}

/**/
func ProbPrimeCheck(n int) bool{

	// using Baillie–PSW primality test https://en.wikipedia.org/wiki/Baillie%E2%80%93PSW_primality_test

	// let n be an odd positive int that we wish to check for primality

	// trial division of small primes upto 100? (use simplePrimeCheck alg)

	// base 2 strong probable prime test. If n is not a strong probable prime base 2, then n is composite; quit. https://en.wikipedia.org/wiki/Strong_pseudoprime

	// Find the first D in the sequence 5, −7, 9, −11, 13, −15, ... [(5+2i)*-1^i, i starts at 0] for which the Jacobi symbol (D/n) is −1. Set P = 1 and Q = (1 − D) / 4.

	// Perform a strong Lucas probable prime test on n using parameters D, P, and Q. If n is not a strong Lucas probable prime, then n is composite. Otherwise, n is almost certainly prime.

	return true
}

/*
returns true if int n is prime, false otherwise
takes a start and max value to check
*/
func SimplePrimeCheck(n int, start int, max int)bool{
	if n <= 3{
		return n > 1
	}else if n % 2 == 0 || n % 3 == 0{
		return false
	}else{

		i := start
		if i < 5{
			i = 5
		}

		for i * i <= n && i <= max{
			if n % i == 0 || n % (i + 2) == 0{
				return false
			}

			i = i + 6
		}
		return true
	}
}
/*
from: https://www.dotnetperls.com/convert-slice-string-go
*/
func IntSliceToString(nums []int)string{
	valuesText := []string{}

	// Create a string slice using strconv.Itoa.
	// ... Append strings to it.
	for i := range nums {
		number := nums[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}

	// Join our string slice.
	result := strings.Join(valuesText, " ")
	return result
}
