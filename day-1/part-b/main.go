package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
	- you are given two lists, each list is unsorted
	- sort each list
	- look at at list_1[i], and list_2[i], get the difference
	- add all differences together

*/

// get left and right number.
// TREE sort them numbers

type Node struct {
	value      *int
	left_node  *Node
	right_node *Node
}

type Column struct {
	lower_bound int
	upper_bound int
}

var ColumnOne = Column{0, 5}
var ColumnTwo = Column{8, 13}

func InsertNode(root *Node, value int) *Node {

	if root == nil {

		return &Node{
			value: &value,
		}

	}

	if value < *root.value {
		root.left_node = InsertNode(root.left_node, value)
	} else {
		root.right_node = InsertNode(root.right_node, value)
	}

	return root
}

func FlattenTree(root *Node, result *[]int, count *map[int]int) {

	if root != nil {
		FlattenTree(root.left_node, result, count)
		*result = append(*result, *root.value)
		value_count, ok := (*count)[*root.value]
		if ok {
			(*count)[*root.value] = value_count + 1
		} else {
			(*count)[*root.value] = 1
		}
		FlattenTree(root.right_node, result, count)
	}

}

// given a file name, and column
// read from file
// return a binary search tree of the column
func ProduceArray(filePath string, column *Column) ([]int, map[int]int, error) {
	//

	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var rootNode *Node
	for scanner.Scan() {

		n, err := strconv.Atoi((scanner.Text()[column.lower_bound:column.upper_bound]))
		if err != nil {
			panic(err)
		}

		rootNode = InsertNode(rootNode, n)

	}

	result := []int{}
	count := map[int]int{}
	FlattenTree(rootNode, &result, &count)

	return result, count, nil
}

func main() {

	listOne, _, err := ProduceArray("./input.txt", &ColumnOne)
	if err != nil {
		panic(err)
	}
	_, countTwo, err := ProduceArray("./input.txt", &ColumnTwo)

	sum := 0
	for i := 0; i < len(listOne); i++ {

		count, ok := countTwo[listOne[i]]
		if !ok {
			count = 0
		}

		sum += listOne[i] * count
	}

	fmt.Println(sum)
}
