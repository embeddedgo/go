# Thumb2 assembler, instruction encoding

## Definitions

### Register specifiers

Rm - first register operand, Rn - second register operand, Rd - destination register, Rdm - combined first operand and destination register, Rdn - combined second operand and destination register, Rt - transfered register.

### Shift type

	vv  <v>  Descr
	-------------------------
	00  <<   logical left
	01  >>   logical right
	10  ->   arithmetic right
	11  @>   rotate right (u5==0: rotate right with extend)
	
### Modified immediate constant (e32)

    xxxx xexx xxxx xxxx  xeee xxxx eeee eeee
          i               iii      abcd efgh
    
    iiiia  const
    ------------------------------------------
    0000x  00000000 00000000 00000000 abcdefgh
    0001x  00000000 abcdefgh 00000000 abcdefgh
    0010x  abcdefgh 00000000 abcdefgh 00000000
    0011x  abcdefgh abcdefgh abcdefgh abcdefgh
    01000  1bcdefgh 00000000 00000000 00000000
    01001  01bcdefg h0000000 00000000 00000000
    01010  001bcdef gh000000 00000000 00000000
    .....  ........ ........ ........ ........
    11111  00000000 00000000 00000001 bcdefgh0
	
## 16 and 32-bit instructions

Bits [15:11] of the least significant half-word determine 16/32-bit instruction. 32-bit instructions have the following format:

	1110 1xxx xxxx xxxx  xxxx xxxx xxxx xxxx
	1111 0xxx xxxx xxxx  xxxx xxxx xxxx xxxx
	1111 1xxx xxxx xxxx  xxxx xxxx xxxx xxxx
	
Almost all 16 and 32-bit instructions require condition suffix inside IT block. Most 16-bit instructions sets flags outside IT block.