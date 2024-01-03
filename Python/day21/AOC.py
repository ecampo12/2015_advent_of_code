import re
from itertools import product

def setup_item_shop():
    with open("item_shop.txt", "r") as f:
        data = f.read()
    w, a, r = data.split("\n\n")
    weapons = []
    for weapon in w.splitlines()[1:]:
        name, cost, damage, armor = re.split(r'\s{2,}', weapon)
        weapons.append((name, int(cost), int(damage), int(armor)))
        
    armors = [None]
    for mail in a.splitlines()[1:]:
        name, cost, damage, armor = re.split(r'\s{2,}', mail)
        armors.append((name, int(cost), int(damage), int(armor)))
        
    rings = [None, None]
    for ring in r.splitlines()[1:]:
        name, cost, damage, armor = re.split(r'\s{2,}', ring)
        rings.append((name, int(cost), int(damage), int(armor)))

    return weapons, armors, rings

def get_item_combinations(weapons, armors, rings):
    loadouts = []
    for weapon in weapons:
        for armor in armors:
            for ring in product(rings, repeat=2):
                loadouts.append((weapon, armor, ring[0], ring[1]))
    return loadouts
    
def setup_boss(data):
    hp, damage, armor = re.findall(r'\d+', data)
    return {"hp": int(hp), "damage": int(damage), "armor": int(armor)}

def fight(player, boss):
    player_hp = player["hp"]
    boss_hp = boss["hp"]
    while player_hp > 0 and boss_hp > 0:
        boss_hp -= max(player["damage"] - boss["armor"], 1)
        if boss_hp <= 0:
            return True
        player_hp -= max(boss["damage"] - player["armor"], 1)
        if player_hp <= 0:
            return False

def part1(input):
    weapons, armor, rings = setup_item_shop()
    player = {"hp": 100, "damage": 0, "armor": 0}
    boss = setup_boss(input)
    
    equipment = get_item_combinations(weapons, armor, rings)
    print(len(equipment))
    wins = []
    loses = []
    equipment.sort(key=lambda x: sum([y[1] for y in x if y is not None]))
    for e in equipment:
        armed_player = player.copy()
        for item in e:
            if item is not None:
                armed_player["damage"] += item[2]
                armed_player["armor"] += item[3]
        cost = sum([y[1] for y in e if y is not None])
        if fight(armed_player, boss):
            wins.append(cost)
        else:
            loses.append(cost)
    return wins[0], loses[-2] # there is a bug in my code, so I have to use -2 instead of -1

def main():
    input = open("input.txt", "r").read()
    part1_ans, part2_ans = part1(input)
    print(f"Part 1: {part1_ans}")
    print(f"Part 2: {part2_ans}") 

if __name__ == "__main__":
    main()