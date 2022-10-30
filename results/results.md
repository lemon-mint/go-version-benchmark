# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.2

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.435000

CacheSize: 49152

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.435000

CacheSize: 49152

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 449.593925 | 11.802017 | 4053.880826 | 8.807523 |
| 1.18 | 453.463027 | 12.896664 | 3640.521238 | 89.268408 |
| 1.19 | 505.412841 | 7.597174 | 3638.554827 | 6.935743 |
| gotip | 3814.871576 | 21.382445 | 3668.447069 | 9.522012 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 120.994900 | 10.629759 | 4862.824142 | 1.311266 |
| 1.18 | 123.363755 | 1.327485 | 4865.636493 | 0.818501 |
| 1.19 | 138.663916 | 6.825462 | 4862.583445 | 0.662964 |
| gotip | 2730.304343 | 16.815931 | 4862.729989 | 0.273228 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 129.063701 | 8.266565 | 5217.432754 | 55.998883 |
| 1.18 | 132.543376 | 2.251869 | 5167.646045 | 35.252967 |
| 1.19 | 137.842676 | 3.230840 | 5192.544174 | 14.981483 |
| gotip | 2745.307653 | 22.474855 | 6291.537102 | 28.750710 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 131.389505 | 5.868576 | 9643.342205 | 21.785247 |
| 1.18 | 137.358726 | 2.631420 | 5563.982211 | 15.022295 |
| 1.19 | 141.100825 | 3.320951 | 9699.868914 | 25.537292 |
| gotip | 2851.135870 | 16.679422 | 9762.426790 | 35.179861 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 125.168666 | 4.378741 | 5876.185196 | 53.729529 |
| 1.18 | 131.268517 | 2.020651 | 5560.541414 | 30.337668 |
| 1.19 | 143.109176 | 2.100324 | 7093.698736 | 56.121371 |
| gotip | 2741.733775 | 11.727736 | 7714.705269 | 32.691974 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 141.255485 | 4.196145 | 6078.228369 | 3.001937 |
| 1.18 | 146.903594 | 2.448957 | 5996.933524 | 14.645058 |
| 1.19 | 147.666463 | 1.748772 | 5926.852050 | 2.610589 |
| gotip | 3365.248789 | 29.932969 | 5983.498759 | 2.511535 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 129.547463 | 2.550073 | 3648.811638 | 7.163245 |
| 1.18 | 138.654039 | 3.741912 | 3563.726326 | 4.585902 |
| 1.19 | 139.775209 | 2.124946 | 360.932814 | 2.574392 |
| gotip | 3071.899952 | 18.798223 | 398.878068 | 1.979128 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 146.377766 | 57.651354 | 4806.783132 | 558.616664 |
| 1.18 | 153.304838 | 56.810546 | 5194.557134 | 896.027876 |
| 1.19 | 154.071639 | 78.801373 | 4779.955804 | 361.591100 |
| gotip | 3832.001017 | 87.591491 | 5019.748610 | 257.244946 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 122.058369 | 3.339438 | 3663.199716 | 8.851480 |
| 1.18 | 128.302175 | 2.527882 | 5216.763654 | 1.967475 |
| 1.19 | 138.996440 | 2.399493 | 2705.141209 | 0.321706 |
| gotip | 2740.185951 | 22.691477 | 2704.952501 | 0.398898 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 154.068642 | 12.798310 | 1083.490528 | 2.409410 |
| 1.18 | 160.463252 | 9.684995 | 1131.939845 | 8.385845 |
| 1.19 | 162.594766 | 9.935870 | 1164.577897 | 8.480419 |
| gotip | 4182.328831 | 33.884608 | 1131.062393 | 5.713482 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 162.064771 | 8.413634 | 2382.457460 | 22.612785 |
| 1.18 | 166.429150 | 2.325731 | 2380.886454 | 39.430958 |
| 1.19 | 168.124222 | 1.731658 | 2250.898232 | 18.155144 |
| gotip | 4198.440432 | 26.277451 | 2257.755776 | 27.376862 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 188.829103 | 13.062482 | 6052.003984 | 59.721453 |
| 1.18 | 190.986370 | 11.929793 | 5958.677246 | 25.519581 |
| 1.19 | 197.366436 | 12.306833 | 3552.654201 | 4.705491 |
| gotip | 5565.308624 | 34.437440 | 3340.239317 | 5.178827 |

![switch_case](./switch_case__725e73000e.png)

