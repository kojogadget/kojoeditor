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
