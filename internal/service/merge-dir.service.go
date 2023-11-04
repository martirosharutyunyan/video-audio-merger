package service

type IMergeVideoAudioDirService interface {
    Merge(source, output string) error
}

type MergeVideoAudioDirService struct {}

func (m MergeVideoAudioDirService) Merge(source, output string) error {

	return nil
}

func NewMergeVideoAudioDirService() *MergeVideoAudioDirService {
	return &MergeVideoAudioDirService{}
}

