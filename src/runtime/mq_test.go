// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "runtime/internal/atomic"

const fillkeys = "ABCDEFGHIJ"

var (
	testms [len(fillkeys)]m // global to avoid GC (mq uses uintptr pointers)
	testm  m
	run    uint32
)

func mqfill(q *mq) {
	for i := range testms {
		m := &testms[i]
		msetkey(m, uintptr(fillkeys[i]))
		q.push(m)
	}
}

func mclfill(q *mcl) {
	for i := range testms {
		m := &testms[i]
		msetkey(m, uintptr(fillkeys[i]))
		q.push(m)
	}
}

func (q *mcl) removebykey(key uintptr) *m {
	m := q.find(key)
	if m != nil {
		q.remove(m)
	}
	return m
}

func (q *msl) pople(val int64) *m {
	m := q.first()
	if m == nil || mval(m) > val {
		return nil
	}
	q.remove(m)
	return m
}

func (q *mq) str() string {
	if q.n == 0 {
		return ""
	}
	var buf [len(testms)]byte
	n := 0
	for m := q.first.ptr(); ; m = mnext1(m) {
		buf[n] = byte(mkey(m))
		n++
		if m == q.last.ptr() {
			break
		}
	}
	return string(buf[:n])
}

func (q *mcl) str() string {
	if q.n == 0 {
		return ""
	}
	var buf [len(testms)]byte
	cur := q.cur.ptr()
	m := cur
	n := 0
	for {
		buf[n] = byte(mkey(m))
		n++
		m = mnext1(m)
		if m == cur {
			break
		}
	}
	return string(buf[:n])
}

func (q *msl) str() string {
	if q.n == 0 {
		return ""
	}
	var buf [len(testms)]byte
	first := q.head.ptr()
	m := first
	n := 0
	for {
		buf[n] = byte(mval(m) >> 30)
		n++
		m = mnext2(m)
		if m == first {
			break
		}
	}
	return string(buf[:n])
}

var mvi int

func mv(val byte) *m {
	m := &testms[mvi]
	mvi++
	msetval(m, int64(val)<<30)
	return m
}

