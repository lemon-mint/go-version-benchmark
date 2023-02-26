# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.6

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.438000

CacheSize: 49152

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.438000

CacheSize: 49152

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 443.968875 | 4.607281 | 3615.835445 | 1.798627 |
| 1.19 | 460.955356 | 4.665028 | 3614.252039 | 21.689167 |
| 1.20 | 397.274045 | 1.925243 | 3641.218518 | 2.781771 |
| gotip | 430.912836 | 3.052753 | 3599.576034 | 3.887352 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 113.116664 | 7.776000 | 4795.286081 | 2.040548 |
| 1.19 | 103.770678 | 4.243801 | 4791.405066 | 1.128046 |
| 1.20 | 107.450581 | 2.748390 | 4790.709405 | 1.291380 |
| gotip | 109.070103 | 1.769400 | 5204.272918 | 1.486070 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 119.708674 | 2.664606 | 5165.258725 | 12.288604 |
| 1.19 | 103.882694 | 9.075974 | 5191.059297 | 269.919388 |
| 1.20 | 108.522038 | 2.741248 | 6113.615624 | 25.291384 |
| gotip | 109.231347 | 2.798495 | 6262.330193 | 18.806809 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 124.602111 | 3.113841 | 5530.720179 | 22.531923 |
| 1.19 | 107.927268 | 10.669770 | 9663.031109 | 20.817323 |
| 1.20 | 110.791858 | 2.648340 | 9727.895829 | 16.867586 |
| gotip | 111.874550 | 2.329656 | 5687.535300 | 17.431551 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 115.367466 | 1.548788 | 5556.197520 | 26.816704 |
| 1.19 | 104.288928 | 13.197473 | 7090.066951 | 46.818402 |
| 1.20 | 112.695361 | 1.359293 | 7654.695058 | 35.493399 |
| gotip | 113.313129 | 4.062053 | 7791.680034 | 44.819389 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 133.153003 | 3.618666 | 5974.206567 | 5.485479 |
| 1.19 | 110.677719 | 2.641719 | 5925.585815 | 2.037539 |
| 1.20 | 119.735187 | 2.373438 | 6007.115283 | 9.853884 |
| gotip | 121.333610 | 2.751586 | 5997.393931 | 1.090509 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 124.731516 | 16.138551 | 3542.216347 | 2.920521 |
| 1.19 | 104.454691 | 4.393971 | 362.686351 | 2.933565 |
| 1.20 | 112.235158 | 3.920096 | 361.624587 | 0.671428 |
| gotip | 112.191607 | 3.953847 | 363.215065 | 0.508178 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 134.630927 | 2.422444 | 5304.141245 | 815.530563 |
| 1.19 | 117.693737 | 15.103233 | 5143.244462 | 685.144454 |
| 1.20 | 120.811750 | 5.287728 | 5232.196615 | 215.511570 |
| gotip | 123.403325 | 4.199521 | 5197.833011 | 343.466975 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 114.033336 | 12.411004 | 5220.753546 | 1.332152 |
| 1.19 | 103.335228 | 3.373770 | 2704.726633 | 0.596695 |
| 1.20 | 107.063807 | 1.513042 | 2705.019380 | 0.839643 |
| gotip | 109.501586 | 2.337041 | 2704.628081 | 0.768543 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 145.645105 | 6.659653 | 1105.225373 | 2.140627 |
| 1.19 | 129.501650 | 4.210401 | 1136.929361 | 7.882950 |
| 1.20 | 130.140720 | 2.493847 | 1095.031846 | 12.318504 |
| gotip | 130.827580 | 3.326474 | 1071.462417 | 9.402654 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 156.931015 | 3.987933 | 2329.206539 | 23.913296 |
| 1.19 | 131.485025 | 1.862211 | 2190.114714 | 27.640779 |
| 1.20 | 136.668886 | 4.128708 | 2161.521689 | 16.667281 |
| gotip | 137.907732 | 2.496697 | 2146.103947 | 20.024775 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 177.351935 | 17.945617 | 5944.134527 | 25.053221 |
| 1.19 | 154.205118 | 6.571846 | 3540.694035 | 3.798130 |
| 1.20 | 159.287624 | 4.237926 | 3521.052872 | 4.625429 |
| gotip | 158.245883 | 4.558120 | 2823.260565 | 2.885928 |

![switch_case](./switch_case__725e73000e.png)

