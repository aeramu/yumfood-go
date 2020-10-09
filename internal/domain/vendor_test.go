package domain

import (
	"reflect"
	"testing"
)

// create name with specific length
func newString(len int) string {
	s := ""
	for i := 0; i < len; i++ {
		s += "a"
	}
	return s
}

func TestNewVendor(t *testing.T) {
	type args struct {
		name        string
		description string
	}
	tests := []struct {
		name string
		args args
		want Vendor
	}{
		// TODO: Add test cases.
		{
			"Basic test",
			args{"basingse", "minangnese restaurant"},
			Vendor{"", "basingse", "minangnese restaurant"},
		},
		{
			"Name validation",
			args{newString(200), "river restaurant"},
			Vendor{"", newString(127), "river restaurant"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVendor(tt.args.name, tt.args.description); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVendor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVendor_Update(t *testing.T) {
	type fields struct {
		ID          string
		Name        string
		Description string
	}
	type args struct {
		name        string
		description string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vendor
	}{
		// TODO: Add test cases.
		{
			"Basic test",
			fields{"", "oldy and rusty", "japanese restaurant"},
			args{"basingse", "minangnese restaurant"},
			&Vendor{"", "basingse", "minangnese restaurant"},
		},
		{
			"Name validation",
			fields{"", "oldy and rusty", "japanese restaurant"},
			args{newString(200), "river restaurant"},
			&Vendor{"", newString(127), "river restaurant"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vendor{
				ID:          tt.fields.ID,
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
			}
			if got := v.Update(tt.args.name, tt.args.description); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vendor.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
