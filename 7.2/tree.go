package main

import "fmt"

type TreeNode interface {
	Size() uint64
	GetParent() TreeNode
	Children() []TreeNode
	AddChild(TreeNode) bool
	Print() string
	GetName() string
}

type DirNode struct {
	Name     string
	children []TreeNode
	Parent   TreeNode
}

type FileNode struct {
	Name     string
	FileSize uint64
	children []TreeNode
	Parent   TreeNode
}


// IMPLEMENT DIRNODE
func (dn DirNode) Size() uint64 {
	var size uint64 = 0
	for _, e := range dn.children {
		size += e.Size()
	}

	return size
}
func (dn DirNode) GetParent() TreeNode {
	return dn.Parent
}
func (dn DirNode) Children() []TreeNode {
	return dn.children
}
func (dn *DirNode) AddChild(node TreeNode) bool {
	dn.children = append(dn.children, node)
	return true
}
func (dn DirNode) Print() string {
	return "- " + dn.Name + " (dir)"
}
func (dn DirNode) GetName() string {
	return dn.Name
}

// IMPLEMENT FILENODE
func (fn FileNode) Size() uint64 {
	return fn.FileSize
}
func (fn FileNode) GetParent() TreeNode {
	return fn.Parent
}
func (fn FileNode) Children() []TreeNode {
	return nil
}
func (fn *FileNode) AddChild(node TreeNode) bool {
	return false
}
func (fn FileNode) Print() string {
	return fmt.Sprintf("- %v (file, size=%v)", fn.Name, fn.FileSize)
}
func (fn FileNode) GetName() string {
	return fn.Name
}
