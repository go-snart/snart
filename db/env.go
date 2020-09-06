package db

import (
	"os"
	"strconv"
	"strings"

	"github.com/go-snart/snart/log"
)

// EnvStrings gets strings for the given type from env vars in the form TYP or TYP_N.
// It will only check N in order above 1 (eg TYP, TYP_1, TYP_2). Skipping an N will end the search.
func EnvStrings(name, typ string) []string {
	strs := []string(nil)

	env := strings.ToUpper(name + "_" + typ)

	for n := 0; ; n++ {
		env := env

		if n > 0 {
			env += "_" + strconv.Itoa(n)
		}

		log.Info.Printf("checking env %s...\n", env)

		str, ok := os.LookupEnv(env)
		if !ok {
			return strs
		}

		strs = append(strs, str)
	}
}
