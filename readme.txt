       ╭╮
╭─┬┬┬──┤└╮
╰┴┴─┴┴┴┴─╯

numb is a tool for working with numbers, units and dates.

examples:

    $ numb
    Enter `q` to quit
    > 3735928559 >> 16 in hex
      0xdead
    > 1920/1080 in rat
      16/9
    > sin(pi/3)
      0.87
    > 3 millilightsecond to mile
      558.85 mile
    > 10 mile/hour to metre/sec
      4.47 metre/sec
    > total_grains = 2^64 - 1
      18446744073709551615
    > 3 GB / (15 Mbps) to minute
      26.67 minute
    > time in London
      13:04
    > today + 10 days
      05 Jun 2022
    > today - {1999-09-09} to months
      272.54 month

install:

    # build with Go >= 1.16
    $ go install github.com/nkanaev/numb@latest

more:

    doc/help.txt     manual, examples
    doc/spec.txt     syntax, grammar
    doc/dump.txt     various design notes
    doc/todo.txt     todo list
