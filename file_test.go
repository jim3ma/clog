// Copyright 2017 Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package clog

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_file_Init(t *testing.T) {
	Convey("Init file logger", t, func() {
		Convey("With mismatched config object", func() {
			err := New(FILE, struct{}{})
			So(err, ShouldNotBeNil)
			_, ok := err.(ErrConfigObject)
			So(ok, ShouldBeTrue)
		})

		Convey("With valid config object", func() {
			So(New(FILE, FileConfig{
				Filename: "test/test.log",
			}), ShouldBeNil)

			Convey("Incorrect level", func() {
				err := New(FILE, FileConfig{
					Level: LEVEL(-1),
				})
				So(err, ShouldNotBeNil)
				_, ok := err.(ErrInvalidLevel)
				So(ok, ShouldBeTrue)
			})

			Convey("Empty file name", func() {
				err := New(FILE, FileConfig{})
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "initFile: OpenFile '': open : no such file or directory")
			})
		})
	})
}
