/*

Author: Princess Ariyibi

This function determines if a given number is a prime number.
You can run this file using 

		go run proth_prime.go

*/


package main
import (
	"fmt"
	"math"
)

func main() {
    fmt.Print("Enter a number you want to check if its a proth prime number: ")
    var p int
	fmt.Scanln(&p)
	number_is_a_proth_number := is_proth(p - 1)
	if number_is_a_proth_number {
		fmt.Print(p, " is a proth number\n")
		if is_proth_prime(p) {
			fmt.Print(p, " is a proth prime number\n")
		} else {
			fmt.Print(p, " is a not a prime number\n")
		}
	} else {
		fmt.Print(p, " is not a proth number\n")
    }
}

func is_power_of_two(x int) bool {
	if (x == 0) {
		return false
	}
	return (byte(x) & byte(x - 1)) == 0
}

func is_proth(p int) bool {
	k := 1
	for k < (p / k) {
		if (p % k == 0) {
			if is_power_of_two(p / k) {
				return true
			}
		}
		k += 2
	}
	return false
}

func is_proth_prime(p int) bool {
	exponent := float64((p - 1) / 2)
	tracker := 0
	for a := 0; a < 10; a++ {
		value := int(math.Pow(float64(a), exponent) + 1)
		is_prime := value % p == 0
		if is_prime {
			tracker += 1
		}
	}
	return tracker > 1
}