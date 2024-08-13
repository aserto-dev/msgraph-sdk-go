package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookTableRow 
type WorkbookTableRow struct {
    Entity
    // The index of the row within the rows collection of the table. Zero-based. Read-only.
    index *int32
}
// NewWorkbookTableRow instantiates a new workbookTableRow and sets the default values.
func NewWorkbookTableRow()(*WorkbookTableRow) {
    m := &WorkbookTableRow{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookTableRowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookTableRowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookTableRow(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookTableRow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["index"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndex(val)
        }
        return nil
    }
    return res
}
// GetIndex gets the index property value. The index of the row within the rows collection of the table. Zero-based. Read-only.
func (m *WorkbookTableRow) GetIndex()(*int32) {
    return m.index
}
// Serialize serializes information the current object
func (m *WorkbookTableRow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("index", m.GetIndex())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIndex sets the index property value. The index of the row within the rows collection of the table. Zero-based. Read-only.
func (m *WorkbookTableRow) SetIndex(value *int32)() {
    m.index = value
}
// WorkbookTableRowable 
type WorkbookTableRowable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIndex()(*int32)
    SetIndex(value *int32)()
}
