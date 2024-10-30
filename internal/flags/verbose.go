package flags

import "github.com/urfave/cli/v2"

var verbosity = int(0)

func NewVerboseFlag() cli.Flag {
	return &cli.BoolFlag{
		Name:    "verbose",
		Aliases: []string{"v"},
		Count:   &verbosity,
		Usage:   "enable progressively more verbose output",
	}
}

func IsVerbose() bool {
	return verbosity > 0
}

func VerboseLevel() int {
	if verbosity > 5 {
		return 5
	}
	return verbosity
}
