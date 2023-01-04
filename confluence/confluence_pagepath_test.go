package confluence

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestPagePathIndex(t *testing.T) {
	index := NewPagePathIndex()

	parent := ContentResult{Id: "1234", Title: "parent"}
	child := ContentResult{Id: "abcde", Title: "child"}
	index.AddParent(parent, child)

	filepath := index.Filepath(child.Id)
	if filepath != parent.Title+"/"+child.Title {
		t.Errorf("filepath=%s", filepath)
	}

	parent = ContentResult{Id: "1234", Title: "parent"}
	middle := ContentResult{Id: "123abcd", Title: "middle"}
	child = ContentResult{Id: "abcde", Title: "child"}
	index.AddParent(parent, middle)
	index.AddParent(middle, child)

	filepath = index.Filepath(child.Id)
	if filepath != parent.Title+"/"+middle.Title+"/"+child.Title {
		t.Errorf("filepath=%s", filepath)
	}

	single := ContentResult{Id: "2sdasg", Title: "single"}
	index.AddPage(single)

	filepath = index.Filepath(single.Id)
	if filepath != single.Title {
		t.Errorf("filepath=%s", filepath)
	}
}
