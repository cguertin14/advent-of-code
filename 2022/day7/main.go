package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type dir struct {
	size     int
	parent   *dir
	children []*dir
}

// navigate directory and apply
// function passed in param.
func (d dir) navigate(fn func(dir)) {
	fn(d)
	for _, c := range d.children {
		c.navigate(fn)
	}
}

func main() {
	lines := strings.Split(input, "\n")

	// build file system from root
	var root dir
	current := &root
	for _, line := range lines[1:] {
		splitted := strings.Split(line, " ")
		if strings.Contains(line, "$") {
			if splitted[1] == "cd" {
				if splitted[2] == ".." {
					current = current.parent
				} else {
					var child dir
					child.parent = current
					current.children = append(current.children, &child)
					current = &child
				}
			}
		} else {
			// check if arg is a number or not
			if splitted[0] == "dir" {
				// do nothing, skip directories
			} else {
				size, _ := strconv.Atoi(splitted[0])
				// update total and parent's
				// total as well while there
				// is no more parent
				update := current
				for update != nil {
					update.size += size
					update = update.parent
				}
			}
		}
	}

	// part 1
	sum := 0
	root.navigate(func(d dir) {
		if d.size <= 100000 {
			sum += d.size
		}
	})
	fmt.Println(sum)

	// part 2
	var smallestDir int
	root.navigate(func(d dir) {
		if root.size-d.size < 70000000-30000000 {
			if smallestDir == 0 || d.size < smallestDir {
				smallestDir = d.size
			}
		}
	})
	fmt.Println(smallestDir)
}
