package model

import (
	"banduslib/internal/common/music_model"
	"banduslib/internal/database/model/file_type"
	"bytes"
	"encoding/gob"
	"fmt"
)

type TuneFile struct {
	TuneID uint64         `gorm:"primaryKey"`
	Type   file_type.Type `gorm:"primaryKey"`
	Data   []byte
}

func (t *TuneFile) MusicModelTune() (*music_model.Tune, error) {
	if t.Type != file_type.MusicModelTune {
		return nil, fmt.Errorf("tune file has type %s not type %s",
			t.Type.String(), file_type.MusicModelTune.String(),
		)
	}

	if t.Data == nil {
		return nil, fmt.Errorf("can't get music model tune from tune file as data is empty")
	}

	buf := bytes.NewBuffer(t.Data)
	dec := gob.NewDecoder(buf)

	tune := &music_model.Tune{}

	if err := dec.Decode(tune); err != nil {
		return nil, err
	}

	return tune, nil
}

func TuneFileFromTune(tune *music_model.Tune) (*TuneFile, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(tune); err != nil {
		return nil, err
	}

	tf := &TuneFile{
		Type: file_type.MusicModelTune,
		Data: buf.Bytes(),
	}

	return tf, nil
}
