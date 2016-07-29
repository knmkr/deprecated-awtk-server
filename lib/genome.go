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
	Id          int
	FilePath    string `json:"filePath" form:"filePath"`
	SampleName  string `json:"sampleName"`
	SampleIndex int    `json:"sampleIndex"`
}

func CreateGenomes(filePath string) ([]Genome, error) {
	tbx, err := bix.New(filePath)
	if err != nil {
		return nil, &GenomeError{fmt.Sprintf("%s", err)}
	}

	vr := tbx.VReader
	sampleNames := vr.Header.SampleNames

	db, dbmap, err := GetDatabaseConnection()
	if err != nil {
		return nil, &GenomeError{fmt.Sprintf("%s", err)}
	}
	defer db.Close()
	defer dbmap.Db.Close()

	tx, err := dbmap.Begin()
	if err != nil {
		return nil, &GenomeError{fmt.Sprintf("%s", err)}
	}

	var genomes []Genome
	for i := range sampleNames {
		genome := &Genome{
			FilePath:    filePath,
			SampleName:  sampleNames[i],
			SampleIndex: i}
		genomes = append(genomes, *genome)

		err = tx.Insert(genome)
		if err != nil {
			msg := fmt.Sprintf("%s", err) //
			err = tx.Rollback()
			return nil, &GenomeError{fmt.Sprintf("%s. %s", msg, err)}
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, &GenomeError{fmt.Sprintf("%s", err)}
	}

	return genomes, nil
}

func GetGenome(genomeId int) (*Genome, error) {
	var genome *Genome

	db, dbmap, err := GetDatabaseConnection()
	if err != nil {
		return nil, &GenomeError{fmt.Sprintf("%s", err)}
	}
	defer db.Close()
	defer dbmap.Db.Close()

	err = dbmap.SelectOne(&genome, "SELECT * FROM genome WHERE id = ?", genomeId)
	if err != nil {
		return nil, &GenomeError{fmt.Sprintf("%s", err)}
	}

	return genome, nil
}
