// +build OMIT

package main

import "golang.org/x/tour/tree"
//import "github.com/golang/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    var doWalk func(t *tree.Tree)
    doWalk = func(t *tree.Tree) {
	if t.Left != nil { doWalk(t.Left) }		    
	ch <- t.Value
	if t.Right != nil { doWalk(t.Right) } 
    }

    doWalk(t)
    close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    c1, c2 := make(chan int), make(chan int)
    go Walk(t1, c1)
    go Walk(t2, c2)
    for {
	v1, ok1 := <- c1
	v2, ok2 := <- c2
	fmt.Print(v1, "-", v2, ", ")
        if ok1 && ok2 && v1 != v2  { return false }
	if !ok1 && !ok2 { return true}
    }
}

func main() {
    ch := make(chan int)
    go Walk(tree.New(1), ch)
    for ok := true; ok; {
	var v int
	v, ok = <- ch
	if ok { fmt.Print(v, ", ") }
    }
    fmt.Println(Same(tree.New(2), tree.New(2)))
}
