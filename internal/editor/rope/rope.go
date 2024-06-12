package rope

import (
	"fmt"
	"unicode/utf8"
)

type Rope struct {
    value   []rune
    weight  int
    length  int
    left    *Rope
    right   *Rope
}

// Constructure
func NewRope(element string) *Rope {
    len := utf8.RuneCountInString(element)
    return &Rope{
        value:  []rune(element),
        weight: len,
        length: len,
    }
}

func (rope *Rope) GetLength() int {
    if rope == nil {
        return 0
    }

    return rope.length
}

func (rope *Rope) GetRune(i int) rune {
    rope.panicIfNil()
    if i < 0 || i >= rope.GetLength() {
        panic("Index is out of range")
    }

    if rope.isLeaf() {
        return rope.value[i - 1]
    } else if i > rope.weight {
        return rope.right.GetRune(i - rope.weight)
    } else {
        return rope.left.GetRune(i)
    }
}

// Merges two ropes
func (rope *Rope) Concat(other *Rope) *Rope {
    if rope == nil {
        return other
    }

    if other == nil {
        return rope
    }

    return &Rope{
        weight: rope.GetLength(),
        length: rope.GetLength() + other.GetLength(),
        left: rope,
        right: other,
    }
}

// Spliting the rope at index (i) and returning two rope strings
func (rope *Rope) Split(i int) (rope1, rope2 *Rope) {
    a := NewRope("test")
    b := NewRope("test")
    return a, b
}

// Inserting a string at index (i) posision
func (rope *Rope) Insert(i int, str string) *Rope {
    return NewRope("HelloWorld")
}

// Delete at posision i for a given length
func (rope *Rope) Delete(i int, length int) *Rope{
    return NewRope("HelloWorld")
}

func (rope *Rope) String() string {
    // Make rope return the string inside the rope
    return ""
}

func (rope *Rope) isLeaf() bool {
    return rope.left == nil
}

func (rope *Rope) panicIfNil() {
    if rope == nil {
        panic(fmt.Sprintf("Operation not permitted on empty rope"))
    }
}
