package main

import (
	"fmt"
	"os"
	"time"

	"github.com/araddon/dateparse"
	"github.com/pkg/errors"
	dfmt "github.com/vigneshuvi/GoDateFormat"
	"gopkg.in/urfave/cli.v2"
)

var (
	version string

	OutFmtFlag = &cli.StringFlag{
		Name:    "output-format",
		Aliases: []string{"out", "of"},
		Value:   "yyyy-mm-dd HH:MM:ss Z",
		Usage:   "output format",
	}
	InFmtFlag = &cli.StringFlag{
		Name:    "input-format",
		Aliases: []string{"in", "if"},
		Usage:   "input format",
	}
)

func main() {
	app := cli.App{
		Name:      "timestamp",
		Copyright: "Copyright © 2018",
		Usage:     "Convert date / timestamps to various formats",
		UsageText: "timestamp [options] [arguments...]",
		Version:   version,
		Flags:     []cli.Flag{InFmtFlag, OutFmtFlag},
		Action:    do,
	}

	app.Run(os.Args)
}

func do(context *cli.Context) error {
	src := context.Args().First()
	infmt := context.String(InFmtFlag.Name)

	var (
		t   time.Time = time.Now()
		err error
	)

	if len(src) > 0 {
		if len(infmt) < 1 {
			t, err = dateparse.ParseStrict(src)
			if err != nil {
				return cli.Exit(errors.Wrap(err, "unable to parse date/time"), 1)
			}
		} else {
			t, err = time.Parse(dfmt.ConvertFormat(infmt), src)
			if err != nil {
				return cli.Exit(errors.Wrap(err, "unable to parse date/time"), 1)
			}
		}
	}

	outfmt := context.String(OutFmtFlag.Name)

	switch outfmt {
	case "unix":
		fmt.Println(t.Unix())

	case "unix.nano":
		fmt.Println(t.UnixNano())

	default:
		fmt.Println(t.Format(dfmt.ConvertFormat(outfmt)))
	}

	return nil
}
