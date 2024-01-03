def parse_input(input):
    dims = []
    for line in input:
        l, w, h = list(map(int, line.split('x')))
        dims.append([l, w, h])
    return dims

def part1(dims):
    total = 0
    for dim in dims:
        l, w, h = dim
        s1 = l * w
        s2 = w * h
        s3 = h * l
        small = min(s1, s2, s3)
        total += 2 * (s1 + s2 + s3) + small
    return total

def part2(dims):
    total = 0
    for dim in dims:
        l, w, h = dim
        r1, r2, _ = sorted(dim)
        total += 2 * r1 + 2 * r2 + l * w * h
    return total

def main():
    input = open("input.txt", "r").read().splitlines()
    part1_sum = part1(parse_input(input))
    part2_sum = part2(parse_input(input))
    print(f"Part 1: {part1_sum}")
    print(f"Part 2: {part2_sum}")

if __name__ == "__main__":
    main()