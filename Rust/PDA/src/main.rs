/*
 Author: Gananath R
 Year: 2021
 PushDown Automata in rust XX^R
 Inspired from: https://homepage.cs.uri.edu/faculty/hamel/courses/2016/spring2016/csc445/lecture-notes/ln445-13.pdf
*/
mod lib;
use lib::finatecontrolunit::FiniteControlUnit;

// Stack Machine XX^R PUSHDOWN AUTOMATA

fn main() {
    println!("\n  PDA Inputs\n");
    /*
     * code: Code should be in string format
              Example: "aabbaaa"
     */
    let code = "aabb";
    /*
     rules: Rules should be in list format.
       The episilon should be added as E.
       Right most string character will enter the stack first, for "SX" X will enter stack first then S.
       [[read,pop,push]]
       Example: rules = [
                         ["a","S", "SX"],
                         ["b","X", "E"],
                         ["E","S", "E"],
                        ]
     */
    let rules = [["a", "S", "SX"], ["b", "X", "E"], ["E", "S", "E"]];
    println!("Code: {}\n",code);
    println!("Rules: {:?}\n",rules);
    let mut vec = Vec::new();

    for (i, a1) in rules.iter().enumerate() {
        vec.push(Vec::new());
        for (_, a2) in a1.iter().enumerate() {
            vec[i].push(a2);
        }
    }
    println!("  PDA Output\n");
    let mut pda = FiniteControlUnit::new(code, vec);
    pda.run();
}
