#!/bin/sh

as -o tmp_main.o tmp_main.s
as -o print.o print.s
ld -o mf tmp_main.o print.o -lc -dynamic-linker /lib64/ld-linux-x86-64.so.2
cp mf ../mf