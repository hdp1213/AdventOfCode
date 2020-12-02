use std::error::Error;

mod utils;

mod day01;

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
    }

    Ok(())
}
