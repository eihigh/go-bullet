#!/bin/bash

# The startup time of the executable file on the Linux file system becomes slow,
# so copy it to the Windows file system and execute it.
cp game.exe /mnt/c/tmp/ && /mnt/c/tmp/game.exe
