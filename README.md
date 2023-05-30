# par

Run commands from stdin in parallel.

Each command is a single line terminated by the newline character (`\n`).

By default, the commands are run in a shell. The default number of
parallel processes equals the number of CPUs (or cores).

The flag `-n` enables dry run mode: commands are printed to `stderr`
but not run.

The flag `-v` enables verbose mode: commands are printed to `stderr`
when run.

## Installation

```shell
$ go install github.com/jabolopes/par
```

## Example

Copy and convert a bunch of CR2 photographs to JPEG in parallel:

```shell
$ for i in *.CR2; do echo convert /run/media/$USER/EOS_DIGITAL/DCIM/100EOS5D/$i ~/fotos/$(basename $i | sed -n "s/CR2$/jpg/p"); done | par -v
```

## Alternatives

> Why not `xargs`?

I can never remember the `xargs` flags by heart and the defaults seem
to be opposite of what I'm trying to achieve in terms of
parallelization. Also, subshelling doesn't work as expected since the
subshelling occurs before `xargs` executes not after, therefore,
subshelling cannot use the items read by `xargs`. There's probably a
solution to that problem but I didn't figure it out.

> Why not GNU `parallel`?

In addition to the same subshelling problem as `xargs`, there's also
the issues [1](https://news.ycombinator.com/item?id=15319715),
[2](https://bugs.launchpad.net/ubuntu/+source/parallel/+bug/1779764),
and [3](https://github.com/NixOS/nixpkgs/issues/110584), etc.
