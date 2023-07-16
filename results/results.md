# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.10

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHz

Cores: 1

Mhz: 2095.078000

CacheSize: 36608

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHz

Cores: 1

Mhz: 2095.078000

CacheSize: 36608

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 578.081968 | 16.384482 | 4215.012152 | 25.555355 |
| 1.19 | 608.306299 | 13.676205 | 4278.168400 | 88.925189 |
| 1.20 | 543.804582 | 23.721669 | 4239.762414 | 24.909073 |
| gotip | 568.891943 | 14.319373 | 4241.474395 | 25.904163 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 140.915673 | 14.690490 | 4811.826444 | 29.067535 |
| 1.19 | 133.067001 | 6.747823 | 4670.128516 | 34.905607 |
| 1.20 | 146.435469 | 8.602721 | 4700.974889 | 67.395977 |
| gotip | 145.589028 | 8.297745 | 5115.080201 | 59.329413 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 162.943858 | 12.001666 | 6968.655178 | 44.135559 |
| 1.19 | 135.832159 | 8.356388 | 6766.382886 | 34.341390 |
| 1.20 | 149.876895 | 6.909461 | 7527.312950 | 65.740010 |
| gotip | 138.885846 | 4.511916 | 7902.328798 | 99.532001 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 151.166961 | 8.223282 | 6196.007296 | 90.186203 |
| 1.19 | 131.841585 | 7.052037 | 10319.819916 | 94.173287 |
| 1.20 | 150.973414 | 10.932608 | 10655.411813 | 49.969725 |
| gotip | 144.938225 | 10.109369 | 6647.191504 | 98.637336 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 152.883908 | 14.533568 | 7377.743109 | 151.718009 |
| 1.19 | 140.938112 | 8.800738 | 9803.694520 | 383.903759 |
| 1.20 | 146.195191 | 10.308663 | 10349.008557 | 108.259559 |
| gotip | 153.070135 | 18.394617 | 10984.124227 | 213.707624 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 171.512356 | 23.086742 | 6043.546094 | 31.915083 |
| 1.19 | 142.230548 | 8.771253 | 6111.551651 | 24.598657 |
| 1.20 | 156.428005 | 7.536233 | 6138.203011 | 87.481925 |
| gotip | 156.600183 | 4.561679 | 6172.048706 | 45.846603 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 161.417218 | 13.744547 | 4291.862251 | 24.136645 |
| 1.19 | 142.254242 | 15.631682 | 478.236966 | 21.415793 |
| 1.20 | 147.435466 | 10.023269 | 457.082818 | 9.262519 |
| gotip | 153.666736 | 7.547628 | 459.410571 | 9.371994 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 166.971315 | 9.341075 | 4301.285168 | 538.851755 |
| 1.19 | 142.938464 | 5.665037 | 4600.535754 | 762.159027 |
| 1.20 | 152.346818 | 13.759479 | 6519.995784 | 808.201865 |
| gotip | 155.798131 | 15.390582 | 6873.158550 | 864.696183 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 159.684358 | 10.543378 | 5218.428514 | 26.016726 |
| 1.19 | 135.617925 | 8.129192 | 2633.169529 | 28.448521 |
| 1.20 | 144.753467 | 6.947198 | 2632.092402 | 25.002544 |
| gotip | 136.385819 | 6.961152 | 5202.784580 | 46.461902 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 198.181803 | 13.345129 | 1525.113202 | 22.833101 |
| 1.19 | 164.574407 | 11.415320 | 1566.269812 | 44.903017 |
| 1.20 | 175.292239 | 9.758776 | 1624.110046 | 52.177869 |
| gotip | 165.156944 | 6.964713 | 1548.661758 | 59.932362 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 205.649308 | 9.431185 | 3083.143408 | 25.684757 |
| 1.19 | 172.061712 | 9.884423 | 2951.551101 | 22.442742 |
| 1.20 | 186.660054 | 10.390697 | 2917.862029 | 38.609581 |
| gotip | 175.972102 | 5.554717 | 2801.320210 | 18.809352 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 245.593819 | 9.040662 | 7014.780961 | 40.191662 |
| 1.19 | 205.570136 | 25.744258 | 3869.430793 | 69.406445 |
| 1.20 | 206.552342 | 10.497080 | 3893.079456 | 64.467518 |
| gotip | 218.644231 | 14.447771 | 3334.120542 | 34.290468 |

![switch_case](./switch_case__725e73000e.png)

