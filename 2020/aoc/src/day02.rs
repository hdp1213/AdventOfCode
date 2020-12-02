use std::error::Error;

use crate::utils;

use regex::Regex;

pub async fn solve() -> Result<(), Box<dyn Error>> {
    let day: i32 = 2;
    utils::load_input(day).await?;
    let input_file = utils::input_dir().join(format!("day{:02}", day));

    println!("loading input for day 02...");
    let lines = utils::read_file_lines(input_file)?;

    let mut rules = Vec::new();

    for line in lines {
        rules.push(line?);
    }

    let valid_passwords = part1(&rules)?;
    println!("valid passwords: {}", valid_passwords);

    Ok(())
}

#[derive(Debug)]
struct PasswordRule {
    min: i32,
    max: i32,
    character: char,
    password: String,
}

impl PasswordRule {
    fn new(line: &str) -> Result<PasswordRule, &'static str> {
        lazy_static! {
            static ref RE: Regex = Regex::new(r"^(\d+)-(\d+) (.): (.*)$").unwrap();
        }

        let caps = RE.captures(line).ok_or_else(|| "bad match")?;
        let min = caps.get(1).unwrap().as_str().parse::<i32>().unwrap();
        let max = caps.get(2).unwrap().as_str().parse::<i32>().unwrap();
        let character = string_to_char(caps.get(3).unwrap().as_str().to_string()).unwrap();
        let password = caps.get(4).unwrap().as_str().to_string();

        Ok(PasswordRule {
            min,
            max,
            character,
            password,
        })
    }

    fn valid_password(&self) -> bool {
        let mut num_matches: i32 = 0;

        for c in self.password.chars() {
            if c == self.character {
                num_matches += 1;
            }
        }

        (num_matches >= self.min) && (num_matches <= self.max)
    }
}

fn string_to_char(string: String) -> Result<char, &'static str> {
    if string.len() != 1 {
        return Err("string is too long");
    }

    let char_vec: Vec<char> = string.chars().collect();
    Ok(char_vec[0])
}

fn part1(input: &Vec<String>) -> Result<i32, Box<dyn Error>> {
    let mut valid_passwords: i32 = 0;

    for line in input {
        let rule = PasswordRule::new(line)?;

        if rule.valid_password() {
            valid_passwords += 1;
        }
    }

    Ok(valid_passwords)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part1_validate_passwords() {
        let input = vec!["1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"]
            .iter()
            .map(|&s| String::from(s))
            .collect();

        let valid_passwords = part1(&input).unwrap();

        assert_eq!(valid_passwords, 2);
    }
}
