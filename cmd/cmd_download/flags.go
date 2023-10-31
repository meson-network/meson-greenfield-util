package cmd_download

import "github.com/urfave/cli/v2"

func GetFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{Name: "data", Required: false},
		&cli.StringFlag{Name: "file_config", Required: false},
		&cli.StringFlag{Name: "dest_dir", Required: false},
		&cli.StringFlag{Name: "thread", Required: false},
		&cli.IntFlag{Name: "retry_times", Required: false},
		&cli.BoolFlag{Name: "no_resume", Required: false},
	}
}

func ReadParam(clictx *cli.Context) (string, string, string, int, int, bool) {
	dataName := clictx.String("data")
	jsonConfigAddress := clictx.String("file_config")
	destDir := clictx.String("dest_dir")
	thread := clictx.Int("thread")
	retry_times := clictx.Int("retry_times")
	noResume := clictx.Bool("no_resume")

	return dataName, jsonConfigAddress, destDir, thread, retry_times, noResume
}
