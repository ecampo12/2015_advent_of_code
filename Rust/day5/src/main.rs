use regex::Regex;

// Since regex does not support backreferences, we need to use a different approach
fn has_double(line: &str) -> bool {
    let chars = line.chars();
    for (i, c) in chars.clone().enumerate() {
        if i > 0 && c == chars.clone().nth(i - 1).unwrap() {
            return true;
        }
    }
    return false;
}

fn part1(input: &String) -> i32 {
    let has_vowel = Regex::new(r"[aeiou]").unwrap();
    let has_bad = Regex::new(r"ab|cd|pq|xy").unwrap();

    input
        .split("\n")
        .map(|line| {
            if has_vowel.captures_iter(line).count() >= 3
                && has_double(line)
                && !has_bad.is_match(line)
            {
                1
            } else {
                0
            }
        })
        .filter(|&x| x == 1)
        .sum()
}

// Again, regex does not support backreferences, so we need to iterate over the string
fn has_pair_with_one_between(line: &str) -> bool {
    let chars = line.chars();
    for (i, c) in chars.clone().enumerate() {
        if i > 1 && c == chars.clone().nth(i - 2).unwrap() {
            return true;
        }
    }
    return false;
}

// Basically we use a sliding window of size 2 and check if the rest of the string contains the pair
fn has_repeating_pair(line: &str) -> bool {
    let chars = line.chars();
    let mut skip = 2;
    let mut pair = chars.clone().take(2).collect::<String>();
    let mut rest = chars.clone().skip(skip);
    loop {
        if rest.clone().collect::<String>().contains(&pair) {
            return true;
        }
        if rest.clone().take(2).collect::<String>().len() < 2 {
            break;
        }
        pair  = format!("{}{}", pair.chars().nth(1).unwrap(), rest.clone().nth(0).unwrap());
        skip += 1;
        rest = chars.clone().skip(skip);
    }
    return false;
}

fn part2(input: &String) -> i32 {
    input
        .split("\n")
        .map(|line| {
            if has_pair_with_one_between(line) && has_repeating_pair(line) {
                1
            } else {
                0
            }
        })
        .filter(|&x| x == 1)
        .sum()
}

fn main() {
    let input = std::fs::read_to_string("./input.txt").expect("Error reading input file");
    println!("Part 1: {}", part1(&input));
    println!("Part 2: {}", part2(&input));
}

#[cfg(test)]
mod tests {
    #[test]
    fn test_part1() {
        let test_input =
            std::fs::read_to_string("./test_input.txt").expect("Error reading test input file");
        assert_eq!(super::part1(&test_input), 2);
    }

    #[test]
    fn test_part2() {
        let test_input =
            std::fs::read_to_string("./test_input2.txt").expect("Error reading test input file");
        assert_eq!(super::part2(&test_input), 2);
    }
}
