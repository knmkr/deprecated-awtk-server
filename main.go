package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

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

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

type loc struct {
	chrom string
	start int
	end   int
}

func (s loc) Chrom() string {
	return s.chrom
}
func (s loc) Start() uint32 {
	return uint32(s.start)
}
func (s loc) End() uint32 {
	return uint32(s.end)
}

func main() {

	f := os.Args[1]
	tbx, err := bix.New(f)
	check(err)

	chrom := os.Args[2]

	s, err := strconv.Atoi(os.Args[3])
	check(err)

	e, err := strconv.Atoi(os.Args[4])
	check(err)

	vr := tbx.VReader

	vals, _ := tbx.Query(loc{chrom, s, e})
	i := 0
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
		response, _ := json.Marshal(record)
		fmt.Println(string(response))

		i++
	}
	tbx.Close()

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
