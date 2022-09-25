# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) CPU E5-2673 v3 @ 2.40GHz

Cores: 1

Mhz: 2394.455000

CacheSize: 30720

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) CPU E5-2673 v3 @ 2.40GHz

Cores: 1

Mhz: 2394.455000

CacheSize: 30720

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 544.474647 | 14.960080 | 4373.644705 | 83.832461 |
| 1.18 | 565.562910 | 21.484682 | 3961.022550 | 70.391931 |
| 1.19 | 632.765821 | 26.195018 | 3918.856753 | 37.292696 |
| gotip | 638.444302 | 14.830121 | 4078.547624 | 50.992344 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 151.080364 | 4.307144 | 4712.095288 | 31.656643 |
| 1.18 | 161.511593 | 12.272339 | 4727.564560 | 20.732806 |
| 1.19 | 178.140462 | 8.049377 | 4735.278455 | 70.768912 |
| gotip | 174.039065 | 5.525472 | 4600.862527 | 74.667537 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 156.776965 | 7.031849 | 6556.771440 | 72.043367 |
| 1.18 | 162.367141 | 8.559007 | 6692.877282 | 36.963282 |
| 1.19 | 171.470056 | 7.922518 | 6601.424481 | 51.383662 |
| gotip | 188.060790 | 8.007114 | 7536.325320 | 54.287541 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 162.043105 | 5.744206 | 5833.396419 | 285.291868 |
| 1.18 | 159.464549 | 8.047042 | 5450.186218 | 40.470229 |
| 1.19 | 166.588931 | 5.720251 | 5582.183080 | 91.840275 |
| gotip | 168.269643 | 3.765479 | 5733.160160 | 95.171632 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 158.439694 | 6.972586 | 7223.002170 | 163.666778 |
| 1.18 | 161.735010 | 5.549351 | 7008.950033 | 175.174822 |
| 1.19 | 181.074581 | 8.537593 | 9815.609667 | 397.485803 |
| gotip | 192.381153 | 4.347584 | 10413.151198 | 356.852070 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 182.338704 | 4.789220 | 5830.697347 | 45.847515 |
| 1.18 | 190.320555 | 2.866312 | 5848.114613 | 48.605637 |
| 1.19 | 198.289222 | 5.369594 | 5957.013893 | 122.343190 |
| gotip | 179.945369 | 5.563051 | 5796.602233 | 63.147218 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 163.389381 | 9.611849 | 4433.414054 | 63.698216 |
| 1.18 | 168.532115 | 3.969349 | 4360.246305 | 43.387898 |
| 1.19 | 176.602242 | 3.557293 | 421.729828 | 8.030607 |
| gotip | 185.876088 | 5.572522 | 446.151219 | 9.012029 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 184.587474 | 9.776270 | 6072.424854 | 821.081107 |
| 1.18 | 191.409308 | 107.270915 | 4490.871382 | 957.864825 |
| 1.19 | 198.363893 | 124.137686 | 4568.970721 | 485.928535 |
| gotip | 205.970431 | 12.540117 | 5411.410233 | 378.319127 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 147.171582 | 4.383622 | 3656.070193 | 41.232878 |
| 1.18 | 163.536667 | 5.745639 | 2639.695037 | 50.855000 |
| 1.19 | 168.106090 | 3.766834 | 2686.577224 | 40.541143 |
| gotip | 174.028320 | 8.963284 | 2612.818586 | 61.970777 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 201.782798 | 6.851202 | 1334.416333 | 25.472065 |
| 1.18 | 217.133742 | 20.157724 | 1456.242815 | 17.594142 |
| 1.19 | 220.158740 | 14.909739 | 1406.898686 | 20.228167 |
| gotip | 211.978253 | 7.160327 | 1263.013071 | 14.435279 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 211.357913 | 11.690388 | 3042.985085 | 26.415579 |
| 1.18 | 226.217123 | 9.684751 | 3000.313766 | 55.902096 |
| 1.19 | 231.572336 | 6.000156 | 2802.241957 | 22.362831 |
| gotip | 229.094495 | 14.034673 | 2766.669473 | 38.410351 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 245.628350 | 5.320355 | 6674.191376 | 80.801285 |
| 1.18 | 246.886503 | 22.681708 | 6869.889430 | 190.366800 |
| 1.19 | 252.089374 | 24.717054 | 3991.164102 | 36.325404 |
| gotip | 275.408779 | 8.373347 | 3648.714906 | 27.026709 |

![switch_case](./switch_case__725e73000e.png)

