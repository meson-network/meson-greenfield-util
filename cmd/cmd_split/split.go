package cmd_split

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/meson-network/meson-greenfield-util/basic"
	"github.com/meson-network/meson-greenfield-util/basic/color"
	"github.com/meson-network/meson-greenfield-util/src/split"
)

func Split(clictx *cli.Context) error {
	fmt.Println(color.Green(basic.Logo))

	originFilePath, destDir, sizeStr, thread := ReadParam(clictx)

	return split.Split(originFilePath, destDir, sizeStr, thread)
}
