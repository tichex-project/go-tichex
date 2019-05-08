package assets

//nolint
const (
	MicroTichexDenom = "uthx"
	MicroTHEURDenom  = "utheur"

	MicroUnit = int64(1e6)
)

// IsValidDenom returns the given denom is valid or not
func IsValidDenom(denom string) bool {
	return denom == MicroTichexDenom ||
		denom == MicroTHEURDenom
}