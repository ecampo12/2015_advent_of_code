package itemShop

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Item struct {
	name   string
	cost   int
	Damage int
	Armor  int
}

type Shop struct {
	weapons []Item
	armors  []Item
	rings   []Item
}

type Loadout struct {
	Weapon    Item
	Armor     Item
	Ring1     Item
	Ring2     Item
	TotalCost int
}

func (s *Shop) SetUpShop(p2 bool) {
	shop := Shop{}
	shopData, err := os.ReadFile("itemShop/itemshop.txt")
	if err != nil {
		fmt.Print("Error reading shop data\n")
		panic(err)
	}

	sd := strings.Split(string(shopData), "\n\n")
	weaponData := strings.Split(sd[0], "\n")[1:]
	armorData := strings.Split(sd[1], "\n")[1:]
	ringData := strings.Split(sd[2], "\n")[1:]

	var name string
	var cost int
	var damage int
	var armor int
	for _, weapon := range weaponData {
		fmt.Sscanf(weapon, "%s %d %d %d", &name, &cost, &damage, &armor)
		shop.weapons = append(shop.weapons, Item{name, cost, damage, armor})
	}

	for _, a := range armorData {
		fmt.Sscanf(a, "%s %d %d %d", &name, &cost, &damage, &armor)
		shop.armors = append(shop.armors, Item{name, cost, damage, armor})
	}

	shop.armors = append(shop.armors, Item{"none", 0, 0, 0})

	var name2 int
	for _, ring := range ringData {
		fmt.Sscanf(ring, "%s %d %d %d %d", &name, &name2, &cost, &damage, &armor)
		name = fmt.Sprintf("%s +%d", name, name2)
		shop.rings = append(shop.rings, Item{name, cost, damage, armor})
	}

	shop.rings = append(shop.rings, Item{"none1", 0, 0, 0})
	shop.rings = append(shop.rings, Item{"none2", 0, 0, 0})

	*s = shop
}

func (s *Shop) AllEquipmentLoadouts() []Loadout {
	var loadouts []Loadout
	for _, weapon := range s.weapons {
		for _, armor := range s.armors {
			for _, ring := range s.rings {
				for _, ring2 := range s.rings {
					if ring2 != ring {
						loadouts = append(loadouts, Loadout{
							weapon,
							armor,
							ring,
							ring2,
							weapon.cost + armor.cost + ring.cost + ring2.cost,
						})
					}
				}
			}
		}
	}

	// sorting the loadouts by cost
	sort.Slice(loadouts, func(i, j int) bool {
		return loadouts[i].TotalCost < loadouts[j].TotalCost
	})

	return loadouts
}
