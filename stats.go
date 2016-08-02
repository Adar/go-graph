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
	"sync"
	"time"
)

type Stats struct {
	At    time.Time
	Value float64
}

type StatsRing struct {
	sync.Mutex
	Values []Stats
	Head   int
}

func NewStatsRing(n int) *StatsRing {
	return &StatsRing{Values: make([]Stats, n)}
}

func (r *StatsRing) Add(s Stats) {
	r.Lock()
	defer r.Unlock()
	r.Values[r.Head] = s
	r.Head = (r.Head + 1) % len(r.Values)
}

func (r *StatsRing) Entries() []Stats {
	r.Lock()
	defer r.Unlock()
	s := make([]Stats, 0, len(r.Values))
	pos := r.Head
	for {
		pos = (pos + 1) % len(r.Values)
		if pos == r.Head {
			return s
		}
		if !r.Values[pos].At.IsZero() {
			s = append(s, r.Values[pos])
		}
	}
}

func (r *StatsRing) GetStats() *StatsRing {
	return r
}

func (r *StatsRing) GetTitle() string {
	return *titleFlag
}
