package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartAxis 
type WorkbookChartAxis struct {
    Entity
    // Represents the formatting of a chart object, which includes line and font formatting. Read-only.
    format WorkbookChartAxisFormatable
    // Returns a gridlines object that represents the major gridlines for the specified axis. Read-only.
    majorGridlines WorkbookChartGridlinesable
    // Returns a Gridlines object that represents the minor gridlines for the specified axis. Read-only.
    minorGridlines WorkbookChartGridlinesable
    // Represents the axis title. Read-only.
    title WorkbookChartAxisTitleable
}
// NewWorkbookChartAxis instantiates a new workbookChartAxis and sets the default values.
func NewWorkbookChartAxis()(*WorkbookChartAxis) {
    m := &WorkbookChartAxis{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartAxisFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartAxisFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookChartAxis(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartAxis) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartAxisFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(WorkbookChartAxisFormatable))
        }
        return nil
    }
    res["majorGridlines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartGridlinesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMajorGridlines(val.(WorkbookChartGridlinesable))
        }
        return nil
    }
    res["minorGridlines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartGridlinesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinorGridlines(val.(WorkbookChartGridlinesable))
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartAxisTitleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val.(WorkbookChartAxisTitleable))
        }
        return nil
    }
    return res
}
// GetFormat gets the format property value. Represents the formatting of a chart object, which includes line and font formatting. Read-only.
func (m *WorkbookChartAxis) GetFormat()(WorkbookChartAxisFormatable) {
    return m.format
}
// GetMajorGridlines gets the majorGridlines property value. Returns a gridlines object that represents the major gridlines for the specified axis. Read-only.
func (m *WorkbookChartAxis) GetMajorGridlines()(WorkbookChartGridlinesable) {
    return m.majorGridlines
}
// GetMinorGridlines gets the minorGridlines property value. Returns a Gridlines object that represents the minor gridlines for the specified axis. Read-only.
func (m *WorkbookChartAxis) GetMinorGridlines()(WorkbookChartGridlinesable) {
    return m.minorGridlines
}
// GetTitle gets the title property value. Represents the axis title. Read-only.
func (m *WorkbookChartAxis) GetTitle()(WorkbookChartAxisTitleable) {
    return m.title
}
// Serialize serializes information the current object
func (m *WorkbookChartAxis) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("format", m.GetFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("majorGridlines", m.GetMajorGridlines())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minorGridlines", m.GetMinorGridlines())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFormat sets the format property value. Represents the formatting of a chart object, which includes line and font formatting. Read-only.
func (m *WorkbookChartAxis) SetFormat(value WorkbookChartAxisFormatable)() {
    m.format = value
}
// SetMajorGridlines sets the majorGridlines property value. Returns a gridlines object that represents the major gridlines for the specified axis. Read-only.
func (m *WorkbookChartAxis) SetMajorGridlines(value WorkbookChartGridlinesable)() {
    m.majorGridlines = value
}
// SetMinorGridlines sets the minorGridlines property value. Returns a Gridlines object that represents the minor gridlines for the specified axis. Read-only.
func (m *WorkbookChartAxis) SetMinorGridlines(value WorkbookChartGridlinesable)() {
    m.minorGridlines = value
}
// SetTitle sets the title property value. Represents the axis title. Read-only.
func (m *WorkbookChartAxis) SetTitle(value WorkbookChartAxisTitleable)() {
    m.title = value
}
// WorkbookChartAxisable 
type WorkbookChartAxisable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFormat()(WorkbookChartAxisFormatable)
    GetMajorGridlines()(WorkbookChartGridlinesable)
    GetMinorGridlines()(WorkbookChartGridlinesable)
    GetTitle()(WorkbookChartAxisTitleable)
    SetFormat(value WorkbookChartAxisFormatable)()
    SetMajorGridlines(value WorkbookChartGridlinesable)()
    SetMinorGridlines(value WorkbookChartGridlinesable)()
    SetTitle(value WorkbookChartAxisTitleable)()
}
