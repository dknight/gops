# gops

The most simple terminal todo utility in the World.

* Simple
* Stupid
* Concrete

![Smart-ID in Go language](https://github.com/dknight/gops/blob/main/files/gopher-gops.svg?raw=true)

## What? Why?

Around the World, there are tons of todo-like software. Why do we need to
make another? Because I desperately need a simple, convenient to-do list
in terminal. The software should be as simple as possible and usable with
only a few keystrokes.

## gops?

The name **gops** is ambiguous:

1. gops might mean 'go-operations' aka 'go-todolist'.
2. gops might be [Gopink](https://en.wikipedia.org/wiki/Gopnik), which I like
more 😀.

## Usage

```
Usage of gops:
  -a    Display also done items.
  -c uint
        Number of the task to complete.
  -f string
        File of stored todo items. (default $HOME/.config/gops/default)
  -n string
        Set name of the new todo task.
  -t    Set list to today's date.
```

### Examples

```sh
gops -n "Make a soup"
gops -n "Drink beer and eat semki"

gops -c 2

gops -a
```

### Some command-line usage examples

Use different list rather than default.

```sh
gops -n "Make training at gym" -f lifestyle
gops -n "Buy healthy food" -f training -f lifestyle

gops -f lifestyle

```

Save and read to file with today's date.

```sh
gops -t -n "Today go to dentist"
gops -t
```

## Testing

```go test```

## Contribution

Any help is appreciated. Found a bug, typo, inaccuracy, etc.? Please do
not hesitate to make a pull request or file an issue.

## License

MIT 2022
