package enum

type PreciseMode string

const (
	Round PreciseMode = "Round"
	Floor PreciseMode = "Floor"
	Ceil  PreciseMode = "Ceil"
)
