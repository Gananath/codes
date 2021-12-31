use std::fmt::{Debug, Display, Formatter, Result};
use std::ops::Index;

// A custom made generic stack data struture in rust

pub struct Stack<T> {
    data: Vec<T>,
    index: usize,
}

impl<T> Stack<T> {
    pub fn new() -> Stack<T> {
        Stack {
            data: Vec::new(),
            index: 0,
        }
    }
    pub fn push(&mut self, i: T) {
        self.data.push(i);
    }
    pub fn pop(&mut self) {
        self.data.pop();
    }
}

// implementing display for stack
impl<T: Debug> Display for Stack<T> {
    fn fmt(&self, f: &mut Formatter) -> Result {
        write!(f, "<{:?}>", self.data)
    }
}

// implementing debug for stack
impl<T: Debug> Debug for Stack<T> {
    fn fmt(&self, f: &mut Formatter) -> Result {
        write!(f, "<{:?}>", self.data)
    }
}

impl<T> Index<usize> for Stack<T> {
    type Output = T;
    fn index(&self, idx: usize) -> &Self::Output {
        let l = self.data.len() - 1;
        &self.data[l - idx]
    }
}

// implementing iterator for stack
impl<T> Iterator for Stack<T>
where
    T: Copy,
{
    type Item = T;

    fn next(&mut self) -> Option<Self::Item> {
        let reverse_idx: Option<usize> = if !self.data.is_empty() {
            Some(self.data.len() - 1)
        } else {
            None
        };

        if let Some(r_idx) = reverse_idx {
            if self.index <= r_idx {
                let i = self.index;
                self.index += 1;
                return Some(self.data[r_idx - i]);
            }
        }

        None
    }
    // finds element by index but consumes it
    fn nth(&mut self, _n: usize) -> Option<Self::Item> {
        self.next()
    }
}

// implementing length for the stack iterator
impl<T: Copy> ExactSizeIterator for Stack<T> {
    fn len(&self) -> usize {
        self.data.len()
    }
}

// implementing double ended iter for reversing
impl<T: Copy> DoubleEndedIterator for Stack<T> {
    fn next_back(&mut self) -> Option<Self::Item> {
        None
    }
}

