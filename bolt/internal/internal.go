package internal

import (
	"github.com/bradgwest/stott"
	"github.com/gogo/protobuf/proto"
)

//go:generate protoc --go_out=. internal.proto

// MarshalStatus encodes a status to binary format.
func MarshalStatus(d *stott.Status) ([]byte, error) {
	return proto.Marshal(&Status{
		ID:     proto.Int64(int64(d.ID)),
		UserID: proto.Int64(int64(d.UserID)),
		Text:   proto.String(d.Text),
	})
}

// UnmarshalStatus decodes a status from a binary data.
func UnmarshalStatus(data []byte, d *stott.Status) error {
	var pb Status
	if err := proto.Unmarshal(data, &pb); err != nil {
		return err
	}

	d.ID = stott.StatusID(pb.GetID())
	d.UserID = stott.UserID(pb.GetUserID())
	d.Text = pb.GetText()

	return nil
}
