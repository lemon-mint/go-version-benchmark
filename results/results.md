# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.6

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) CPU E5-2673 v3 @ 2.40GHz

Cores: 1

Mhz: 2397.221000

CacheSize: 30720

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) CPU E5-2673 v3 @ 2.40GHz

Cores: 1

Mhz: 2397.221000

CacheSize: 30720

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 563.225872 | 15.331610 | 4151.773270 | 31.079329 |
| 1.19 | 597.598540 | 12.804932 | 4137.684199 | 37.504778 |
| 1.20 | 509.303239 | 17.405073 | 4219.518012 | 41.158635 |
| gotip | 550.538988 | 14.496102 | 4257.153416 | 60.894271 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 149.236664 | 4.617837 | 4780.461479 | 77.266546 |
| 1.19 | 140.086477 | 7.917560 | 4814.966658 | 49.321765 |
| 1.20 | 151.753646 | 7.125159 | 4772.118772 | 47.661846 |
| gotip | 143.546277 | 7.731783 | 4886.809387 | 55.844929 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 159.703026 | 17.448061 | 6942.426509 | 47.304162 |
| 1.19 | 137.219120 | 7.319538 | 6857.005850 | 43.046976 |
| 1.20 | 150.353841 | 5.310820 | 7758.681769 | 114.826947 |
| gotip | 143.962899 | 9.487542 | 7791.280721 | 46.964283 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 167.936320 | 25.218064 | 6238.645567 | 93.540866 |
| 1.19 | 135.908266 | 6.343146 | 6381.325630 | 76.831314 |
| 1.20 | 153.188972 | 9.639389 | 6457.489851 | 63.563469 |
| gotip | 147.629098 | 3.614055 | 6398.583545 | 63.173894 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 158.922196 | 8.579495 | 7311.622451 | 58.904123 |
| 1.19 | 138.746059 | 11.598652 | 9508.071748 | 99.639893 |
| 1.20 | 149.005092 | 5.413734 | 10227.485007 | 54.377898 |
| gotip | 150.601184 | 5.859084 | 10385.918959 | 95.573609 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 176.929351 | 7.282707 | 5888.608736 | 43.060000 |
| 1.19 | 145.415435 | 4.685941 | 6021.978933 | 41.202984 |
| 1.20 | 158.511021 | 7.216188 | 6046.068713 | 55.992132 |
| gotip | 160.572554 | 6.801137 | 6236.088293 | 59.274143 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 166.287468 | 6.123527 | 4538.525229 | 44.819017 |
| 1.19 | 142.467720 | 5.805867 | 451.246574 | 11.239883 |
| 1.20 | 152.914941 | 6.949313 | 449.924934 | 6.827927 |
| gotip | 149.574312 | 5.906750 | 427.381833 | 9.970143 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 176.069515 | 9.787733 | 4692.825125 | 782.362646 |
| 1.19 | 148.822302 | 6.877547 | 5059.322073 | 1033.164676 |
| 1.20 | 168.539119 | 10.279857 | 6662.163251 | 998.595908 |
| gotip | 167.681555 | 7.873613 | 7397.505451 | 730.557311 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 151.846083 | 10.683277 | 2709.714255 | 60.714631 |
| 1.19 | 135.576057 | 6.001948 | 2698.989518 | 34.502332 |
| 1.20 | 147.176086 | 6.351669 | 2674.446381 | 41.170335 |
| gotip | 152.193619 | 10.553799 | 2773.856421 | 38.506960 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 196.356828 | 17.075230 | 1410.505903 | 10.365727 |
| 1.19 | 166.480413 | 13.815642 | 1443.398386 | 8.964214 |
| 1.20 | 170.671056 | 4.550759 | 1439.878253 | 23.334864 |
| gotip | 174.967782 | 6.862928 | 1381.572660 | 19.741014 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 209.409669 | 6.943072 | 3022.898433 | 43.302611 |
| 1.19 | 174.137101 | 11.423550 | 2853.808232 | 36.092222 |
| 1.20 | 180.404327 | 7.762342 | 2774.575438 | 36.440751 |
| gotip | 183.188335 | 6.660948 | 2761.992171 | 45.662832 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 237.306504 | 18.432035 | 6955.145502 | 50.610172 |
| 1.19 | 202.118992 | 16.738608 | 4099.207233 | 58.311159 |
| 1.20 | 216.847062 | 12.743803 | 3711.938838 | 35.900666 |
| gotip | 220.723069 | 8.621122 | 3478.404646 | 53.147091 |

![switch_case](./switch_case__725e73000e.png)

