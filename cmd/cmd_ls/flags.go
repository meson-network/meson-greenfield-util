package cmd_ls

import "github.com/urfave/cli/v2"

func GetFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{Name: "type", Required: true},
		&cli.StringFlag{Name: "detail", Required: false},
	}
}

func ReadParam(clictx *cli.Context) (string,bool) {
	dataType := clictx.String("type")
	showDetail := clictx.Bool("detail")
	return dataType,showDetail
}