package database

import (
	"context"
	"time"

	"cyclonedx-enrich/utils"

	getter "github.com/hashicorp/go-getter/v2"
)

func download() error {

	filename := utils.Getenv("DATABASE_FILE", "database.db")
	rawURL := utils.Getenv("DOWNLOAD_DATABASE_URL", "https://raw.githubusercontent.com/fnxpt/cyclonedx-enrich-poc/main/database.db")
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
