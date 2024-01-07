use regex::Regex;

fn part1(input: &String) -> i32 {
    let mut lights = vec![vec![0; 1000]; 1000];
    let digits = Regex::new(r"\d+").unwrap();
    let lines = input.split("\n");
    for line in lines{
        let nums = digits.captures_iter(line).map(|x| x[0].parse::<usize>().unwrap()).collect::<Vec<usize>>();
        let (x1, y1, x2, y2) = (nums[0], nums[1], nums[2], nums[3]);
        for x in x1..x2+1 {
            for y in y1..y2+1 {
                if line.contains("turn on") {
                    lights[x][y] = 1;
                } else if line.contains("turn off") {
                    lights[x][y] = 0;
                } else if line.contains("toggle") {
                    lights[x][y] = 1 - lights[x][y];
                }
            }
        }
    }
    return lights.iter().map(|x| x.iter().sum::<i32>()).sum();
}

fn part2(input: &String) -> i32 {
    let mut lights = vec![vec![0; 1000]; 1000];
    let digits = Regex::new(r"\d+").unwrap();
    let lines = input.split("\n");
    for line in lines {
        let nums = digits.captures_iter(line).map(|x| x[0].parse::<usize>().unwrap()).collect::<Vec<usize>>();
        let (x1, y1, x2, y2) = (nums[0], nums[1], nums[2], nums[3]);
        for x in x1..x2+1 {
            for y in y1..y2+1 {
                if line.contains("turn on") {
                    lights[x][y] += 1;
                } else if line.contains("turn off") {
                    lights[x][y] = std::cmp::max(0, lights[x][y] - 1);
                } else if line.contains("toggle") {
                    lights[x][y] += 2;
                }
            }
        }
    }
    return lights.iter().map(|x| x.iter().sum::<i32>()).sum();
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
        assert_eq!(super::part1(&test_input), 998996);
    }

    #[test]
    fn test_part2() {
        let test_input =
            std::fs::read_to_string("./test_input2.txt").expect("Error reading test input file");
        assert_eq!(super::part2(&test_input), 2000001);
    }
}