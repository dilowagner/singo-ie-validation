package validators

// Validatable interface
type Validatable interface {
	IsValid(inscricao string) bool
}
