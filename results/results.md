# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.10

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz

Cores: 1

Mhz: 2593.905000

CacheSize: 36608

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz

Cores: 1

Mhz: 2593.905000

CacheSize: 36608

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 478.212343 | 19.666503 | 3517.524403 | 2.216340 |
| 1.19 | 504.944242 | 6.282071 | 3542.577144 | 1.686432 |
| 1.20 | 433.002609 | 5.647959 | 3554.368756 | 1.970503 |
| gotip | 474.414729 | 5.656117 | 3534.452063 | 3.922543 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 121.829796 | 3.557348 | 4012.928917 | 1.601287 |
| 1.19 | 113.676796 | 9.437526 | 3904.049533 | 3.094333 |
| 1.20 | 118.215672 | 4.092624 | 3908.455230 | 10.661480 |
| gotip | 118.442553 | 3.852758 | 4245.451598 | 1.214026 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 130.878865 | 3.976556 | 5769.778640 | 28.207492 |
| 1.19 | 116.130317 | 5.005747 | 5576.127950 | 25.194370 |
| 1.20 | 119.432794 | 3.514832 | 6236.855955 | 51.767796 |
| gotip | 119.327332 | 3.502796 | 6585.803419 | 29.327262 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 136.410170 | 4.673607 | 5364.267784 | 38.092085 |
| 1.19 | 118.832940 | 7.265803 | 9089.601096 | 27.825270 |
| 1.20 | 123.086960 | 4.235549 | 9104.294125 | 34.767093 |
| gotip | 120.689853 | 5.050006 | 5539.739550 | 74.535924 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 125.398529 | 3.369237 | 6052.389108 | 41.185833 |
| 1.19 | 114.399256 | 12.358387 | 7697.386153 | 55.119264 |
| 1.20 | 123.183922 | 2.110439 | 8823.306053 | 41.648398 |
| gotip | 122.381713 | 3.281457 | 9140.144616 | 45.852736 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 143.520481 | 5.312232 | 5029.463816 | 1.996974 |
| 1.19 | 124.299471 | 6.269869 | 5085.244881 | 3.558871 |
| 1.20 | 131.538580 | 3.749068 | 5082.276299 | 1.771320 |
| gotip | 137.064231 | 4.482275 | 5134.337032 | 1.570154 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 134.047296 | 4.904948 | 3576.024720 | 1.612817 |
| 1.19 | 117.993932 | 12.044247 | 370.357313 | 0.955198 |
| 1.20 | 122.692411 | 4.223974 | 369.815458 | 0.603869 |
| gotip | 121.458678 | 2.778992 | 374.731020 | 0.780589 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 146.701129 | 6.717542 | 4831.970885 | 565.806245 |
| 1.19 | 131.740255 | 26.727841 | 4912.528323 | 380.827657 |
| 1.20 | 132.702299 | 3.304475 | 5477.995448 | 184.292077 |
| gotip | 135.892201 | 5.190033 | 5698.105812 | 270.024694 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 123.766786 | 3.699421 | 4348.612319 | 1.678795 |
| 1.19 | 113.297514 | 2.894457 | 2191.114293 | 1.296567 |
| 1.20 | 121.622468 | 3.545637 | 2191.495184 | 0.605280 |
| gotip | 117.476960 | 3.570960 | 4348.678387 | 1.445136 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 158.665824 | 16.894994 | 1198.478269 | 10.258770 |
| 1.19 | 142.523855 | 6.198718 | 1212.284136 | 5.373931 |
| 1.20 | 148.021330 | 5.080745 | 1200.130310 | 11.232292 |
| gotip | 143.291399 | 3.556953 | 1167.284061 | 5.719905 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 167.452842 | 2.903367 | 2411.083410 | 7.837066 |
| 1.19 | 147.021569 | 2.952782 | 2290.293576 | 13.436083 |
| 1.20 | 152.443458 | 4.089922 | 2258.139711 | 4.776184 |
| gotip | 149.522174 | 3.080629 | 2193.367241 | 6.194443 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 194.373986 | 5.282266 | 5839.070847 | 2.735961 |
| 1.19 | 166.046126 | 7.543584 | 3219.942116 | 2.223243 |
| 1.20 | 174.178705 | 4.762693 | 3219.890797 | 2.650671 |
| gotip | 170.413057 | 4.170525 | 2771.817154 | 2.312653 |

![switch_case](./switch_case__725e73000e.png)

