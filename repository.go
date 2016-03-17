package git

type Repository struct {
	name      string // 仓库名称
	localPath string //仓库本地路径
}

func NewRepository(name string, path string) *Repository {
	return &Repository{name: name, localPath: path}
}

// 将仓库编程Git可以管理的仓库
func (this *Repository) Init() error {
	// git init
	return nil
}

// 将文件加到git仓库
// fileName: 添加的文件名
func (this *Repository) Add(fileName string) error {
	// 判断是否在仓库中
	// git add
	return nil
}

// 把文件提交到仓库
func (this *Repository) Commit(fileName string, record string) error {
	// git commit -m "record"
	return nil
}

// 将本地仓库与远程仓库关联
func (this *Repository) Associate() error {
	// git remote add origin git@github.com:GitUserName/repositoryName.git
	return nil
}

// 将本地仓库的所有内存推送到远程库上
func (this *Repository) Push() error {
	// 第一次推送master分支的所有内容 git push -u origin master
	// 之后推送就可以使用 git push origin master
	return nil
}

// 从远程仓库克隆
// 参数userName: Git用户名, repoName远程仓库名
func (this *Repository) Clone(userName string, repoName string) error {
	// git clone git@github.com:userName/repoName.git
	return nil
}

// 查看仓库的当前状态
func (this *Repository) Status() (string, error) {
	// git status
	return "", nil
}

// 查看修改内容
func (this *Repository) Diff() (string, error) {
	// git diff
}
