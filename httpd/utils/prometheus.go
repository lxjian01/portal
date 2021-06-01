package utils

import (
	"errors"
	"fmt"
	"net/http"
)

func PrometheusReload(url string) error {
	reloadUrl := fmt.Sprintf("%s%s", url, "/-/reload")
	resp, err := http.Post(reloadUrl,"application/json",nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New("reload prometheus error")
	}
	return nil
}
