# brainfuck

Brainfuck interpreter in Go. Infinite tape of 8-bit cells, negative pointers supported. EOF on input is `0x00`. `0x0D` (CR) is converted to `0x0A` (LF) on input.

## install

```bash
$ go install github.com/sportshead/brainfuck@latest
```

## usage

```bash
$ brainfuck examples/hello.bf
$ brainfuck examples/bitwidth.bf
```
