.section .rodata
okWord:
.asciz "ok" 
p1:
.asciz "<"
p2:
.asciz ">"
.section .text
.global _start
_start:
pushq %rbp
movq %rsp , %rbp
pushq $5
pushq $2
pushq $3
pushq $56
pushq $76
pushq $23
pushq $65
pushq $5
pushq $4
popq %rax
popq %rbx
addq %rax, %rbx
pushq %rbx
popq %rsi
call print
call printSpace
movq $okWord, %rsi
call printwln
pushq $6
pushq $7
popq %rax
popq %rbx
imulq %rax, %rbx
pushq %rbx
popq %rsi
call print
call printSpace
movq $okWord, %rsi
call printwln
pushq $1360
pushq $23
popq %rax
popq %rbx
subq %rax, %rbx
pushq %rbx
popq %rsi
call print
call printSpace
movq $okWord, %rsi
call printwln
movq $8 , %rcx
xor %rdx, %rdx
movq %rbp , %rax
subq %rsp , %rax
idiv %rcx
movq %rax , %r12
movq %rbp , %r14
subq $8 , %r14
movq $p1 , %rsi
call printw
movq %r12 , %rsi
call print
movq $p2 , %rsi
call printw
call printSpace
loop:
movq (%r14), %rsi
call print
call printSpace
subq $8, %r14
subq $1 ,%r12
cmp $0 , %r12
jne loop
movq $okWord, %rsi
call printwln
movq %rbp , %rsp
popq %rbp
exit:
mov $60, %rax
xor %rdi, %rdi
syscall
