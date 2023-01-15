package prototype

import "testing"

func TestClone(t *testing.T) {
	file1 := &File{name: "file1"}
	file2 := &File{name: "file2"}
	file3 := &File{name: "file3"}

	folder1 := &Folder{
		name:     "Folder1",
		children: []INode{file1, file2, file3},
	}

	folder2 := &Folder{
		name:     "Folder2",
		children: []INode{folder1, file2, file3},
	}
	folder2.Print("  ")

	cloneFolder := folder2.Clone()

	cloneFolder.Print("  ")
}
