package core

import (
	"fmt"
	"io"
	"net/http"

	"github.com/tidwall/pretty"
)

func Display(response *http.Response) error {
	fmt.Println(response.Status)

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return err
	}

	prettyJSON := pretty.Color([]byte(body), nil)
	fmt.Println(string(prettyJSON[:]))

	return nil
}

func Save(response *http.Response) error {
	return nil
}
