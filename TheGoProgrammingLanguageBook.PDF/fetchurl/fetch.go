package fetchurl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func GetURL() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(string(url), "https://") {
			resp, err := http.Get(url)
			if err != nil {
			}
			b, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: reading %s : %v\n", string(url), err)
				os.Exit(1)
			}
			fmt.Printf("%s", b)
		}
	}
}
