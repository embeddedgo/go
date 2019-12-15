// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Some syscalls are alowed to be called by low priority interrupt handlers.
//
// Other syscalls, that can't run concurently with the scheduler, are available
// only to threads.

#define SYS_nanotime      0
#define SYS_walltime      1
#define SYS_setwalltime   2
#define SYS_irqctl        3
#define SYS_setprivlevel  4
#define SYS_write         5

#define SYS_LAST_LOWISR   5 

#define SYS_setsystim1    6
#define SYS_newosproc     7
#define SYS_exitThread    8
#define SYS_futexsleep    9
#define SYS_futexwakeup  10
#define SYS_osyield      11
#define SYS_nanosleep    12

#define SYS_NUM          13
