       ╭╮
╭─┬┬┬──┤└╮
╰┴┴─┴┴┴┴─╯

numb
    ... is a tool for working with numbers and units.
    ... supports arbitrary-precision arithmetic.
    ... provides mathematical functions.
    ... is an ugly child of `bc` & `units` unix commands.
    ... is MIT licensed.

examples:

    $ numb
    enter `q` to quit
    > 2 + 2
      4
    > 15 inches to centimetre
      38.1 cm
    > (0b1000 - 1) xor 2 as hex
      0x5

install:

    $ go get -u github.com/nkanaev/numb/cmd/numb
