import re
from math import sqrt, floor
from time import perf_counter

def get_next_coordinate(coord):
    if coord[0] == 1:
        return (coord[1] + 1, 1)
    else:
        return (coord[0] - 1, coord[1] + 1)

# I know that the way nmubers are generated is a triangle number, but I don't know how to calculate it.
# I'm sure theres a way to calculate repeated multiplication and modulos, but I don't know how to do that.
def part1(start, start_coord, end_coord):
    while (start_coord) != end_coord:
        start_coord = get_next_coordinate(start_coord)
        start = (start * 252533) % 33554393

    return start

def main():
    input = open("input.txt", "r").read()
    start = 27995004
    start_coord = (6, 6)
    row, col = re.findall(r"(\d+)", input)
    end_coord = (int(row), int(col))

    t1 = perf_counter()
    part1_ans = part1(start, start_coord, end_coord)
    t2 = perf_counter()
    print(f"Part 1: {part1_ans}")
    print(f'Time: {t2 - t1:.4f} seconds')

if __name__ == "__main__":
    main()