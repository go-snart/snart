package help

import (
	"fmt"
	"runtime"
	"time"

	dg "github.com/bwmarrin/discordgo"
	"github.com/pbnjay/memory"

	"github.com/go-snart/snart/log"
	"github.com/go-snart/snart/route"
)

// About gives statistics about the given Bot.
func (h *Help) About(ctx *route.Ctx) error {
	err := ctx.Flag.Parse()
	if err != nil {
		err = fmt.Errorf("flag parse: %w", err)
		log.Warn.Println(err)

		return err
	}

	rep := ctx.Reply()
	rep.Embed = &dg.MessageEmbed{
		Fields: []*dg.MessageEmbedField{
			h.aboutGeneral(),
			h.aboutVersions(),
			h.aboutRuntime(),
		},
	}

	return rep.Send()
}

func (h *Help) aboutGeneral() *dg.MessageEmbedField {
	guilds := len(h.Session.State.Guilds)
	chans := 0
	users := 0
	routes := len(h.Handler.Routes)
	uptime := time.Since(h.Startup).Truncate(time.Second)

	for _, g := range h.Session.State.Guilds {
		chans += len(g.Channels)
		users += len(g.Members)
	}

	return &dg.MessageEmbedField{
		Name: "General",
		Value: fmt.Sprintf(
			"Users: `%d`\nChannels: `%d`\nGuilds: `%d`\nRoutes: `%d`\nUptime: `%s`",
			users, chans, guilds, routes, uptime,
		),
		Inline: false,
	}
}

func (h *Help) aboutVersions() *dg.MessageEmbedField {
	snartV := "unknown"

	return &dg.MessageEmbedField{
		Name: "Version",
		Value: fmt.Sprintf(
			"Go: `%s`\nSnart: `%s`\nDiscordGo: `%s`\nAPI: `%s`",
			runtime.Version(), snartV, dg.VERSION, dg.APIVersion,
		),
		Inline: false,
	}
}

func (h *Help) aboutRuntime() *dg.MessageEmbedField {
	cpus := runtime.NumCPU()
	ram := int(memory.TotalMemory())

	ms := runtime.MemStats{}
	runtime.ReadMemStats(&ms)

	mem := int(ms.Sys)

	return &dg.MessageEmbedField{
		Name: "Runtime",
		Value: fmt.Sprintf(
			"CPUs: `%d`\nMemory: `%s`\nRAM: `%s`",
			cpus, iec(mem), iec(ram),
		),
		Inline: false,
	}
}

func iec(b int) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}

	div, exp := int64(unit), 0

	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf(
		"%.1f %ciB",
		float64(b)/float64(div),
		"KMGTPE"[exp],
	)
}
