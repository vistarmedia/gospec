// Copyright © 2009-2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package gospec

import (
	"container/list"
	"fmt"
	"math"
	"testing"
)


func Test__Positive_assertation_failures_are_reported_with_the_positive_message(t *testing.T) {
	log := new(spyErrorLogger)
	
	log.Then(1).Should.Equal(1)
	log.ShouldHaveNoErrors(t)
	
	log.Then(1).Should.Equal(2)
	log.ShouldHaveTheError("Expected '2' but was '1'", t)
}

func Test__Negative_assertation_failures_are_reported_with_the_negative_message(t *testing.T) {
	log := new(spyErrorLogger)
	
	log.Then(1).ShouldNot.Equal(2)
	log.ShouldHaveNoErrors(t)
	
	log.Then(1).ShouldNot.Equal(1)
	log.ShouldHaveTheError("Did not expect '1' but was '1'", t)
}

func Test__Errors_in_asserts_are_reported_with_the_error_message(t *testing.T) {
	log := new(spyErrorLogger)
	
	log.Then(1).Should.BeNear(1.0, 0.001)
	log.ShouldHaveTheError("Expected a float, but was '1' of type 'int'", t)
}


// "Equal" matcher

func Test__String_should_EQUAL_string(t *testing.T) {
	log := new(spyErrorLogger)

	log.Then("apple").Should.Equal("apple")
	log.ShouldHaveNoErrors(t)

	log.Then("apple").Should.Equal("orange")
	log.ShouldHaveTheError("Expected 'orange' but was 'apple'", t)
}

func Test__String_should_NOT_EQUAL_string(t *testing.T) {
	log := new(spyErrorLogger)

	log.Then("apple").ShouldNot.Equal("orange")
	log.ShouldHaveNoErrors(t)

	log.Then("apple").ShouldNot.Equal("apple")
	log.ShouldHaveTheError("Did not expect 'apple' but was 'apple'", t)
}

func Test__Int_should_EQUAL_int(t *testing.T) {
	log := new(spyErrorLogger)

	log.Then(42).Should.Equal(42)
	log.ShouldHaveNoErrors(t)

	log.Then(42).Should.Equal(999)
	log.ShouldHaveTheError("Expected '999' but was '42'", t)
}

func Test__Struct_should_EQUAL_struct(t *testing.T) {
	log := new(spyErrorLogger)

	log.Then(DummyStruct{42, 1}).Should.Equal(DummyStruct{42, 2})
	log.ShouldHaveNoErrors(t)

	log.Then(DummyStruct{42, 1}).Should.Equal(DummyStruct{999, 2})
	log.ShouldHaveTheError("Expected 'DummyStruct999' but was 'DummyStruct42'", t)
}

func Test__Struct_pointer_should_EQUAL_struct_pointer(t *testing.T) {
	log := new(spyErrorLogger)

	log.Then(&DummyStruct{42, 1}).Should.Equal(&DummyStruct{42, 2})
	log.ShouldHaveNoErrors(t)

	log.Then(&DummyStruct{42, 1}).Should.Equal(&DummyStruct{999, 2})
	log.ShouldHaveTheError("Expected 'DummyStruct999' but was 'DummyStruct42'", t)
}


// "Be" matcher

func Test__Object_should_BE_some_expression(t *testing.T) {
	log := new(spyErrorLogger)
	value := 42
	
	log.Then(value).Should.Be(value > 40)
	log.ShouldHaveNoErrors(t)
	
	log.Then(value).Should.Be(value > 999)
	log.ShouldHaveTheError("Criteria not satisfied by '42'", t)
}

func Test__Object_should_NOT_BE_some_expression(t *testing.T) {
	log := new(spyErrorLogger)
	value := 42
	
	log.Then(value).ShouldNot.Be(value < 40)
	log.ShouldHaveNoErrors(t)
	
	log.Then(value).ShouldNot.Be(value < 999)
	log.ShouldHaveTheError("Criteria not satisfied by '42'", t)
}


// "BeNear" matcher

