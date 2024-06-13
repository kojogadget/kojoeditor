package rope

import (
	"fmt"
	"unicode/utf8"
)

type Rope struct {
    value   []rune
    weight  int
    length  int
    height  int
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
        height: 1,
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

    maxLen := max(rope.GetLength(), other.GetLength())
    minLen := min(rope.GetLength(), other.GetLength())

    if maxLen - minLen <= 1 {
        return &Rope{
            weight: rope.GetLength(),
            length: rope.GetLength() + other.GetLength(),
            left: rope,
            right: other,
            height: maxLen + 1,
        }
    } else if rope.height == maxLen {
        leftRope := rope.left
        rightRope := rope.right.Concat(other)

        return leftRope.Concat(rightRope)
    } else {
        leftRope := rope.Concat(other.left)
        rightRope := other.right

        return leftRope.Concat(rightRope)
    }
}

// Spliting the rope at index (i) and returning two rope strings
func (rope *Rope) Split(i int) (ropeA, ropeB *Rope) {
    rope.panicIfNil()
    if i < 0 || i > rope.GetLength() {
        panic("Index is out of range")
    }

    var leftRope *Rope
    var rightRope *Rope

    if i == rope.weight {
        if rope.isLeaf() {
            leftRope = rope
        } else {
            leftRope = rope.left
        }
        return leftRope, rope.right

    } else if i > rope.weight {
        leftRope, rightRope := rope.right.Split(i - rope.weight)
        return rope.left.Concat(leftRope), rightRope

    } else {
        if rope.isLeaf() {
            leftRope = &Rope{
                value: rope.value[0:i],
                weight: i,
                length: i,
            }
            rightRope = &Rope{
                value: rope.value[i:rope.weight],
                weight: rope.GetLength() - i,
                length: rope.GetLength() - i,
            }

            return leftRope, rightRope

        } else {
            leftRope, rightRope := rope.left.Split(i)
            return leftRope, rightRope.Concat(rope.right)
        }
    }
}

// Inserting a string at index (i) posision
func (rope *Rope) Insert(i int, str string) *Rope {
    addRope := NewRope(str)
    leftRope, rightRope := rope.Split(i)

    return leftRope.Concat(addRope).Concat(rightRope)
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

