package composite

import "fmt"

type FileSystemNode interface {
	search(keyword string)
}

type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)

}

type Folder struct {
	name  string
	nodes []FileSystemNode
}

func (f *Folder) addNode(node FileSystemNode) {
	f.nodes = append(f.nodes, node)
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, node := range f.nodes {
		node.search(keyword)
	}
}
