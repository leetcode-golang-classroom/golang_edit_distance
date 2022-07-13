package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	word1 := "intention"
	word2 := "execution"
	for idx := 0; idx < b.N; idx++ {
		minDistance(word1, word2)
	}
}
func Test_minDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "word1 = \"horse\", word2 = \"ros\"",
			args: args{word1: "horse", word2: "ros"},
			want: 3,
		},
		{
			name: "word1 = \"intention\", word2 = \"execution\"",
			args: args{word1: "intention", word2: "execution"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
