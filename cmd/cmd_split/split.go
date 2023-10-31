package cmd_split

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/meson-network/meson-file-util/basic"
	"github.com/meson-network/meson-file-util/basic/color"
	"github.com/meson-network/meson-file-util/src/split"
)

func Split(clictx *cli.Context) error {
	fmt.Println(color.Green(basic.Logo))

	originFilePath, destDir, sizeStr, thread := ReadParam(clictx)

	return split.Split(originFilePath, destDir, sizeStr, thread)
}
