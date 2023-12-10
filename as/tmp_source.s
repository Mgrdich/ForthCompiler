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
addq %rax , %rbx
pushq %rbx
popq %rsi
movq %rsi , %r15
call print
call printSpace
pushq %r15
movq $okWord, %rsi
call printwln
popq %rsi
call print
call printSpace
movq $okWord, %rsi
call printwln
pushq $6
pushq $7
popq %rax
popq %rbx
imulq %rax , %rbx
pushq %rbx
popq %rsi
call print
call printSpace
movq $okWord, %rsi
call printwln
popq %rsi
movq %rsi , %r15
call print
call printSpace
pushq %r15
movq $okWord, %rsi
call printwln
pushq $1360
pushq $23
popq %rax
popq %rbx
subq %rax , %rbx
pushq %rbx
popq %rsi
call print
call printSpace
movq $okWord, %rsi
call printwln
popq %rsi
movq %rsi , %r15
call print
call printSpace
pushq %r15
movq $okWord, %rsi
call printwln
movq %rbp , %rsp
popq %rbp
exit:
movq $60 , %rax
xorq %rdi , %rdi
syscall