package dirls

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type node struct {
	fullPath string
	file     file
	children []*node
	parent   *node
}

type nodes []*node

func (n nodes) Len() int           { return len(n) }
func (n nodes) Less(i, j int) bool { return n[i].file.size < n[j].file.size }
func (n nodes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

type nodeMap map[string]*node

type file struct {
	path string
	size int64
}

type files []file

func (f files) Len() int           { return len(f) }
func (f files) Less(i, j int) bool { return f[i].size < f[j].size }
func (f files) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

func walkDir(ff *files, n nodeMap) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		f := file{
			path: path,
			size: info.Size(),
		}
		*ff = append(*ff, f)

		if n != nil {
			n[path] = &node{
				fullPath: path,
				file:     f,
				children: make([]*node, 0),
			}
		}

		return nil
	}
}

// DirLs lists the directories and files inside the given path. Returns the
// collection of the files/paths
func DirLs(dir string) (files, *node, error) {
	ff := make(files, 0)

	t := make(map[string]*node, 0)

	err := filepath.Walk(dir, walkDir(&ff, t))
	if err != nil {
		return nil, nil, err
	}
	return ff, generateNodeTree(t), nil
}

type sortBy int

const (
	ASC sortBy = iota
	DESC
)

// SortFiles sorts the files either ASC or DESC
func SortFiles(ff files, sb sortBy) {
	switch sb {
	case ASC:
		sort.Sort(ff)
	case DESC:
		sort.Sort(sort.Reverse(ff))
	}
}

// SortNodes sorts the nodes either ASC or DESC
func SortNodes(ff nodes, sb sortBy) {
	switch sb {
	case ASC:
		sort.Sort(ff)
	case DESC:
		sort.Sort(sort.Reverse(ff))
	}
}

// generateNodeTree ...
func generateNodeTree(n nodeMap) *node {
	var result *node
	paths := make([]string, len(n))
	i := 0
	for k := range n {
		paths[i] = k
		i++
	}
	// need to sort so that the dot is the root directory
	sort.Strings(paths)
	for _, path := range paths {
		node := n[path]
		parentPath := filepath.Dir(node.fullPath)
		parent, exists := n[parentPath]
		// if it's the "." then it's the parent directory
		if path == "." || !exists {
			result = node
		} else {
			node.parent = parent
			parent.children = append(parent.children, node)
		}
	}
	return result
}

func PrintFiles(ff files, sb sortBy) {
	SortFiles(ff, sb)
	for i := range ff {
		fmt.Printf("%s %d\n", ff[i].path, ff[i].size)
	}
}

func PrintTree(n *node, space string) {
	fmt.Printf("%s%s %d\n", space, n.fullPath, n.file.size)
	if n.children != nil && len(n.children) > 0 {
		SortNodes(n.children, ASC)
		for _, c := range n.children {
			PrintTree(c, space+"-")
		}
	}
}
