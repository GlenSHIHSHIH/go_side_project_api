package cmd

import "github.com/urfave/cli/v2"

func BuildUpFlag(flags ...[]cli.Flag) []cli.Flag {
	var f []cli.Flag

	for _, value := range flags {
		f = append(f, value...)
	}

	return f
}
