import re

def part1(input):
    start = input.split('\n')[-1]
    input = input.split('\n')[:-2]
    molecules = set()
    for line in input:
        m, replace = line.split(" => ")
        for i in range(len(start)):
            if start[i:i+len(m)] == m:
                molecules.add(start[:i] + replace + start[i+len(m):])
    return len(molecules)
    
# Works on the real input, but not on the test input
def part2(input):
    start = input.split('\n')[-1][::-1] # reverse string
    replacements = {}
    for m in re.findall(r'(\w+) => (\w+)', input):
        replacements[m[1][::-1]] = m[0][::-1]
    def rep(x):
        return replacements[x.group()]
    steps = 0
    while start != 'e':
        start = re.sub('|'.join(replacements.keys()), rep, start, 1)
        steps += 1
    return steps

def main():
    input = open("input.txt", "r").read()
    part1_sum = part1(input)
    print(f"Part 1: {part1_sum}")
    part2_sum = part2(input)
    print(f"Part 2: {part2_sum}")

if __name__ == "__main__":
    main()