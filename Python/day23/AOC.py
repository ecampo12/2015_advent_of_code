def part1(input, part2=False):
    registers = {"a": 0, "b": 0}
    if part2:
        registers["a"] = 1
    i = 0
    while i < len(input):
        inst = input[i].split()
        match inst[0]:
            case "hlf":
                registers[inst[1]] /= 2
                i += 1
            case "tpl":
                registers[inst[1]] *= 3
                i += 1
            case "inc":
                registers[inst[1]] += 1
                i += 1
            case "jmp":
                i += int(inst[1])
                continue
            case "jie":
                if registers[inst[1][0]] % 2 == 0:
                    i += int(inst[2])
                    continue
                else:
                    i += 1
            case "jio":
                if registers[inst[1][0]] == 1:
                    i += int(inst[2])
                    continue
                else:
                    i += 1
    return registers

def part2(input):
    return part1(input, part2=True)

def main():
    input = open("input.txt", "r").read().splitlines()
    part1_ans = part1(input)
    print(f"Part 1: {part1_ans['b']}")
    part2_ans = part2(input)
    print(f"Part 2: {part2_ans['b']}")

if __name__ == "__main__":
    main()