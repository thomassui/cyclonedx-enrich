package database

import (
	"context"
	"os"
	"time"

	getter "github.com/hashicorp/go-getter/v2"
)

func download() error {

	filename := os.Getenv("DATABASE_FILE")
	rawURL := os.Getenv("DOWNLOAD_DATABASE_URL")
	defaultProgressBar := &ProgressBar{}

	ctx := context.Background()
	req := &getter.Request{
		Src:              rawURL,
		Dst:              filename,
		GetMode:          getter.ModeFile,
		ProgressListener: defaultProgressBar,
	}
	client := &getter.Client{
		Getters: []getter.Getter{
			&getter.HttpGetter{
				ReadTimeout: 60 * time.Second,
			},
		},
	}

	_, err := client.Get(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
