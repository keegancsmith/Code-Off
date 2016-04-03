package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct{ r, c int }

var maze [][]byte
var parent = map[Pos]Pos{}
var delta = []Pos{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func seen(p Pos) bool {
	_, seen := parent[p]
	return seen
}

func canVisit(p Pos) bool {
	// bounds
	if p.r < 0 || p.r >= len(maze) || p.c < 0 || p.c >= len(maze[0]) {
		return false
	}
	// prevent back-tracking
	if seen(p) {
		return false
	}
	// wall or start point
	if maze[p.r][p.c] == '#' || maze[p.r][p.c] == '@' {
		return false
	}
	return true
}

func bfs(p Pos) Pos {
	q := []Pos{p}
	for len(q) > 0 {
		p, q = q[0], q[1:]
		if maze[p.r][p.c] == 'U' {
			return p
		}
		for _, d := range delta {
			n := Pos{p.r + d.r, p.c + d.c}
			if canVisit(n) {
				parent[n] = p
				q = append(q, n)
			}
		}
	}
	panic("no path found")
}

func path(p Pos) {
	for seen(p) {
		maze[p.r][p.c] = '.'
		p = parent[p]
	}
}

func start() Pos {
	for r, l := range maze {
		for c, v := range l {
			if v == '@' {
				return Pos{r, c}
			}
		}
	}
	panic("no start")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		maze = append(maze, []byte(scanner.Text()))
	}
	s := start()
	e := bfs(s)
	path(parent[e])
	for _, l := range maze {
		fmt.Println(string(l))
	}
}
