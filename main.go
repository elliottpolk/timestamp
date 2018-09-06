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
		Copyright: "Copyright © 2018",
		Usage:     "Convert date / timestamps to various formats",
		Version:   version,
		Flags:     []cli.Flag{OutFmtFlag, InFmtFlag},
		Action:    do,
	}

	app.Run(os.Args)
}

func do(context *cli.Context) error {
	src := context.Args().First()
	if len(src) < 1 {
		return cli.Exit(errors.New("a valid date/time must be provided"), 1)
	}

	infmt := context.String(InFmtFlag.Names()[0])

	var (
		t   time.Time
		err error
	)
	if len(infmt) < 1 {
		t, err = dateparse.ParseAny(src)
		if err != nil {
			return cli.Exit(errors.Wrap(err, "unable to parse date/time"), 1)
		}
	}

	outfmt := context.String(OutFmtFlag.Names()[0])
	fmt.Println(t.Format(dfmt.ConvertFormat(outfmt)))

	return nil
}