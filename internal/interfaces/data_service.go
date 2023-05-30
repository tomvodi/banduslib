package interfaces

import (
	"banduslib/internal/api/apimodel"
	"banduslib/internal/common/music_model"
	"banduslib/internal/database/model"
	"banduslib/internal/database/model/file_type"
)

//go:generate mockgen -source data_service.go -destination ./mocks/mock_data_service.go

type DataService interface {
	Tunes() ([]*apimodel.Tune, error)
	CreateTune(tune apimodel.CreateTune) (*apimodel.Tune, error)
	GetTune(id uint64) (*apimodel.Tune, error)
	UpdateTune(id uint64, tune apimodel.UpdateTune) (*apimodel.Tune, error)
	DeleteTune(id uint64) error

	AddFileToTune(tuneId uint64, tFile *model.TuneFile) error
	DeleteFileFromTune(tuneId uint64, fType file_type.Type) error
	GetTuneFile(tuneId uint64, fType file_type.Type) (*model.TuneFile, error)
	GetTuneFiles(tuneId uint64) ([]*model.TuneFile, error)

	MusicSets() ([]*apimodel.MusicSet, error)
	CreateMusicSet(tune apimodel.CreateSet) (*apimodel.MusicSet, error)
	GetMusicSet(id uint64) (*apimodel.MusicSet, error)
	UpdateMusicSet(id uint64, tune apimodel.UpdateSet) (*apimodel.MusicSet, error)
	DeleteMusicSet(id uint64) error

	AssignTunesToMusicSet(setId uint64, tuneIds []uint64) (*apimodel.MusicSet, error)

	ImportMusicModel(muMo music_model.MusicModel, filename string) ([]*apimodel.ImportTune, error)
}
