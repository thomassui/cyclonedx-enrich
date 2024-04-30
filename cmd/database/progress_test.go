package database

import (
	"io"
	"reflect"
	"testing"

	"github.com/cheggaaa/pb"
)

func TestProgressBarConfig(t *testing.T) {
	type args struct {
		bar    *pb.ProgressBar
		prefix string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ProgressBarConfig(tt.args.bar, tt.args.prefix)
		})
	}
}

func TestProgressBar_TrackProgress(t *testing.T) {
	type args struct {
		src         string
		currentSize int64
		totalSize   int64
		stream      io.ReadCloser
	}
	tests := []struct {
		name string
		cpb  *ProgressBar
		args args
		want io.ReadCloser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cpb.TrackProgress(tt.args.src, tt.args.currentSize, tt.args.totalSize, tt.args.stream); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProgressBar.TrackProgress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readCloser_Close(t *testing.T) {
	tests := []struct {
		name    string
		c       *readCloser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Close(); (err != nil) != tt.wantErr {
				t.Errorf("readCloser.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
