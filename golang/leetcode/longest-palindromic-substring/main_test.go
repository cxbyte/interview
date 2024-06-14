package longest_palindromic_substring

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{s: `babad`},
			want: "bab",
		},
		{
			name: "2",
			args: args{s: `cbbd`},
			want: "bb",
		},
		{
			name: "3",
			args: args{s: `ac`},
			want: "a",
		},
		{
			name: "4",
			args: args{s: `abb`},
			want: "bb",
		},
		{
			name: "4",
			args: args{s: `abbbb`},
			want: "bbbb",
		},
		{
			name: "5",
			args: args{s: `bbbba`},
			want: "bbbb",
		},
		{
			name: "6",
			args: args{s: `a`},
			want: "a",
		},
		{
			name: "7",
			args: args{s: `bb`},
			want: "bb",
		},
		{
			name: "8",
			args: args{s: `zudfweormatjycujjirzjpyrmaxurectxrtqedmmgergwdvjmjtstdhcihacqnothgttgqfywcpgnuvwglvfiuxteopoyizgehkwuvvkqxbnufkcbodlhdmbqyghkojrgokpwdhtdrwmvdegwycecrgjvuexlguayzcammupgeskrvpthrmwqaqsdcgycdupykppiyhwzwcplivjnnvwhqkkxildtyjltklcokcrgqnnwzzeuqioyahqpuskkpbxhvzvqyhlegmoviogzwuiqahiouhnecjwysmtarjjdjqdrkljawzasriouuiqkcwwqsxifbndjmyprdozhwaoibpqrthpcjphgsfbeqrqqoqiqqdicvybzxhklehzzapbvcyleljawowluqgxxwlrymzojshlwkmzwpixgfjljkmwdtjeabgyrpbqyyykmoaqdambpkyyvukalbrzoyoufjqeftniddsfqnilxlplselqatdgjziphvrbokofvuerpsvqmzakbyzxtxvyanvjpfyvyiivqusfrsufjanmfibgrkwtiuoykiavpbqeyfsuteuxxjiyxvlvgmehycdvxdorpepmsinvmyzeqeiikajopqedyopirmhymozernxzaueljjrhcsofwyddkpnvcvzixdjknikyhzmstvbducjcoyoeoaqruuewclzqqqxzpgykrkygxnmlsrjudoaejxkipkgmcoqtxhelvsizgdwdyjwuumazxfstoaxeqqxoqezakdqjwpkrbldpcbbxexquqrznavcrprnydufsidakvrpuzgfisdxreldbqfizngtrilnbqboxwmwienlkmmiuifrvytukcqcpeqdwwucymgvyrektsnfijdcdoawbcwkkjkqwzffnuqituihjaklvthulmcjrhqcyzvekzqlxgddjoir`},
			want: "gykrkyg",
		},
		{
			name: "9",
			args: args{s: `aacabdkacaa`},
			want: "aca",
		},
		{
			name: "10",
			args: args{s: `tattarrattat`},
			want: "tattarrattat",
		},
		{
			name: "11",
			args: args{s: `tattarrrattat`},
			want: "tattarrrattat",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
