package main

import (
	"bufio"
	"fmt"
	"os"
)

var maze [][]byte
var seen [][]bool
var delta = []struct{ r, c int }{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func dfs(r, c int) bool {
	if r < 0 || r >= len(maze) || c < 0 || c >= len(maze[0]) || seen[r][c] {
		return false
	}
	seen[r][c] = true
	switch maze[r][c] {
	case '#':
		return false
	case '@':
		return true
	}
	for _, d := range delta {
		if !dfs(r+d.r, c+d.c) {
			continue
		}
		if maze[r][c] == ' ' {
			maze[r][c] = '.'
		}
		return true
	}
	return false
}

func start() (int, int) {
	for r, l := range maze {
		for c, v := range l {
			if v == 'U' {
				return r, c
			}
		}
	}
	panic("no start")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		maze = append(maze, []byte(scanner.Text()))
		seen = append(seen, make([]bool, len(maze[0])))
	}
	r, c := start()
	dfs(r, c)
	for _, l := range maze {
		fmt.Println(string(l))
	}
}
