package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CallRecording 
type CallRecording struct {
    Entity
    // The content of the recording. Read-only.
    content []byte
    // Date and time at which the recording was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The unique identifier of the onlineMeeting related to this recording. Read-only.
    meetingId *string
    // The identity information of the organizer of the onlineMeeting related to this recording. Read-only.
    meetingOrganizer IdentitySetable
    // The URL that can be used to access the content of the recording. Read-only.
    recordingContentUrl *string
}
// NewCallRecording instantiates a new callRecording and sets the default values.
func NewCallRecording()(*CallRecording) {
    m := &CallRecording{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCallRecordingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallRecordingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallRecording(), nil
}
// GetContent gets the content property value. The content of the recording. Read-only.
func (m *CallRecording) GetContent()([]byte) {
    return m.content
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time at which the recording was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *CallRecording) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallRecording) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["meetingId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingId(val)
        }
        return nil
    }
    res["meetingOrganizer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingOrganizer(val.(IdentitySetable))
        }
        return nil
    }
    res["recordingContentUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordingContentUrl(val)
        }
        return nil
    }
    return res
}
// GetMeetingId gets the meetingId property value. The unique identifier of the onlineMeeting related to this recording. Read-only.
func (m *CallRecording) GetMeetingId()(*string) {
    return m.meetingId
}
// GetMeetingOrganizer gets the meetingOrganizer property value. The identity information of the organizer of the onlineMeeting related to this recording. Read-only.
func (m *CallRecording) GetMeetingOrganizer()(IdentitySetable) {
    return m.meetingOrganizer
}
// GetRecordingContentUrl gets the recordingContentUrl property value. The URL that can be used to access the content of the recording. Read-only.
func (m *CallRecording) GetRecordingContentUrl()(*string) {
    return m.recordingContentUrl
}
// Serialize serializes information the current object
func (m *CallRecording) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("meetingId", m.GetMeetingId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("meetingOrganizer", m.GetMeetingOrganizer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recordingContentUrl", m.GetRecordingContentUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The content of the recording. Read-only.
func (m *CallRecording) SetContent(value []byte)() {
    m.content = value
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time at which the recording was created. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *CallRecording) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetMeetingId sets the meetingId property value. The unique identifier of the onlineMeeting related to this recording. Read-only.
func (m *CallRecording) SetMeetingId(value *string)() {
    m.meetingId = value
}
// SetMeetingOrganizer sets the meetingOrganizer property value. The identity information of the organizer of the onlineMeeting related to this recording. Read-only.
func (m *CallRecording) SetMeetingOrganizer(value IdentitySetable)() {
    m.meetingOrganizer = value
}
// SetRecordingContentUrl sets the recordingContentUrl property value. The URL that can be used to access the content of the recording. Read-only.
func (m *CallRecording) SetRecordingContentUrl(value *string)() {
    m.recordingContentUrl = value
}
// CallRecordingable 
type CallRecordingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()([]byte)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMeetingId()(*string)
    GetMeetingOrganizer()(IdentitySetable)
    GetRecordingContentUrl()(*string)
    SetContent(value []byte)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMeetingId(value *string)()
    SetMeetingOrganizer(value IdentitySetable)()
    SetRecordingContentUrl(value *string)()
}
