/*
By Ross Wagner

Finds prime numbers less than 10^8 using a parallelized sieve.

outputs <execution time>  <total number of primes found>  <sum of all primes found> < top ten maximum primes, listed in order from lowest to highest> into primes.txt

*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)






func main(){



	var numTh int  = 8

	var upTo int = 100000000


	GoPrime(numTh, upTo)



}





/*runs some helper functions and prints results fo primes.txt*/
func GoPrime(numTh int, upTo int) {

	err := ioutil.WriteFile("primes.txt", []byte(""), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

	primesInfo ,err := os.OpenFile("primes.txt", os.O_RDWR,0666)

	if err != nil{
		log.Fatal(err)
	}
	//clear file and set cursor to beginning
	primesInfo.Truncate(0)
	primesInfo.Seek(0,0)


	// start timer
	start := time.Now()
	primes := ParallelSieveOfEratosthenes(upTo,numTh)
	// end timer
	elapsed := time.Since(start)
	finalSum := sumSlice(primes)
	finalCount:= len(primes)
	n:=10

	if finalCount < 10{
		n =finalCount
	}



	nLargest := primes[finalCount-n:]


	// write info to file
	primesInfo.WriteString(fmt.Sprintf("Primes up to: %d\n",upTo) )
	primesInfo.WriteString(fmt.Sprintf("Run time with %d threads: %s\n",numTh,elapsed) )
	primesInfo.WriteString(fmt.Sprintf("Sum: %d\n",finalSum) )
	primesInfo.WriteString(fmt.Sprintf("Count: %d\n",finalCount) )
	primesInfo.WriteString(fmt.Sprintf("%d Largest:\n",n) )

	for _, num := range nLargest{
		primesInfo.WriteString(fmt.Sprintf("\t%d\n",num) )
	}

}

func Pf(done chan<-bool,start int,p int,integers []bool, upTo int, numTh int){
	for i := p * (2+start); i <= upTo; i += p*numTh {
		integers[i] = false
	}
	done<-true
}
func ParallelFilter( p int,integers []bool, upTo int, numTh int){

	done := make(chan bool, numTh)
	// make numTh threads
	for j:=0;j<numTh;j++ {


		go Pf(done, j, p, integers, upTo, numTh)

	}

	// make sure filter is done
	for i:=0;i<numTh;i++{
		<-done
	}
}


/*
From: https://siongui.github.io/2017/04/17/go-sieve-of-eratosthenes/
modified for parallelism
*/
func ParallelSieveOfEratosthenes(upTo int, numTh int) []int {
	// Create a boolean array "prime[0..n]" and initialize
	// all entries it as true. A value in prime[i] will
	// finally be false if i is Not a prime, else true.
	integers := make([]bool, upTo+1)
	for i := 2; i < upTo+1; i++ {
		integers[i] = true
	}



	for p := 2; p*p <= upTo; p++ {
		// If integers[p] is not changed, then it is a prime
		if integers[p] == true {
			// Update all multiples of p
			ParallelFilter(p, integers, upTo, numTh)

		}


	}

	// get prime numbers
	var primes []int
	for p := 2; p <= upTo; p++ {
		if integers[p] == true {
			primes = append(primes, p)
		}
	}

	return primes

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


/*takes an int slice and returns a uint64*/
func sumSlice(nums []int)uint64{
	var sum uint64 = 0
	for _, num :=range nums{
		sum += uint64(num)
	}
	return sum
}

