.section .rodata
digitln:
   .asciz "%d\n"
digit:
   .asciz "%d"
wordln:
    .asciz "%s\n"
word:
    .asciz "%s"
eol:
   .asciz "\n"
space:
    .asciz " "

.section .text

.globl print
print:
 push %rbp
 mov %rsp, %rbp
 mov $digit, %rdi
 xor %rax, %rax
 xor %rdx, %rdx
 call printf
 mov %rbp, %rsp
 pop %rbp
 ret

.globl printSpace
printSpace:
 push %rbp
 mov %rsp, %rbp
 mov $word, %rdi
 mov $space, %rsi
 xor %rax, %rax
 xor %rdx, %rdx
 call printf
 mov %rbp, %rsp
 pop %rbp
 ret

.globl printeol
printeol:
 push %rbp
 mov %rsp, %rbp
 mov $eol, %rdi
 xor %rax, %rax
 xor %rdx, %rdx
 call printf
 mov %rbp, %rsp
 pop %rbp
 ret


.globl println
println:
 push %rbp
 mov %rsp, %rbp
 mov $digitln, %rdi
 xor %rax, %rax
 xor %rdx, %rdx
 call printf
 mov %rbp, %rsp
 pop %rbp
 ret

.globl printwln
printwln:
 push %rbp
 mov %rsp, %rbp
 mov $wordln, %rdi
 xor %rax, %rax
 xor %rdx, %rdx
 call printf
 mov %rbp, %rsp
 pop %rbp
 ret

.globl printw
printw:
 push %rbp
 mov %rsp, %rbp
 mov $word, %rdi
 xor %rax, %rax
 xor %rdx, %rdx
 call printf
 mov %rbp, %rsp
 pop %rbp
 ret