package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/urfave/cli/v2"
)

var (
	// wg        sync.WaitGroup
	shopeeId  string
	skipCount string
)

type ActionFunc func(shopeeId, skipCount int) error

func SetShopeeDataByCli(actionFn ActionFunc) *cli.App {
	app := &cli.App{
		Name:  "Setting config",
		Usage: "Setting shopee's Id and page skip count",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "shopeeId",
				Usage:       "shopee id",
				Aliases:     []string{"Id", "I"},
				Value:       "32286362",
				Destination: &shopeeId,
			},
			&cli.StringFlag{
				Name:        "skipCount",
				Usage:       "page skip count",
				Aliases:     []string{"Skip", "S"},
				Value:       "0",
				Destination: &skipCount,
			},
		},
		Action: func(c *cli.Context) error {

			id, err := strconv.Atoi(shopeeId)
			if err != nil {
				log.Fatal(err)
			}

			skip, err := strconv.Atoi(skipCount)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("shopeeId:%v \n", id)
			fmt.Printf("skipCount:%v \n", skip)

			return actionFn(id, skip)
			// return nil
			// return runAction(id, skip)
		},
	}

	return app
}
