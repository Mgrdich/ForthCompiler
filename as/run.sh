as -o main.o main.s
as -o print.o print.s
ld -o main main.o print.o -lc -dynamic-linker /lib64/ld-linux-x86-64.so.2