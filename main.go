package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"sort"
	"time"
	"errors"
	"encoding/gob"
	"bytes"
)

func main() {
	ErrorEncode()
}

func ErrorEncode() {

	errToSend := errors.New("this is an error")

	var network bytes.Buffer // Stand-in for the network.

	// Create an encoder and send a value.
	enc := gob.NewEncoder(&network)
	err := enc.Encode(errToSend)
	if err != nil {
		log.Fatal("encode:", err)
	}	

	// Create a decoder and receive a value.
	dec := gob.NewDecoder(&network)
	var errReceived error
	err = dec.Decode(&errReceived)
	if err != nil {
		log.Fatal("decode:", err)
	}	

	fmt.Println(&errToSend, &errReceived)
	fmt.Println(errReceived)

}	

func students() {
	students := readJson("ex1.json")
	mentors := readJson("ex2.json")
	var beg, end time.Time
	for limit := 1; limit < 1000; limit <<= 1 {
		fmt.Println("limit =", limit)
		beg = time.Now()
		a := maxCompatibilitySum(students[:limit], mentors[:limit])
		end = time.Now()
		fmt.Println("adam", a, end.Sub(beg))
		beg = time.Now()
		s := maxCompatibilitySumOther(students[:limit], mentors[:limit])
		end = time.Now()
		fmt.Println("slow", s, end.Sub(beg))
		fmt.Println()
	}
}

func readJson(fn string) [][]int {
	// Let's first read the `config.json` file
	content, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	var payload [][]int
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return payload
}

func xorChain(x int) {
	var xor int
	for i := 1; i <= x; i++ {
		xor ^= i
		if i&1 == 1 {
			fmt.Println(i, xor, ((i&2)>>1)^1)
		}
	}
}

func perms(n int) chan []int {

	ch := make(chan []int)

	go func() {
		defer close(ch)

		perm := make([]int, n)
		for i := range perm {
			perm[i] = i + 1
		}

		for {

			send := make([]int, n)
			copy(send, perm)
			ch <- send

			nextPermutation(perm)
			if sort.IntsAreSorted(perm) {
				break
			}
		}
	}()

	return ch
}

func encode(arr []int) []int {
	var ans []int
	for i := 0; i+1 < len(arr); i++ {
		ans = append(ans, arr[i]^arr[i+1])
	}
	return ans
}

func BinSearchPrimes() {
	var start time.Time
	var sTime, wTime time.Duration

	var breaks []int
	limit := 50

	for x := 0; x < 10_000; x++ {

		b := sort.Search(limit, func(limit int) bool {

			if limit < 10 {
				return false
			}

			start = time.Now()
			sieve(limit)
			sTime = time.Since(start)

			start = time.Now()
			willans(limit)
			wTime = time.Since(start)

			return sTime < wTime

		})

		breaks = append(breaks, b)
		limit = b << 1
	}

	var sum int
	for _, x := range breaks {
		sum += x
	}

	fmt.Println(float64(sum) / float64(len(breaks)))
}

func sieve(limit int) {

	comp := map[int][]int{}

	for p, found := 2, 0; found <= limit; p++ {
		if fctrs := comp[p]; len(fctrs) == 0 {
			comp[2*p] = append(comp[2*p], p)
			found++
		} else {
			for _, f := range fctrs {
				comp[p+f] = append(comp[p+f], f)
			}
		}
		delete(comp, p)
	}
}

func willans(limit int) {

	one := big.NewInt(1)
	zero := big.NewInt(0)

	fact := big.NewInt(1)
	mod := big.NewInt(0)
	top := big.NewInt(0)

	j := big.NewInt(1)
	for found := 0; found < limit; {

		// multiply in current j
		fact.Mul(fact, j)
		top.Add(fact, one)

		// increment j
		j.Add(j, one)

		if mod.Mod(top, j); mod.Cmp(zero) == 0 {
			found++
		}
	}
}
