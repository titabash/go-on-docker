package utility

import (
	. "app/utility/logger"
	"io/ioutil"
	"net/http"
)

func GetImageFromWeb(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	Log.Debugf("StatusCode: %v, Status: %s", resp.StatusCode, resp.Status)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
