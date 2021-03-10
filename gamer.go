package snart

import (
	"log"
	"math/rand"
	"time"

	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
)

// Gamer generates an Activity.
type Gamer interface {
	Activity() discord.Activity
}

func (b *Bot) selectGamer(i int) Gamer {
	g := b.Gamers[i]

	// example: 5 gamers, selecting #2
	// 01234
	copy(b.Gamers[i:], b.Gamers[i+1:])
	// 01344
	copy(b.Gamers[len(b.Gamers)-1:], []Gamer{g})
	// 01342

	return g
}

func (b *Bot) randomGamer() Gamer {
	//nolint:gomnd,gosec
	// selectGamer pushes to the end, so select from the first half (older)
	i := rand.Intn((len(b.Gamers) + 1) / 2)

	return b.selectGamer(i)
}

// CycleGamers continually updates the Bot's status using a random Gamer, on an interval.
func (b *Bot) CycleGamers() {
	const interval = time.Second * 12

	tick := time.NewTicker(interval)

	for {
		g := b.randomGamer()
		a := g.Activity()

		//nolint:exhaustivestruct // discord
		data := gateway.UpdateStatusData{
			Activities: []discord.Activity{a},
		}

		err := b.State.Gateway.UpdateStatus(data)
		if err != nil {
			log.Printf("update status (%#v): %s", g, err)
		}

		<-tick.C
	}
}

// GamerFunc is a func that is a Gamer.
type GamerFunc func() discord.Activity

// Activity simply calls the GamerFunc and returns the returned Activity.
func (f GamerFunc) Activity() discord.Activity {
	return f()
}

// GamerStatic is an Activity that is a Gamer.
type GamerStatic discord.Activity

// Activity simply returns the GamerStatic as an Activity.
func (s GamerStatic) Activity() discord.Activity {
	return discord.Activity(s)
}

// GamerTimer is a Time that is a Gamer.
type GamerTimer time.Time

// Activity returns an Activity that describes the duration since the GamerTimer.
func (t GamerTimer) Activity() discord.Activity {
	since := time.Since(time.Time(t)).Round(time.Second)

	//nolint:exhaustivestruct // discord
	return discord.Activity{
		Name: "for " + since.String(),
		Type: discord.GameActivity,
	}
}
