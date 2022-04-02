package IdUtils

import (
	"fmt"
	"testing"
)

func TestGetId(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "TestGetId",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := GetId()
			if got := id; got > 0 != tt.want {
				t.Errorf("GetId() = %v, want %v", got, tt.want)
			} else {
				fmt.Println(fmt.Sprintf("TestGetId:%d", id))
			}
		})
	}
}
