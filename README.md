# awtk

Embeddable RESTful API server for whole genome apps

[![CircleCI](https://circleci.com/gh/AWAKENS-dev/awtk.svg?style=svg)](https://circleci.com/gh/AWAKENS-dev/awtk)

## Motivation

Although there are useful Bioinformatics tools to query genotype information from VCF genome data, such as vcftools, bcftools, or plink, it is not straightforward for **non-bioinformatics software engineers** to utilize them.

So, we implemented `awtk`, an embeddable RESTful API server, which provides query APIs to retrieve genotype information from VCF in RESTful manner.


## How it works

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

Once developers embed `awtk` server in their applications as a middleware, they could utilize genome data as JSON with simple API calls.

`awtk` is implemented in Go, so it works on exisiting application platforms (Windows, OSX, Linux) thus developers could create apps on their desirable platforms including non-web based solutions.


### Requirements

Queries to retrieve genotype are indexed search by tabix. So, input VCFs need to be gzipped by [bzgip](https://github.com/samtools/htslib) and [tabix](http://www.htslib.org/doc/tabix.html) index files (.tbi) need to be created like VCFs distributed in 1000 Genomes Project: ftp://ftp.1000genomes.ebi.ac.uk/vol1/ftp/release/20130502

**NOTE**: Sorry for inconvenience! We are planning to automate this preprocessing within `awtk`.


## How to use

Run following on your application

```
$ awtk runserver
```

Then call API endpoints at `localhost:1323`.

If you want to change the port, run following

```
$ awtk runserver --addr localhost:8888
```

See example applicaiton for embedding and calling API in Electron (Node.js) at: https://github.com/AWAKENS-dev/example-app-electron


### API Endpoints

#### POST /v1/genomes

E.g.

```bash
$ curl -X POST --data "filePath=test/data/test.vcf41.vcf.gz" "localhost:1323/v1/genomes"
```

#### GET /v1/genomes

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

#### GET /v1/genomes/\<id\>

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

#### GET /v1/genomes/\<id\>/genotypes

##### ?locations=\<chrom\>:\<pos\>,\<chrom\>:\<pos\>,...

E.g.

```bash
$ curl "localhost:1323/v1/genomes/1/genotypes?locations=20:14370,20:17330"
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
      ],
      "reference": "G"
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
      ],
      "reference": "T"
    }
  ]
}
```

##### ?range=\<chrom\>:\<pos\>-\<pos\>

E.g.

```bash
$ curl "localhost:1323/v1/genomes/1/genotypes?range=20:14370-14375"
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
      ],
      "reference": "G"
    }
  ]
}
```

##### ?range=\<chrom\>:\<pos\>-\<pos\>&fmt=seq

E.g.

```bash
$ curl "localhost:1323/v1/genomes/1/genotypes?range=20:14370-14375&fmt=seq"
```

```json
{
  "chrom": "20",
  "start": 14370,
  "end": 14375,
  "reference": [
    "G",
    "N",
    "N",
    "N",
    "N",
    "N"
  ],
  "haplotype_1": [
    "G",
    "N",
    "N",
    "N",
    "N",
    "N"
  ],
  "haplotype_2": [
    "G",
    "N",
    "N",
    "N",
    "N",
    "N"
  ]
}
```

#### [WIP] GET /v1/evidences/\<id\>

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
