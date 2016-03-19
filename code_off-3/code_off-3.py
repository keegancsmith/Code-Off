#!/usr/bin/env python

import fileinput
from collections import defaultdict

grid = {}
component = {}

def assign(v, p):
    if p in component or p not in grid or grid[p] == '#':
        return
    component[p] = v
    for r, c in ((0, 1), (1, 0), (-1, 0), (0, -1)):
        np = (p[0] + r, p[1] + c)
        assign(v, np)

def main():
    for r, line in enumerate(fileinput.input()):
        for c, b in enumerate(line):
            grid[(r, c)] = b
    nRows, nCols = r + 1, c

    # Populate component, which we can use to tell if two positions are in the
    # same connected component
    for i, p in enumerate(grid):
        assign(i, p)

    # Find the size of each connected component
    sizes = defaultdict(int)
    for v in component.values():
        sizes[v] += 1

    # Find the maximal sized components
    maxSize = max(sizes.values())
    maxComponents = set(k for k, v in sizes.items() if v == maxSize)

    for r in range(nRows):
        line = []
        for c in range(nCols):
            p = (r, c)
            v = grid[p]
            if component.get(p) in maxComponents:
                v = '*'
            line.append(v)
        print ''.join(line)

if __name__ == '__main__':
    main()
