package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSpacerItem struct {
	QLayoutItem
}

type QSpacerItem_ITF interface {
	QLayoutItem_ITF
	QSpacerItem_PTR() *QSpacerItem
}

func PointerFromQSpacerItem(ptr QSpacerItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSpacerItem_PTR().Pointer()
	}
	return nil
}

func NewQSpacerItemFromPointer(ptr unsafe.Pointer) *QSpacerItem {
	var n = new(QSpacerItem)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QSpacerItem_") {
		n.SetObjectNameAbs("QSpacerItem_" + qt.Identifier())
	}
	return n
}

func (ptr *QSpacerItem) QSpacerItem_PTR() *QSpacerItem {
	return ptr
}

func NewQSpacerItem(w int, h int, hPolicy QSizePolicy__Policy, vPolicy QSizePolicy__Policy) *QSpacerItem {
	defer qt.Recovering("QSpacerItem::QSpacerItem")

	return NewQSpacerItemFromPointer(C.QSpacerItem_NewQSpacerItem(C.int(w), C.int(h), C.int(hPolicy), C.int(vPolicy)))
}

func (ptr *QSpacerItem) ChangeSize(w int, h int, hPolicy QSizePolicy__Policy, vPolicy QSizePolicy__Policy) {
	defer qt.Recovering("QSpacerItem::changeSize")

	if ptr.Pointer() != nil {
		C.QSpacerItem_ChangeSize(ptr.Pointer(), C.int(w), C.int(h), C.int(hPolicy), C.int(vPolicy))
	}
}

func (ptr *QSpacerItem) ExpandingDirections() core.Qt__Orientation {
	defer qt.Recovering("QSpacerItem::expandingDirections")

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QSpacerItem_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSpacerItem) IsEmpty() bool {
	defer qt.Recovering("QSpacerItem::isEmpty")

	if ptr.Pointer() != nil {
		return C.QSpacerItem_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSpacerItem) SpacerItem() *QSpacerItem {
	defer qt.Recovering("QSpacerItem::spacerItem")

	if ptr.Pointer() != nil {
		return NewQSpacerItemFromPointer(C.QSpacerItem_SpacerItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSpacerItem) DestroyQSpacerItem() {
	defer qt.Recovering("QSpacerItem::~QSpacerItem")

	if ptr.Pointer() != nil {
		C.QSpacerItem_DestroyQSpacerItem(ptr.Pointer())
	}
}

func (ptr *QSpacerItem) ObjectNameAbs() string {
	defer qt.Recovering("QSpacerItem::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSpacerItem_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSpacerItem) SetObjectNameAbs(name string) {
	defer qt.Recovering("QSpacerItem::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QSpacerItem_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}