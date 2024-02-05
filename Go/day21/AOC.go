package main

import (
	"day21/itemShop"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Character struct {
	hp, damage, armor int
}

func fight(player, boss Character) bool {
	playerDamage := int(math.Max(float64(player.damage-boss.armor), 1))
	bossDamage := int(math.Max(float64(boss.damage-player.armor), 1))
	for {
		boss.hp -= playerDamage
		if boss.hp <= 0 {
			return true
		}

		player.hp -= bossDamage
		if player.hp <= 0 {
			return false
		}
	}
}

func setBoss(input []string) Character {
	hp, _ := strconv.Atoi(strings.Split(input[0], " ")[2])
	damage, _ := strconv.Atoi(strings.Split(input[1], " ")[1])
	armor, _ := strconv.Atoi(strings.Split(input[2], " ")[1])

	return Character{hp, damage, armor}
}

func (c *Character) equip(loadout itemShop.Loadout) {
	c.damage = loadout.Weapon.Damage
	c.armor = loadout.Armor.Armor

	if loadout.Ring1.Damage > 0 {
		c.damage += loadout.Ring1.Damage
	} else {
		c.armor += loadout.Ring1.Armor
	}

	if loadout.Ring2.Damage > 0 {
		c.damage += loadout.Ring2.Damage
	} else {
		c.armor += loadout.Ring2.Armor
	}
}

func (c *Character) reset() {
	c.hp = 100
	c.damage = 0
	c.armor = 0
}

func part1(input []string) int {
	shop := itemShop.Shop{}
	shop.SetUpShop(false)
	loadouts := shop.AllEquipmentLoadouts()

	player := Character{100, 0, 0}
	boss := setBoss(input)

	for _, loadout := range loadouts {
		player.reset()
		player.equip(loadout)
		if fight(player, boss) {
			return loadout.TotalCost
		}
	}

	return 0
}

func part2(input []string) int {
	shop := itemShop.Shop{}
	shop.SetUpShop(true)
	loadouts := shop.AllEquipmentLoadouts()
	player := Character{100, 0, 0}
	boss := setBoss(input)

	// going through the loadouts in reverse order to find the most expensive loadout that still loses
	for i := len(loadouts) - 1; i >= 0; i-- {
		player.reset()
		player.equip(loadouts[i])
		if !fight(player, boss) {
			return loadouts[i].TotalCost
		}
	}

	return 0
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error reading file\n")
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	part1_ans := part1(lines)
	fmt.Printf("Part 1: %d\n", part1_ans)

	part2_ans := part2(lines)
	fmt.Printf("Part 2: %d\n", part2_ans)
}
