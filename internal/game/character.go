package game

// This is a struct that outlines a blueprint for what a character is
type Character struct {
	Name            string
	CurrentHealth   int
	MaxHealth       int
	EquippedWeapons []string
}
