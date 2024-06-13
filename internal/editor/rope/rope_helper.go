package rope

import "fmt"

func (rope *Rope) isLeaf() bool {
    return rope.left == nil
}

func (rope *Rope) getLength() int {
    if rope == nil {
        return 0
    }

    return rope.length
}

func (rope *Rope) getHeight() int {
    if rope == nil {
        return 0
    }

    return rope.height
}

func (rope *Rope) balance() *Rope {
    if rope == nil {
        return nil
    }

    balanceFactor := rope.getBalanceFactor()

    if balanceFactor > 1 {
        if rope.left.getBalanceFactor() < 0 {
            rope.left = rope.left.rotateLeft()
        }
        return rope.rotateRight()
    }

    if balanceFactor < -1 {
        if rope.right.getBalanceFactor() > 0 {
            rope.right = rope.right.rotateRight()
        }
        return rope.rotateLeft()
    }

    return rope
}

func (rope *Rope) getBalanceFactor() int {
    if rope == nil {
        return 0
    }
    return rope.left.getHeight() - rope.right.getHeight()
}

func (rope *Rope) rotateRight() *Rope {
    if rope.isLeaf() {
        return rope
    }

    newRoot := rope.left
    tmp := newRoot.right
    newRoot.right = rope
    rope.left = tmp

    rope.weight = rope.left.getLength() 
    rope.length = rope.left.getLength() + rope.right.getLength()
    rope.height = max(rope.left.getHeight(), rope.right.getHeight()) + 1

    newRoot.weight = newRoot.left.getLength()
    newRoot.length = newRoot.left.getLength() + newRoot.right.getLength()
    newRoot.height = max(newRoot.left.getHeight(), newRoot.right.getHeight()) + 1

    return newRoot
}

func (rope *Rope) rotateLeft() *Rope {
    if rope.isLeaf() {
        return rope
    }

    newRoot := rope.right
    tmp := newRoot.left
    newRoot.left = rope
    rope.right = tmp

    rope.weight = rope.left.getLength() 
    rope.length = rope.left.getLength() + rope.right.getLength()
    rope.height = max(rope.left.getHeight(), rope.right.getHeight()) + 1

    newRoot.weight = newRoot.left.getLength()
    newRoot.length = newRoot.left.getLength() + newRoot.right.getLength()
    newRoot.height = max(newRoot.left.getHeight(), newRoot.right.getHeight()) + 1

    return newRoot
}

func (rope *Rope) content(from, to int) []rune {
    if rope.isLeaf() {
        return rope.value[from:to]
    } 

    var res []rune
    if from < rope.weight {
        end := min(to, rope.weight)
        res = append(res, rope.left.content(from, end)...)
    }
    if to > rope.weight {
        start := max(0, from - rope.weight)
        res = append(res, rope.right.content(start, to - rope.weight)...)
    }

    return res
}

func (rope *Rope) panicIfNil() {
    if rope == nil {
        panic(fmt.Sprintf("Operation not permitted on empty rope"))
    }
}
