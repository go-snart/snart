package gamer

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/log"
)

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

// Cycle continually updates status on the session.
func (q Queue) Cycle(ctx context.Context, ses *dg.Session) {
	for {
		select {
		case <-ctx.Done():
			err := ctx.Err()
			if err != nil {
				err = fmt.Errorf("ctx err: %w", err)
				log.Warn.Println(err)

				return
			}
		default:
		}

		game := q.Random().Game()
		log.Debug.Printf("%v\n", game)

		for {
			err := ses.UpdateStatusComplex(
				dg.UpdateStatusData{Game: game},
			)
			if err == nil {
				break
			}

			if !errors.Is(err, dg.ErrWSNotFound) {
				err = fmt.Errorf("update status: %w", err)
				log.Warn.Println(err)
			}

			time.Sleep(time.Second / 10)
		}

		time.Sleep(time.Second * 12)
	}
}
