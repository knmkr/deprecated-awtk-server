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
