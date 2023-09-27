// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

TEXT runtime·save_g(SB),NOSPLIT|NOFRAME,$0-0
	RET

TEXT runtime·load_g(SB),NOSPLIT|NOFRAME,$0-0
	RET
