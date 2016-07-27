# wgx

whole genome query engine for desktop apps

[![CircleCI](https://circleci.com/gh/knmkr/wgx.svg?style=svg)](https://circleci.com/gh/knmkr/wgx) 

```
       +--------------------------------+
       |       whole genome data        |
       +-------------+---+--------------+
                     ^   |
indexed search, etc. |   |
                     |   v
       +-------------+---+--------------+
       |              wgx               |
       |         (query engine)         |
       +-------------+---+--------------+
                     ^   |
       query via API |   | results as json
                     |   v
       +-------------+---+--------------+
       |        your desktop app        |
       +--------------------------------+
```

## Endpoints

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
      "id": "rs6054257",
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
