use crate::lib::stack::Stack;

// Finite control unit for PushDownAutomata

#[derive(Debug)]
pub struct FiniteControlUnit<'a> {
    rules: Vec<Vec<&'a &'a str>>,
    code_stack: Stack<char>,
    machine_stack: Stack<char>,
}

impl<'a> FiniteControlUnit<'a> {
    pub fn new(code: &str, rules: Vec<Vec<&'a &'a str>>) -> Self {
        let code = code.chars().collect::<Vec<char>>();
        let mut fcu = Self {
            rules,
            code_stack: Stack::new(),
            machine_stack: Stack::new(),
        };

        fcu.machine_stack.push('S');
        Self::load_values(&mut fcu, &code);
        fcu
    }

    fn load_values(fcu: &mut FiniteControlUnit, code: &[char]) {
        for i in (0..code.len()).rev() {
            // code will be pushed in reverse order
            // that is first line of code first
            fcu.code_stack.push(code[i]);
        }
    }

    fn get_values(fcu: &mut FiniteControlUnit<'a>) -> Option<Vec<&'a &'a str>> {
        let code_stack_val = &fcu.code_stack;
        let machine_stack_val = &fcu.machine_stack;
        if code_stack_val.len() > 0 && machine_stack_val.len() > 0 {
            for i in &fcu.rules {
                if *i[0] == code_stack_val[0].to_string()
                    && *i[1] == machine_stack_val[0].to_string()
                {
                    return Some(i.to_vec());
                }
            }
        }

        None
    }

    pub fn run(&mut self) {
        while self.code_stack.len() > 0 {
            let val = Self::get_values(self);

            if let Some(v) = val {                
                self.machine_stack.pop();
                let ch: Vec<char> = v[2].chars().rev().collect(); // converts string to char list and reverse it
                for i in ch {                    
                    self.machine_stack.push(i);
                }
                self.code_stack.pop();
            } else {
                self.machine_stack.pop();                
            }
        }
        println!("{}", self.code_stack);
        println!("{}\n", self.machine_stack);
    }
}

