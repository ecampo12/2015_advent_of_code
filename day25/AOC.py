import re
from time import perf_counter

def get_triangle_code(coord):
    row, col = coord
    triangle = (row + col-1) * (row + col) // 2
    return triangle - row + 1

# I'm sure theres a way to calculate repeated multiplication and modulos, but I don't know how to do that.
def part1(end_coord):
    iter = get_triangle_code(end_coord)
    start = 20151125
    for _ in range(iter-1):
        start = (start * 252533) % 33554393

    return start

def main():
    input = open("input.txt", "r").read()
    row, col = re.findall(r"(\d+)", input)
    end_coord = (int(row), int(col))

    t1 = perf_counter()
    part1_ans = part1(end_coord)
    t2 = perf_counter()
    print(f"Part 1: {part1_ans}")
    print(f'Time: {t2 - t1:.4f} seconds')

if __name__ == "__main__":
    main()