package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct{ r, c int }

var maze [][]byte

func val(p Pos) byte {
	if p.r < 0 || p.r >= len(maze) || p.c < 0 || p.c >= len(maze[0]) {
		return '#'
	}
	return maze[p.r][p.c]
}

func countAdjacent(p Pos) int {
	seen := map[Pos]bool{p: true}
	q := []Pos{p}
	for len(q) > 0 {
		p, q = q[0], q[1:]
		for dr := -1; dr <= 1; dr++ {
			for dc := -1; dc <= 1; dc++ {
				if dr == 0 && dc == 0 {
					continue
				}
				n := Pos{p.r + dr, p.c + dc}
				if val(n) != '#' && !seen[n] {
					seen[n] = true
					q = append(q, n)
				}
			}
		}
	}
	return len(seen)
}

func explode(p Pos, radius int) map[Pos]bool {
	cover := map[Pos]bool{}
	for d := -radius; d <= radius; d++ {
		cover[Pos{p.r + d, p.c}] = true
		cover[Pos{p.r, p.c + d}] = true
	}
	return cover
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		maze = append(maze, []byte(scanner.Text()))
	}

	bombs := map[Pos]int{}
	for r, row := range maze {
		for c, v := range row {
			if '0' <= v && v <= '9' {
				p := Pos{r, c}
				bombs[p] = countAdjacent(p) + int(v-'0') - 1
			}
		}
	}

	exploded := map[Pos]bool{}
	for p, radius := range bombs {
		e := explode(p, radius)
		for n := range e {
			exploded[n] = true
		}
	}

	for r, row := range maze {
		for c := range row {
			if exploded[Pos{r, c}] {
				maze[r][c] = '*'
			}
		}
	}
	for _, l := range maze {
		fmt.Println(string(l))
	}
}
