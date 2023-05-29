package main

func decode(encoded []int) []int {

	n := len(encoded)

	// Space complexity is obscured here...
	// If cap(encoded) > len(encoded) then this space is O(1)
	// If they are equal, Go will double the cap and space is O(n)
	encoded = append(encoded, 0)

	// track difference between bits we have and bits we need
	// we need n+1 and have a 0 that we just appended; neither are in the loop
	xor := n + 1

	for i := n; i > 0; i-- {
		// at each step, encoded[i:] has the correct value of perm
		// encoded[i-1] = encoded[i-1] ^ encoded[i]; save to encoded
		encoded[i-1] ^= encoded[i]
		xor ^= i ^ encoded[i-1]
	}

	for i := range encoded {
		encoded[i] ^= xor
	}

	return encoded
}
