# Collatz

Print Collatz series for given integers.

## Installation

### Unix users (Linux, BSDs and MacOSX)

Unix users may download and install latest *collatz* release with command:

```bash
sh -c "$(curl http://sweetohm.net/dist/collatz/install)"
```

If *curl* is not installed on you system, you might run:

```bash
sh -c "$(wget -O - http://sweetohm.net/dist/collatz/install)"
```

**Note:** Some directories are protected, even as *root*, on **MacOSX** (since *El Capitan* release), thus you can't install *collatz* in */usr/bin* for instance.

### Binary package

Otherwise, you can download latest binary archive at <https://github.com/c4s4/collatz/releases>. Unzip the archive, put the binary of your platform somewhere in your *PATH* and rename it *collatz*.

## Usage

To print Collatz series, type:

```bash
$ collatz 7
7 22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1 (16) [52]
```

This prints Collatz series, number of steps (here *16*) and maximum value (here *52*).

*Enjoy:*
