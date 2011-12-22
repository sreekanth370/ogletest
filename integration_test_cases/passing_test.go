// Copyright 2011 Aaron Jacobs. All Rights Reserved.
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

package oglematchers_test

import (
	. "github.com/jacobsa/oglematchers"
	. "github.com/jacobsa/ogletest"
	"testing"
)

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

type PassingTest struct {
}

func init()                     { RegisterTestSuite(&PassingTest{}) }
func TestOgletest(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func (t *PassingTest) EmptyTestMethod() {
}

func (t *PassingTest) SuccessfullMatches() {
	ExpectThat(17, Equals(17.0))
	ExpectThat(16.9, LessThan(17))
	ExpectThat("taco", HasSubstr("ac"))
}

func (t *PassingTest) ExpectationAliases() {
	ExpectEq(17, 17.0)

	ExpectLe(17, 17.0)
	ExpectLe(17, 18.0)
	ExpectLt(17, 18.0)

	ExpectGe(17, 17.0)
	ExpectGe(17, 16.0)
	ExpectGt(17, 16.0)

	ExpectNe(17, 18.0)

	ExpectTrue(true)
	ExpectFalse(false)
}
