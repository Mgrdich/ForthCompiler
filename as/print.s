.section .rodata
digitln:
   .asciz "%d\n"
digit:
   .asciz "%d"
wordln:
    .asciz "%s\n"
eol:
   .asciz "\n"

.section .text

.globl print
print:
 push %rax
 mov $digit, %rdi
 xor %rax, %rax
 call printf
 pop %rax
 ret


.globl println
println:
 push %rax
 mov $digitln, %rdi
 xor %rax, %rax
 call printf
 pop %rax
 ret

.globl printwln
printwln:
 push %rax
 mov $wordln, %rdi
 xor %rax, %rax
 call printf
 pop %rax
 ret
