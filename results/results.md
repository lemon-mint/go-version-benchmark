# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.3

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.437000

CacheSize: 49152

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.437000

CacheSize: 49152

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 474.924015 | 41.831372 | 4020.732617 | 8.752731 |
| 1.18 | 477.107463 | 8.758290 | 3621.196120 | 3.092047 |
| 1.19 | 532.859338 | 12.241738 | 3616.003496 | 6.674207 |
| gotip | 3943.807726 | 65.562194 | 3635.980372 | 1.924597 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 130.624685 | 4.940801 | 4795.700533 | 7.012390 |
| 1.18 | 134.465405 | 11.869702 | 4800.308902 | 5.769091 |
| 1.19 | 144.874786 | 3.264265 | 4795.662696 | 3.226776 |
| gotip | 2800.085168 | 26.411140 | 4795.406553 | 2.709633 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 138.810723 | 2.878282 | 5214.039664 | 66.957810 |
| 1.18 | 141.212336 | 3.644331 | 5178.784676 | 229.171141 |
| 1.19 | 142.966207 | 2.909503 | 5194.363691 | 26.863437 |
| gotip | 2807.287610 | 25.279581 | 6118.943414 | 27.666037 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 139.773728 | 2.864753 | 10421.496962 | 23.572499 |
| 1.18 | 145.645739 | 4.683536 | 6315.644115 | 108.933618 |
| 1.19 | 147.076313 | 3.957665 | 10493.612537 | 72.733139 |
| gotip | 2895.674808 | 33.048088 | 10616.048577 | 81.101110 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 133.159686 | 4.764751 | 6932.376853 | 108.800156 |
| 1.18 | 140.405696 | 3.172937 | 5781.086031 | 184.802184 |
| 1.19 | 145.957967 | 2.339650 | 7318.241662 | 64.384449 |
| gotip | 2714.450817 | 21.460367 | 8114.849101 | 72.306793 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 150.290216 | 43.897872 | 6072.758268 | 5.526459 |
| 1.18 | 154.654251 | 5.228117 | 5981.564258 | 6.458291 |
| 1.19 | 155.461401 | 3.448203 | 5916.069111 | 10.865438 |
| gotip | 3260.161882 | 32.591082 | 5979.275067 | 19.105047 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 137.229007 | 4.035816 | 3644.032498 | 28.042126 |
| 1.18 | 139.337292 | 2.766773 | 3544.410851 | 1.922518 |
| 1.19 | 143.769567 | 2.072839 | 360.786206 | 0.421050 |
| gotip | 3065.141076 | 26.753574 | 361.848100 | 33.806390 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 157.607521 | 67.518074 | 4679.412061 | 328.514227 |
| 1.18 | 155.438420 | 80.904500 | 4957.668737 | 676.948896 |
| 1.19 | 161.334864 | 91.467299 | 5146.108040 | 561.109862 |
| gotip | 3788.754336 | 59.911269 | 5147.049393 | 340.127249 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 127.343933 | 78.808105 | 3732.978934 | 21.202717 |
| 1.18 | 139.747222 | 5.119202 | 5217.617215 | 0.679391 |
| 1.19 | 141.306757 | 7.479718 | 2707.395903 | 0.856298 |
| gotip | 2788.823520 | 42.067302 | 2706.862933 | 0.774617 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 159.883555 | 15.983432 | 1067.735430 | 4.562536 |
| 1.18 | 164.511907 | 12.222751 | 1111.001352 | 14.263377 |
| 1.19 | 173.251016 | 13.041241 | 1149.226277 | 15.598577 |
| gotip | 3904.457234 | 36.971380 | 1094.492712 | 11.069501 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 168.754280 | 5.113916 | 2352.339310 | 25.191042 |
| 1.18 | 173.382389 | 3.381930 | 2338.404855 | 21.644746 |
| 1.19 | 173.891286 | 3.704778 | 2200.300791 | 29.431863 |
| gotip | 3895.082294 | 38.454451 | 2181.536466 | 202.989902 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 203.358109 | 14.159432 | 5954.748923 | 31.206542 |
| 1.18 | 199.542343 | 15.675928 | 5941.235581 | 31.329664 |
| 1.19 | 205.552688 | 13.592069 | 3552.167968 | 6.264865 |
| gotip | 5111.237986 | 24.410056 | 3505.873139 | 4.800230 |

![switch_case](./switch_case__725e73000e.png)

