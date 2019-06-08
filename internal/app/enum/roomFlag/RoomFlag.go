package roomFlag

type Flag string

const (
	Road           Flag = "road"
	Unfordable     Flag = "unfordable"
	Trees          Flag = "trees"
	OreProbability Flag = "ore_probability"

	Forest   Flag = "forest"
	Sand     Flag = "sand"
	Plain    Flag = "plain"
	Hill     Flag = "hill"
	Mountain Flag = "mountain"
	Sea      Flag = "sea"
)

func BiomFlags() []Flag {
	return []Flag{
		Forest,
		Sand,
		Plain,
		Hill,
		Mountain,
		Sea,
	}
}
