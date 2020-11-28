	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 10, 15	sdk_version 10, 15
	.intel_syntax noprefix
	.globl	__Z10AddTwoIntsiiPi     ## -- Begin function _Z10AddTwoIntsiiPi
	.p2align	4, 0x90
__Z10AddTwoIntsiiPi:                    ## @_Z10AddTwoIntsiiPi
## %bb.0:
	push	rbp
	mov	rbp, rsp
	add	edi, esi
	mov	dword ptr [rdx], edi
	pop	rbp
	ret
                                        ## -- End function
	.globl	__Z10SqrtAvx256PfS_     ## -- Begin function _Z10SqrtAvx256PfS_
	.p2align	4, 0x90
__Z10SqrtAvx256PfS_:                    ## @_Z10SqrtAvx256PfS_
## %bb.0:
	push	rbp
	mov	rbp, rsp
	vsqrtps	ymm0, ymmword ptr [rdi]
	vmovups	ymmword ptr [rsi], ymm0
	pop	rbp
	vzeroupper
	ret
                                        ## -- End function
.subsections_via_symbols
