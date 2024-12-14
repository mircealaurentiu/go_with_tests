package generics

import "testing"

func TestAssertFunction(t *testing.T) {

	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "notHello")
	})

	//AssertEqual(t, 1, "1") // this does not work, it is supposed to not work

}

func TestStack(t *testing.T) {
	t.Run("integers stack", func(t *testing.T) {

		myStackOfInt := new(StackOfInts)

		//check if stack is empty
		AssertTrue(t, myStackOfInt.IsEmpty())

		// add item, then check it is not empty
		myStackOfInt.Push(11)
		AssertFalse(t, myStackOfInt.IsEmpty())

		// add another item, then pop it
		myStackOfInt.Push(22)
		popValue, _ := myStackOfInt.Pop()
		AssertEqual(t, popValue, 22)

		popValue, _ = myStackOfInt.Pop()
		AssertEqual(t, popValue, 11)

		AssertTrue(t, myStackOfInt.IsEmpty())
	})

	t.Run("string stack", func(t *testing.T) {

		myStackOfInt := new(StackOfStrings)

		//check if stack is empty
		AssertTrue(t, myStackOfInt.IsEmpty())

		// add item, then check it is not empty
		myStackOfInt.Push("11")
		AssertFalse(t, myStackOfInt.IsEmpty())

		// add another item, then pop it
		myStackOfInt.Push("22")
		popValue, _ := myStackOfInt.Pop()
		AssertEqual(t, popValue, "22")

		popValue, _ = myStackOfInt.Pop()
		AssertEqual(t, popValue, "11")

		AssertTrue(t, myStackOfInt.IsEmpty())
	})

	t.Run("integer stack with generics", func(t *testing.T) {
		myStackOfInt := new(Stack[int])

		//check if stack is empty
		AssertTrue(t, myStackOfInt.IsEmpty())

		// add item, then check it is not empty
		myStackOfInt.Push(11)
		AssertFalse(t, myStackOfInt.IsEmpty())

		// add another item, then pop it
		myStackOfInt.Push(22)
		popValue, _ := myStackOfInt.Pop()
		AssertEqual(t, popValue, 22)

		popValue, _ = myStackOfInt.Pop()
		AssertEqual(t, popValue, 11)

		AssertTrue(t, myStackOfInt.IsEmpty())
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {

	t.Helper()
	if got != want {
		t.Errorf("\ngot  %+v \nwant %+v", got, want)
	}

}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {

	t.Helper()
	if got == want {
		t.Errorf("\ngot  %+v \nwant %+v", got, want)
	}

}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("\ngot  %v \nwant true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("\ngot  %v \nwant false", got)
	}
}
