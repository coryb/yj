package main

import (
	//"bufio"
	"io/ioutil"
	"fmt"
	"os"
	"gopkg.in/coryb/yaml.v2"
	"encoding/json"
)

func yamlFixup(data interface{}) (interface{}, error) {
	switch d := data.(type) {
	case map[interface{}]interface{}:
		// need to copy this map into a string map so json can encode it
		copy := make(map[string]interface{})
		for key, val := range d {
			switch k := key.(type) {
			case string:
				if fixed, err := yamlFixup(val); err != nil {
					return nil, err
				} else {
					copy[k] = fixed
				}
			default:
				err := fmt.Errorf("YAML: key %s is type '%T', require 'string'", key, k)
				return nil, err
			}
		}
		return copy, nil

	case map[string]interface{}:
		copy := make(map[string]interface{})
		for k, v := range d {
			if fixed, err := yamlFixup(v); err != nil {
				return nil, err
			} else {
				copy[k] = fixed
			}
		}
		return copy, nil
	case []interface{}:
		copy := make([]interface{}, 0, len(d))
		for _, val := range d {
			if fixed, err := yamlFixup(val); err != nil {
				return nil, err
			} else {
				copy = append(copy, fixed)
			}
		}
		return copy, nil
	case string:
		if d == "" || d == "\n" {
			return nil, nil
		}
		return d, nil
	default:
		return d, nil
	}
}

func main() {
	// reader := bufio.NewReader(os.Stdin)
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("Failed to read from STDIN: %s\n", err);
		os.Exit(1)
	}
	var data interface{}
	if err := yaml.Unmarshal(in, &data); err != nil {
		fmt.Printf("Failed to parse input as YAML: %s\n", err)
		os.Exit(1)
	}
	fixed, err := yamlFixup(data)
	if err != nil {
		fmt.Printf("Failed to prepare YAML data for JSON: %s\n", err)
		os.Exit(1)
	}
	if out,err := json.MarshalIndent(fixed, "", "  "); err != nil {
		fmt.Printf("Failed to create JSON from %v: %s\n", data, err)
		os.Exit(1)
	} else {
		fmt.Print(string(out))
	}
}
