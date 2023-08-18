package datastructure

import (
	"fmt"
	"log"
)

type SegmentTree struct {
	data []int
	Tree []int
}

// 返回一个
func NewSegmentTree(data []int) *SegmentTree {
	S := new(SegmentTree)
	S.data = data
	S.Tree = make([]int, len(data)*4)
	S.buildSegmentTree(0, 0, len(S.data)-1)
	return S
}

// 获取线段树的左孩子
func (S *SegmentTree) leftChild(index int) int {
	return index*2 + 1
}

// 获取线段树的右孩子
func (S *SegmentTree) rightChild(index int) int {
	return index*2 + 2
}

// 递归构建线段树
func (S *SegmentTree) buildSegmentTree(treeIndex, l, r int) {
	// base case: 递归到叶子节点
	if l == r {
		S.Tree[treeIndex] = S.data[l]
		return
	}
	left := S.leftChild(treeIndex)
	right := S.rightChild(treeIndex)
	// mid := (l + r) / 2 // 容易超范围
	mid := l + (r-l)/2
	S.buildSegmentTree(left, l, mid)
	S.buildSegmentTree(right, mid+1, r)
	S.Tree[treeIndex] = S.Tree[left] * S.Tree[right]
}

// 查询区间 [l-r] 的值
func (S *SegmentTree) Query(l, r int) int {
	if l < 0 || l > len(S.data)-1 || r < 0 || r > len(S.data)-1 {
		log.Fatal(fmt.Errorf("the index is invalid"))
	}
	return S.query(0, 0, len(S.data)-1, l, r)
}

func (S *SegmentTree) query(treeIndex, treeL, treeR, l, r int) int {
	if l == treeL && r == treeR {
		return S.Tree[treeIndex]
	}
	mid := treeL + (treeR-treeL)/2
	if mid < l {
		return S.query(S.rightChild(treeIndex), mid+1, treeR, l, r)
	} else if mid >= r {
		return S.query(S.leftChild(treeIndex), treeL, mid, l, r)
	} else {
		leftResult := S.query(S.leftChild(treeIndex), treeL, mid, l, mid)
		rightResult := S.query(S.rightChild(treeIndex), mid+1, treeR, mid+1, r)
		return leftResult * rightResult
	}
}
