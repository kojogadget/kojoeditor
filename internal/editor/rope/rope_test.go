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

    if rope.getLength() != 4 {
        t.Error("Error in getLength: ", rope.getLength(), " != 4")
    }

    rope = NewRope("another")

    if rope.getLength() != 7 {
        t.Error("Error in getLength: ", rope.getLength(), " != 7")
    }
}

func TestGetHeight(t *testing.T) {
    rope := NewRope("test")

    if rope.getHeight() != 1 {
        t.Error("Error in getLength: ", rope.getHeight(), " != 1")
    }

    newRope := NewRope("another")
    highRope := rope.Concat(newRope)

    if highRope.getHeight() != 2 {
        t.Error("Error in getLength: ", highRope.getHeight(), " != 2")
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

    if rope.GetRune(1) != 'e' {
        t.Errorf("expected rune %v, got %v", 'e', rope.GetRune(2))
    }

    if rope.GetRune(4) != 'i' {
        t.Errorf("expected rune %v, got %v", 'i', rope.GetRune(5))
    }
}

func TestConcat(t *testing.T) {
    strA := "test"
    strB := " other"
    ropeA := NewRope(strA)
    ropeB := NewRope(strB)
    var ropeC *Rope

    rope := ropeA.Concat(ropeB)

    if rope.String() != "test other" {
        t.Errorf("expected string %v, got %v", "test other", rope.String())
    }

    rope = ropeA.Concat(nil)
    
    if rope.String() != strA {
        t.Errorf("expected string %v, got %v", strA, rope.String())
    }

    rope = ropeC.Concat(ropeB)
    
    if rope.String() != strB {
        t.Errorf("expected string %v, got %v", strB, rope.String())
    }
}

func TestConcatBalanced(t *testing.T) {
    var rope1 *Rope
    var rope2 *Rope

    for i := 0; i < 10; i++ {
        rope1 = NewRope("hello")
        rope2 = NewRope("world")

        concatenated := rope1.Concat(rope2)

        if concatenated.getBalanceFactor() < -1 || concatenated.getBalanceFactor() > 1 {
            t.Errorf("expected balanced rope, got unbalanced rope with balance factor %d", concatenated.getBalanceFactor())
        }
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