func Test__Float_should_BE_NEAR_float(t *testing.T) {
	log := new(spyErrorLogger)
	value := float64(3.141)
	pi := float64(math.Pi)
	
	log.Then(value).Should.BeNear(pi, 0.001)
	log.ShouldHaveNoErrors(t)
	
	log.Then(value).Should.BeNear(pi, 0.0001)
	log.ShouldHaveTheError(fmt.Sprintf("Expected '%v' ± 0.0001 but was '3.141'", pi), t)
}

func Test__Float_should_NOT_BE_NEAR_float(t *testing.T) {
	log := new(spyErrorLogger)
	value := float64(3.15)
	pi := float64(math.Pi)
	
	log.Then(value).ShouldNot.BeNear(pi, 0.001)
	log.ShouldHaveNoErrors(t)
	
	log.Then(value).ShouldNot.BeNear(pi, 0.01)
	log.ShouldHaveTheError(fmt.Sprintf("Did not expect '%v' ± 0.01 but was '3.15'", pi), t)
}

func Test__Int_should_BE_NEAR_float_IS_NOT_ALLOWED(t *testing.T) {
	log := new(spyErrorLogger)
	
	log.Then(int(3)).Should.BeNear(math.Pi, 0.2)
	log.ShouldHaveTheError("Expected a float, but was '3' of type 'int'", t)
}


// "Contain" matcher

func Test__Array_should_CONTAIN_a_value(t *testing.T) {
	log := new(spyErrorLogger)
	values := [...]string{"one", "two", "three"}
	
	log.Then(values).Should.Contain("one")
	log.Then(values).Should.Contain("two")
	log.Then(values).Should.Contain("three")
	log.ShouldHaveNoErrors(t)
	
	log.Then(values).Should.Contain("four")
	log.ShouldHaveTheError("Expected 'four' to be in '[one two three]' but it was not", t)
}

func Test__Array_should_NOT_CONTAIN_a_value(t *testing.T) {
	log := new(spyErrorLogger)
	values := [...]string{"one", "two", "three"}
	
	log.Then(values).ShouldNot.Contain("four")
	log.ShouldHaveNoErrors(t)
	
	log.Then(values).ShouldNot.Contain("one")
	log.ShouldHaveTheError("Did not expect 'one' to be in '[one two three]' but it was", t)
}

func Test__Iterable_should_CONTAIN_a_value(t *testing.T) {
	log := new(spyErrorLogger)
	values := list.New()
	values.PushBack("one")
	values.PushBack("two")
	values.PushBack("three")
	
	log.Then(values.Iter()).Should.Contain("one")
	log.Then(values.Iter()).Should.Contain("two")
	log.Then(values.Iter()).Should.Contain("three")
	log.ShouldHaveNoErrors(t)
	
	log.Then(values.Iter()).Should.Contain("four")
	log.ShouldHaveTheError("Expected 'four' to be in '[one two three]' but it was not", t)
}

func Test__Iterable_should_NOT_CONTAIN_a_value(t *testing.T) {
	log := new(spyErrorLogger)
	values := list.New()
	values.PushBack("one")
	values.PushBack("two")
	values.PushBack("three")
	
	log.Then(values.Iter()).ShouldNot.Contain("four")
	log.ShouldHaveNoErrors(t)
	
	log.Then(values.Iter()).ShouldNot.Contain("one")
	log.ShouldHaveTheError("Did not expect 'one' to be in '[one two three]' but it was", t)
}


// Utilities

type spyErrorLogger struct {
	failures  int
	lastError *Error
}

func (log *spyErrorLogger) Then(actual interface{}) *MatcherBuilder {
	return newMatcherBuilder(actual, callerLocation(), log)
}

func (log *spyErrorLogger) AddError(error *Error) {
	log.failures++
	log.lastError = error
}

func (log *spyErrorLogger) AddFatalError(error *Error) {
	log.AddError(error)
}

func (log *spyErrorLogger) Reset() {
	log.failures = 0
	log.lastError = nil
}

func (log *spyErrorLogger) ShouldHaveNoErrors(t *testing.T) {
	assertEquals(0, log.failures, t)
	log.Reset()
}

func (log *spyErrorLogger) ShouldHaveTheError(message string, t *testing.T) {
	assertEquals(1, log.failures, t)
	lastMessage := ""
	if log.lastError != nil {
		lastMessage = log.lastError.Message
	}
	assertEquals(message, lastMessage, t)
	log.Reset()
}

