package tracker

import "fmt"

type WorldTracker struct {
	Dimension     string
	DimensionType *DimensionType
}

func (receiver Tracker) ResetWorldTracker(dimName string, dimTypeName string) error {
	dimensionType, ok := receiver.ServerInfo.DimensionTypes[dimTypeName]
	if !ok {
		return fmt.Errorf("unknown dimenstion type %v of dimension %v", dimTypeName, dimName)
	}
	receiver.WorldTracker = WorldTracker{
		Dimension:     dimName,
		DimensionType: &dimensionType,
	}
	return nil
}
