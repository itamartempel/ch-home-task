# CH Home task Solution

## Overview
* This is the solution for the home task provided by CH.
* The solution is implemented entirely in GoLang.
* The solution aims to be efficient and minimize memory consumption as much as possible.

## Explanation
* My solution is compose from 2 entirely decopled component:
  * **[ValuedUrlIterator](pkg/valuedurliterator/valuedurliterator.go#L12)** - This component takes a string buffer reader and iterates through the buffer line by line, deserializing each line into a valued URL (as described in the task description) efficiently, without loading everything into memory at once.
  * **[TopNValuedUrls](pkg/topnvaluedurls/topnvaluedurls.go#L8)** - This component maintains a top **N** list (in our case, top 10) of valued URLs. It uses a **Heap MAX Tree** (similar to a PriorityQueue) to manage a slice of valued URLs with a given size. This is the most efficient way to maintain an ordered list given the nature of my code, which deals with a data stream of valued URLs. For each valued URL that enters the list:
    *  The **Heap MAX Tree** puts the URL in the right place in the list with a complexity of O(log n).
    * If the length of the list exceeds **N** (the initial size of the top list), the smallest valued URL is popped from the list while keeping the order intact, with a complexity of O(1) (since the smallest valued URL is always the last item in the slice).

* The [main function](main.go#L13) scan the `stdin` for a file path containing valued URLs, and then utilizes the combination of the two components mentioned above to initiate the process of printing the top 10 valued URLs.
* All the critical-path functionality is tested, by executing:
  ```sh
  $> make test
  ```
## Usage
> Make sure that you have go language install on your meching (e.g. `brew install go` on mac)

### Installion
```sh
$> go install github.com/itamartempel/ch-home-task@latest
```
### Or Build manually


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