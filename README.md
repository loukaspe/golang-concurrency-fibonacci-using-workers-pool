# Golang Fibonacci Solution using Concurrency Worker Pool Design Pattern

## General

This is a solution to the Fibonacci sequence problem using golang concurrency and the
worker pool design pattern.

Source for the pool implementation: https://github.com/godoylucase/workers-pool

## Results 
For requested calculation of the first 40 fibonacci numbers, a simple recursive loop
took almost 12s while the concurrent one with 4 workers took 5s.