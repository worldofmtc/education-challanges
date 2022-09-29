package ortalamahesaplama_test

import (
	"github.com/go-turk/education-challanges/ortalamahesaplama"
	"testing"
)

func TestOrtalamaHesapla(t *testing.T) {
	type args struct {
		vizeNot   int
		finalNot  int
		insiyatif float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "OrtalamaHesapla",
			args: args{
				vizeNot:   60,
				finalNot:  70,
				insiyatif: 0.3,
			},
			want: 66,
		},
		{
			name: "OrtalamaHesapla #2",
			args: args{
				vizeNot:   50,
				finalNot:  50,
				insiyatif: 0.2,
			},
			// (50*0.4)+(50*0.6) = 50
			// 50*(1+0.2) = 60
			want: 50,
		},
		{
			name: "OrtalamaHesapla #3",
			args: args{
				vizeNot:   20,
				finalNot:  70,
				insiyatif: 0.1,
			},
			want: 50,
		},
		{
			name: "OrtalamaHesapla #4",
			args: args{
				vizeNot:   100,
				finalNot:  20,
				insiyatif: 0.5,
			},
			want: 65,
		},
		{
			name: "OrtalamaHesapla #5",
			args: args{
				vizeNot:   30,
				finalNot:  80,
				insiyatif: 0.2,
			},
			want: 65,
		},
		{
			name: "OrtalamaHesapla #6",
			args: args{
				vizeNot:   70,
				finalNot:  23,
				insiyatif: 0.35,
			},
			want: 41.8,
		},
		{
			name: "OrtalamaHesapla #7",
			args: args{
				vizeNot:   100,
				finalNot:  100,
				insiyatif: 0.5,
			},
			want: 100,
		},
		{
			name: "OrtalamaHesapla #8",
			args: args{
				vizeNot:   0,
				finalNot:  100,
				insiyatif: 0.5,
			},
			want: 65,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ortalamahesaplama.OrtalamaHesapla(tt.args.vizeNot, tt.args.finalNot, tt.args.insiyatif); got != tt.want {
				t.Errorf("Girdiler %d %d %.2f olan %s test'i başarısız olmuştur, çıktı: %.2f beklenen: %.2f", tt.args.vizeNot, tt.args.finalNot, tt.args.insiyatif, tt.name, got, tt.want)
			}
		})
	}
}
