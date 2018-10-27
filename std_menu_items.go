package webapp

import (
	"fmt"

	"github.com/richardwilkes/toolbox/cmdline"
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/webapp/keys"
)

// NewCloseKeyWindowItem creates the standard "Close" menu item that will
// close the current key window when chosen.
func NewCloseKeyWindowItem() *MenuItem {
	item := NewMenuItemWithKey(i18n.Text("Close"), keys.VirtualKeyW)
	item.Validator = func() bool {
		wnd := KeyWindow()
		return wnd != nil && wnd.Closable()
	}
	item.Handler = func() {
		if wnd := KeyWindow(); wnd != nil && wnd.Closable() {
			wnd.AttemptClose()
		}
	}
	return item
}

// NewMinimizeWindowItem creates the standard "Minimize" menu item.
func NewMinimizeWindowItem() *MenuItem {
	item := NewMenuItemWithKey(i18n.Text("Minimize"), keys.VirtualKeyM)
	item.Validator = func() bool {
		w := KeyWindow()
		return w != nil && w.Minimizable()
	}
	item.Handler = func() {
		if wnd := KeyWindow(); wnd != nil {
			wnd.Minimize()
		}
	}
	return item
}

// NewZoomWindowItem creates the standard "Zoom" menu item.
func NewZoomWindowItem() *MenuItem {
	item := NewMenuItemWithKeyAndModifiers(i18n.Text("Zoom"), keys.VirtualKeyZ, keys.ShiftModifier|keys.PlatformMenuModifier())
	item.Validator = func() bool {
		w := KeyWindow()
		return w != nil && w.Resizable()
	}
	item.Handler = func() {
		if wnd := KeyWindow(); wnd != nil {
			wnd.Zoom()
		}
	}
	return item
}

// NewBringAllWindowsToFrontItem creates the standard "Bring All to Front"
// item.
func NewBringAllWindowsToFrontItem() *MenuItem {
	item := NewMenuItem(i18n.Text("Bring All to Front"))
	item.Handler = AllWindowsToFront
	return item
}

// NewQuitItem creates the standard "Quit" item.
func NewQuitItem() *MenuItem {
	item := NewMenuItemWithKey(fmt.Sprintf(i18n.Text("Quit %s"), cmdline.AppName), keys.VirtualKeyQ)
	item.Handler = AttemptQuit
	return item
}

// NewCutItem creates the standard "Cut" item.
func NewCutItem() *MenuItem {
	item := NewMenuItemWithKey(i18n.Text("Cut"), keys.VirtualKeyX)
	// RAW: Implement validator and handler
	item.Validator = func() bool { return false }
	return item
}

// NewCopyItem creates the standard "Copy" item.
func NewCopyItem() *MenuItem {
	item := NewMenuItemWithKey(i18n.Text("Copy"), keys.VirtualKeyC)
	// RAW: Implement validator and handler
	item.Validator = func() bool { return false }
	return item
}

// NewPasteItem creates the standard "Paste" item.
func NewPasteItem() *MenuItem {
	item := NewMenuItemWithKey(i18n.Text("Paste"), keys.VirtualKeyV)
	// RAW: Implement validator and handler
	item.Validator = func() bool { return false }
	return item
}

// NewDeleteItem creates the standard "Delete" item.
func NewDeleteItem() *MenuItem {
	item := NewMenuItemWithKeyAndModifiers(i18n.Text("Delete"), keys.VirtualKeyBackspace, 0)
	// RAW: Implement validator and handler
	item.Validator = func() bool { return false }
	return item
}

// NewSelectAllItem creates the standard "Select All" item.
func NewSelectAllItem() *MenuItem {
	item := NewMenuItemWithKey(i18n.Text("Select All"), keys.VirtualKeyA)
	// RAW: Implement validator and handler
	item.Validator = func() bool { return false }
	return item
}
