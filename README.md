# wilks

[![Build Status](https://travis-ci.org/KyleBanks/wilks.svg?branch=master)](https://travis-ci.org/KyleBanks/wilks) &nbsp;
[![GoDoc](https://godoc.org/github.com/KyleBanks/wilks?status.svg)](https://godoc.org/github.com/KyleBanks/wilks) &nbsp;
[![Go Report Card](https://goreportcard.com/badge/github.com/KyleBanks/wilks)](https://goreportcard.com/report/github.com/KyleBanks/wilks) &nbsp;
[![Coverage Status](https://coveralls.io/repos/github/KyleBanks/wilks/badge.svg?branch=master)](https://coveralls.io/github/KyleBanks/wilks?branch=master)
 
A Go implementation of the [Wilks Formula](https://en.wikipedia.org/wiki/Wilks_Coefficient) for comparing powerlifters across gender and weight class.

## Install

```sh
$ go get github.com/KyleBanks/wilks
```

## Usage

Simply import the `wilks` package and use the `Score` function:

```go
import (
	"fmt"
	"github.com/KyleBanks/wilks"
)

func main() {
	var isFemale bool
	var bw float64
	var total float64 
	
	score := wilks.Score(isFemale, bw, total)
	fmt.Printf("%.02f\n", score)
}
```

`wilks.Score` requires values in kilograms. If you are using values in pounds, there is a convenience function for conversion:

```go
kg := wilks.PoundsToKilos(lbs)
```

## Command-Line 

This repository also comes with a command-line tool for calculating Wilks scores.

You can download the appropriate binary from the [Releases](https://github.com/KyleBanks/wilks/releases) page, or if you have a working Go installation, you can install like so:

```sh
$ go get github.com/KyleBanks/wilks/cmd/wilks
```

After installing the command-line tool, you can invoke it and supply a `-bw` and `-total` along with a `-gender` of `m` for a male lifter or `f` for a female lifter:

```sh
$ wilks -bw 83 -total 500 -gender m
333.75
$ wilks -bw 52 -total 400 -gender f
498.65
```

By default `wilks` assumes you are supplying values in kilograms. If you prefer to use pounds, supply the `-lbs` flag:

```sh
$ wilks -lbs -bw 180 -total 1100 -gender m
```

You can also view usage using `--help`:

```sh
$ wilks --help
```

## License

```
MIT License

Copyright (c) 2017 Kyle Banks

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```