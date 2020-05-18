package randomizer

// IRandomizer interface for radnomizer components
type IRandomizer interface {
	RandInt(min int, max int) int
}
