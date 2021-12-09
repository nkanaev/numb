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
    > total_grains = 2^64 - 1
      18,446,744,073,709,551,615
    > 10 GB / (15 MB/s) to minute
      11.11 min
    > 10 Kbit/sec * 30 days to megabyte
      3,240 MB

install:

    $ go install github.com/nkanaev/numb@latest
