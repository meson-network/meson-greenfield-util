package data

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const DefaultIndexBaseUrl = "https://greenfield-sp.bnbchain.org/download/meson-greenfield-data/data-index/"

type DataIndexInfo struct {
	Name           string `json:"name"`
	Config_address string `json:"config_address"`
	Size           int64  `json:"size"`
	Description    string `json:"description"`
}

func GetDataIndex(targetPath string, timeoutSec int) ([]*DataIndexInfo, error) {
	// download json
	client := &http.Client{
		Timeout: time.Duration(timeoutSec) * time.Second,
	}

	req, err := http.NewRequest("GET", targetPath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var dataIndex = []*DataIndexInfo{}
	err = json.Unmarshal(content, &dataIndex)
	if err != nil {
		return nil, err
	}

	return dataIndex, nil
}
