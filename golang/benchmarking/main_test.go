package benchmarking

import (
	"fmt"
	"testing"
)

func Test_test(t *testing.T) {
	type args struct {
		loop int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{loop: 10},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loops2(tt.args.loop); got != tt.want {
				t.Errorf("loops2() = %v, want %v", got, tt.want)
			}
		})
	}
}

var table = []struct {
	input int
}{
	{input: 10},
	{input: 100},
	{input: 1000},
}

func BenchmarkLoops2(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				loops2(v.input)
			}
		})
	}
}

func BenchmarkLoops3(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				loops3(v.input)
			}
		})
	}
}

/*
go test -bench=Loops2 -count 5 | tee loops2.txt
go test -bench=Loops3 -count 5 | tee loops3.txt
benchstat loops2.txt loops3.txt
*/
