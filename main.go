package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	appName  = "filewatcher"
	usageMsg = `%s watch to exist a file at regular intervals.

A loop of watching does not stop when find a file.

Usage:

	%s <file> <interval> [<loop>]

Each arguments are:

	<file>
		a watching file.
		wildcards can use to file name.
	<interval>
		a watching interval time.
		the value range is from 1(ms) to 60000(one hour).
	[<loop>]
		a watching loop count.
		the value range is none or 0 and from 1 to 1000.
		none or 0 is infinity loop.
`
)

func init() {
	log.SetFlags(log.Lmicroseconds)
}

func main() {
	ln := len(os.Args)
	if ln != 3 && ln != 4 {
		fmt.Fprintf(os.Stderr, usageMsg, appName, appName)
		os.Exit(1)
	}

	fn := os.Args[1]

	ms, err := strconv.Atoi(os.Args[2])
	if err != err {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	if ms < 1 || ms > 60*1000 {
		fmt.Fprintf(os.Stderr, usageMsg, appName, appName)
		os.Exit(1)
	}

	lp := 0
	if ln == 4 {
		lp, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		if lp > 1000 {
			fmt.Fprintf(os.Stderr, usageMsg, appName, appName)
			os.Exit(1)
		}
		if lp < 0 {
			lp = 0
		}
	}

	// file finding function.
	find := func(fn string) {
		if fs, err := filepath.Glob(fn); err != nil {
			log.Println(err)
		} else if len(fs) == 0 {
			log.Println("not found file or directory")
		} else {
			log.Println("found file or directory")
		}
	}

	find(fn)
	if lp != 1 {
		tck := time.Tick(time.Duration(ms) * time.Millisecond)
		if lp == 0 {
			// infinity loop.
			for range tck {
				find(fn)
			}
		} else {
			// limited loop.
			cnt := 2
			for range tck {
				find(fn)
				if cnt >= lp {
					break
				}
				cnt++
			}
		}
	}
}
