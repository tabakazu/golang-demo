package entity

import (
	"testing"
)

func TestUserFullName(t *testing.T) {
	t.Parallel()
	type fields struct {
		GivenName  string
		FamilyName string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"", fields{"John", "Snow"}, "JohnSnow"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := &User{
				GivenName:  tt.fields.GivenName,
				FamilyName: tt.fields.FamilyName,
			}
			if got := u.FullName(); got != tt.want {
				t.Errorf("User.FullName() = %v, want %v", got, tt.want)
			}
		})
	}
}
