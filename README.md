# awtk

AWAKENS toolkit for whole genome desktop apps

[![CircleCI](https://circleci.com/gh/AWAKENS-dev/awtk.svg?style=svg)](https://circleci.com/gh/AWAKENS-dev/awtk)

```
                         +--------------------------------+
                         |        your desktop app        |
                         +-------------+---+--------------+
                                       |   ^
                         query via API |   | data as json
                                       v   |
       +-------------------------------+---+---------------------------------+
       |                               awtk                                  |
       +-------------+---+--------------------------------+---+--------------+
                     |   ^                                |   ^
      indexed search |   | genotype data            query |   | evidence data
                     v   |                                v   |
       +-------------+---+--------------+   +-------------+---+--------------+
       |       whole genome data        |   |       evidence database        |
       +--------------------------------+   +--------------------------------+
```

## API Endpoints

### POST /v1/genomes

E.g.

```bash
$ curl -X POST --data "filePath=test/data/test.vcf41.vcf.gz" "localhost:1323/v1/genomes"
```

### GET /v1/genomes

E.g.

```bash
$ curl "localhost:1323/v1/genomes"
```

```json
[
  {"id":1,"filePath":"test/data/test.vcf41.vcf.gz","sampleName":"NA00001","sampleIndex":0},
  {"id":2,"filePath":"test/data/test.vcf41.vcf.gz","sampleName":"NA00002","sampleIndex":1},
  {"id":3,"filePath":"test/data/test.vcf41.vcf.gz","sampleName":"NA00003","sampleIndex":2}
]
```

### GET /v1/genomes/\<id\>

E.g.

```bash
$ curl "localhost:1323/v1/genomes/1"
```

```json
{
  "id": 1,
  "filePath": "test/data/test.vcf41.vcf.gz",
  "sampleName": "NA00001",
  "sampleIndex": 0
}
```

### GET /v1/genomes/\<id\>/genotypes

E.g.

```bash
$ curl "localhost:1323/v1/genomes/1/genotypes?locations=20:14370"
```

```json
{
  "sampleName": "NA00001",
  "genotypes": [
    {
      "chrom": "20",
      "position": 14370,
      "snpId": "rs6054257",
      "genotype": [
        "G",
        "G"
      ],
      "alleles": [
        "G",
        "A"
      ]
    }
  ]
}
```

E.g.

```bash
$ curl "localhost:1323/v1/genomes/1/genotypes?range=20:10000-20000"
```

```json
{
  "sampleName": "NA00001",
  "genotypes": [
    {
      "chrom": "20",
      "position": 14370,
      "snpId": "rs6054257",
      "genotype": [
        "G",
        "G"
      ],
      "alleles": [
        "G",
        "A"
      ]
    },
    {
      "chrom": "20",
      "position": 17330,
      "snpId": ".",
      "genotype": [
        "T",
        "T"
      ],
      "alleles": [
        "T",
        "A"
      ]
    }
  ]
}
```


### GET /v1/evidences/\<id\>

```bash
$ curl "localhost:1323/v1/evidences/1"
```

```json
{
  "id": 1,
  "phenotype": "Coffe consumption",
  "model": "additive",
  "unit": "mg/day decrease",
  "population": "European",
  "variants": [
    {
      "name": "effective homo",
      "criteria": "rs4410790(T;T)",
      "effectSize": 0.3,
      "Effect": "",
      "reference": "https://www.ebi.ac.uk/gwas/search?query=rs4410790"
    },
    {
      "name": "effective hetero",
      "criteria": "rs4410790(T;C)",
      "effectSize": 0.15,
      "Effect": "",
      "reference": "https://www.ebi.ac.uk/gwas/search?query=rs4410790"
    },
    {
      "name": "non-effective homo",
      "criteria": "rs4410790(C;C)",
      "effectSize": 0,
      "Effect": "",
      "reference": "https://www.ebi.ac.uk/gwas/search?query=rs4410790"
    }
  ]
}
```
