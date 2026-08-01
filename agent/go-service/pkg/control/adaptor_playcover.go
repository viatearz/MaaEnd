// Copyright (c) 2026 MaaEnd Contributors
package control

import (
	maa "github.com/MaaXYZ/maa-framework-go/v4"
)

type PlayCoverControlAdaptor struct {
	*ADBControlAdaptor
}

func newPlayCoverControlAdaptor(ctx *maa.Context, ctrl *maa.Controller, w, h int) *PlayCoverControlAdaptor {
	return &PlayCoverControlAdaptor{
		ADBControlAdaptor: newADBControlAdaptor(ctx, ctrl, w, h),
	}
}
