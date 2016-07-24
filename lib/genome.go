package wgx

import (
	"fmt"
	// log "github.com/Sirupsen/logrus"
	"github.com/brentp/bix"
)

//
type GenomeError struct {
	Msg string
}

//
func (err *GenomeError) Error() string {
	return fmt.Sprintf("%s", err.Msg)
}

type Genome struct {
	FilePath    string `json:"filePath" form:"filePath"`
	SampleName  string `json:"sampleName"`
	SampleIndex int    `json:"sampleIndex"`
}

func CreateGenomes(filePath string) ([]Genome, error) {
	var genomes []Genome

	tbx, err := bix.New(filePath)
	if err != nil {
		return nil, &GenomeError{"Invalid file or path."}
	}

	vr := tbx.VReader
	sampleNames := vr.Header.SampleNames

	for i := range sampleNames {
		genomes = append(genomes, Genome{filePath, sampleNames[i], i})
	}

	// TODO: save genomes in db

	return genomes, nil
}
