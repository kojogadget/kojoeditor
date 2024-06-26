package rope

import "unicode/utf8"


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

func (rope *Rope) GetRune(i int) rune {
    rope.panicIfNil()
    if i < 0 || i >= rope.getLength() {
        panic("Index is out of range")
    }

    if rope.isLeaf() {
        return rope.value[i]
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

    newRope := &Rope{
        left: rope,
        right: other,
        weight: rope.getLength(),
        length: rope.getLength() + other.getLength(),
        height: max(rope.getHeight(), other.getHeight()) + 1,
    }

    return newRope.balance()
}

// Spliting the rope at index (i) and returning two rope strings
func (rope *Rope) Split(i int) (ropeA, ropeB *Rope) {
    rope.panicIfNil()
    if i < 0 || i > rope.getLength() {
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
        leftRope, rightRope = rope.right.Split(i - rope.weight)
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
                weight: rope.getLength() - i,
                length: rope.getLength() - i,
            }

            return leftRope, rightRope

        } else {
            leftRope, rightRope = rope.left.Split(i)
            return leftRope, rightRope.Concat(rope.right)
        }
    }
}

// Inserting a string at index (i) posision
func (rope *Rope) Insert(i int, str string) *Rope {
    addRope := NewRope(str)
    leftRope, rightRope := rope.Split(i)

    rope = leftRope.Concat(addRope).Concat(rightRope)
    return rope
}

// Delete at posision i for a given length
func (rope *Rope) Delete(i int, length int) *Rope{
    ropeA, ropeB := rope.Split(i)
    _, ropeC := ropeB.Split(length)
    return ropeA.Concat(ropeC)
}

func (rope *Rope) String() string {
    str := rope.content(0, rope.getLength())
    return string(str)
}

func (rope *Rope) SubString(from, to int) string {
    str := rope.content(from, to)
    return string(str)
}

