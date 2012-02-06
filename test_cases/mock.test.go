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
	"github.com/jacobsa/oglemock"
	"github.com/jacobsa/ogletest/test_cases/mock_image"
	"image/color"
	"testing"
)

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

type MockTest struct {
	controller oglemock.Controller
	image mock_image.MockImage
}

func init()                     { RegisterTestSuite(&MockTest{}) }
func TestOgletest(t *testing.T) { RunTests(t) }

func (t *MockTest) SetUp(i *TestInfo) {
	t.controller = i.MockController
	t.image = mock_image.NewMockImage(t.controller, "some mock image")
}

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func (t *MockTest) ExpectationSatisfied() {
	// TODO(jacobsa): Replace this hand-spun expectation with a call to a more
	// convenient ExpectCall function when one is available. See issue #8.
	t.controller.ExpectCall(
		t.image,
		"At",
		"blah_test.go",
		117)(11, GreaterThan(19)).WillOnce(
			oglemock.Return(color.Gray{0}))

	ExpectThat(t.image.At(11, 23), IdenticalTo(color.Gray{0}))
}

func (t *MockTest) MockExpectationNotSatisfied() {
	// TODO(jacobsa): Replace this hand-spun expectation with a call to a more
	// convenient ExpectCall function when one is available. See issue #8.
	t.controller.ExpectCall(
		t.image,
		"At",
		"blah_test.go",
		117)(11, GreaterThan(19))
}

func (t *MockTest) ExpectCallForUnknownMethod() {
	// TODO(jacobsa): Replace this hand-spun expectation with a call to a more
	// convenient ExpectCall function when one is available. See issue #8.
	t.controller.ExpectCall(
		t.image,
		"FooBar",
		"blah_test.go",
		117)(11)
}

func (t *MockTest) UnexpectedCall() {
	t.image.At(11, 23)
}
