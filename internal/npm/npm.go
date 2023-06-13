package npm

import (
	"encoding/json"
	"fmt"
	"github.com/digi-wolk/oss-license-auditor/internal/definitions"
	"github.com/digi-wolk/oss-license-auditor/internal/types"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// UpdatePackageFromNpm updates NPM package information
func UpdatePackageFromNpm(npmPackage *types.Package) error {
	packageVersion := npmPackage.Version

	const baseUrl = "https://registry.npmjs.org/"
	urlPath := fmt.Sprintf("%s/%s/%s", npmPackage.Owner, npmPackage.Name, packageVersion)
	// Parse and validate the URL
	parsedURL, err := url.Parse(baseUrl)
	if err != nil {
		fmt.Println("Failed to parse the base URL:", err)
		return err
	}
	parsedURL.Path += urlPath

	resp, err := http.Get(parsedURL.String())
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("Error occurred while closing response body:", err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var packageInfo types.Package
	// More regular format
	err = json.Unmarshal(body, &packageInfo)
	if err != nil || packageInfo.License == "" {
		if strings.Contains(strings.ToLower(string(body)), "not found") {
			log.Printf("WARNING: Package %s (version %s) not found on NPM.", npmPackage.Name, packageVersion)
			npmPackage.License = "UNKNOWN (not found)"
			npmPackage.IsLicenseRiskyWarn = true
			// Do not fail, instead just add UNKNOWN license
			return nil
		}
		// Less regular format
		lessRegularLicenseFormat := types.PackageInfoObjectLicense{}
		err = json.Unmarshal(body, &lessRegularLicenseFormat)
		if err != nil {
			log.Print(string(body))
			npmPackage.License = "UNKNOWN"
			npmPackage.IsLicenseRiskyWarn = true
			// Do not fail, instead just add UNKNOWN license
			return nil
		}
		// Now only one license is supported
		// TODO: Handle multiple licenses later
		npmPackage.License = lessRegularLicenseFormat.Licenses[0].Type
		if npmPackage.License == "" {
			npmPackage.License = "UNKNOWN (empty)"
			npmPackage.IsLicenseRiskyWarn = true
			// Do not fail, instead just add UNKNOWN license
			return nil
		}
		npmPackage.IsLicenseRiskyFail = definitions.IsLicenseRiskyFail(packageInfo.License)
		npmPackage.IsLicenseRiskyWarn = definitions.IsLicenseRiskyWarn(packageInfo.License)
		return nil
	}
	npmPackage.License = packageInfo.License
	npmPackage.IsLicenseRiskyFail = definitions.IsLicenseRiskyFail(packageInfo.License)
	npmPackage.IsLicenseRiskyWarn = definitions.IsLicenseRiskyWarn(packageInfo.License)

	return nil
}
