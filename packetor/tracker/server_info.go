package tracker

var DimensionTypeDefault = DimensionType{
	MinY:   -64,
	Height: 384,
}

type (
	DimensionType struct {
		MinY   int32
		Height int32
	}
	ServerInfo struct {
		DimensionTypes map[string]DimensionType
	}
)
