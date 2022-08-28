package bing

import (
	"reflect"
	"testing"
)

func TestFileUtils_ReadBing(t *testing.T) {
	type fields struct {
		readmePath string
		bingPath   string
		monthPath  string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*Image
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "1", fields: fields{bingPath: "./bing-wallpaper.md"}, want: nil, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FileUtils{
				readmePath: tt.fields.readmePath,
				bingPath:   tt.fields.bingPath,
				monthPath:  tt.fields.monthPath,
			}
			got, err := f.ReadBing()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadBing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadBing() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}
