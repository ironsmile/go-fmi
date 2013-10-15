package main

import "path"

func parsePath(input_path string) string {
	var out = path.Clean(path.Join("/", input_path))
	if len(out) == 1 {
		return out
	}
	return out + "/"
}
