package composite

import "testing"

func TestNode(t *testing.T) {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}
	folder1 := &Folder{
		name: "Folder1",
	}
	folder1.addNode(file1)
	folder2 := &Folder{
		name: "Folder2",
	}
	folder2.addNode(file2)
	folder2.addNode(file3)
	folder2.addNode(folder1)
	folder2.search("rose")
}
