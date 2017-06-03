package validators

// Validatable interface
type Validatable interface {
	IsValid(insc string) bool
}
