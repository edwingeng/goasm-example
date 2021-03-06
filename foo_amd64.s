//+build !noasm !appengine
// AUTO-GENERATED BY C2GOASM -- DO NOT EDIT

TEXT ·_AddTwoInts(SB), $0-24

    MOVQ a+0(FP), DI
    MOVQ b+8(FP), SI
    MOVQ c+16(FP), DX

    WORD $0xf701                 // add    edi, esi
    WORD $0x3a89                 // mov    dword [rdx], edi
    RET




TEXT ·_SqrtAvx256(SB), $0-16

    MOVQ in+0(FP), DI
    MOVQ out+8(FP), SI

    LONG $0x0751fcc5             // vsqrtps    ymm0, yword [rdi]
    LONG $0x0611fcc5             // vmovups    yword [rsi], ymm0
    VZEROUPPER
    RET
