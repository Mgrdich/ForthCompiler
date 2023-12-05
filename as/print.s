.section .rodata
fmt:
   .asciz "%d\n"
eol:
   .asciz "\n"

.section .text
.globl print
print:
 push %rax
 #mov  $42, %rsi
 mov $fmt, %rdi
 xor %rax, %rax
 call printf
 pop %rax
 ret