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

operations:

    +     addition
    -     subtraction
    *     multiplication
    /     division
    mod   modulo
    pow   exponent

    <<    shift left
    >>    shift right
    and   bitwise and
    or    bitwise or
    xor   bitwise xor

formats:

    dec   decimal
    hex   hexadecimal
    oct   octal
    bin   binary
    rat   rational (x/y)
    exp   scientific (-1.23e+45)

install:

    $ go get -u github.com/nkanaev/numb/cmd/numb
