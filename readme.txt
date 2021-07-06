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
    > 10 mile/hour to metre/sec
      4.47 m/s
    > total_grains = 2 pow 64 - 1
      18,446,744,073,709,551,615

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
    sci   scientific (-1.23e+45)

install:

    $ go get -u github.com/nkanaev/numb/cmd/numb
