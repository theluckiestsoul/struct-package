package npcs

// Power describes teh attch and defense power of an NPC
type Power struct {
	Attack  int
	Defense int
}

// Location describes where in the virual world an NPC exists
type Location struct {
	X float64
	Y float64
	Z float64
}

// NonPlayerCharacter represents metadata for an in-game creature
type NonPlayerCharacter struct {
	Name  string
	Speed int
	HP    int
	Power Power
	Loc   Location
}
