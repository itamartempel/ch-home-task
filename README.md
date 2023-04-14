# CH Home task Solution

## Overview
* This is the solution for the home task provided by CH.
* The solution is implemented entirely in GoLang.
* The solution aims to be efficient and minimize memory consumption as much as possible.

## 

## Usage
### Build
> Make sure that you have go language install on your meching (e.g. `brew install go` on mac)
>

Clone and build the repo:
```sh
$>  git clone https://github.com/itamartempel/ch-home-task && cd ch-home-task
$> make build
go fmt ./...
go vet ./...
go test ./...
?       github.com/itamartempel/ch-home-task/cli        [no test files]
ok      github.com/itamartempel/ch-home-task/pkg/topnvaluedurls 0.937s
ok      github.com/itamartempel/ch-home-task/pkg/valuedurliterator      0.681s
go build -o bin/ch-home-task cli/main.go
```
Than the binary will available under `./bin/ch-home-task` in the repo dir.

### Execution
As stated in the task description, the command receives a file path from stdin. Therefore, the execution should follow this format:
```sh
$> echo "testdata/exampleData.txt" | ./bin/ch-home-task
http://api.tech.com/item/122345
http://api.tech.com/item/124345
http://api.tech.com/item/125345
http://api.tech.com/item/123345
http://api.tech.com/item/121345
```

I have also created a script that generates test data files of sizes 1MB, 10MB, and 100MB.:
```sh
$> make gen-testdata
go run ./testdata/gen-testdata.go
testdata/generated/1MB.txt Created successfully
testdata/generated/10MB.txt Created successfully
testdata/generated/100MB.txt Created successfully
```

Than you can test on some big files:
```sh
$> time echo "testdata/generated/100MB.txt" | ./bin/ch-home-task
http://api.tech.com/item/132113
http://api.tech.com/item/897548
http://api.tech.com/item/996768
http://api.tech.com/item/658828
http://api.tech.com/item/383411
http://api.tech.com/item/132212
http://api.tech.com/item/555821
http://api.tech.com/item/910476
http://api.tech.com/item/968206
http://api.tech.com/item/724148
echo "testdata/generated/100MB.txt"  0.00s user 0.00s system 47% cpu 0.002 total
./bin/ch-home-task  0.65s user 0.06s system 49% cpu 1.426 total
```