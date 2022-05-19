# Algorithms and Datastructures

![test](https://github.com/idexter/algorithms-and-datastructures/workflows/test/badge.svg)
[![codecov](https://codecov.io/gh/idexter/algorithms-and-datastructures/branch/master/graph/badge.svg?token=LHA4hN5boC)](codecov)
[![Go Report Card](https://goreportcard.com/badge/github.com/idexter/algorithms-and-datastructures?master)](https://goreportcard.com/report/github.com/idexter/algorithms-and-datastructures)

This repository stores my solutions for problems from sites like [LeetCode](https://leetcode.com), [HackerRank](https://hackerrank.com), etc.

Some of them can be not efficient by default. I didn't set the goal to make the most effective solution. So be careful if you decide to use them.

There are two reasons I hae made this repository:
- The first one is to improve my own algorithmic and problem-solving skills. Mostly because I self-taught software engineer.
- The second one, is try to avoid algorithmic interviews. 🥲

I think knowledge of algorithms and data structures is pretty useful, but at the same time a lot of companies overuse them on technical interviews.
Especially when it happens due interview for senior and lead positions. Anyway. Enjoy the code, improve yourself and don't stop learn new things.

NOTE: I don't have all tests from platforms described below, so I have made my own tests for some solutions.

## Index

### LeetCode

- Introduction to Data Structure
    - [Arrays 101](https://leetcode.com/explore/learn/card/fun-with-arrays/): [Solutions](./leetcode/explore/learn/card/fun-with-arrays)
- Easy Collection
    - [Top Interview Questions](https://leetcode.com/explore/featured/card/top-interview-questions-easy/): [Solutions](./leetcode/explore/featured/card/top-interview-questions-easy)

### HackerRank

- Prepare
  - Problem Solving
    - [Algorithms](https://www.hackerrank.com/domains/algorithms)
      - [Solve Me First](https://www.hackerrank.com/challenges/solve-me-first/problem) : [Solution](./hackerrank/prepare/alogrithms/warm-up/solve_me_first.go) 
      - [Simple Array Sum](https://www.hackerrank.com/challenges/simple-array-sum/problem) : [Solution](./hackerrank/prepare/alogrithms/warm-up/simple_array_sum.go) 
      - [Two Strings](https://www.hackerrank.com/challenges/two-strings/problem) : [Solution](./hackerrank/prepare/alogrithms/strings/two_strings.go) 
      - [New Year Chaos](https://www.hackerrank.com/challenges/new-year-chaos/problem) : [Solution](./hackerrank/prepare/alogrithms/constructive-algorithms/new_year_chaos.go) 
    - [Data Structures](https://www.hackerrank.com/domains/data-structures)
      - [Arrays - DS](https://www.hackerrank.com/challenges/arrays-ds/problem) : [Solution](./hackerrank/prepare/data-structures/arrays/arrays_ds.go) 
      - [2D Array - DS](https://www.hackerrank.com/challenges/2d-array/problem) : [Solution](./hackerrank/prepare/data-structures/arrays/arrays_ds_2d.go)
      - [Print the Elements of a Linked List](https://www.hackerrank.com/challenges/print-the-elements-of-a-linked-list/problem) : [Solution](./hackerrank/prepare/data-structures/linked-lists/print_elements.go) 
      - [Insert a node at a specific position in a linked list](https://www.hackerrank.com/challenges/insert-a-node-at-a-specific-position-in-a-linked-list/problem) : [Solution](./hackerrank/prepare/data-structures/linked-lists/insert_node_at_position.go)
      - [Reverse a doubly linked list](https://www.hackerrank.com/challenges/reverse-a-doubly-linked-list/problem) : [Solution](./hackerrank/prepare/data-structures/linked-lists/reverse_doubly_linked_list.go)
  - [Interview Preparation Kit](https://www.hackerrank.com/interview/interview-preparation-kit)
    - [Dictionaries and Hash Maps](https://www.hackerrank.com/interview/interview-preparation-kit/dictionaries-hashmaps/challenges)
      - [Hash Tables: Ransom Note](https://www.hackerrank.com/challenges/ctci-ransom-note/problem) : [Solution](./hackerrank/prepare/interview-preparation-kit/dictionaries-and-hashmaps/hash_table_ransom_note.go)
      - [Two Strings](https://www.hackerrank.com/challenges/two-strings/problem) : [Solution](./hackerrank/prepare/alogrithms/strings/two_strings.go)
    - [Arrays](https://www.hackerrank.com/interview/interview-preparation-kit/arrays/challenges)
      - [Arrays: Left Rotation](https://www.hackerrank.com/challenges/ctci-array-left-rotation/problem) : [Solution](./hackerrank/prepare/data-structures/arrays/arrays_left_rotation.go)

## Usage

Just use default make target.
It will run `golint` and tests:

```bash
make
```
