// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#define SYS_nanotime      0
#define SYS_setsystim1    1
#define SYS_walltime      2
#define SYS_setwalltime   3
#define SYS_irqctl        4
#define SYS_setprivlevel  5
#define SYS_write         6

#define SYS_LAST_NONBLOCK 6

#define SYS_newosproc     7
#define SYS_exitThread    8
#define SYS_futexsleep    9
#define SYS_futexwakeup  10
#define SYS_osyield      11
#define SYS_nanosleep    12

#define SYS_NUM          13
