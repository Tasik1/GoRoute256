package main

import "testing"

func Test_stringOpener(t *testing.T) {
	type args struct {
		toOpen string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				toOpen: "a4bc2d5e",
			},
			want:    "aaaabccddddde",
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				toOpen: "abcd",
			},
			want:    "abcd",
			wantErr: false,
		},
		{
			name: "3",
			args: args{
				toOpen: "3abc",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "4",
			args: args{
				toOpen: "45",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "5",
			args: args{
				toOpen: "aaa10b",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "6",
			args: args{
				toOpen: "aaa0b",
			},
			want:    "aab",
			wantErr: false,
		},
		{
			name: "7",
			args: args{
				toOpen: "",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "8",
			args: args{
				toOpen: "d\n5abc",
			},
			want:    "d\n\n\n\n\nabc",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := stringOpener(tt.args.toOpen)
			if (err != nil) != tt.wantErr {
				t.Errorf("stringOpener() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("stringOpener() got = %v, want %v", got, tt.want)
			}
		})
	}
}
