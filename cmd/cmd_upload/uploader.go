package cmd_upload

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/meson-network/meson-file-util/basic"
	"github.com/meson-network/meson-file-util/basic/color"
	"github.com/meson-network/meson-file-util/src/uploader/uploader_r2"
)

func Uploader(clictx *cli.Context) error {

	fmt.Println(color.Green(basic.Logo))

	originDir, thread, bucketName, additional_path,
		accountId, accessKeyId, accessKeySecret, retry_times := ReadParam(clictx)

	return uploader_r2.Upload_r2(originDir, thread, bucketName, additional_path,
		accountId, accessKeyId, accessKeySecret, retry_times)
}
