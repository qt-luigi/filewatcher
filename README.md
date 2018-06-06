# filewatcher

filewatcher watch to exist a file at regular intervals.

A loop of watching does not stop when find a file.

## Installation

When you have installed the Go, Please execute following `go get` command:

```sh
go get -u github.com/qt-luigi/filewatcher
```

## Usage

```sh
$ filewatcher 
filewatcher watch to exist a file at regular intervals.

A loop of watching does not stop when find a file.

Usage:

	filewatcher <file> <interval> [<loop>]

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
```

