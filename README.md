# ls-like utilite

This is my first Go program that recursively prints the contents of a specified directory.


## Getting Started

To run this proram you need installed Go

If you still haven’t it, use the following link: 
https://golang.org/doc/install

## Usage
Print subdirectories recursively
```
go run main.go directory 
```
Print subdirectories and files recursively
```
go run main.go directory -f
```

## Examples
```
go run main.go . -f
├───main.go (1881b)
├───main_test.go (1318b)
└───testdata
    ├───project
    │   ├───file.txt (19b)
    │   └───gopher.png (70372b)
    ├───static
    │   ├───css
    │   │   └───body.css (28b)
    │   ├───html
    │   │   └───index.html (57b)
    │   └───js
    │   └───site.js (10b)
    ├───zline
    │   └───empty.txt (empty)
    └───zzfile.txt (empty)
go run main.go .
└───testdata
    ├───project
    ├───static
    │   ├───css
    │   ├───html
    │   └───js
    └───zline
  ```
  
  ## Running the tests
  
  ### Test example
  
  ```
go test -v
=== RUN   TestTreeFull
--- PASS: TestTreeFull (0.00s)
=== RUN   TestTreeDir
--- PASS: TestTreeDir (0.00s)
PASS
ok  	_/path/to/project/ls_go	0.011s
```

### Test in docker container

As you know formatting on different systems and different terminals is not the same.
To solve this problem you can run the test in a container that will provide you with a standardized environment.

if you have never heard of docker, you can read a short article at the following link:
https://www.docker.com/resources/what-container

To run the test using docker container enter the following commands in the root of repository:
```
docker build -t test_ls_go .
docker run test_ls_go
```
