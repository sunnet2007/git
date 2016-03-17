package git

type Branch struct {
	name string // 分支名字
}

func NewBranch(name string) *Branch {
	return &Branch{name: name}
}

// 创建并切换
func (this *Branch) CheckOut() error {
	// git checkout -b branchName
	return nil
}

// 查看当前分支
func (this *Branch) Branch() (string, error) {
	// git branch
	return "", nil
}

// 切换分支
func (this *Branch) Switch() error {
	// git checkout
}
