The Thumb2 assembler is intended to generate code for ARMv7-M ISA. It can be used for other variants of ARMv7 ISA as long as the access to special registers or coprocessors is not needed (however, you can always use WORD, HWORD to generate not supported instructions).

For the rest of this document thumb means the Thumb2 assembler, arm means the ARM assembler.

### Differences in relation to arm

1. The lack of .S suffix does not mean that the generated data processing instruction does not modify flags. In such case thumb generates shortest instruction encoding that may (but may not) modify the flags. Use .P suffix to preserve flags.

2. Thumb uses R7 as a temporary register (REGTMP) to promote shortest encodings (arm uses R11).

3. Thumb requires support for hardware division. ARMv7-M ISA supports it, but in case of ARMv7-A the hardware division is optional.

4. Thumb requires support for unalligned memory access by MOVW and MOVH instructions (ARMv7 supports it if ARMv7-A SCTLR.A=0, ARMv7-M if CCR.UNALIGN_TRP=0).

5. Thumb does not support .IB (increment before), .DA (decrement after) suffixes.

### Thumb only features

1. As long as a function has no any explicit IT instruction thumb automatically adds IT instructions before group of 1 to 4 instructions with condition suffixes to make them conditional. Any explicit IT instruction disables this auto-IT mode for the entrie function.

2. In auto-IT mode thumb converts `B.cond label` to `Bcond label` (if not the last instruction in the IT block). In case of B.cond the condition comms from IT instruction, whereas Bcond repersents a conditional branch that does not require IT.

### Portable code

To write portable code for arm and thumb:

1. In thumb code always ues .S suffix if you want to set flags like in arm code.

2. Check condition flags immediately after setting them, do not use .P to preserve flags (arm does not support .P for arithmetic instructions).

3. Be careful when using R7, R11 (temporary) registers.

4. Do not use IT instruction (rely on auto-IT mode instead).

5. Do not use thumb-only instructions like CBZ, CBNZ.

