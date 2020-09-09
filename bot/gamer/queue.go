package gamer

import "math/rand"

// Queue provides a pseudo-random queue of Gamers.
type Queue []Gamer

// Select returns the Gamer at position i in the Queue, updating the order accordingly.
func (q Queue) Select(i int) Gamer {
	g := q[i]
	copy(q[i:], q[i+1:])
	copy(q[len(q)-1:], []Gamer{g})

	return g
}

// Random calls Select on an index in [0, (len(q) + 1) / 2).
func (q Queue) Random() Gamer {
	// nolint:gosec
	// gamer queue doesn't need to be secure lmao
	return q.Select(rand.Intn((len(q) + 1) / 2))
}
