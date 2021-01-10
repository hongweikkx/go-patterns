package composite

import (
	"fmt"
	"testing"
)

func TestComposite(t *testing.T) {
	root := NewComposite("root")
	root.Add(NewLeaf("Leaf A"))
	root.Add(NewLeaf("Leaf B"))

	compositeX := NewComposite("compositeX")
	compositeX.Add(NewLeaf("leaf XA"))
	compositeX.Add(NewLeaf("leaf XB"))
	root.Add(compositeX)
	//
	compositeXY := NewComposite("compositeXY")
	compositeXY.Add(NewLeaf("leaf XYA"))
	compositeXY.Add(NewLeaf("leaf XYB"))
	compositeX.Add(compositeXY)

	root.Add(NewLeaf("Leaf C"))
	leafD := NewLeaf("Leaf D")
	root.Add(leafD)
	root.Remove(leafD)

	fmt.Println(root.Display(1))
}
