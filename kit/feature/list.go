// Code generated by the feature package; DO NOT EDIT.

package feature

var backendExample = MakeBoolFlag(
	"Backend Example",
	"backendExample",
	"Gavin Cabbage",
	false,
	Permanent,
	false,
)

// BackendExample - A permanent backend example boolean flag
func BackendExample() BoolFlag {
	return backendExample
}

var frontendExample = MakeIntFlag(
	"Frontend Example",
	"frontendExample",
	"Gavin Cabbage",
	42,
	Temporary,
	true,
)

// FrontendExample - A temporary frontend example integer flag
func FrontendExample() IntFlag {
	return frontendExample
}

var newAuth = MakeBoolFlag(
	"New Auth Package",
	"newAuth",
	"Alirie Gray",
	false,
	Temporary,
	false,
)

// NewAuthPackage - Enables the refactored authorization api
func NewAuthPackage() BoolFlag {
	return newAuth
}

var sessionService = MakeBoolFlag(
	"Session Service",
	"sessionService",
	"Lyon Hill",
	false,
	Temporary,
	true,
)

// SessionService - A temporary switching system for the new session system
func SessionService() BoolFlag {
	return sessionService
}

var all = []Flag{
	backendExample,
	frontendExample,
	newAuth,
	sessionService,
}

var byKey = map[string]Flag{
	"backendExample":  backendExample,
	"frontendExample": frontendExample,
	"newAuth":         newAuth,
	"sessionService":  sessionService,
}
