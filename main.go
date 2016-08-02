/*
go-graph - Graphs from shell STDIN.

Copyright (c) 2016 Christian Senkowski

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const (
	VERSION = "0.1"
	// default refresh interval in seconds
	DEFAULT_REFRESH = 5
	// for 24 hours
	HISTORY_LENGTH = 60 * 60 * 24 / DEFAULT_REFRESH
)

var (
	allStats  *StatsRing
	portFlag  = flag.Int("port", 8080, "The listening port.")
	titleFlag = flag.String("title", "Title", "The graph title.")
)

func usage(code int) {
	fmt.Printf(
		`go-graph %s - (c) 2016 Christian Senkowski - MIT Licensed - http://e-cs.co/

After invoking, web UI will be available on http://localhost:8080/. Stats will
be collected every 5 seconds and graphs will refresh every 10 seconds. Graphs
will show 24 hours of history.
`, VERSION)
	os.Exit(code)
}

func main() {

	flag.Parse()

	if len(flag.Args()) < 0 {
		usage(1)
	}

	log.SetPrefix("go-graph: ")
	log.SetFlags(0)

	allStats = NewStatsRing(HISTORY_LENGTH)
	// start open input reader
	go doHost()

	// start the web server
	go startWeb()

	// wait for ^C
	ch := make(chan os.Signal, 1)
	signal.Notify(ch)
	for s := range ch {
		if s == syscall.SIGTERM || s == os.Interrupt {
			break
		} else {
		}
	}
	signal.Stop(ch)
	close(ch)
}

func doHost() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			if ival, err := strconv.ParseFloat(text, 64); err == nil {
				ival = ival / 1024 / 1024
				stats := Stats{At: time.Now()}
				if ival > 1 {
					stats.Value = ival
					allStats.Add(stats)
					fmt.Println(ival)
				}
			}
		}
	}
}
