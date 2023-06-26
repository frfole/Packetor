package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type AwardStatisticsEntry struct {
	CategoryID  int32
	StatisticID int32
	Value       int32
}

type AwardStatistics struct {
	Statistics []AwardStatisticsEntry
}

func (p AwardStatistics) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	entries := make([]AwardStatisticsEntry, count)
	for i := int32(0); i < count; i++ {
		cid, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
		sid, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
		value, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
		entries[i] = AwardStatisticsEntry{
			CategoryID:  cid,
			StatisticID: sid,
			Value:       value,
		}
	}
	return AwardStatistics{Statistics: entries}, nil
}

func (p AwardStatistics) IsValid() (reason error) {
	for i, statistic := range p.Statistics {
		if statistic.CategoryID < 0 || 8 < statistic.CategoryID {
			return fmt.Errorf("CategoryID must be in range <0; 8> was %d at index %d", statistic.StatisticID, i)
		}
		switch statistic.CategoryID {
		// TODO: validate other categories
		case 8:
			if statistic.StatisticID < 0 || 73 < statistic.StatisticID {
				return fmt.Errorf("StatisticID for category custom must be in range <0; 73> was %d at index %d", statistic.StatisticID, i)
			}
		}
	}
	return nil
}
