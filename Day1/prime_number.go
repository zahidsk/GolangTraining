package main

import "fmt"

func main()  {
	fmt.Println("Printing prime no.till 100")
	prime_array := []int{}
	for n:=0; n<=100; n++{
		if IsPrime(n) {
			prime_array = append(prime_array, n)
		}
	}
	fmt.Println(prime_array)
	fmt.Println("================================")
	var n int
	fmt.Println("You want to test, if your number is prime ..please enter no.")
	fmt.Scanf("%d",&n)
	if IsPrime(n) {
		fmt.Printf("Your number %d is prime ",n)
	} else { fmt.Printf("Your number %d is not prime ",n)}
}

func IsPrime(n int) bool{
	if n <= 1{ return false}
	if n <= 3 {return true}
	if (n%2 == 0 || n%3 == 0 || n%5 == 0) {return false}

	for i:=5; i*i<n; i=i+6{
		if n%i == 0 || n%(i+2)==0 {
			return false
		}
	}
	return true
}