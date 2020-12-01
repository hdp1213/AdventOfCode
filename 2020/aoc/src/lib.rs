use std::error::Error;

mod day01;

pub struct Config {
  day: i32,
}

impl Config {
  pub fn new(args: &[String]) -> Result<Config, &'static str> {
    if args.len() < 2 {
      return Err("not enough arguments")
    }
    let day = args[1].clone().parse::<i32>().expect("invalid day");
    Ok(Config { day })
  }
}

pub fn run(config: Config) -> Result<(), Box<dyn Error>> {
  if config.day == 1 {
    day01::solve_day01_part1()
  }

  Ok(())
}
