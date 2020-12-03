use std::collections::HashSet;
use std::error::Error;

use crate::utils;

pub async fn solve() -> Result<(), Box<dyn Error>> {
    let day: i32 = 3;
    utils::load_input(day).await?;
    let input_file = utils::input_dir().join(format!("day{:02}", day));

    println!("loading input for day 03...");
    let lines = utils::read_file_lines(input_file)?;

    let mut map_lines = Vec::new();

    for line in lines {
        map_lines.push(line?);
    }

    let num_trees = part1(&map_lines)?;
    println!("number of trees for first pass: {}", num_trees);

    Ok(())
}

struct TreeMap {
    width: usize,
    height: usize,
    trees: HashSet<(usize, usize)>,
}

impl TreeMap {
    fn new(input: &Vec<String>) -> Result<TreeMap, &'static str> {
        let first_line = input.iter().next().unwrap();

        let height = input.len();
        let width = first_line.len();
        let mut trees = HashSet::new();

        for (i, line) in input.iter().enumerate() {
            for (j, c) in line.char_indices() {
                if c == '#' {
                    trees.insert((i, j));
                }
            }
        }

        Ok(TreeMap {
            width,
            height,
            trees,
        })
    }
}

fn part1(input: &Vec<String>) -> Result<usize, Box<dyn Error>> {
    let map = TreeMap::new(input)?;
    let mut num_trees: usize = 0;

    for i in 0..map.height {
        let j = (3 * i) % map.width;

        if map.trees.contains(&(i, j)) {
            num_trees += 1;
        }
    }

    return Ok(num_trees);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn new_tree_map() {
        let input = vec![".##.", "....", "#.#."]
            .iter()
            .map(|&s| String::from(s))
            .collect();

        let map = TreeMap::new(&input).unwrap();
        println!("{:?}", map.trees);

        assert_eq!(map.height, 3);
        assert_eq!(map.width, 4);
        assert!(map.trees.contains(&(0, 1)));
        assert!(map.trees.contains(&(0, 2)));
        assert!(map.trees.contains(&(2, 0)));
        assert!(map.trees.contains(&(2, 2)));
    }

    #[test]
    fn part1_tree_path() {
        let input = vec![
            "..##.......",
            "#...#...#..",
            ".#....#..#.",
            "..#.#...#.#",
            ".#...##..#.",
            "..#.##.....",
            ".#.#.#....#",
            ".#........#",
            "#.##...#...",
            "#...##....#",
            ".#..#...#.#",
        ]
        .iter()
        .map(|&s| String::from(s))
        .collect();

        let num_trees = part1(&input).unwrap();

        assert_eq!(num_trees, 7);
    }
}
