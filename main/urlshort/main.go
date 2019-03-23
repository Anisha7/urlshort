package urlshort

import (
	"flag"
	"fmt"
	"net/http"
	"reflect"

	"gopkg.in/yaml.v2"
)

var path = flag.String("-path", "", "path to get")

func main() {
	flag.Parse()
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	if *path != "" {
		url, prs := pathsToUrls[*path]
		if prs == true {
			// you got the url
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("HERE1")
				fmt.Fprintln(w, url)
			})
		}
		// else, path not in map
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("HERE2")
			fmt.Fprintln(w, fallback)
		})

	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("HERE3")
		fmt.Fprintln(w, fallback)
	})
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yaml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yml []byte) (map[string]string, error) { // note: check what type to return
	out := ""
	item := yaml.Unmarshal(yml, &out)
	fmt.Println(out)
	fmt.Println(item)
	fmt.Println(reflect.TypeOf(item))
	return nil, fmt.Errorf("parsed yaml")
}

func buildMap(parsedYaml map[string]string) map[string]string { // specify type
	return nil
}
