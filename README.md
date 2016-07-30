# wgx

whole genome query engine for desktop apps

[![CircleCI](https://circleci.com/gh/knmkr/wgx.svg?style=svg)](https://circleci.com/gh/knmkr/wgx) 

```
       +--------------------------------+
       |        your desktop app        |
       +-------------+---+--------------+
                     ^   |
    genotype as json |   | query via API
                     |   v
       +-------------+---+--------------+
       |              wgx               |
       |         (query engine)         |
       +-------------+---+--------------+
                     ^   |
       genotype data |   | indexed search
                     |   v
       +-------------+---+--------------+
       |       whole genome data        |
       +--------------------------------+
```

## API Endpoints

### POST /v1/genomes

E.g.

```bash
$ curl -X POST --data "filePath=test/data/test.vcf41.vcf.gz" "localhost:1323/v1/genomes"
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
$ curl "localhost:1323/v1/genomes/1/genotypes?locations=20-14370"
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
