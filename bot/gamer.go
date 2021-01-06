package bot

import (
	"fmt"
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

	log.Printf("select gamer %d: %v", i, g)

	copy(b.Gamers[i:], b.Gamers[i+1:])
	copy(b.Gamers[len(b.Gamers)-1:], []Gamer{g})

	return g
}

func (b *Bot) randomGamer() (int, Gamer) {
	//nolint:gomnd,gosec
	// 2 is for selecting the first half
	// this doesn't need to be secure lmao
	i := rand.Intn((len(b.Gamers) + 1) / 2)

	log.Printf("random gamer %d", i)

	return i, b.selectGamer(i)
}

// CycleGamers continually updates the Bot's status using a random Gamer, on an interval.
func (b *Bot) CycleGamers() {
	log.Printf("cycling gamers")

	tick := time.NewTicker(time.Second * 12).C

	for {
		log.Printf("gamer cycle")

		i, g := b.randomGamer()
		a := g.Activity()

		//nolint:exhaustivestruct
		// discord types are excessive
		data := gateway.UpdateStatusData{
			Activities: &[]discord.Activity{a},
		}

		err := b.State.Gateway.UpdateStatus(data)
		if err != nil {
			log.Printf("update status %d: %s", i, err)
		}

		log.Printf("updated status")

		<-tick
	}
}

// GamerFunc is a func that is a Gamer.
type GamerFunc func() discord.Activity

// Activity simply calls the GamerFunc and returns the returned Activity.
func (f GamerFunc) Activity() discord.Activity {
	log.Printf("calling gamer func")

	return f()
}

// GamerStatic is an Activity that is a Gamer.
type GamerStatic discord.Activity

// Activity simply returns the GamerStatic as an Activity.
func (s GamerStatic) Activity() discord.Activity {
	log.Printf("returning gamer static")

	return discord.Activity(s)
}

// GamerTimer is a Time that is a Gamer.
type GamerTimer time.Time

// String returns the GamerTimer formatted as "since Jan _2 15:04:05".
func (t GamerTimer) String() string {
	return "since " + time.Time(t).Format(time.Stamp)
}

// Activity returns an Activity that describes the duration since the GamerTimer.
func (t GamerTimer) Activity() discord.Activity {
	since := time.Since(time.Time(t)).Round(time.Second)
	log.Printf("timer gamer: %s", since)

	//nolint:exhaustivestruct
	// discord types are excessive
	return discord.Activity{
		Name: fmt.Sprintf("for %s", since),
		Type: discord.GameActivity,
	}
}
