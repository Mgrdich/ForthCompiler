.section .data
word:
    .asciz "Hello"

.section .text
.globl _start
_start:
    mov $42, %rsi
    call print
    mov $word, %rsi
    call printwln
    mov $42, %rsi
    call println
exit:
    mov $60, %rax
    xor %rdi, %rdi
    syscall
