package sort

import (
	"testing"
)

func TestQuicWiki(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		//want []int
	}{
		// TODO: Add test cases.
		{"wiki_1",
			args{[]int{23, 45, 17, 11, 13, 89, 72, 26, 3, 17, 11, 13}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := QuicWiki(tt.args.data)
			t.Log(res)
		})
	}
}

func TestQuickSort1(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		//want []int
	}{
		// TODO: Add test cases.
		{"quick_sort1",
			args{[]int{23, 45, 17, 11, 13, 89, 72, 26, 3, 17, 11, 13}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := QuickSort1(tt.args.data)
			t.Log(res)
		})
	}
}

func TestQuickSort2(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		//want []int
	}{
		// TODO: Add test cases.
		{"quick_sort2",
			args{[]int{23, 45, 17, 11, 13, 89, 72, 26, 3, 17, 11, 13}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := QuickSort2(tt.args.data)
			t.Log(res)
		})
	}
}
