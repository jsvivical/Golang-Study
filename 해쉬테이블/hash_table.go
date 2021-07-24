/*****************************************************************************************************************************************
 * 해쉬테이블 : 한 개 이상의 키와 값의 쌍을 저장하는 데이터 구조로, 버킷이나 슬롯배열에서 원하는 값이 담긴 위치를 계산할 때, 해시함수를 사용
*****************************************************************************************************************************************/

package main

import "fmt"

const SIZE int = 15

type Node struct {
	Value int
	Next  *Node
}
type hashTable struct {
	Table map[int]*Node
	Size  int
}

func hashFunction(i, size int) int {
	return (i % size)
}

func insert(hash *hashTable, value int) int {
	index := hashFunction(value, hash.Size)
	element := Node{Value: value, Next: hash.Table[index]}
	hash.Table[index] = &element
	return index
}

func traverse(hash *hashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%3d - > ", t.Value)
				t = t.Next
			}
			fmt.Println()
		}
	}
}

func main() {
	table := make(map[int]*Node, SIZE)
	hash := &hashTable{Table: table, Size: SIZE}
	fmt.Println("Number of spaces : ", hash.Size)
	for i := 0; i < 120; i++ {
		insert(hash, i)
	}
	traverse(hash)
}
