# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.7

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.451000

CacheSize: 49152

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.451000

CacheSize: 49152

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 467.795793 | 11.235349 | 3622.185693 | 2.343864 |
| 1.19 | 482.776207 | 6.975737 | 3618.015510 | 5.510259 |
| 1.20 | 419.072841 | 4.859209 | 3630.322382 | 2.112688 |
| gotip | 462.813020 | 33.821322 | 3559.369567 | 2.591850 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 115.340554 | 10.346308 | 4799.148993 | 1.214862 |
| 1.19 | 110.265397 | 5.741864 | 4794.871017 | 1.499817 |
| 1.20 | 113.025709 | 4.072258 | 4794.939711 | 1.081862 |
| gotip | 117.625620 | 5.641022 | 5208.579823 | 1.340094 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 127.108546 | 2.385632 | 5172.135852 | 23.731754 |
| 1.19 | 110.990371 | 2.931840 | 5211.727825 | 25.177615 |
| 1.20 | 115.508769 | 3.641731 | 6098.019857 | 249.521415 |
| gotip | 118.893223 | 2.486693 | 6480.918111 | 25.183950 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 128.305342 | 2.301683 | 6380.478636 | 95.261485 |
| 1.19 | 114.945497 | 2.340144 | 10517.832874 | 54.737906 |
| 1.20 | 118.399464 | 4.174469 | 10641.837854 | 88.344843 |
| gotip | 122.804819 | 4.582902 | 6607.436793 | 98.424631 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 122.793015 | 3.711480 | 5867.491959 | 85.261011 |
| 1.19 | 110.504500 | 4.024284 | 7488.707315 | 172.405417 |
| 1.20 | 118.745478 | 3.142368 | 8052.495748 | 116.234672 |
| gotip | 119.747773 | 3.430840 | 8148.255540 | 98.045809 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 142.191678 | 9.044943 | 5978.393450 | 6.654498 |
| 1.19 | 121.847147 | 3.948649 | 5925.194777 | 2.018900 |
| 1.20 | 127.319694 | 3.161913 | 6010.957351 | 5.311809 |
| gotip | 131.638427 | 5.495141 | 6159.909810 | 2.768477 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 130.568122 | 4.226248 | 3545.256826 | 2.444428 |
| 1.19 | 112.926658 | 3.119393 | 361.258597 | 0.434483 |
| 1.20 | 120.659893 | 2.937731 | 361.469548 | 0.506342 |
| gotip | 123.393278 | 3.832993 | 375.335643 | 0.212707 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 146.269194 | 4.386853 | 4812.238670 | 959.690229 |
| 1.19 | 128.520333 | 13.979936 | 5009.679276 | 549.957043 |
| 1.20 | 131.184413 | 3.761023 | 5216.043228 | 283.625531 |
| gotip | 130.983618 | 4.077971 | 5127.526415 | 157.372606 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 119.604232 | 3.797316 | 5218.228551 | 1.500210 |
| 1.19 | 107.431624 | 14.636715 | 2708.150129 | 2.415477 |
| 1.20 | 117.740640 | 3.708166 | 2707.338653 | 3.564840 |
| gotip | 116.422577 | 3.683506 | 2707.096685 | 0.665966 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 154.024072 | 15.337548 | 1124.251363 | 5.137798 |
| 1.19 | 139.073365 | 4.262040 | 1147.001361 | 9.924343 |
| 1.20 | 139.852962 | 3.422718 | 1101.656227 | 7.885818 |
| gotip | 142.087874 | 3.651276 | 1095.800484 | 7.222141 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 163.191208 | 15.091672 | 2337.949371 | 81.089385 |
| 1.19 | 142.033776 | 3.283347 | 2228.828227 | 15.208904 |
| 1.20 | 142.508878 | 7.970287 | 2173.160722 | 45.589595 |
| gotip | 148.717301 | 4.538967 | 2117.058314 | 30.174268 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 192.911577 | 8.004907 | 5877.006154 | 20.652754 |
| 1.19 | 171.974430 | 7.424559 | 3504.205717 | 9.534690 |
| 1.20 | 173.447789 | 3.838579 | 3522.948575 | 5.567524 |
| gotip | 167.455104 | 5.349621 | 3184.085331 | 1.362457 |

![switch_case](./switch_case__725e73000e.png)

