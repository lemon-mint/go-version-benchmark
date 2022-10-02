# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.1

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
| 1.17 | 458.143083 | 33.847130 | 4019.129548 | 2.387172 |
| 1.18 | 463.696691 | 10.919485 | 3617.636313 | 298.634929 |
| 1.19 | 521.002616 | 9.663585 | 3616.188765 | 469.567514 |
| gotip | 524.756013 | 6.844128 | 3617.751278 | 1.939220 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 123.270724 | 5.774985 | 4791.960896 | 2.575524 |
| 1.18 | 127.859887 | 10.448103 | 4795.901986 | 72.456629 |
| 1.19 | 138.978437 | 1.622906 | 4791.581462 | 1.952111 |
| gotip | 144.261743 | 3.271134 | 4795.997765 | 1.789125 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 129.707831 | 3.469667 | 5224.179160 | 58.146553 |
| 1.18 | 137.100156 | 5.202916 | 5167.600646 | 330.595037 |
| 1.19 | 139.718235 | 1.958524 | 5206.189719 | 260.422920 |
| gotip | 144.021723 | 2.629978 | 6226.754132 | 19.976098 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 134.320617 | 2.237170 | 9967.355933 | 59.424204 |
| 1.18 | 139.567759 | 4.112252 | 5811.556091 | 167.691766 |
| 1.19 | 140.809539 | 3.627405 | 10125.056870 | 53.920036 |
| gotip | 147.692996 | 3.877715 | 5949.679546 | 47.444376 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 126.791638 | 3.963875 | 5843.624837 | 37.790321 |
| 1.18 | 131.444142 | 2.324956 | 5565.107177 | 25.403638 |
| 1.19 | 142.009312 | 6.508090 | 7096.780489 | 82.962099 |
| gotip | 152.436670 | 4.578707 | 7857.187612 | 45.536414 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 143.417849 | 2.569190 | 6063.832647 | 3.008056 |
| 1.18 | 152.012228 | 4.487564 | 5974.200853 | 9.896262 |
| 1.19 | 149.758344 | 2.616617 | 5920.676230 | 2.299453 |
| gotip | 155.524794 | 3.975048 | 5954.775728 | 12.654346 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 130.067622 | 3.028200 | 3639.274785 | 6.915320 |
| 1.18 | 137.216365 | 2.458388 | 3544.462146 | 450.189181 |
| 1.19 | 142.956038 | 4.096835 | 360.740307 | 3.154804 |
| gotip | 145.467230 | 1.900483 | 400.686889 | 2.669323 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 147.766756 | 4.517246 | 4631.105358 | 656.805662 |
| 1.18 | 152.967606 | 68.727404 | 5031.015114 | 1015.072020 |
| 1.19 | 157.860278 | 90.706394 | 4750.069510 | 300.093840 |
| gotip | 162.853285 | 44.432996 | 5095.316483 | 324.656560 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 121.816607 | 56.510197 | 3662.248495 | 17.499634 |
| 1.18 | 127.476056 | 2.808580 | 5220.388015 | 0.612481 |
| 1.19 | 137.332725 | 2.693249 | 2705.648982 | 1.035722 |
| gotip | 144.263025 | 2.182957 | 5220.445215 | 1.889145 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 158.359419 | 10.276153 | 1064.675107 | 3.239921 |
| 1.18 | 160.279951 | 9.832980 | 1105.726608 | 4.762709 |
| 1.19 | 165.714046 | 9.709889 | 1138.292990 | 4.370843 |
| gotip | 170.153266 | 2.974903 | 1054.799547 | 9.474928 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 163.016378 | 2.589237 | 2335.727674 | 52.214575 |
| 1.18 | 169.035225 | 3.430936 | 2330.335225 | 50.453871 |
| 1.19 | 172.043002 | 2.878069 | 2229.673728 | 48.108940 |
| gotip | 176.337362 | 3.098591 | 2171.959763 | 23.082881 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 191.530152 | 15.007239 | 5964.073178 | 43.163538 |
| 1.18 | 196.543731 | 14.444029 | 5938.502854 | 26.171613 |
| 1.19 | 196.368744 | 13.538476 | 3560.817481 | 2.583115 |
| gotip | 205.783935 | 6.730569 | 3313.777989 | 2.789576 |

![switch_case](./switch_case__725e73000e.png)

