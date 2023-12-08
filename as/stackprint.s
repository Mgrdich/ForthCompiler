        movq $8 , %rcx

        xor %rdx, %rdx

        movq %rbp , %rax
        subq %rsp , %rax

        idiv %rcx

        movq %rax , %r12


        movq %rbp , %r14
        subq $8 , %r14

        # print stack count
        moveq $p1 , %rsi
        call printw

        moveq %r12 , %rsi
        call print

        moveq $p2 , %rsi
        call printw

loop:
        mov (%r14), %rsi
        call print
        call printSpace
        sub $8, %r14
        sub $1 ,%r12
        cmp $0 , %r12
        jne loop

        call printeol