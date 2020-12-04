use std::collections::{HashMap, HashSet};
use std::error::Error;

use crate::utils;

// parse numbers from string
// split numbers into two groups: ones <= to half the total, and ones > half the total
// store the difference of number and total in first group, store only the number in second group
// if two numbers exist s.t. they sum to total, the intersection of these two sets will give the pair

pub async fn solve() -> Result<(), Box<dyn Error>> {
    let numbers = utils::start_day(1, &(|l| l.parse::<i32>().unwrap())).await?;
    let total = 2020;

    let (a1, a2) = part1(&numbers, total)?;
    println!("{} * {} = {}", a1, a2, a1 * a2);

    let (b1, b2, b3) = part2(&numbers, total)?;
    println!("{} * {} * {} = {}", b1, b2, b3, b1 * b2 * b3);

    Ok(())
}

fn part1(input: &Vec<i32>, total: i32) -> Result<(i32, i32), Box<dyn Error>> {
    let half_total = total / 2;

    let mut lower_numbers: HashSet<i32> = HashSet::new();
    let mut higher_numbers: HashSet<i32> = HashSet::new();

    for &number in input {
        if number <= half_total {
            lower_numbers.insert(total - number);
        } else if number > half_total {
            higher_numbers.insert(number);
        }
    }

    for &number in lower_numbers.intersection(&higher_numbers) {
        let (a, b) = (number, total - number);
        return Ok((a, b));
    }

    Ok((0, 0))
}

fn part2(input: &Vec<i32>, total: i32) -> Result<(i32, i32, i32), Box<dyn Error>> {
    let half_total = total / 2;

    let mut lower_numbers: HashSet<i32> = HashSet::new();
    let mut higher_numbers: HashSet<i32> = HashSet::new();
    let mut tuples: HashMap<i32, (i32, i32, i32)> = HashMap::new();

    for i in 0..input.len() {
        let a = input[i];

        // Make sure to add all numbers individually to the HashSets
        if a <= half_total {
            lower_numbers.insert(total - a);
        } else {
            higher_numbers.insert(a);
        }

        for j in i..input.len() {
            let b = input[j];

            // Also add the pair's sum
            let sub_sum = a + b;
            tuples.insert(sub_sum, sort_triplet((a, b, total - sub_sum)));

            if sub_sum <= half_total {
                lower_numbers.insert(total - sub_sum);
            } else if sub_sum < total {
                higher_numbers.insert(sub_sum);
            }
        }
    }

    for &number in lower_numbers.intersection(&higher_numbers) {
        if tuples.contains_key(&number) {
            let (a, b, c) = tuples[&number];
            return Ok((a, b, c));
        }
    }

    Ok((0, 0, 0))
}

fn sort_triplet(tuple: (i32, i32, i32)) -> (i32, i32, i32) {
    let mut vector = vec![tuple.0, tuple.1, tuple.2];
    vector.sort();
    (vector[0], vector[1], vector[2])
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part1_pair_found() {
        let input = vec![1721, 979, 366, 299, 675, 1456];
        let total = 2020;

        let (a, b) = part1(&input, total).unwrap();

        assert_eq!(a, 1721);
        assert_eq!(b, 299);
    }

    #[test]
    fn part2_triplet_found() {
        let input = vec![1721, 979, 366, 299, 675, 1456];
        let total = 2020;

        let (a, b, c) = part2(&input, total).unwrap();

        assert_eq!(a, 366);
        assert_eq!(b, 675);
        assert_eq!(c, 979);
    }
}
