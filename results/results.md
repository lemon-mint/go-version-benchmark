# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.18.5

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

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 122.431456 | 3.043247 | 4792.198211 | 1.469026 |
| 1.18 | 127.882559 | 1.617745 | 4795.965920 | 0.610418 |
| 1.19 | 179.817652 | 20.009112 | 4792.584853 | 0.417937 |
| gotip | 139.078312 | 6.872770 | 4796.771619 | 2.075413 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 131.437930 | 5.233318 | 5174.748800 | 85.389340 |
| 1.18 | 138.114391 | 7.383490 | 5174.692205 | 19.437204 |
| 1.19 | 142.180320 | 4.347551 | 5193.561308 | 20.041242 |
| gotip | 143.093812 | 3.333504 | 5203.251847 | 25.736512 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 132.245362 | 4.095117 | 9926.956914 | 83.746306 |
| 1.18 | 140.358686 | 2.260641 | 5866.638554 | 58.013335 |
| 1.19 | 143.552923 | 2.951009 | 10110.139754 | 90.903160 |
| gotip | 144.443982 | 6.178292 | 6014.430600 | 96.891459 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 127.660300 | 2.060021 | 5986.680535 | 69.964207 |
| 1.18 | 134.471548 | 3.155450 | 5795.779505 | 28.642864 |
| 1.19 | 144.001399 | 3.559533 | 7256.920609 | 88.243303 |
| gotip | 144.420430 | 4.135073 | 8030.237488 | 99.831461 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 144.268928 | 2.742135 | 6088.970464 | 12.376644 |
| 1.18 | 149.149658 | 1.742563 | 5978.815594 | 6.608663 |
| 1.19 | 152.658665 | 5.501448 | 5929.741694 | 2.321795 |
| gotip | 157.761534 | 4.738518 | 5940.844488 | 4.374187 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 134.914700 | 2.521045 | 3634.971276 | 5.602219 |
| 1.18 | 143.111687 | 2.586616 | 3550.022795 | 1.876985 |
| 1.19 | 141.035387 | 3.684845 | 362.418573 | 2.189899 |
| gotip | 148.173366 | 3.388616 | 391.342402 | 38.467667 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 145.188425 | 4.604965 | 5093.277125 | 572.003026 |
| 1.18 | 154.074190 | 68.024729 | 4954.909440 | 798.853102 |
| 1.19 | 155.022047 | 120.809541 | 4907.149586 | 568.912556 |
| gotip | 156.639672 | 7.307709 | 5159.405406 | 362.240483 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 123.199655 | 3.417359 | 3680.141238 | 9.451914 |
| 1.18 | 127.804678 | 4.450180 | 5214.589499 | 0.699758 |
| 1.19 | 141.891017 | 4.851267 | 2705.099140 | 0.385615 |
| gotip | 139.410204 | 2.371565 | 5215.020365 | 0.810426 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 161.831983 | 10.057210 | 1076.309855 | 3.905964 |
| 1.18 | 167.935945 | 11.891199 | 1113.359690 | 6.528766 |
| 1.19 | 169.290859 | 13.096906 | 1147.368069 | 13.586750 |
| gotip | 169.179290 | 2.665402 | 1153.797640 | 10.979785 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 167.923748 | 2.175885 | 2345.486047 | 9.787200 |
| 1.18 | 172.057567 | 3.175127 | 2342.220119 | 27.414229 |
| 1.19 | 175.728643 | 2.192216 | 2219.940477 | 18.849974 |
| gotip | 174.176925 | 4.534882 | 2182.104201 | 18.868622 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 203.560901 | 12.518426 | 5930.805077 | 21.668180 |
| 1.18 | 199.048452 | 14.891589 | 5912.771113 | 19.477635 |
| 1.19 | 202.537041 | 19.454415 | 3569.612209 | 6.106752 |
| gotip | 204.774270 | 4.552200 | 3404.949796 | 3.578990 |

![switch_case](./switch_case__725e73000e.png)

