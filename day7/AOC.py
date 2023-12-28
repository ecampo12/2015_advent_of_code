import re

def part1(input: [str], override = None):
    wires = {}
    while input:
        wait = []
        for line in input:
            if "NOT" in line:
                _, source, _, wire = line.split(" ")
                if source not in wires:
                    wait.append(line)
                else:
                    wires[wire] = abs(wires[source] - 65535)
            elif re.search(r"AND|OR|LSHIFT|RSHIFT", line):
                s1, OP, s2, _, wire = line.split(" ")
                if (not s1.isnumeric() and s1 not in wires) or (not s2.isnumeric() and s2 not in wires):
                    wait.append(line)
                else:
                    match(OP):
                        case "AND":
                            if s1.isnumeric():
                                wires[wire] = int(s1) & wires[s2]
                            else:
                                wires[wire] = wires[s1] & wires[s2]
                        case "OR": wires[wire] = wires[s1] | wires[s2]
                        case "LSHIFT": wires[wire] = wires[s1] << int(s2)
                        case "RSHIFT": wires[wire] = wires[s1] >> int(s2)
            else:
                source, _, wire = line.split(" ")
                if override and wire == "b":
                    wires[wire] = override
                elif source.isnumeric():
                    wires[wire] = int(source)
                elif source in wires:
                    wires[wire] = wires[source]
                else:
                    wait.append(line)
        input = wait
    return wires

def part2(input, override):
    return part1(input, override)

def main():
    input = open("input.txt", "r").read().splitlines()
    part1_sum = part1(input)["a"]
    part2_sum = part2(input, part1_sum)["a"]
    print(f"Part 1: {part1_sum}")
    print(f"Part 2: {part2_sum}")

if __name__ == "__main__":
    main()