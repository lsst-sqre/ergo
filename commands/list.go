package commands

import (
	"fmt"
	"github.com/cristianoliveira/ergo/proxy"
)

func List(config *proxy.Config) {
	fmt.Println("Ergo Proxy current list: ")
	for _, s := range config.Services {
		localUrl := `http://` + s.Name + config.Domain
		fmt.Printf(" - %s -> %s \n", localUrl, s.Url)
	}
}
