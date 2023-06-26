package decode

import "testing"

func TestPosition(t *testing.T) {
	tests := []struct {
		name  string
		p     Position
		wantX int32
		wantY int32
		wantZ int32
	}{
		{
			name:  "position decode",
			p:     0b01000110000001110110001100_10110000010101101101001000_001100111111,
			wantX: 18357644,
			wantY: 831,
			wantZ: -20882616,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.X(); got != tt.wantX {
				t.Errorf("X() = %v, want %v", got, tt.wantX)
			}
			if got := tt.p.Y(); got != tt.wantY {
				t.Errorf("Y() = %v, want %v", got, tt.wantY)
			}
			if got := tt.p.Z(); got != tt.wantZ {
				t.Errorf("Z() = %v, want %v", got, tt.wantZ)
			}
		})
	}
}
