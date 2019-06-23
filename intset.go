// Copyright (c) 2019 Glib Smaga
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package intset

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// Set ...
type Set []int

// Add ...
func (s *Set) Add(i int) {
	*s = append(*s, i)
}

// String returns the short hunman readable representation of the entire set.
// Multipe consecutive ints get aggregated together into one range, i.e.
// `{1,2,3,4,5}` gets represented as `1-5`.
func (s Set) String() string {
	if len(s) == 0 {
		return ""
	}

	sort.Ints(s)

	var (
		rangeStart *int = nil
		prev            = math.MaxInt64

		buf bytes.Buffer
	)

	w := func(i int) { fmt.Fprint(&buf, i) }

	// finalize the previous range
	fin := func(i int) {
		if i > *rangeStart {
			fmt.Fprint(&buf, "-")
		}
		if i != *rangeStart {
			w(i)
		}
	}

	// start a new range
	start := func(i int) {
		rangeStart = &i
		w(i)
	}

	for _, i := range s {
		// start a new year range if there isn't one already open
		if rangeStart == nil {
			start(i)
		}

		// more than one int gap between the last known int in the range and this
		// current -- close our the previous range and start a new one
		if i-prev > 1 {
			fin(prev)
			fmt.Fprint(&buf, ",")
			start(i)
		}

		prev = i
	}

	// close out the last range
	fin(prev)

	return buf.String()
}

// Parse a human readable format and construct an underlying set, i.e.
// "2012-2014,2020" results in []int{2012,2013,2014,2020}.
func Parse(str string) (Set, error) {
	s := Set{}

	for _, chunk := range strings.Split(str, ",") {
		// ignore whitespace in the beginning and end of each chunk
		chunk = strings.TrimSpace(chunk)

		if strings.Contains(chunk, "-") {
			span := strings.Split(chunk, "-")
			if len(span) > 2 {
				return nil, fmt.Errorf("invalid int span: %v", span)
			}

			start, err := strconv.Atoi(span[0])
			if err != nil {
				return nil, fmt.Errorf("could not convert %v to int", span[0])
			}
			end, err := strconv.Atoi(span[1])
			if err != nil {
				return nil, fmt.Errorf("could not convert %v to int", span[1])
			}

			for i := start; i <= end; i++ {
				s.Add(i)
			}

			continue
		}

		year, err := strconv.Atoi(chunk)
		if err != nil {
			return nil, fmt.Errorf("could not convert %v to int", chunk)
		}

		s.Add(year)
	}

	return s, nil
}

// Must is a helper when error handling is not desired and panic is preferred.
func Must(r Set, err error) Set {
	if err != nil {
		panic(err)
	}
	return r
}
