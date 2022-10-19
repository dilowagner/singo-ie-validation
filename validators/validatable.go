package validators

// Validator interface
type Validator interface {
	IsValid(insc string) bool
}
