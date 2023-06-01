package definitions

import "testing"

// Test IsLicenseRiskyFail should not fail for BSD-3-Clause
func TestIsLicenseRiskyFailForBSD3Clause(t *testing.T) {
	if IsLicenseRiskyFail("BSD-3-Clause") {
		t.Errorf("IsLicenseRisky was incorrect, got: %t, want: %t.", true, false)
	}
}

// Test IsLicenseRiskyFail for AGPL-3.0
func TestIsLicenseRiskyFailForAGPL30(t *testing.T) {
	if !IsLicenseRiskyFail("AGPL-3.0") {
		t.Errorf("IsLicenseRisky was incorrect, got: %t, want: %t.", false, true)
	}
}

// Test IsLicenseRiskyFail for GPL-3.0
func TestIsLicenseRiskyFailForGPL30(t *testing.T) {
	if !IsLicenseRiskyFail("GPL-3.0") {
		t.Errorf("IsLicenseRisky was incorrect, got: %t, want: %t.", false, true)
	}
}

// Test IsLicenseRiskyWarn should warn for EUPL-1.0
func TestIsLicenseRiskyWarnForEUPL10(t *testing.T) {
	if !IsLicenseRiskyWarn("EUPL-1.0") {
		t.Errorf("IsLicenseRisky was incorrect, got: %t, want: %t.", false, true)
	}
}

// Test IsLicenseRiskyWarn should not warn for GPL-1.0-only
func TestIsLicenseRiskyWarnForGPL10Only(t *testing.T) {
	if IsLicenseRiskyWarn("GPL-1.0-only") {
		t.Errorf("IsLicenseRisky was incorrect, got: %t, want: %t.", true, false)
	}
}

// Test IsLicenseRiskyWarn should warn for BSD-3-Clause
func TestIsLicenseRiskyWarnForBSD3Clause(t *testing.T) {
	if !IsLicenseRiskyWarn("BSD-3-Clause") {
		t.Errorf("IsLicenseRisky was incorrect, got: %t, want: %t.", false, true)
	}
}
