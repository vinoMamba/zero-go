package main

import (
	"fmt"

	"github.com/vinoMamba/zero/pkg/config"
)

func main() {
	// init config
	conf := config.NewConfig()

	fmt.Println(conf.GetString("env"))
}
