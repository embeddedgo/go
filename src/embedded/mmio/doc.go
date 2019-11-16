// Package mmio provides data types and methods that can be used to define and
// access memory mapped registers of peripherals in embedded systems.
//
// You can not use ordinary load/store operations to access memory mapped
// peripheral registers because these operations can be reordered or optimized
// out by the compiler. This is unacceptable because in case of MMIO the order
// of accessing registers matters and even reading from register can have side
// effects.
//
// This package has some support from compiler to ensure order and speed of
// I/O operations.
package mmio