package awtk

import (
	"bytes"
	"encoding/json"
	// log "github.com/Sirupsen/logrus"
	"github.com/brentp/bix"
	"github.com/brentp/irelate/interfaces"
)

type Genotype struct {
	Chrom    string   `json:"chrom"`
	Position int      `json:"position"`
	SnpId    string   `json:"snpId"`
	Genotype []string `json:"genotype"`
	Alleles  []string `json:"alleles"`
}

type Genotypes struct {
	SampleName string     `json:"sampleName"`
	Genotypes  []Genotype `json:"genotypes"`
}

func (genotypes *Genotypes) AddGenotype(genotype Genotype) []Genotype {
	genotypes.Genotypes = append(genotypes.Genotypes, genotype)
	return genotypes.Genotypes
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

func QueryGenotypes(f string, idx int, locs []Location) ([]byte, error) {
	var genotypes Genotypes
	var sampleName string

	tbx, err := bix.New(f)
	if err != nil {
		return nil, err
	}

	vr := tbx.VReader

	for i := range locs {
		vals, _ := tbx.Query(locs[i])

		for {
			v, err := vals.Next()

			if err != nil {
				break
			}

			// Parse sample names
			line := []byte(v.(interfaces.IVariant).String())
			fields := makeFields(line)
			variant := vr.Parse(fields)
			vr.Header.ParseSamples(variant)
			sampleNames := vr.Header.SampleNames
			samples := variant.Samples

			chrom := v.(interfaces.IPosition).Chrom()
			pos := v.(interfaces.IPosition).End()
			snpId := v.(interfaces.IVariant).Id()
			// info := v.(interfaces.IVariant).Info()

			// Parse alleles
			ref := v.(interfaces.IVariant).Ref()
			alt := v.(interfaces.IVariant).Alt()
			alleles := []string{}
			alleles = append(alleles, ref)

			if len(alt) == 1 && alt[0] == "." {
				// no ALTs (= ALT is ["."])
			} else {
				alleles = append(alleles, alt...)
			}

			// Get genotypes of 1st sample
			sample := samples[idx]
			sampleName = sampleNames[idx]

			genotype := []string{}
			gt := sample.GT

			for j := range gt {
				if gt[j] == -1 {
					// no GTs (= GT is missing value: -1)
					genotype = append(genotype, ".")
				} else {
					genotype = append(genotype, alleles[gt[j]])
				}

			}

			genotypes.AddGenotype(Genotype{chrom,
				int(pos),
				snpId,
				genotype,
				alleles})
		}
	}

	tbx.Close()

	genotypes.SampleName = sampleName
	response, err := json.Marshal(genotypes)
	if err != nil {
		return nil, err
	}

	return response, nil
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
