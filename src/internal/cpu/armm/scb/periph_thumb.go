// Instances:
//  SCB  0xE000ED00  -  -  System Control Block
// Registers:
//  0x00  32  CPUID  CPUID Base Register
//  0x04  32  ICSR   Interrupt Control and State Register
//  0x08  32  VTOR   Vector Table Offset Register
//  0x0C  32  AIRCR  Application Interrupt and Reset Control Register
//  0x10  32  SCR    System Control Register
//  0x14  32  CCR    Configuration and Control Register
//  0x18  32  SHPR1  System Handler Priority Register 1
//  0x1C  32  SHPR2  System Handler Priority Register 2
//  0x20  32  SHPR3  System Handler Priority Register 3
//  0x24  32  SHCSR  System Handler Control and State Register
//  0x28  32  CFSR   Configurable Fault Status Register
//  0x2C  32  HFSR   HardFault Status Register
//  0x34  32  MMFR   MemManage Fault Address Register
//  0x38  32  BFAR   BusFault Address Register
//  0x3C  32  AFSR   Auxiliary Fault Status Register
package scb

const (
	Revision    CPUID = 0xf << 0   //+
	PartNo      CPUID = 0xfff << 4 //+
	Constant    CPUID = 0xf << 16  //+
	Variant     CPUID = 0xf << 20  //+
	Implementer CPUID = 0xff << 24 //+
)

const (
	Revisionn    = 0
	PartNon      = 4
	Constantn    = 16
	Variantn     = 20
	Implementern = 24
)

const (
	VECTACTIVE  ICSR = 0x1ff << 0  //+ Active exception number (0: thread mode)
	RETTOBASE   ICSR = 1 << 11     //+ No preempted active exceptions
	VECTPENDING ICSR = 0x1ff << 12 //+ Highest priority pending exception number
	ISRPENDING  ICSR = 1 << 22     //+ Int. pending flag, excluding NMI, Faults
	PENDSTCLR   ICSR = 1 << 25     //+ SysTick exception clear-pending bit
	PENDSTSET   ICSR = 1 << 26     //+ SysTick exception set-pending bit
	PENDSVCLR   ICSR = 1 << 27     //+ PendSV clear-pending bit
	PENDSVSET   ICSR = 1 << 28     //+ PendSV set-pending bit
	NMIPENDSET  ICSR = 1 << 31     //+ NMI set-pending bit
)

const (
	VECTACTIVEn  = 0
	VECTPENDINGn = 12
)

const (
	TBLOFF VTOR = 0x1ffffff << 7 //+
)

const (
	VECTRESET     AIRCR = 1 << 0       //+
	VECTCLRACTIVE AIRCR = 1 << 1       //+
	SYSRESETREQ   AIRCR = 1 << 2       //+
	PRIGROUP      AIRCR = 7 << 8       //+
	ENDIANNESS    AIRCR = 1 << 15      //+
	VECTKEY       AIRCR = 0xffff << 16 //+
)

const (
	VECTRESETn     = 0
	VECTCLRACTIVEn = 1
	SYSRESETREQn   = 2
	PRIGROUPn      = 8
	ENDIANNESSn    = 15
	VECTKEYn       = 16
)

const (
	SLEEPONEXIT SCR = 1 << 1 //+
	SLEEPDEEP   SCR = 1 << 2 //+
	SEVONPEND   SCR = 1 << 4 //+
)

const (
	NONBASETHRDENA CCR = 1 << 0  //+
	USERSETMPEND   CCR = 1 << 1  //+
	UNALIGN_TRP    CCR = 1 << 3  //+
	DIV_0_TRP      CCR = 1 << 4  //+
	BFHFNMIGN      CCR = 1 << 8  //+
	STKALIGN       CCR = 1 << 9  //+ Stack 8 byte aligned on exception entry
	DC             CCR = 1 << 16 //+ Enable data cache
	IC             CCR = 1 << 17 //+ Enable instruction cache
	BP             CCR = 1 << 18 //+ Branch prediction is enabled
)

const (
	PRI_MemManage  SHPR1 = 0xff << 0  //+
	PRI_BusFault   SHPR1 = 0xff << 8  //+
	PRI_UsageFault SHPR1 = 0xff << 16 //+
)

const (
	PRI_MemManagen  = 0
	PRI_BusFaultn   = 8
	PRI_UsageFaultn = 16
)

const (
	PRI_SVCall SHPR2 = 0xff << 24 //+
)

const (
	PRI_SVCalln = 24
)

const (
	PRI_PendSV  SHPR3 = 0xff << 16 //+
	PRI_SysTick SHPR3 = 0xff << 24 //+
)

const (
	PRI_PendSVn  = 16
	PRI_SysTickn = 24
)

const (
	MEMFAULTACT    SHCSR = 1 << 0  //+
	BUSFAULTACT    SHCSR = 1 << 1  //+
	USGFAULTACT    SHCSR = 1 << 3  //+
	SVCALLACT      SHCSR = 1 << 7  //+
	MONITORACT     SHCSR = 1 << 8  //+
	PENDSVACT      SHCSR = 1 << 10 //+
	SYSTICKACT     SHCSR = 1 << 11 //+
	USGFAULTPENDED SHCSR = 1 << 12 //+
	MEMFAULTPENDED SHCSR = 1 << 13 //+
	BUSFAULTPENDED SHCSR = 1 << 14 //+
	SVCALLPENDED   SHCSR = 1 << 15 //+
	MEMFAULTENA    SHCSR = 1 << 16 //+
	BUSFAULTENA    SHCSR = 1 << 17 //+
	USGFAULTENA    SHCSR = 1 << 18 //+
)

const (
	// MFSR
	IACCVIOL  CFSR = 1 << 0 //+
	DACCVIOL  CFSR = 1 << 1 //+
	MUNSTKERR CFSR = 1 << 3 //+
	MSTKERR   CFSR = 1 << 4 //+
	MLSPERR   CFSR = 1 << 5 //+
	MMARVALID CFSR = 1 << 7 //+

	// BFSR
	IBUSERR     CFSR = 1 << 8  //+
	PRECISERR   CFSR = 1 << 9  //+
	IMPRECISERR CFSR = 1 << 10 //+
	UNSTKERR    CFSR = 1 << 11 //+
	STKERR      CFSR = 1 << 12 //+
	LSPERR      CFSR = 1 << 13 //+
	BFARVALID   CFSR = 1 << 15 //+

	// UFSR
	UNDEFINSTR CFSR = 1 << 16 //+
	INVSTATE   CFSR = 1 << 17 //+
	INVPC      CFSR = 1 << 18 //+
	NOCP       CFSR = 1 << 19 //+
	UNALIGNED  CFSR = 1 << 24 //+
	DIVBYZERO  CFSR = 1 << 25 //+
)

const (
	VECTTBL  HFSR = 1 << 1  //+
	FORCED   HFSR = 1 << 30 //+
	DEBUGEVT HFSR = 1 << 31 //+
)
