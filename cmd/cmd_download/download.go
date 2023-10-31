package cmd_download

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/meson-network/meson-file-util/basic"
	"github.com/meson-network/meson-file-util/basic/color"
	"github.com/meson-network/meson-file-util/src/download"
)

func Download(clictx *cli.Context) error {
	fmt.Println(color.Green(basic.Logo))

	dataName, jsonConfigAddress, destDir, thread, retryNum, noResume := ReadParam(clictx)
	return download.Download(dataName, jsonConfigAddress, destDir, thread, retryNum, noResume)
}
