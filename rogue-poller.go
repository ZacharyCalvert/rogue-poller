package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"strings"

	"gopkg.in/yaml.v2"
)

/*
Plan of attack:
x Define struct for content of site
x Pull HTML
x Emit Logs
- Figure out process to notify via email
- Figure out process to notify via twitter
- Poller configuration
- Read YML file including list of product pages
*/

func main() {
	configYmlPtr := flag.String("config", "config.yml", "YML configuration for products to monitor.")
	config := loadProductConfig(*configYmlPtr)
	products := config["products"].([]interface{})

	htmlCache := make(map[string][]byte)

	for _, entry := range products {
		product := entry.(map[interface{}]interface{})
		page := product["page"].(string)
		label := product["label"].(string)
		id := product["id"].(int)
		if inspectPageForProduct(page, fmt.Sprintf("%d", id), htmlCache) {
			fmt.Printf("https://www.roguefitness.com/%s has stock available of %s\n", page, label)
		}
	}
}

/*
Fun discovery that makes this so easy.  Rogue's HTML renders a single input, per product when
available.  That product button will NOT be hidden.
There is a product page and an ID to search for on that page
*/
func inspectPageForProduct(page string, identifier string, cache map[string][]byte) bool {
	target := "https://www.roguefitness.com/" + page
	resp, err := http.Get(target)
	// handle the error if there is one
	if err != nil {
		fmt.Println("Error obtaining", target, err)
		return false
	}

	defer resp.Body.Close()

	if _, ok := cache[page]; !ok {
		html, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response from", target, err)
			return false
		}

		cache[page] = html
	}

	html := cache[page]

	re := regexp.MustCompile("\n([^\n]*)super_group\\[" + identifier + "\\]([^\n]*)\n")

	found := re.Find([]byte(html))

	result := found != nil && len(found) > 0 && !strings.Contains(string(found), "type=\"hidden\"")

	return result
}

func loadProductConfig(config string) map[interface{}]interface{} {
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

	return configMap
}
