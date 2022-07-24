func findRedundantConnection(edges [][]int) []int {

	edgeNum := len(edges)
	dSet := newDisjointSet(edgeNum)
	stack := make(Stack, 0)
	for i, edge := range edges {
		dSet.union(i, edge[0]-1, edge[1]-1, &stack)
	}
	return edges[stack.pop()]

}

type DisjointSet struct {

	parents []int
	ranks []int
	count int

}

func newDisjointSet(edgeNum int) *DisjointSet {

	parents := make([]int, edgeNum)
	for i := range parents{
		parents[i] = i
	}
	ranks := make([]int, edgeNum)
	return &DisjointSet{
		parents: parents,
		ranks: ranks,
		count: edgeNum,
	}

}

type Stack []int

func (s *Stack) push(i int){
	*s = append(*s, i)
}

func (s *Stack) pop() int {
	l := len(*s)
	ans := (*s)[l-1]
	*s = (*s)[:l-1]
	return ans
}



func (dSet *DisjointSet) find(index int) int {

	parents := dSet.parents
	if index != parents[index] {
		parents[index] = dSet.find(parents[index]) // path compression
	}
	return parents[index]

}

func (dSet *DisjointSet) union(row, child1, child2 int, stack *Stack) {

	root1 := dSet.find(child1)
	root2 := dSet.find(child2)
	if root1 == root2 {
		stack.push(row)
	}
	ranks := dSet.ranks
	parents := dSet.parents
	if ranks[root1] < ranks[root2] {
		parents[root1] = root2
	} else if ranks[root1] > ranks[root2] {
		parents[root2] = root1
	} else {
		parents[root2] = root1
		ranks[root1]++
	}

}
