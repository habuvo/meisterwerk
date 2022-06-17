package handlers

import "testing"

func Test_timeParamsChecker(t *testing.T) {
	type args struct {
		f string
		t string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"empty", args{"", ""}, false},
		{"half empty", args{"", "2022-12-10 12:00"}, false},
		{"half empty another", args{"2022-12-10 12:00", ""}, false},
		{"bad", args{"2022-15-10 12:00", ""}, true},
		{"another bad", args{"", "2022-15-10 12:00"}, true},
		{"bad both", args{"2022-15-10 12:00", "2022-15-10 12:00"}, true},
		{"overlap", args{"2022-10-01 12:00", "2022-01-10 12:00"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := timeParamsChecker(tt.args.f, tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("timeParamsChecker() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
