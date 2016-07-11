package wgx

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/brentp/bix"
	"github.com/brentp/irelate/interfaces"
)

type Genotype struct {
	SampleName string   `json:"sampleName"`
	Chrom      string   `json:"chrom"`
	Position   int      `json:"position"`
	Id         string   `json:"id"`
	Genotype   []string `json:"genotype"`
	Alleles    []string `json:"alleles"`
}

// FIXME
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

type Location struct {
	chrom string
	start int
	end   int
}

func (s Location) Chrom() string {
	return s.chrom
}
func (s Location) Start() uint32 {
	return uint32(s.start)
}
func (s Location) End() uint32 {
	return uint32(s.end)
}

func NewLocation(chrom string, start int, end int) Location {
	return Location{chrom, start, end}
}

func QueryGenotypes(f string, loc Location) []byte {
	var response []byte

	tbx, err := bix.New(f)
	check(err)

	vr := tbx.VReader

	vals, _ := tbx.Query(loc)

	for {
		v, err := vals.Next()
		if err != nil {
			break
		}

		chrom := v.(interfaces.IPosition).Chrom()
		pos := v.(interfaces.IPosition).End()
		id_ := v.(interfaces.IVariant).Id()
		// info_ := v.(interfaces.IVariant).Info()

		// Parse alleles
		ref := v.(interfaces.IVariant).Ref()
		alt := v.(interfaces.IVariant).Alt()
		alleles := []string{}
		alleles = append(alleles, ref)
		alleles = append(alleles, alt...)

		// Parse samples
		line := []byte(v.(interfaces.IVariant).String())
		fields := makeFields(line)
		variant := vr.Parse(fields)
		vr.Header.ParseSamples(variant)
		sampleNames := vr.Header.SampleNames
		samples := variant.Samples

		// Get genotypes
		idx := 0
		sample := samples[idx]
		sampleName := sampleNames[idx]

		genotype := []string{}
		gt := sample.GT
		for j := range gt {
			genotype = append(genotype, alleles[gt[j]])
		}

		// jsonify
		record := &Genotype{
			SampleName: sampleName,
			Chrom:      chrom,
			Position:   int(pos),
			Id:         id_,
			Genotype:   genotype,
			Alleles:    alleles}
		response, err = json.Marshal(record)

	}
	tbx.Close()

	// FIXME
	return response
}

// copied from github.com/brentp/bix
func makeFields(line []byte) [][]byte {
	fields := make([][]byte, 9)
	copy(fields[:8], bytes.SplitN(line, []byte{'\t'}, 8))
	s := 0
	for i, f := range fields {
		if i == 7 {
			break
		}
		s += len(f) + 1
	}
	e := bytes.IndexByte(line[s:], '\t')
	if e == -1 {
		e = len(line)
	} else {
		e += s
	}

	fields[7] = line[s:e]
	if len(line) > e+1 {
		fields[8] = line[e+1:]
	} else {
		fields = fields[:8]
	}

	return fields
}
