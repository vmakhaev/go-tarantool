package tarantool

type Ping struct {
}

var _ Query = (*Ping)(nil)

func (q Ping) GetCommandID() uint {
	return PingCommand
}

func (q Ping) PackMsg(data *packData, b []byte) (o []byte, err error) {
	return b, nil
}

// MarshalBinary implements encoding.BinaryMarshaler
func (q *Ping) MarshalBinary() (data []byte, err error) {
	return q.PackMsg(nil, nil)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (q *Ping) UnmarshalBinary(data []byte) (err error) {
	_, err = q.UnmarshalMsg(data)
	return err
}

// UnmarshalMsg implements msgp.Unmarshaller
func (q *Ping) UnmarshalMsg([]byte) (buf []byte, err error) {
	return buf, nil
}
