# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.8

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz

Cores: 1

Mhz: 2593.908000

CacheSize: 36608

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz

Cores: 1

Mhz: 2593.908000

CacheSize: 36608

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 461.555422 | 3.602988 | 3516.700217 | 2.463000 |
| 1.19 | 482.888013 | 11.123842 | 3541.831673 | 3.157122 |
| 1.20 | 418.819022 | 5.047263 | 3553.959666 | 1.730366 |
| gotip | 464.886379 | 9.164642 | 3567.980941 | 3.586629 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 119.672538 | 3.457246 | 4013.141973 | 1.507622 |
| 1.19 | 106.587413 | 3.713838 | 3901.948992 | 1.726935 |
| 1.20 | 120.938579 | 6.299500 | 3940.540424 | 7.406733 |
| gotip | 122.616729 | 3.963420 | 4244.704321 | 2.613153 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 127.515929 | 3.224061 | 5741.736867 | 30.340265 |
| 1.19 | 112.554421 | 4.602261 | 5582.336353 | 25.925167 |
| 1.20 | 115.999432 | 3.634574 | 6221.818711 | 11.686960 |
| gotip | 120.919230 | 3.487406 | 6557.064903 | 29.658270 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 132.595458 | 3.060371 | 5362.215939 | 54.134286 |
| 1.19 | 112.137215 | 2.904136 | 8920.822909 | 53.951908 |
| 1.20 | 120.562128 | 2.056372 | 8830.981840 | 102.220187 |
| gotip | 122.431423 | 4.443851 | 5432.045832 | 30.057281 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 122.523016 | 4.042122 | 5937.881501 | 33.702022 |
| 1.19 | 110.022033 | 6.018216 | 7586.738060 | 54.519118 |
| 1.20 | 116.515350 | 3.252847 | 8644.813320 | 65.790946 |
| gotip | 125.954962 | 3.358543 | 9108.449664 | 171.174178 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 139.946139 | 5.761592 | 5053.272348 | 3.790664 |
| 1.19 | 115.760152 | 2.926208 | 5083.089512 | 3.588795 |
| 1.20 | 128.216176 | 4.674762 | 5079.659854 | 1.415433 |
| gotip | 137.597872 | 5.714454 | 5196.842705 | 1.597706 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 129.304287 | 4.384418 | 3575.002578 | 1.519435 |
| 1.19 | 112.734632 | 8.730150 | 370.164559 | 0.720915 |
| 1.20 | 117.679132 | 3.394023 | 369.688929 | 0.491681 |
| gotip | 119.927318 | 5.544551 | 374.544400 | 0.934540 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 143.939639 | 3.638713 | 4694.452412 | 517.588274 |
| 1.19 | 123.517928 | 6.290037 | 5128.975961 | 498.547851 |
| 1.20 | 131.101499 | 2.870583 | 5385.888575 | 471.373966 |
| gotip | 137.639286 | 4.026594 | 5256.841203 | 299.229924 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 120.754794 | 11.447447 | 4347.911253 | 2.058157 |
| 1.19 | 108.614824 | 2.505554 | 2191.474330 | 0.632040 |
| 1.20 | 118.354144 | 3.325169 | 2191.252930 | 0.775388 |
| gotip | 125.162873 | 4.770231 | 4348.532915 | 5.842647 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 156.306669 | 3.443108 | 1152.342051 | 6.458564 |
| 1.19 | 134.067169 | 3.843611 | 1164.703487 | 9.219461 |
| 1.20 | 139.979626 | 4.291211 | 1162.129534 | 11.367301 |
| gotip | 145.490942 | 6.679684 | 1126.272397 | 7.537390 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 162.782671 | 17.926660 | 2365.595822 | 14.470858 |
| 1.19 | 138.112178 | 4.067761 | 2245.180772 | 9.526896 |
| 1.20 | 144.794221 | 3.149164 | 2231.537232 | 17.753902 |
| gotip | 149.304965 | 3.517584 | 2113.405116 | 9.453449 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 186.118048 | 6.287233 | 5724.728815 | 4.034810 |
| 1.19 | 157.409459 | 11.544441 | 3223.472370 | 1.647435 |
| 1.20 | 169.097130 | 4.495086 | 3224.146874 | 1.596714 |
| gotip | 171.610195 | 2.812493 | 2688.147706 | 3.674208 |

![switch_case](./switch_case__725e73000e.png)

