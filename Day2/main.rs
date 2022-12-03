use std::io;
use std::io::prelude::*;
use std::collections::HashMap;
use std::borrow::Borrow;
fn main() -> io::Result<()> {
    let mut map = HashMap::new();
    map.insert("Rock", 1);
    map.insert("Paper", 2);
    map.insert("Scissor", 3);
    map.insert("Win", 6);
    map.insert("Draw", 3);
    map.insert("Lose", 0);
    let mut total_score_1 = 0;
    let mut total_score_2 = 0;
    for line in io::stdin().lock().lines() {
        let line = line?;
        let letters = line.split(" ");
        let vec: Vec<&str> = letters.collect();
        let mut curr_score_1 = 0;
        let mut curr_score_2 = 0;
        // First Part
        match vec[0] {
                // Rock - ?
            "A" => {
                match vec[1] {
                    "X" => {
                        curr_score_1 = get_score(&map, "Rock", "Draw");
                        curr_score_2 = get_score(&map, "Scissor", "Lose");
                    },
                    "Y" => {
                        curr_score_1 = get_score(&map, "Paper", "Win");
                        curr_score_2 = get_score(&map, "Rock", "Draw");
                    },
                    "Z" => {
                        curr_score_1 = get_score(&map, "Scissor", "Lose");
                        curr_score_2 = get_score(&map, "Paper", "Win");
                    },
                    _ => println!("Invalid input"),
                }
            },
            // Paper - ?
            "B" => {
                match vec[1] {
                    "X" => {
                        curr_score_1 = get_score(&map, "Rock", "Lose");
                        curr_score_2 = get_score(&map, "Rock", "Lose");
                    }
                    "Y" => {
                        curr_score_1 = get_score(&map, "Paper", "Draw");
                        curr_score_2 = get_score(&map, "Paper", "Draw");
                    },
                    "Z" => {
                        curr_score_1 = get_score(&map, "Scissor", "Win");
                        curr_score_2 = get_score(&map, "Scissor", "Win");
                    },
                    _ => println!("Invalid input"),
                }
            },
            // Scissors - ?
            "C" => {
                match vec[1] {
                    "X" => {
                        curr_score_1 = get_score(&map, "Rock", "Win");
                        curr_score_2 = get_score(&map, "Paper", "Lose");
                    },
                    "Y" => {
                        curr_score_1 = get_score(&map, "Paper", "Lose");
                        curr_score_2 = get_score(&map, "Scissor", "Draw");
                    },
                    "Z" => {
                        curr_score_1 = get_score(&map, "Scissor", "Draw");
                        curr_score_2 = get_score(&map, "Rock", "Win");
                    },
                     _ => println!("Invalid input"),
                }
            },
            _ => println!("Invalid input"),
        }

        total_score_1 += curr_score_1;
        total_score_2 += curr_score_2;
        println!("Total score 1: {}", total_score_1);
        println!("Total score 2: {}", total_score_2);
    }
    println!("Total score 1: {}", total_score_1);
    println!("Total score 2: {}", total_score_2);
    Ok(())
}

fn get_score(c: &HashMap<&str, i32>, pick: &str, result: &str) -> i32 {
    let map = c.borrow().clone();
    return map.get(pick).unwrap() + map.get(result).unwrap();
}