.section .text
.globl _start
_start:
    mov $42, %rsi
    call print
exit:
    mov $60, %rax
    xor %rdi, %rdi
    syscall