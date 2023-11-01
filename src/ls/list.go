package ls

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	humanize "github.com/dustin/go-humanize"
	"github.com/meson-network/meson-greenfield-util/src/data"
)

func ListIndex(dataType string, showDetail bool) error {
	// url
	if dataType == "" {
		fmt.Println("[ERROR] data type error")
		return errors.New("data type error")
	}

	targetUrl := data.DefaultIndexBaseUrl + strings.ToLower(dataType) + ".json"

	fmt.Println("getting data index ...")

	datasets, err := data.GetDataIndex(targetUrl, 10)
	if err != nil {
		fmt.Println("[ERROR]", err.Error())
		return err
	}

	// sort with name
	sort.Slice(datasets, func(i, j int) bool {
		return datasets[i].Name < datasets[j].Name
	})

	for _, v := range datasets {
		if showDetail {
			fmt.Println("--", "Name:", v.Name, "\t", "Size:", humanize.Bytes(uint64(v.Size)), "\t", "Config address:", v.Config_address, "\t", "Description:", v.Description)
		} else {
			fmt.Println("--", "Name:", v.Name, "\t", "Size:", humanize.Bytes(uint64(v.Size)))
		}

	}
	fmt.Println("")

	return err
}
