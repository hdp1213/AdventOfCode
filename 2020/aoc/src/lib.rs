use std::error::Error;

#[macro_use]
extern crate lazy_static;

mod utils;

mod day01;
mod day02;
mod day03;
mod day04;

pub struct Args {
    day: i32,
}

impl Args {
    pub fn new(args: &[String]) -> Result<Args, &'static str> {
        if args.len() < 2 {
            return Err("not enough arguments");
        }

        let day = args[1].clone().parse::<i32>().expect("invalid day");
        Ok(Args { day })
    }
}

pub async fn run(config: Args) -> Result<(), Box<dyn Error>> {
    if config.day == 1 {
        day01::solve().await?;
    } else if config.day == 2 {
        day02::solve().await?;
    } else if config.day == 3 {
        day03::solve().await?;
    } else if config.day == 4 {
        day04::solve().await?;
    }

    Ok(())
}
