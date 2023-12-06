.section .rodata
digitln:
   .asciz "%d\n"
digit:
   .asciz "%d"
word:
    .asciz "%s\n"
eol:
   .asciz "\n"

.section .text

.globl println
.globl prints
.globl print

print:
 push %rax
 mov $digit, %rdi
 xor %rax, %rax
 call printf
 pop %rax
 ret

println:
 push %rax
 mov $digitln, %rdi
 xor %rax, %rax
 call printf
 pop %rax
 ret

prints:
 push %rax
 mov $word, %rdi
 xor %rax, %rax
 call printf
 pop %rax
 ret
