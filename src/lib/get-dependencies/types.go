package get_dependencies

type GithubRepoDetailsLicense struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type GoModPackageInfo struct {
	Provider string `json:"provider"`
	FullName string `json:"full_name"`
	Owner    string `json:"owner"`
	Name     string `json:"name"`
	Version  string `json:"version"`
}

type GithubRepoDetails struct {
	FullName string                   `json:"full_name"`
	License  GithubRepoDetailsLicense `json:"license"`
}

type PackageJsonLock struct {
	Packages map[string]PackageJsonLockPackage `json:"packages"`
}

type PackageJsonLockPackage struct {
	Version         string            `json:"version"`
	Resolved        string            `json:"resolved"`
	Dev             bool              `json:"dev"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}
