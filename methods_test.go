// Copyright 2012 Aaron Jacobs. All Rights Reserved.
// Author: aaronjjacobs@gmail.com (Aaron Jacobs)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ogletest

import (
	. "github.com/jacobsa/oglematchers"
	"reflect"
	"testing"
)

func TestRegisterMethodsTest(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////////////////

type MethodsTest struct {
}

func init() { RegisterTestSuite(&MethodsTest{}) }

type OneMethodType int
func (x OneMethodType) Foo() {}

////////////////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////////////////

func (t *MethodsTest) NoMethods() {
	type foo int

	methods := getMethodsInSourceOrder(reflect.TypeOf(foo(17)))
	ExpectThat(methods, ElementsAre())
}

func (t *MethodsTest) OneMethod() {
	methods := getMethodsInSourceOrder(reflect.TypeOf(OneMethodType(17)))
	AssertThat(methods, ElementsAre(Any()))

	ExpectEq("Foo", methods[0].Name)
}

func (t *MethodsTest) MultipleMethods() {
	ExpectEq("TODO", "")
}

func (t *MethodsTest) MultipleMethodsOnSingleLine() {
	ExpectEq("TODO", "")
}
