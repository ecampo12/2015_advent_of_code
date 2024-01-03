from Game import Game, Character
from time import perf_counter

# Takes a while to run
def part1(input):
    player = Character(50, mana=500)
    game = Game(player, input)
    return game.find_lowest_mana_use()

# Runs faster, probably because the player dies more often
def part2(input):
    player = player = Character(50, mana=500)
    game = Game(player, input)
    return game.find_lowest_mana_use(True)

def main():
    print("Don't worry the program just ttakes a while to run...")
    input = open("input.txt", "r").read()
    t1 = perf_counter()
    part1_ans = part1(input)
    t2 = perf_counter()
    print(f"Part 1: {part1_ans}")
    print(f"Part 1 took {t2 - t1:.4f} seconds\n")
    
    t1 = perf_counter()
    part2_ans = part2(input)
    t2 = perf_counter()
    print(f"Part 2: {part2_ans}")
    print(f"Part 2 took {t2 - t1:.4f} seconds")

if __name__ == "__main__":
    main()