# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.18.5

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz

Cores: 1

Mhz: 2593.906000

CacheSize: 36608

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz

Cores: 1

Mhz: 2593.906000

CacheSize: 36608

Microcode: 0xffffffff

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 127.838211 | 5.476270 | 3900.821972 | 0.801537 |
| 1.18 | 130.014994 | 3.230051 | 4011.118841 | 1.311301 |
| 1.19 | 139.036907 | 2.966615 | 3901.606673 | 1.780159 |
| gotip | 146.433561 | 3.155666 | 4011.445145 | 0.794918 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 132.989355 | 3.905868 | 5693.309196 | 26.633661 |
| 1.18 | 136.052430 | 10.563583 | 5675.955487 | 16.426043 |
| 1.19 | 140.870306 | 4.418939 | 5514.761986 | 32.320245 |
| gotip | 143.935591 | 1.742503 | 6344.642129 | 26.700486 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 134.237665 | 5.439653 | 8551.504131 | 32.663887 |
| 1.18 | 140.333927 | 1.784794 | 4874.807251 | 24.070726 |
| 1.19 | 141.449844 | 8.064192 | 8480.689232 | 35.251138 |
| gotip | 144.917048 | 2.624616 | 5139.949496 | 26.838624 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 129.577201 | 4.647208 | 5845.177547 | 69.132843 |
| 1.18 | 133.816154 | 3.433259 | 5602.266769 | 27.664798 |
| 1.19 | 142.811120 | 1.705550 | 7267.282312 | 24.928082 |
| gotip | 149.761887 | 3.109961 | 8644.935664 | 38.371118 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 143.618279 | 2.503810 | 5012.632826 | 2.311833 |
| 1.18 | 151.274976 | 6.661108 | 5027.642889 | 3.277869 |
| 1.19 | 150.877214 | 3.329590 | 5080.699156 | 1.569033 |
| gotip | 155.254207 | 3.592337 | 5020.564984 | 1.956943 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 133.820086 | 2.602434 | 3713.467769 | 5.602530 |
| 1.18 | 140.653191 | 1.433595 | 3574.686483 | 0.894836 |
| 1.19 | 141.969278 | 2.527433 | 369.143304 | 0.319414 |
| gotip | 149.448229 | 2.243091 | 369.613422 | 0.627723 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 146.687477 | 12.207740 | 5097.363611 | 468.906861 |
| 1.18 | 158.601427 | 74.308146 | 4090.103836 | 502.348952 |
| 1.19 | 158.807097 | 11.875775 | 4996.621332 | 551.799337 |
| gotip | 161.588267 | 2.974665 | 5310.589610 | 164.990999 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 126.986538 | 4.557114 | 3163.092619 | 11.098293 |
| 1.18 | 132.827845 | 2.927548 | 4348.013526 | 2.142168 |
| 1.19 | 141.869997 | 5.598149 | 2190.604596 | 1.661617 |
| gotip | 143.605570 | 2.562201 | 4347.478675 | 0.629446 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 156.819656 | 5.487758 | 1087.441083 | 11.526423 |
| 1.18 | 168.453137 | 9.948730 | 1145.448643 | 11.976925 |
| 1.19 | 166.290744 | 6.150342 | 1154.904754 | 8.064775 |
| gotip | 170.153649 | 3.481353 | 1073.420208 | 3.170219 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 163.428253 | 2.334653 | 2353.789736 | 17.577726 |
| 1.18 | 173.299599 | 3.355269 | 2353.301538 | 9.747755 |
| 1.19 | 170.652152 | 4.955858 | 2230.896983 | 5.227371 |
| gotip | 179.531545 | 2.989732 | 2204.471470 | 9.212132 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 192.412284 | 3.401928 | 5200.829983 | 2.603761 |
| 1.18 | 197.936736 | 12.612372 | 5722.885702 | 2.714444 |
| 1.19 | 202.933844 | 7.625682 | 3221.800181 | 1.425857 |
| gotip | 206.480579 | 3.046314 | 3272.218410 | 0.904618 |

![switch_case](./switch_case__725e73000e.png)

