// Copyright (c) 2021 Takayuki NAGATOMI
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package takolabel

import (
	"reflect"
	"testing"
)

func TestParseDelete(t *testing.T) {
	d := Delete{}
	err := d.ParseDelete([]byte(`repositories:
  - some-owner/some-owner-repo-1
  - some-owner/some-owner-repo-2
  - another-owner/another-owner-repo-1
labels:
  - Label 1
  - Label 2
  - Label 3
`))
	if err != nil {
		t.Fatalf("error: %q", err)
	}

	want := DeleteTarget{
		Repositories: Repositories{
			{"some-owner", "some-owner-repo-1"},
			{"some-owner", "some-owner-repo-2"},
			{"another-owner", "another-owner-repo-1"},
		},
		Labels: []string{
			"Label 1",
			"Label 2",
			"Label 3",
		},
	}

	if !reflect.DeepEqual(d.Target, want) {
		t.Errorf("got %v want %v", d.Target, want)
	}
}
