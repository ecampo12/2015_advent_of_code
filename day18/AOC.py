from time import perf_counter

def parse_input(input):
    grid = []
    for line in input:
        grid.append(list(line))
    return grid
def neighbors(grid, r, c):
    hood = [(-1, 0), (-1, 1), (0, 1), (1, 1), (1, 0), (-1, -1), (0, -1), (1, -1)]
    neighbors = 0
    for dr, dc in hood:
        if r + dr < 0 or r + dr >= len(grid) or c + dc < 0 or c + dc >= len(grid[0]):
            continue
        if grid[r + dr][c + dc] == '#':
            neighbors += 1
    return neighbors

def turn_on(grid, p2=[]):
    on = []
    for r, row in enumerate(grid):
        for c, col in enumerate(row):
            n = neighbors(grid, r, c)
            if col == '#':
                if n == 2 or n == 3:
                    on.append((r, c))
            elif col == '.':
                if n == 3:
                    on.append((r, c))
    if p2:
        for r, c in p2:
            on.append((r, c))
    
    return on
def part1(input, step=4, always_on=[]):
    grid = parse_input(input)
    new_grid = []
    
    for r, c in always_on:
        grid[r][c] = '#'
        
    for _ in range(len(grid)):
        new_grid.append(['.' for _ in range(len(grid[0]))])
        
    for _ in range(step):
        on = turn_on(grid, always_on)
        for r, c in on:
            new_grid[r][c] = '#'
        grid = new_grid
        new_grid = []
        for _ in range(len(grid)):
            new_grid.append(['.' for _ in range(len(grid[0]))])
    return sum([row.count('#') for row in grid])

def part2(input, step=5):
    always_on = [(0, 0), 
                 (0, len(input[0]) - 1), 
                 (len(input) - 1, 0), 
                 (len(input) - 1, len(input[0]) - 1)
            ]
    return part1(input, step, always_on)
    
def main():
    input = open("input.txt", "r").read().splitlines()
    t1 = perf_counter()
    part1_sum = part1(input, 100)
    t2  = perf_counter()
    print(f"Part 1: {part1_sum}")
    print(f'Part 1 took {t2 - t1} seconds\n')
    
    t1 = perf_counter()
    part2_sum = part2(input, 100)
    t2 = perf_counter()
    print(f"Part 2: {part2_sum}")
    print(f'Part 2 took {t2 - t1} seconds')

if __name__ == "__main__":
    main()