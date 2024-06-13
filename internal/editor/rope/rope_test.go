package rope

import (
	"reflect"
	"testing"
	"unicode/utf8"
)

func TestNewRope(t *testing.T) {
    str := "test"
    expectedValue := []rune(str)
    expectedWeight := utf8.RuneCountInString(str)
    expectedLength := utf8.RuneCountInString(str)
    rope := NewRope(str)

    if !reflect.DeepEqual(rope.value, expectedValue) {
        t.Errorf("expected value %v, got %v", expectedValue, rope.value)
    }

    if !reflect.DeepEqual(rope.weight, expectedWeight) {
        t.Errorf("expected weight %v, got %v", expectedWeight, rope.weight)
    }

    if !reflect.DeepEqual(rope.length, expectedLength) {
        t.Errorf("expected value %v, got %v", expectedLength, rope.length)
    }
}

func TestGetLength(t *testing.T) {
    rope := NewRope("test")

    if rope.GetLength() != 4 {
        t.Error("Error in GetLength: ", rope.GetLength(), " != 4")
    }

    rope = NewRope("another")

    if rope.GetLength() != 7 {
        t.Error("Error in GetLength: ", rope.GetLength(), " != 7")
    }
}

func TestIsLeaf(t *testing.T) {
    rope := &Rope{}
    if !rope.isLeaf() {
        t.Errorf("expected true, got false")
    }

    rope.left = &Rope{}
    if rope.isLeaf() {
        t.Errorf("expected false, got true")
    }
}

func TestPanicIfNil(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("expected panic, got none")
        }
    }()

    var rope *Rope
    rope.panicIfNil() 
}

func TestPanicIfNilNoPanic(t *testing.T) {
    rope := &Rope{}
    defer func() {
        if r := recover(); r != nil {
            t.Errorf("expected no panic, got panic")
        }
    }()

    rope.panicIfNil()
}

func TestGetRune(t *testing.T) {
    rope := NewRope("testing")

    if rope.GetRune(2) != 'e' {
        t.Errorf("expected rune %v, got %v", 'e', rope.GetRune(2))
    }

    if rope.GetRune(5) != 'i' {
        t.Errorf("expected rune %v, got %v", 'e', rope.GetRune(5))
    }
}

func TestConcat(t *testing.T) {
    strA := "test"
    expectedValueA := []rune(strA)
    expectedWeightA := utf8.RuneCountInString(strA)
    expectedLengthA := utf8.RuneCountInString(strA)
    strB := "other"
    expectedValueB := []rune(strB)
    expectedWeightB := utf8.RuneCountInString(strB)
    expectedLengthB := utf8.RuneCountInString(strB)
    ropeA := NewRope(strA)
    ropeB := NewRope(strB)

    rope := ropeA.Concat(ropeB)

    if !reflect.DeepEqual(rope.weight, expectedLengthA) {
        t.Errorf("expected weight %v, got %v", expectedLengthA, rope.weight)
    }

    if !reflect.DeepEqual(rope.GetLength(), expectedLengthA + expectedLengthB) {
        t.Errorf("expected weight %v, got %v", expectedLengthA + expectedLengthB, rope.GetLength())
    }

    if !reflect.DeepEqual(rope.left.value, expectedValueA) {
        t.Errorf("expected value %v, got %v", expectedValueA, rope.left.value)
    }

    if !reflect.DeepEqual(rope.left.weight, expectedWeightA) {
        t.Errorf("expected weight %v, got %v", expectedWeightA, rope.left.weight)
    }

    if !reflect.DeepEqual(rope.left.length, expectedLengthA) {
        t.Errorf("expected value %v, got %v", expectedLengthA, rope.left.length)
    }

    if !reflect.DeepEqual(rope.right.value, expectedValueB) {
        t.Errorf("expected value %v, got %v", expectedValueB, rope.right.value)
    }

    if !reflect.DeepEqual(rope.right.weight, expectedWeightB) {
        t.Errorf("expected weight %v, got %v", expectedWeightB, rope.right.weight)
    }

    if !reflect.DeepEqual(rope.right.length, expectedLengthB) {
        t.Errorf("expected value %v, got %v", expectedLengthB, rope.right.length)
    }
}

func TestSplit(t *testing.T) {
    rope := NewRope("test")
    leftRope, rightRope := rope.Split(2)

    if !reflect.DeepEqual(leftRope.value, []rune{'t', 'e'}) {
        t.Errorf("expected weight %v, got %v", []rune{'t', 'e'}, leftRope.value)
    }

    if !reflect.DeepEqual(rightRope.value, []rune{'s', 't'}) {
        t.Errorf("expected value %v, got %v", []rune{'s', 't'}, rightRope.value)
    }
}

func TestInsert(t *testing.T) {
    rope := NewRope("test")
    addRope := rope.Insert(rope.GetLength(), "ing")

    if !reflect.DeepEqual(addRope.length, 7) {
        t.Errorf("expected weight %v, got %v", 7, addRope.length)
    }

    if rope.String() != "test" {
        t.Errorf("expected weight %v, got %v", "test", rope.String())
    }

    if addRope.String() != "testing" {
        t.Errorf("expected weight %v, got %v", "testing", addRope.String())
    }
}

func TestString(t *testing.T) {
    str := "test"
    rope := NewRope(str)
    
    if rope.String() != str {
        t.Errorf("expected weight %v, got %v", str, rope.String())
    }
}
