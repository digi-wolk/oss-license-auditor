package cli

type Arguments struct {
	Path              string
	OnlyRiskyLicenses bool
	Verbose           bool
	FailOnRisky       bool
	Ci                bool
}
