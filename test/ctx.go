package test

import (
	"strings"

	"github.com/go-snart/snart/logs"
	"github.com/go-snart/snart/route"
)

var (
	CtxRouter  = Router()
	CtxPrefix  = Prefix()
	CtxSession = Session()
)

func Ctx(content string) *route.Ctx {
	logs.Info.Println("enter")
	defer logs.Info.Println("exit")

	return CtxRouter.Ctx(
		CtxPrefix,
		CtxSession,
		Message(content),
		strings.Split(content, "\n")[0],
	)
}
