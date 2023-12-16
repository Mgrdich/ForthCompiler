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
    movq $1, %rax
    xor %rdx, %rdx

check_stack_alignment_print:
    movq %rsp, %rax      # Move stack pointer to %rax
    andq $15, %rax       # Perform bitwise "and" with 15 (binary 1111)
    testq %rax, %rax
    jz aligned_print
    xor %rax, %rax
    push %rax

aligned_print:
    call printf

    cmp $0, %rax
    jne return_print
    pop %rax
return_print:
    mov %rbp, %rsp
    pop %rbp
    ret

.globl printSpace
printSpace:
    push %rbp
    mov %rsp, %rbp

    mov $word, %rdi
    mov $space, %rdi
    movq $1, %rax
    xor %rdx, %rdx

check_stack_alignment_printSpace:
    movq %rsp, %rax      # Move stack pointer to %rax
    andq $15, %rax       # Perform bitwise "and" with 15 (binary 1111)
    testq %rax, %rax
    jz aligned_printSpace
    xor %rax, %rax
    push %rax

aligned_printSpace:
    call printf

    cmp $0, %rax
    jne return_printSpace
    pop %rax
return_printSpace:
    mov %rbp, %rsp
    pop %rbp
    ret

.globl printeol
printeol:
    push %rbp
    mov %rsp, %rbp

    mov $eol, %rdi
    movq $1, %rax
    xor %rdx, %rdx

check_stack_alignment_printeol:
    movq %rsp, %rax      # Move stack pointer to %rax
    andq $15, %rax       # Perform bitwise "and" with 15 (binary 1111)
    testq %rax, %rax
    jz aligned_printeol
    xor %rax, %rax
    push %rax

aligned_printeol:
    call printf

    cmp $0, %rax
    jne return_printeol
    pop %rax
return_printeol:
    mov %rbp, %rsp
    pop %rbp
    ret


.globl println
println:
    push %rbp
    mov %rsp, %rbp

    mov $digitln, %rdi
    movq $1, %rax
    xor %rdx, %rdx

check_stack_alignment_println:
    movq %rsp, %rax      # Move stack pointer to %rax
    andq $15, %rax       # Perform bitwise "and" with 15 (binary 1111)
    testq %rax, %rax
    jz aligned_println
    xor %rax, %rax
    push %rax

aligned_println:
    call printf

    cmp $0, %rax
    jne return_println
    pop %rax
return_println:
    mov %rbp, %rsp
    pop %rbp
    ret

.globl printwln
printwln:
    push %rbp
    mov %rsp, %rbp

    mov $wordln, %rdi
    movq $1, %rax
    xor %rdx, %rdx

check_stack_alignment_printwln:
    movq %rsp, %rax      # Move stack pointer to %rax
    andq $15, %rax       # Perform bitwise "and" with 15 (binary 1111)
    testq %rax, %rax
    jz aligned_printwln
    xor %rax, %rax
    push %rax

aligned_printwln:
    call printf

    cmp $0, %rax
    jne return_printwln
    pop %rax
return_printwln:
    mov %rbp, %rsp
    pop %rbp
    ret


.globl printw
printw:
    push %rbp
    mov %rsp, %rbp

    mov $word, %rdi
    movq $1, %rax
    xor %rdx, %rdx

check_stack_alignment_printw:
    movq %rsp, %rax      # Move stack pointer to %rax
    andq $15, %rax       # Perform bitwise "and" with 15 (binary 1111)
    testq %rax, %rax
    jz aligned_printw
    xor %rax, %rax
    push %rax

aligned_printw:
    call printf

    cmp $0, %rax
    jne return_printw
    pop %rax
return_printw:
    mov %rbp, %rsp
    pop %rbp
    ret