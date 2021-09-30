package main

import "testing"

func Test_suma(t *testing.T) {
	type args struct {
		numberOne int
		numberTwo int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Test_sumaVerdadero",
			args: args{
				numberOne: 3,
				numberTwo: 4,
			},
			want: 7,
			wantErr: false,
		},
		{
			name: "Test_sumaNegativo",
			args: args{
				numberOne: 0,
				numberTwo: 4,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := suma(tt.args.numberOne, tt.args.numberTwo)
			if (err != nil) != tt.wantErr {
				t.Errorf("suma() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("suma() got = %v, want %v", got, tt.want)
			}
		})
	}
}
