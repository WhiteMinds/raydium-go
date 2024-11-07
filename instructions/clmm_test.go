package instructions

import (
	"reflect"
	"testing"
)

func TestDecodeRaydiumData(t *testing.T) {
	type args struct {
		data []byte
	}
	type testCase[T any] struct {
		name    string
		args    args
		want    T
		wantErr bool
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeRaydiumData(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeRaydiumData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeRaydiumData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
