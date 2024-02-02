package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Ingredient struct {
	name                                            string
	capacity, durability, flavor, texture, calories int
}

type Recipe struct {
	ingredients []Ingredient
	amounts     []int
}

func parseInput(input []string) []Ingredient {
	Ingredients := []Ingredient{}

	var name string
	for _, line := range input {
		var data []int
		name = strings.Split(line, ":")[0]
		nums := regexp.MustCompile(`-?\d`).FindAllString(line, -1)

		for _, num := range nums {
			x, _ := strconv.Atoi(num)
			data = append(data, x)
		}
		Ingredients = append(Ingredients, Ingredient{name, data[0], data[1], data[2], data[3], data[4]})
	}

	return Ingredients
}

func generateIngredientPerm(elements []Ingredient) [][]Ingredient {
	var result [][]Ingredient

	var generate func(int)
	generate = func(n int) {
		if n == 1 {
			temp := make([]Ingredient, len(elements))
			copy(temp, elements)
			result = append(result, temp)
			return
		}

		for i := 0; i < n; i++ {
			elements[i], elements[n-1] = elements[n-1], elements[i]
			generate(n - 1)
			elements[i], elements[n-1] = elements[n-1], elements[i]
		}
	}

	generate(len(elements))
	return result
}

func generateAmountPerm(targetSum, x, start int, currentSum int, currentCombo []int, result *[][]int) {
	if len(currentCombo) == x {
		if currentSum == targetSum {
			temp := make([]int, len(currentCombo))
			copy(temp, currentCombo)
			*result = append(*result, temp)
		}
		return
	}

	for i := start; i <= targetSum-currentSum; i++ {
		generateAmountPerm(targetSum, x, i, currentSum+i, append(currentCombo, i), result)
	}
}

func generateRecipes(ingredients []Ingredient) []Recipe {
	r := []Recipe{}
	combos := [][]int{}
	ingredientsPerm := generateIngredientPerm(ingredients)
	generateAmountPerm(100, len(ingredients), 1, 0, []int{}, &combos)

	for _, combo := range combos {
		for _, ingredientPerm := range ingredientsPerm {
			recipe := Recipe{ingredientPerm, combo}
			r = append(r, recipe)
		}
	}

	return r
}

func negativeToZero(x int) int {
	if x < 0 {
		return 0
	}
	return x
}

func part1(input []string, calcount bool) int {
	ingredients := parseInput(input)
	recipes := generateRecipes(ingredients)
	prods := []int{}

	for _, recipe := range recipes {
		cookie := Ingredient{"cookie", 0, 0, 0, 0, 0}

		for i, ingredient := range recipe.ingredients {
			cookie.capacity += recipe.amounts[i] * ingredient.capacity
			cookie.durability += recipe.amounts[i] * ingredient.durability
			cookie.flavor += recipe.amounts[i] * ingredient.flavor
			cookie.texture += recipe.amounts[i] * ingredient.texture
			cookie.calories += recipe.amounts[i] * ingredient.calories
		}

		cookie.capacity = negativeToZero(cookie.capacity)
		cookie.durability = negativeToZero(cookie.durability)
		cookie.flavor = negativeToZero(cookie.flavor)
		cookie.texture = negativeToZero(cookie.texture)

		if calcount && cookie.calories != 500 {
			continue
		}

		x := cookie.capacity * cookie.durability * cookie.flavor * cookie.texture
		prods = append(prods, x)

	}
	slices.Sort(prods)
	return prods[len(prods)-1]
}

func part2(input []string) int {
	return part1(input, true)
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error reading file\n")
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	part1_ans := part1(lines, false)
	fmt.Printf("Part 1: %d\n", part1_ans)

	part2_ans := part2(lines)
	fmt.Printf("Part 2: %d\n", part2_ans)
}
