package roomFlag

type Flag string

const (
	Road            Flag = "road"
	Unfordable      Flag = "unfordable"
	Trees           Flag = "trees"
	OreProbability  Flag = "ore_probability"
	FishProbability Flag = "fish_probability"
	GemProbability  Flag = "gem_probability"
	CaveProbability Flag = "cave_probability"
)

func (flag Flag) String() string {
	return string(flag)
}
