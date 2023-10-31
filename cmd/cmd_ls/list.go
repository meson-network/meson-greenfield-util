package cmd_ls

import (
	"github.com/urfave/cli/v2"

	"github.com/meson-network/meson-file-util/src/ls"
)

func ListIndex(clictx *cli.Context) error {
	dataType, showDetail := ReadParam(clictx)
	return ls.ListIndex(dataType, showDetail)
}