func MQTest() string {
	if !atomic.Cas(&run, 0, 1) {
		return "" // run only once
	}

	q := new(mq)

	mqfill(q)
	if q.str() != fillkeys {
		return "mq: fill"
	}
	for i := 0; i < len(fillkeys); i++ {
		m := q.pop()
		if mkey(m) != uintptr(fillkeys[i]) || q.str() != fillkeys[i+1:] {
			return "mq: pop"
		}
	}
	if m := q.pop(); m != nil && q.str() != "" {
		return "mq: pop from empty"
	}
	mqfill(q)
	if q.str() != fillkeys {
		return "mq fill"
	}
	order := fillkeys
	for i := 0; i < len(fillkeys); i++ {
		q.push(q.pop())
		order = order[1:] + order[:1]
		if q.str() != order {
			return "mq: push(pop)"
		}
	}
	if order != fillkeys {
		return "mq: order != fillkeys"
	}

	cl := new(mcl)

	mclfill(cl)
	if cl.str() != fillkeys {
		return "mcl fill"
	}

	var m *m
	if m = cl.removebykey('x'); m != nil || cl.str() != "ABCDEFGHIJ" {
		return "mcl: removebykey, key unknown"
	}
	if m = cl.removebykey('A'); mkey(m) != 'A' || cl.str() != "BCDEFGHIJ" {
		return "mcl: removebykey A"
	}
	if m = cl.removebykey('E'); mkey(m) != 'E' || cl.str() != "FGHIJBCD" {
		return "mcl: removebykey E"
	}
	if m = cl.removebykey('D'); mkey(m) != 'D' || cl.str() != "FGHIJBC" {
		return "mcl: removebykey D"
	}
	if cl.push(m); cl.str() != "FGHIJBCD" {
		return "mcl: push D"
	}
	msetkey(&testm, 'G')
	if cl.push(&testm); cl.str() != "FGHIJBCDG" {
		return "mcl: push G"
	}
	if m = cl.removebykey('G'); mkey(m) != 'G' || cl.str() != "HIJBCDGF" {
		return "mcl: removebykey G"
	}
	if m = cl.removebykey('G'); mkey(m) != 'G' || cl.str() != "FHIJBCD" {
		return "mcl: removebykey G"
	}
	if cl.push(&testm); cl.str() != "FHIJBCDG" {
		return "mcl: push G"
	}
	if m = cl.removebykey('D'); mkey(m) != 'D' || cl.str() != "GFHIJBC" {
		return "mcl: removebykey D"
	}
	if cl.remove(&testm); cl.str() != "FHIJBC" {
		return "mcl: remove G"
	}
	if cl.push(m); cl.str() != "FHIJBCD" {
		return "mcl: push D"
	}
	if m = cl.removebykey('H'); mkey(m) != 'H' || cl.str() != "IJBCDF" {
		return "mcl: removebykey"
	}
	if m = cl.removebykey('I'); mkey(m) != 'I' || cl.str() != "JBCDF" {
		return "mcl: removebykey"
	}
	if m = cl.removebykey('C'); mkey(m) != 'C' || cl.str() != "DFJB" {
		return "mcl: removebykey"
	}
	if m = cl.removebykey('B'); mkey(m) != 'B' || cl.str() != "DFJ" {
		return "mcl: removebykey"
	}
	if cl.push(m); cl.str() != "DFJB" {
		return "mcl: push"
	}
	if cl.push(&testm); cl.str() != "DFJBG" {
		return "mcl: push G"
	}
	if cl.remove(&testm); cl.str() != "DFJB" {
		return "mcl: remove G"
	}
	if m = cl.removebykey('F'); mkey(m) != 'F' || cl.str() != "JBD" {
		return "mcl: removebykey"
	}
	if m = cl.removebykey('B'); mkey(m) != 'B' || cl.str() != "DJ" {
		return "mcl: removebykey"
	}
	if m = cl.removebykey('J'); mkey(m) != 'J' || cl.str() != "D" {
		return "mcl: removebykey"
	}
	if cl.push(m); cl.str() != "DJ" {
		return "mcl: push"
	}
	if m = cl.removebykey('D'); mkey(m) != 'D' || cl.str() != "J" {
		return "mcl: removebykey"
	}
	if cl.push(m); cl.str() != "JD" {
		return "mq: push"
	}
	if m = cl.removebykey('x'); m != nil || cl.str() != "JD" {
		return "mcl: removebykey, key unknown"
	}
	if m = cl.removebykey('J'); mkey(m) != 'J' || cl.str() != "D" {
		return "mcl: removebykey"
	}
	if m = cl.removebykey('x'); m != nil || cl.str() != "D" {
		return "mcl: removebykey, key unknown"
	}
	if m = cl.removebykey('D'); mkey(m) != 'D' || cl.str() != "" {
		return "mcl: removebykey"
	}
	if m = cl.removebykey('D'); m != nil || cl.str() != "" {
		return "mcl: removebykey, key unknown"
	}

	sl := new(msl)

	sl.insertbyval(mv('3'))
	sl.insertbyval(mv('0'))
	sl.insertbyval(mv('9'))
	sl.insertbyval(mv('1'))
	sl.insertbyval(mv('7'))
	sl.insertbyval(mv('5'))
	sl.insertbyval(mv('2'))
	sl.insertbyval(mv('8'))
	sl.insertbyval(mv('4'))
	sl.insertbyval(mv('6'))
	if sl.str() != "0123456789" {
		return "msl: insertbyval"
	}
	if m = sl.pople('2' << 30); mval(m) != '0'<<30 || sl.str() != "123456789" {
		return "msl: pople"
	}
	if m = sl.pople('2' << 30); mval(m) != '1'<<30 || sl.str() != "23456789" {
		return "msl: pople"
	}
	if m = sl.pople('2' << 30); mval(m) != '2'<<30 || sl.str() != "3456789" {
		return "msl: pople"
	}
	if m = sl.pople('2' << 30); m != nil || sl.str() != "3456789" {
		println(m, sl.str())
		return "msl: pople"
	}
	if m = sl.pople('8' << 30); mval(m) != '3'<<30 || sl.str() != "456789" {
		return "msl: pople"
	}
	if m = sl.pople('8' << 30); mval(m) != '4'<<30 || sl.str() != "56789" {
		return "msl: pople"
	}
	if m = sl.pople('8' << 30); mval(m) != '5'<<30 || sl.str() != "6789" {
		return "msl: pople"
	}
	msetval(&testm, '7'<<30)
	if sl.insertbyval(&testm); sl.str() != "67789" {
		return "msl: insertbyval 7"
	}
	if m = sl.pople('8' << 30); mval(m) != '6'<<30 || sl.str() != "7789" {
		return "msl: pople"
	}
	if m = sl.pople('8' << 30); mval(m) != '7'<<30 || sl.str() != "789" {
		return "msl: pople"
	}
	if sl.remove(&testm); sl.str() != "89" {
		return "msl: remove 7"
	}
	if m = sl.pople('8' << 30); mval(m) != '8'<<30 || sl.str() != "9" {
		return "msl: pople"
	}
	if m = sl.pople('8' << 30); m != nil || sl.str() != "9" {
		return "msl: pople"
	}
	if m = sl.pople('9' << 30); mval(m) != '9'<<30 || sl.str() != "" {
		return "msl: pople"
	}
	if m = sl.pople('9' << 30); m != nil || sl.str() != "" {
		return "msl: pople"
	}
	return ""
}
