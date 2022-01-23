
## Description

Based on a pep talk generator that [@VVh1sp3r](https://twitter.com/VVh1sp3r) tweeted. [(see image)](https://pbs.twimg.com/media/FJgwQhBXMAUR34M?format=jpg&name=medium).  

I thought it would be a fun little project to use as my first Go program.  

It reandomly selects a phrase each from 4 columns, or 2D arrarys, and combines them to create a simple pep talk sentence. The random function checks the size of each array (column) passed to it, so the arrays don't have to be equal size if you want to add or remove phrases.  

## How To Run Locally

1. copy the code  
1. `go run generator.go`  
