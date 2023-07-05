package tracker

type (
	DimensionType struct {
		MinY   int32
		Height int32
	}
	ServerInfo struct {
		DimensionTypes map[string]DimensionType
	}
)
