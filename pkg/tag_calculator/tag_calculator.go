package tag_calculator

type AgentTag string

const REPRODUCTION_THRESHOLD = 80

const (
	Death        AgentTag = "Death"
	Reproduction          = "Reproduction"
	Fight                 = "Fight"
)

type TagCalculator struct {
}

func (tc TagCalculator) DummyCalculate(energy int) AgentTag {
	if energy == 0 {
		return Death
	} else if energy > REPRODUCTION_THRESHOLD {
		return Reproduction
	}
	return Fight
}
