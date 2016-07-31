package awtk

import (
	"fmt"
	// log "github.com/Sirupsen/logrus"
	"encoding/json"
	"gopkg.in/gorp.v1"
)

//
type EvidenceError struct {
	Msg string
}

//
func (err *EvidenceError) Error() string {
	return fmt.Sprintf("%s", err.Msg)
}

type Evidence struct {
	Id         int       `json:"id"`
	Phenotype  string    `json:"phenotype"`
	Model      string    `json:"model"`
	Unit       string    `json:"unit"`
	Population string    `json:"population"`
	Variants   []Variant `db:"-" json:"variants"`
}

type Variant struct {
	Id            int     `json:"-"`
	EvidenceId    int     `json:"-"`
	Name          string  `json:"name"`
	CriteriaBySnp string  `json:"criteriaBySnp"`
	CriteriaByLoc string  `json:"criteriaByLoc"`
	EffectSize    float64 `json:"effectSize"`
	Effect        string  `json:"effect`
	Reference     string  `json:"reference"`
}

// InitEvidence initializes evidences table records.
func InitEvidence(dbmap *gorp.DbMap) error {
	// E.g. additive model from GWAS
	// TODO: load from spreadsheets
	evidence := &Evidence{
		Id:         -1,
		Phenotype:  "Coffe consumption",
		Model:      "additive",
		Unit:       "mg/day decrease",
		Population: "European",
	}
	err := dbmap.Insert(evidence)
	if err != nil {
		return &EvidenceError{fmt.Sprintf("%s", err)}
	}

	variantA1 := &Variant{
		Id:            -1,
		EvidenceId:    evidence.Id,
		Name:          "effective homo",
		CriteriaBySnp: "rs4410790(T;T)",
		CriteriaByLoc: "chr7:17284577(T;T)",
		EffectSize:    0.3,
		Reference:     "https://www.ebi.ac.uk/gwas/search?query=rs4410790",
	}
	variantA2 := &Variant{
		Id:            -1,
		EvidenceId:    evidence.Id,
		Name:          "effective hetero",
		CriteriaBySnp: "rs4410790(T;C)",
		CriteriaByLoc: "chr7:17284577(T;C)",
		EffectSize:    0.15,
		Reference:     "https://www.ebi.ac.uk/gwas/search?query=rs4410790",
	}
	variantA3 := &Variant{
		Id:            -1,
		EvidenceId:    evidence.Id,
		Name:          "non-effective homo",
		CriteriaBySnp: "rs4410790(C;C)",
		CriteriaByLoc: "chr7:17284577(C;C)",
		EffectSize:    0.0,
		Reference:     "https://www.ebi.ac.uk/gwas/search?query=rs4410790",
	}
	err = dbmap.Insert(variantA1, variantA2, variantA3)
	if err != nil {
		return &EvidenceError{fmt.Sprintf("%s", err)}
	}

	// E.g. SNP combination model
	// TODO: load from spreadsheets
	evidence = &Evidence{
		Id:         -1,
		Phenotype:  "Alzheimer's risk (variations in ApoE)",
		Model:      "or",
		Unit:       "",
		Population: "European",
	}
	err = dbmap.Insert(evidence)
	if err != nil {
		return &EvidenceError{fmt.Sprintf("%s", err)}
	}

	variantB1 := &Variant{Id: -1, EvidenceId: evidence.Id,
		Name:          "e1/e1",
		CriteriaBySnp: "and(rs429358(C;C),rs7412(T;T))",
		CriteriaByLoc: "and(chr19:45411941(C;C),chr19:45412079(T;T))",
		Effect:        "the rare missing allele",
		Reference:     "http://snpedia.com/index.php/APOE"}
	variantB2 := &Variant{Id: -1, EvidenceId: evidence.Id,
		Name:          "e1/e2",
		CriteriaBySnp: "and(rs429358(C;T),rs7412(T;T))",
		CriteriaByLoc: "and(chr19:45411941(C;T),chr19:45412079(T;T))",
		Effect:        "",
		Reference:     "http://snpedia.com/index.php/APOE"}
	variantB3 := &Variant{Id: -1, EvidenceId: evidence.Id,
		Name:          "e1/e3",
		CriteriaBySnp: "and(rs429358(C;T),rs7412(C;T))",
		CriteriaByLoc: "and(chr19:45411941(C;T),chr19:45412079(C;T))",
		Effect:        "ambiguous with e2/e4",
		Reference:     "http://snpedia.com/index.php/APOE"}
	variantB4 := &Variant{Id: -1, EvidenceId: evidence.Id,
		Name:          "e1/e4",
		CriteriaBySnp: "and(rs429358(C;C),rs7412(C;T))",
		CriteriaByLoc: "and(chr19:45411941(C;C),chr19:45412079(C;T))",
		Effect:        "",
		Reference:     "http://snpedia.com/index.php/APOE"}
	variantB5 := &Variant{Id: -1, EvidenceId: evidence.Id,
		Name:          "e2/e2",
		CriteriaBySnp: "and(rs429358(T;T),rs7412(T;T))",
		CriteriaByLoc: "and(chr19:45411941(T;T),chr19:45412079(T;T))",
		Effect:        "",
		Reference:     "http://snpedia.com/index.php/APOE"}
	variantB6 := &Variant{Id: -1, EvidenceId: evidence.Id,
		Name:          "e2/e3",
		CriteriaBySnp: "and(rs429358(T;T),rs7412(C;T))",
		CriteriaByLoc: "and(chr19:45411941(T;T),chr19:45412079(C;T))",
		Effect:        "",
		Reference:     "http://snpedia.com/index.php/APOE"}
	variantB7 := &Variant{Id: -1, EvidenceId: evidence.Id,
		Name:          "e2/e4",
		CriteriaBySnp: "and(rs429358(C;T),rs7412(C;T))",
		CriteriaByLoc: "and(chr19:45411941(C;T),chr19:45412079(C;T))",
		Effect:        "ambiguous with e1/e3",
		Reference:     "http://snpedia.com/index.php/APOE"}
	variantB8 := &Variant{Id: -1, EvidenceId: evidence.Id,
		Name:          "e3/e3",
		CriteriaBySnp: "and(rs429358(T;T),rs7412(C;C))",
		CriteriaByLoc: "and(chr19:45411941(T;T),chr19:45412079(C;C))",
		Effect:        "the most common",
		Reference:     "http://snpedia.com/index.php/APOE"}
	variantB9 := &Variant{Id: -1, EvidenceId: evidence.Id,
		Name:          "e3/e4",
		CriteriaBySnp: "and(rs429358(C;T),rs7412(C;C))",
		CriteriaByLoc: "and(chr19:45411941(C;T),chr19:45412079(C;C))",
		Effect:        "",
		Reference:     "http://snpedia.com/index.php/APOE"}
	variantB10 := &Variant{Id: -1, EvidenceId: evidence.Id,
		Name:          "e4/e4",
		CriteriaBySnp: "and(rs429358(C;C),rs7412(C;C))",
		CriteriaByLoc: "and(chr19:45411941(C;C),chr19:45412079(C;C))",
		Effect:        "~11x increased Alzheimer's risk",
		Reference:     "http://snpedia.com/index.php/APOE"}
	err = dbmap.Insert(variantB1, variantB2, variantB3, variantB4, variantB5, variantB6, variantB7, variantB8, variantB9, variantB10)
	if err != nil {
		return &EvidenceError{fmt.Sprintf("%s", err)}
	}

	return nil
}

func GetEvidence(evidenceId int) ([]byte, error) {
	var evidence *Evidence

	db, dbmap, err := GetDatabaseConnection()
	if err != nil {
		return nil, &EvidenceError{fmt.Sprintf("%s", err)}
	}
	defer db.Close()
	defer dbmap.Db.Close()

	err = dbmap.SelectOne(&evidence, "SELECT * FROM evidences WHERE id = ?", evidenceId)
	if err != nil {
		return nil, &EvidenceError{fmt.Sprintf("%s", err)}
	}

	var variants []Variant
	_, err = dbmap.Select(&variants, "SELECT * FROM variants WHERE evidenceid = ?", evidence.Id)
	if err != nil {
		return nil, &EvidenceError{fmt.Sprintf("%s", err)}
	}

	evidence.Variants = variants

	response, err := json.Marshal(evidence)
	if err != nil {
		return nil, err
	}

	return response, nil
}
