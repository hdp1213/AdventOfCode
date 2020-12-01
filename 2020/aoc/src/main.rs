use std::env;
use std::process;

use aoc::Args;

#[tokio::main]
async fn main() {
    let args: Vec<String> = env::args().collect();

    let config = Args::new(&args).unwrap_or_else(|err| {
        println!("Problem parsing args: {}", err);
        process::exit(1);
    });

    if let Err(e) = aoc::run(config).await {
        println!("Application error: {}", e);
        process::exit(1);
    }
}
