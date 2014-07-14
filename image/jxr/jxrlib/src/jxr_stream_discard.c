// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "jxr.h"
#include "jxr_private.h"

static
ERR CloseWS_Discard(struct WMPStream** ppWS)
{
	ERR err = WMP_errSuccess;

	Call(WMPFree((void**)ppWS));

Cleanup:
	return err;
}

static
Bool EOSWS_Discard(struct WMPStream* pWS)
{
	return pWS->state.buf.cbBuf <= pWS->state.buf.cbCur;
}

static
ERR ReadWS_Discard(struct WMPStream* pWS, void* pv, size_t cb)
{
	ERR err = WMP_errSuccess;

	if(pWS->state.buf.cbBuf < pWS->state.buf.cbCur) {
		return WMP_errBufferOverflow;
	}
	if(pWS->state.buf.cbBuf < pWS->state.buf.cbCur + cb) {
		cb = pWS->state.buf.cbBuf - pWS->state.buf.cbCur;
	}
	// memcpy(pv, pWS->state.buf.pbBuf + pWS->state.buf.cbCur, cb);
	memset(pv, 0, cb);
	pWS->state.buf.cbCur += cb;

Cleanup:
	return err;
}

static
ERR WriteWS_Discard(struct WMPStream* pWS, const void* pv, size_t cb)
{
	ERR err = WMP_errSuccess;

	FailIf(pWS->state.buf.cbCur + cb < pWS->state.buf.cbCur, WMP_errBufferOverflow);
	FailIf(pWS->state.buf.cbBuf < pWS->state.buf.cbCur + cb, WMP_errBufferOverflow);

	// memcpy(pWS->state.buf.pbBuf + pWS->state.buf.cbCur, pv, cb);
	pWS->state.buf.cbCur += cb;

Cleanup:
	return err;
}

static
ERR SetPosWS_Discard(struct WMPStream* pWS, size_t offPos)
{
	// While the following condition is possibly useful, failure occurs
	// at the end of a file since packets beyond the end may be accessed
	// FailIf(pWS->state.buf.cbBuf < offPos, WMP_errBufferOverflow);
	pWS->state.buf.cbCur = offPos;
	return WMP_errSuccess;
}

static
ERR GetPosWS_Discard(struct WMPStream* pWS, size_t* poffPos)
{
	*poffPos = pWS->state.buf.cbCur;
	return WMP_errSuccess;
}

ERR CreateWS_Discard(struct WMPStream** ppWS)
{
	ERR err = WMP_errSuccess;
	struct WMPStream* pWS = NULL;

	Call(WMPAlloc((void** )ppWS, sizeof(**ppWS)));
	pWS = *ppWS;

	pWS->state.buf.pbBuf = NULL;
	pWS->state.buf.cbBuf = (1<<30); // 1GB
	pWS->state.buf.cbCur = 0;

	pWS->Close = CloseWS_Discard;
	pWS->EOS = EOSWS_Discard;

	pWS->Read = ReadWS_Discard;
	pWS->Write = WriteWS_Discard;

	pWS->SetPos = SetPosWS_Discard;
	pWS->GetPos = GetPosWS_Discard;

Cleanup:
	return err;
}
