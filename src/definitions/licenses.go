package definitions

import (
	"strings"
)

// DefaultRiskyLicensesFail List of risky open source licenses
var DefaultRiskyLicensesFail = []string{
	"AGPL-1.0-only",
	"AGPL-1.0-or-later",
	"AGPL-3.0-only",
	"AGPL-3.0-or-later",
	"CC-BY-1.0",
	"CC-BY-2.0",
	"CC-BY-2.5",
	"CC-BY-NC-1.0",
	"CC-BY-NC-2.0",
	"CC-BY-NC-2.5",
	"CC-BY-NC-3.0",
	"CC-BY-NC-4.0",
	"CC-BY-NC-ND-1.0",
	"CC-BY-NC-ND-2.0",
	"CC-BY-NC-ND-2.5",
	"CC-BY-NC-ND-3.0",
	"CC-BY-NC-ND-4.0",
	"CC-BY-NC-SA-1.0",
	"CC-BY-NC-SA-2.0",
	"CC-BY-NC-SA-2.5",
	"CC-BY-NC-SA-3.0",
	"CC-BY-NC-SA-4.0",
	"CC-BY-ND-1.0",
	"CC-BY-ND-2.0",
	"CC-BY-ND-2.5",
	"CC-BY-ND-3.0",
	"CC-BY-ND-4.0",
	"CC-BY-SA-1.0",
	"CC-BY-SA-2.0",
	"CC-BY-SA-2.5",
	"CC-BY-SA-3.0",
	"CC-BY-SA-4.0",
	"CC-PDDC",
	"CC BY-NC 2.0",
	"CC BY-NC 2.5",
	"CC BY-NC 3.0",
	"CC BY-NC 4.0",
	"CC-NC-1.0",
	"CC-NC-2.0",
	"CC-NC-2.5",
	"CC-NC-3.0",
	"CC-NC-4.0",
	"LGPL-2.0-only",
	"LGPL-2.0-or-later",
	"LGPL-2.1-only",
	"LGPL-2.1-or-later",
	"LGPL-3.0-only",
	"LGPL-3.0-or-later",
	"LGPLLR",
	"AGPL-1.0",
	"AGPL-3.0",
	"GFDL-1.1",
	"GFDL-1.2",
	"GFDL-1.3",
	"GPL-1.0+",
	"GPL-1.0",
	"GPL-2.0+",
	"GPL-2.0-with-GCC-exception",
	"GPL-2.0-with-autoconf-exception",
	"GPL-2.0-with-bison-exception",
	"GPL-2.0-with-classpath-exception",
	"GPL-2.0-with-font-exception",
	"GPL-2.0",
	"GPL-3.0",
	"GPL-3.0+",
	"GPL-3.0-with-GCC-exception",
	"GPL-3.0-with-autoconf-exception",
	"GPL-3.0",
	"LGPL-2.0+",
	"LGPL-2.0",
	"LGPL-2.1+",
	"LGPL-2.1",
	"LGPL-3.0+",
	"LGPL-3.0",
	"Nunit",
}

// DefaultRiskyLicensesWarn List of risky open source licenses
var DefaultRiskyLicensesWarn = []string{
	"EUPL-1.0",
	"EUPL-1.1",
	"EUPL-1.2",
	"BSD-1-Clause",
	"BSD-2-Clause-FreeBSD",
	"BSD-2-Clause-NetBSD",
	"BSD-2-Clause-Patent",
	"BSD-3-Clause",
	"BSD-3-Clause-Attribution",
	"BSD-3-Clause-Clear",
	"BSD-3-Clause-LBNL",
	"BSD-3-Clause-No-Nuclear-License-2014",
	"BSD-3-Clause-No-Nuclear-License",
	"BSD-3-Clause-No-Nuclear-Warranty",
	"BSD-3-Clause-Open-MPI",
	"BSD-4-Clause-UC",
	"BSD-4-Clause",
	"BSD-Protection",
	"BSD-Source-Code",
	"MIT-0",
	"MIT-CMU",
	"MIT-advertising",
	"MIT-enna",
	"MIT-feh",
	"MITNFA",
}

// IsLicenseRiskyFail Checks if a license is risky if license contains one of the risky licenses (case-insensitive)
func IsLicenseRiskyFail(license string) bool {
	var riskyLicenses []string
	// TODO: Add ability to load from a json config file || default to DefaultRiskyLicensesFail
	riskyLicenses = DefaultRiskyLicensesFail
	for _, riskyLicense := range riskyLicenses {
		if strings.TrimSpace(strings.ToLower(license)) == strings.TrimSpace(strings.ToLower(riskyLicense)) {
			return true
		}
	}

	return false
}

// IsLicenseRiskyWarn Checks if a license is risky if license contains one of the risky licenses (case-insensitive)
func IsLicenseRiskyWarn(license string) bool {
	var riskyLicenses []string
	// TODO: Add ability to load from a json config file || default to DefaultRiskyLicensesWarn
	riskyLicenses = DefaultRiskyLicensesWarn
	for _, riskyLicense := range riskyLicenses {
		if strings.TrimSpace(strings.ToLower(license)) == strings.TrimSpace(strings.ToLower(riskyLicense)) {
			return true
		}
	}

	return false
}
