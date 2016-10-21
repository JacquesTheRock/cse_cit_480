package entity

type TreeNode struct {
	Parents  []TreeNode `json:"parents"`
	Self     Cross      `json:"data"`
	Children []TreeNode `json:"children"`
}
