package main

// countCollisions returns the total number of collisions on the road.
// Cars that forever move outward (leading L's, trailing R's) never collide,
// so we trim them and count every remaining moving car (non-'S') as a collision.
func countCollisions(directions string) int {
	bytes := []byte(directions)
	n := len(bytes)

	left := 0
	for left < n && bytes[left] == 'L' {
		left++
	}

	right := n - 1
	for right >= left && bytes[right] == 'R' {
		right--
	}

	collisions := 0
	for i := left; i <= right; i++ {
		if bytes[i] != 'S' {
			collisions++
		}
	}

	return collisions
}
