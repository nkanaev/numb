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
    Enter `q` to quit
    > 0b1001100110111 as hex
      0x1337
    > 3 millilightsecond to mile
      558.85 mi
    > 12 parsec to lightyear
      39.14 ly

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
