package main

type Node struct {
	val  int
	next *Node
}

func newNode(val int) *Node {
	return &Node{
		val:  val,
		next: nil,
	}
}

func createListFromArray(arr []int, n int) *Node {
	head := newNode(arr[0])
	curr := head
	for i := 1; i < n; i++ {
		head.next = newNode(arr[i])
		head = head.next
	}

	return curr
}

func readInt() int {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		log.Fatal("could not read input number", err)
	}
	return n
}

func readString() string {
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Fatal("could not read input string", err)
	}
	return s
}

func readDouble() float64 {
	var f float64
	_, err := fmt.Scanf("%f", &f)
	if err != nil {
		log.Fatal("could not read input float", err)
	}
	return f
}

func readChar() rune {
	var r rune
	_, err := fmt.Scanf("%c", &r)
	if err != nil {
		log.Fatal("Could not read input char", err)
	}
	return r
}

func merge_linkList(head1 *Node, head2 *Node) *Node {

}
