        movq $8 , %rcx

        xor %rdx, %rdx

        movq %rbp , %rax
        subq %rsp , %rax

        idiv %rcx

        movq %rax , %r12


        movq %rbp , %r14
        subq $8 , %r14

        # print stack count
        movq $p1 , %rsi
        call printw

        movq %r12 , %rsi
        call print

        movq $p2 , %rsi
        call printw

        call printSpace
loop:
        mov (%r14), %rsi
        call print
        call printSpace
        sub $8, %r14
        sub $1 ,%r12
        cmp $0 , %r12
        jne loop

        mov $okWord, %rsi
        call printw

        call printeol