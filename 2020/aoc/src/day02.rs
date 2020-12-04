use std::error::Error;

use crate::utils;

use regex::Regex;

pub async fn solve() -> Result<(), Box<dyn Error>> {
    let rules = utils::start_day(2, &(|l| l)).await?;

    let valid_passwords = part1(&rules)?;
    println!("valid passwords: {}", valid_passwords);

    let new_valid_passwords = part2(&rules)?;
    println!("new valid passwords: {}", new_valid_passwords);

    Ok(())
}

#[derive(Debug)]
struct PasswordRule {
    num1: i32,
    num2: i32,
    character: char,
    password: String,
}

impl PasswordRule {
    fn new(line: &str) -> Result<PasswordRule, &'static str> {
        lazy_static! {
            static ref RE: Regex = Regex::new(r"^(\d+)-(\d+) (.): (.*)$").unwrap();
        }

        let caps = RE.captures(line).ok_or_else(|| "bad match")?;
        let num1 = caps.get(1).unwrap().as_str().parse::<i32>().unwrap();
        let num2 = caps.get(2).unwrap().as_str().parse::<i32>().unwrap();
        let character = string_to_char(caps.get(3).unwrap().as_str().to_string()).unwrap();
        let password = caps.get(4).unwrap().as_str().to_string();

        Ok(PasswordRule {
            num1,
            num2,
            character,
            password,
        })
    }

    fn is_valid_password_old(&self) -> bool {
        // Old rule assumes num1 and num2 are min and max values for
        // character to appear under
        let mut num_matches: i32 = 0;

        for c in self.password.chars() {
            if c == self.character {
                num_matches += 1;
            }
        }

        (num_matches >= self.num1) && (num_matches <= self.num2)
    }

    fn is_valid_password_new(&self) -> bool {
        // New rule assumes num1 and num2 are indices for character
        // to appear exactly once at either index (XOR)
        let ind1 = (self.num1 - 1) as usize;
        let ind2 = (self.num2 - 1) as usize;
        let chars: Vec<char> = self.password.chars().collect();
        (chars[ind1] == self.character) ^ (chars[ind2] == self.character)
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

        if rule.is_valid_password_old() {
            valid_passwords += 1;
        }
    }

    Ok(valid_passwords)
}

fn part2(input: &Vec<String>) -> Result<i32, Box<dyn Error>> {
    let mut valid_passwords: i32 = 0;

    for line in input {
        let rule = PasswordRule::new(line)?;

        if rule.is_valid_password_new() {
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

    #[test]
    fn part2_validate_passwords() {
        let input = vec!["1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"]
            .iter()
            .map(|&s| String::from(s))
            .collect();

        let valid_passwords = part2(&input).unwrap();

        assert_eq!(valid_passwords, 1);
    }
}
