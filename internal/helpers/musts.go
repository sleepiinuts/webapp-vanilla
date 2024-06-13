package helpers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sleepiinuts/webapp-plain/configs"
)

func MustParseInt(i string) int {
	n, err := strconv.Atoi(i)
	if err != nil {
		panic(fmt.Sprintf("not a valid integer: %s\n", i))
	}
	return n
}

func MustParseTime(t string) time.Time {
	tt, err := time.Parse(configs.DateFormat, t)
	if err != nil {
		panic(fmt.Sprintf("not a valid time: %s\n", t))
	}
	return tt
}
