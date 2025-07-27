package config

import "github.com/jessevdk/go-flags"

type Options struct {
	Config string `short:"c" long:"config" description:"Path for config" required:"true"`
}

var Opt Options
var Parser = flags.NewParser(&Opt, flags.Default)
