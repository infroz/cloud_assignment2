package commits

type commits struct {
	Repos []repos
	Auth  bool
}
type reposTmp struct {
	Path    string `json:"path_with_namespace"`
	ID      int    `json:"id"`
	Commits int
}

type repos struct {
	Repository string
	Commits    int
}
