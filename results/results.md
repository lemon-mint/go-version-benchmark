# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.18.3

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

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 125.068177 | 2.838065 | 4795.909855 | 1.210382 |
| 1.18 | 137.986398 | 3.920857 | 4799.783905 | 1.246505 |
| 1.19beta1 | 284.149111 | 2.358779 | 4799.464163 | 0.365003 |
| gotip | 155.635180 | 5.997427 | 4795.011216 | 1.005233 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 173.417170 | 2.449661 | 2359.017092 | 18.035459 |
| 1.18 | 180.466321 | 3.466851 | 2353.600590 | 16.551511 |
| 1.19beta1 | 323.257260 | 3.617312 | 2199.637982 | 19.945891 |
| gotip | 194.932976 | 3.962434 | 2196.503429 | 23.426557 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 140.254125 | 2.373813 | 6080.190922 | 7.263415 |
| 1.18 | 158.890505 | 8.244108 | 5980.333444 | 5.193746 |
| 1.19beta1 | 301.093734 | 4.259693 | 5943.813103 | 10.283197 |
| gotip | 172.858020 | 4.076289 | 5935.083094 | 8.097334 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 156.075010 | 4.205225 | 4703.603031 | 566.299395 |
| 1.18 | 164.952137 | 86.289950 | 5078.412025 | 651.140038 |
| 1.19beta1 | 307.128731 | 545.527030 | 5064.179481 | 597.085655 |
| gotip | 183.992158 | 3.894670 | 5307.289929 | 495.338028 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 132.940003 | 2.941082 | 6583.476492 | 79.721776 |
| 1.18 | 140.036421 | 6.856763 | 6354.435693 | 178.453811 |
| 1.19beta1 | 290.558716 | 7.654146 | 7791.740028 | 51.243325 |
| gotip | 162.029278 | 6.534249 | 7830.564127 | 61.442884 |

![MergeSort](./MergeSort__619024e898.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 166.408430 | 10.264985 | 1125.023351 | 5.748273 |
| 1.18 | 176.140146 | 11.599717 | 1141.036787 | 11.805026 |
| 1.19beta1 | 316.274137 | 4.847446 | 1149.313431 | 3.464486 |
| gotip | 188.924679 | 5.402669 | 1158.979092 | 18.459495 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 137.512453 | 2.306301 | 3650.068880 | 22.718368 |
| 1.18 | 145.773920 | 2.309998 | 3545.644677 | 2.002888 |
| 1.19beta1 | 288.397813 | 3.604086 | 396.876941 | 31.410006 |
| gotip | 159.252066 | 2.941302 | 361.364945 | 34.432544 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 204.656284 | 9.390758 | 5926.432596 | 16.537688 |
| 1.18 | 209.929968 | 17.287355 | 5913.953936 | 31.494736 |
| 1.19beta1 | 353.570769 | 10.010858 | 3418.865117 | 1.667413 |
| gotip | 236.353565 | 4.090678 | 3568.619340 | 1.815157 |

![switch_case](./switch_case__725e73000e.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 131.654659 | 3.507225 | 3883.079068 | 34.950368 |
| 1.18 | 139.332235 | 2.003572 | 5219.487404 | 1.336616 |
| 1.19beta1 | 286.159951 | 5.821695 | 5219.221849 | 1.048697 |
| gotip | 155.695146 | 2.485318 | 2707.949250 | 0.456128 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

