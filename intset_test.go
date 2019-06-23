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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	tests := []struct {
		in  string
		out []int
	}{
		{
			in:  "1",
			out: []int{1},
		},
		{
			in:  "1-4",
			out: []int{1, 2, 3, 4},
		},
		{
			in:  "2018,2020",
			out: []int{2018, 2020},
		},
		{
			in:  "2010-2015,2017-2019,2022",
			out: []int{2010, 2011, 2012, 2013, 2014, 2015, 2017, 2018, 2019, 2022},
		},
	}
	for _, tt := range tests {
		out, err := Parse(tt.in)
		require.NoError(t, err)
		assert.Equal(t, Set(tt.out), out)
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		in  []int
		out string
	}{
		{
			in:  []int{},
			out: "",
		},
		{
			in:  []int{3, 3, 2, 2, 1},
			out: "1-3",
		},
		{
			in:  []int{1, 1, 1, 5, 5},
			out: "1,5",
		},
		{
			in:  []int{2018},
			out: "2018",
		},
		{
			in:  []int{2018, 2019},
			out: "2018-2019",
		},
		{
			in:  []int{2018, 2019, 2022},
			out: "2018-2019,2022",
		},
		{
			in:  []int{2012, 2013, 2017, 2018, 2019, 2022},
			out: "2012-2013,2017-2019,2022",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.out, Set(tt.in).String())
	}
}
