package types

import "github.com/digi-wolk/oss-license-auditor/internal/cli"

type Package struct {
	Name               string `json:"name"`
	Owner              string `json:"owner"`
	Version            string `json:"version"`
	Dev                bool
	PackageManagerFile string `json:"packageManagerFile"`
	License            string `json:"license"`
	IsLicenseRiskyFail bool   `json:"isLicenseRiskyFail"`
	IsLicenseRiskyWarn bool   `json:"isLicenseRiskyWarn"`
}

type Dependencies struct {
	Packages            []Package     `json:"packages"`
	HasRiskyFailLicense bool          `json:"hasRiskyFailLicense"`
	PackageManagerFile  string        `json:"packageManagerFile"`
	CliArguments        cli.Arguments `json:"cliArguments"`
}

type GithubRepo struct {
	Owner     string `json:"owner"`
	Repo      string `json:"repo"`
	LicenseId string `json:"licenseId"`
}

type NpmPackage struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	License string `json:"license"`
}

type PackageInfoObjectLicense struct {
	License ObjectLicense `json:"license"`
}

type PackageInfoObjectLicenses struct {
	Licenses []ObjectLicense `json:"licenses"`
}

type ObjectLicense struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}
