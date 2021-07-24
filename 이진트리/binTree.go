package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	left  *Tree
	value int
	right *Tree
}

const MAX_QUEUE int = 100

var front, rear int

var queue [MAX_QUEUE]*Tree

func Inorder(t *Tree) {
	if t == nil {
		return
	}
	Inorder(t.left)
	fmt.Printf("%5d", t.value)
	Inorder(t.right)
}

func Preorder(t *Tree) {
	if t == nil {
		return
	}
	fmt.Printf("%5d", t.value)
	Preorder(t.left)
	Preorder(t.right)
}

func Postorder(t *Tree) {
	if t == nil {
		return
	}
	Postorder(t.left)
	Postorder(t.right)
	fmt.Printf("%5d", t.value)
}

func addQ(v *Tree) {
	rear++
	queue[rear] = v
}

func deleteQ() *Tree {
	front++
	temp := queue[front]
	return temp
}

func LevelOrder(t *Tree) {
	front, rear = 0, 0
	if t == nil {
		return
	}
	addQ(t)
	for {
		t = deleteQ()
		if t != nil {
			fmt.Printf("%5d", t.value)
			if t.left != nil {
				addQ(t.left)
			}
			if t.right != nil {
				addQ(t.right)
			}
		} else {
			break
		}
	}
}

func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	fmt.Print("Random nums : ")
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
		fmt.Printf("%5d", temp)
		t = insert(t, temp)
	}
	fmt.Println("\n")
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}

	if v == t.value {
		return t
	}

	if v < t.value {
		t.left = insert(t.left, v)
		return t
	}
	t.right = insert(t.right, v)
	return t
}

func main() {
	tree := create(10)
	fmt.Println("The value of the root of the tree is", tree.value)
	fmt.Print("Preorder : ")
	Preorder(tree)
	fmt.Println()
	fmt.Println()

	fmt.Print("Inorder : ")
	Inorder(tree)
	fmt.Println()
	fmt.Println()

	fmt.Print("Postorder : ")
	Postorder(tree)
	fmt.Println()
	fmt.Println()

	fmt.Print("Levelorder : ")
	LevelOrder(tree)
	fmt.Println()
	fmt.Println()

	fmt.Println("the value of the root of the tree is ", tree.value)

}
