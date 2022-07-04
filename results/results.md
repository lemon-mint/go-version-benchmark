# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.18.3

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) CPU E5-2673 v3 @ 2.40GHz

Cores: 1

Mhz: 2394.454000

CacheSize: 30720

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) CPU E5-2673 v3 @ 2.40GHz

Cores: 1

Mhz: 2394.454000

CacheSize: 30720

Microcode: 0xffffffff

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 152.769634 | 5.094989 | 3808.754565 | 16.019188 |
| 1.18 | 168.180763 | 7.587596 | 2799.520648 | 27.600790 |
| 1.19beta1 | 354.552025 | 6.205499 | 2831.775724 | 38.231295 |
| gotip | 192.345397 | 4.751234 | 2822.114920 | 22.112726 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 152.905013 | 3.902122 | 4887.253201 | 14.932316 |
| 1.18 | 164.778058 | 4.988308 | 4854.567630 | 24.539919 |
| 1.19beta1 | 350.173609 | 8.559356 | 4864.264884 | 16.923454 |
| gotip | 189.363870 | 2.486640 | 4886.786936 | 28.252427 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 159.196673 | 6.874878 | 7382.125745 | 90.454731 |
| 1.18 | 166.717578 | 10.876994 | 7167.515649 | 66.387325 |
| 1.19beta1 | 353.217884 | 11.802585 | 8927.651870 | 41.411523 |
| gotip | 196.534557 | 5.178131 | 8922.048277 | 74.795194 |

![MergeSort](./MergeSort__619024e898.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 202.097884 | 14.220042 | 1357.096288 | 10.490736 |
| 1.18 | 207.852733 | 17.185337 | 1407.558423 | 10.248117 |
| 1.19beta1 | 386.097058 | 6.240390 | 1427.380785 | 9.834049 |
| gotip | 231.911598 | 5.141619 | 1430.591436 | 11.959175 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 205.386473 | 4.756802 | 3089.407320 | 13.275654 |
| 1.18 | 213.542941 | 6.061014 | 3072.132673 | 24.009496 |
| 1.19beta1 | 389.928684 | 10.965728 | 2862.782732 | 12.908579 |
| gotip | 241.450090 | 4.957268 | 2870.999381 | 21.037060 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 179.083742 | 4.583882 | 6006.896722 | 17.156137 |
| 1.18 | 188.267437 | 4.054074 | 5997.530907 | 17.593515 |
| 1.19beta1 | 371.505243 | 4.675935 | 6103.139526 | 19.041876 |
| gotip | 209.489779 | 6.157213 | 6103.156625 | 28.436702 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 167.100168 | 7.463616 | 4663.615317 | 25.271601 |
| 1.18 | 173.954985 | 5.999051 | 4607.529471 | 24.110457 |
| 1.19beta1 | 352.702849 | 6.246018 | 460.856449 | 4.610951 |
| gotip | 191.418161 | 8.053209 | 454.662677 | 4.305545 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 184.574183 | 7.300493 | 6435.295292 | 686.692999 |
| 1.18 | 196.921003 | 110.752796 | 4462.308826 | 719.315731 |
| 1.19beta1 | 373.225054 | 496.443553 | 4474.258008 | 165.885350 |
| gotip | 215.291153 | 11.410997 | 4702.194360 | 944.262159 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 248.881595 | 9.306435 | 7062.914213 | 16.479319 |
| 1.18 | 245.952950 | 24.371101 | 7050.487434 | 28.432784 |
| 1.19beta1 | 425.465909 | 13.617057 | 4097.545872 | 15.095458 |
| gotip | 280.974545 | 7.327745 | 4093.620334 | 13.500361 |

![switch_case](./switch_case__725e73000e.png)

