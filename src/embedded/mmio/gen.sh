#!/bin/sh

sed 's/32/16/g' mmio32.go >mmio16.go
sed 's/32/8/g' mmio32.go >mmio8.go
sed 's/32/64/g' mmio32.go >mmio64.go