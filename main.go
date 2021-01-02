package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/araddon/dateparse"
	"github.com/pkg/errors"
	cli "github.com/urfave/cli/v2"
	dfmt "github.com/vigneshuvi/GoDateFormat"
)

var (
	version  string
	compiled string = fmt.Sprint(time.Now().Unix())
	githash  string

	outFmtFlag = &cli.StringFlag{
		Name:    "output-format",
		Aliases: []string{"out", "of"},
		Value:   "yyyy-mm-dd HH:MM:ss Z",
		Usage:   "output format",
	}
	inFmtFlag = &cli.StringFlag{
		Name:    "input-format",
		Aliases: []string{"in", "if"},
		Usage:   "input format",
	}
)

func main() {
	ct, err := strconv.ParseInt(compiled, 0, 0)
	if err != nil {
		panic(err)
	}

	app := cli.App{
		Name:      "timestamp",
		Copyright: fmt.Sprintf("Copyright © 2018-%s Elliott Polk", time.Now().Format("2006")),
		Version:   fmt.Sprintf("%s | compiled %s | commit %s", version, time.Unix(ct, -1).Format(time.RFC3339), githash),
		Compiled:  time.Unix(ct, -1),
		Usage:     "Convert date / timestamps to various formats",
		UsageText: "timestamp [options] [arguments...]",
		Flags: []cli.Flag{
			inFmtFlag,
			outFmtFlag,
		},
		Action: func(context *cli.Context) error {
			src := context.Args().First()
			infmt := context.String(inFmtFlag.Name)

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

			outfmt := context.String(outFmtFlag.Name)

			switch outfmt {
			case "unix":
				fmt.Println(t.Unix())

			case "unix.milli":
				fmt.Println(t.Unix() * 1000)

			case "unix.nano":
				fmt.Println(t.UnixNano())

			default:
				fmt.Println(t.Format(dfmt.ConvertFormat(outfmt)))
			}

			return nil
		},
	}

	app.Run(os.Args)
}
