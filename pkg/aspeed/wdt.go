// SPDX-License-Identifier: BSD-3-Clause

package aspeed

const (
	WDT1ReloadControl uintptr = 0x1e785004
	WDT1Restart       uintptr = 0x1e785008
	WDT1Control       uintptr = 0x1e78500c
	WDT2Control       uintptr = 0x1e78502c
	WDT2TimeoutClear  uintptr = 0x1e785034

	WDTRestartPassword uint32 = 0x4755
)

func DisableWDT() error {
	if err := writeMem(WDT1Control, 0); err != nil {
		return err
	}

	return writeMem(WDT2Control, 0)
}

func EnableWDT() error {
	// 0x1 - Reset boot code source select
	if err := writeMem(WDT2TimeoutClear, 0x1); err != nil {
		return err
	}
	// 0x80 - Use second boot code whenever WDT reset
	// 0x2  - Reset system after timeout
	return writeMem(WDT2Control, 0x82)

	// Old WDT1 is not saved so nothing to restore.
}

func ResetCPU() error {
	// - Load 16 into WDT00 when reset/restart
	// 16 is a small value that will quickly trigger
	if err := writeMem(WDT1ReloadControl, 16); err != nil {
		return err
	}
	// 0x4755 - Restart password
	if err := writeMem(WDT1Restart, WDTRestartPassword); err != nil {
		return err
	}
	// 0x2 - Reset system after timeout
	// 0x1 - WDT enable
	return writeMem(WDT1Control, 0x3)
}
