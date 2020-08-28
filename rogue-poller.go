package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

/*
Plan of attack:
- Define struct for content of site
- Pull HTML
- Emit Logs
- Figure out process to notify via email
- Figure out process to notify via twitter
- Poller configuration
- Read YML file including list of product pages
*/

func main() {
	configYmlPtr := flag.String("config", "config.yml", "YML configuration for products to monitor.")
	loadProductConfig(*configYmlPtr)
}

func loadProductConfig(config string) {
	dat, err := ioutil.ReadFile(config)
	if err != nil {
		fmt.Println("Configuration failed to load from", config)
		panic(err)
	}
	configMap := make(map[interface{}]interface{})
	err = yaml.Unmarshal(dat, &configMap)

	if err != nil {
		fmt.Println("Configuration failed to load from", config)
		panic(err)
	}

	fmt.Printf("--- m:\n%v\n\n", configMap)
}
