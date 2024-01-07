
fn look_ans_say(input: &String) -> String{
    let chars = input.chars().collect::<Vec<char>>();
    let mut output = String::new();
    let mut count = 0;
    let mut look_at = chars[0];
    // print!("{}\n", input);
    chars.iter().for_each(|c| {
        if *c == look_at {
            count += 1;
        } else {
            output.push_str(&count.to_string());
            output.push(look_at);
            count = 1;
            look_at = *c;
        }
    });
    output.push_str(&count.to_string());
    output.push(look_at);
    return output;
}

fn part1(input: String) -> String {
    let mut input = input;
    for _ in 0..40 {
        input = look_ans_say(&input);
    }
    input
}

fn part2(input: String) -> String {
    let mut input = input;
    for _ in 0..50 {
        input = look_ans_say(&input);
    }
    input
}

fn main() {
    let input = std::fs::read_to_string("./input.txt").expect("Error reading input file");
    println!("Part 1: {}", part1(input.clone()).len());
    println!("Part 2: {}", part2(input).len());
}

#[cfg(test)]
mod tests {
    #[test]
    fn test_look_ans_say() {
        let test_input =
            std::fs::read_to_string("./test_input.txt").expect("Error reading test input file");
        let mut results = Vec::<String>::new();
        for line in test_input.lines() {
            results.push(super::look_ans_say(&line.to_string()));
        }
        let expected = vec![
            "11".to_string(),
            "21".to_string(),
            "1211".to_string(),
            "111221".to_string(),
            "312211".to_string(),
        ];
        assert_eq!(results, expected);
    }
}
