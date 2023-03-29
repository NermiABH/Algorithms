package main

import (
	"errors"
	"fmt"
)

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

func (t *Tree) Insert(x int) error {
	if t == nil {
		return errors.New("tree is nil")
	} else if t.Val == x {
		return errors.New("this node value already exists")
	} else if t.Val > x {
		if t.Left == nil {
			t.Left = &Tree{Val: x}
			return nil
		}
		return t.Left.Insert(x)
	}
	if t.Right == nil {
		t.Right = &Tree{Val: x}
		return nil
	}
	return t.Right.Insert(x)
}

func (t *Tree) Exist(x int) bool {
	if t == nil {
		return false
	} else if t.Val == x {
		return true
	}
	return t.Left.Exist(x) || t.Right.Exist(x)
}

//func (t *Tree) Delete(x int) error {
//	if t == nil{
//		return errors.New("tree is nil")
//	}else if t.Val
//}

func (t *Tree) FindMin() int {
	if t.Left == nil {
		return t.Val
	}
	return t.Left.FindMin()
}

func (t *Tree) FindMax() int {
	if t.Right == nil {
		return t.Val
	}
	return t.Right.FindMax()
}

func (t *Tree) PrintInorder() {
	if t == nil {
		return
	}
	t.Left.PrintInorder()
	fmt.Print(t.Val, " ")
	t.Right.PrintInorder()
}

func main() {
	t := &Tree{Val: 8}
	t.Insert(3)
	t.Insert(4)
	t.Insert(10)
	t.Insert(19)
	fmt.Println(t.Exist(10), t.FindMin(), t.FindMax())
	t.PrintInorder()
}
