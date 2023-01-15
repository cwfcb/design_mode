package prototype

import "fmt"

type INode interface {
	Print(indentation string)
	Clone() INode
}

type File struct {
	name string
}

func (f *File) Print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) Clone() INode {
	return &File{name: f.name + "_clone"}
}

type Folder struct {
	name     string
	children []INode
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.name)
	newIndentation := indentation + indentation
	for _, child := range f.children {
		child.Print(newIndentation)
	}
}

func (f *Folder) Clone() INode {
	newFolder := &Folder{name: f.name + "_clone"}
	var children []INode
	for _, child := range f.children {
		children = append(children, child.Clone())
	}
	newFolder.children = children
	return newFolder
}
