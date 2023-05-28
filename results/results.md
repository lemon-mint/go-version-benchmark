# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.9

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.439000

CacheSize: 49152

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.439000

CacheSize: 49152

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 459.721961 | 6.055547 | 3620.257077 | 4.124084 |
| 1.19 | 479.920747 | 6.980730 | 3615.824535 | 1.319227 |
| 1.20 | 412.023652 | 6.947599 | 3633.700130 | 2.883263 |
| gotip | 469.136799 | 12.679092 | 3578.851762 | 316.644209 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 120.197359 | 6.323316 | 4799.304411 | 1.780278 |
| 1.19 | 110.970386 | 7.333482 | 4795.987780 | 0.776646 |
| 1.20 | 117.121070 | 3.868416 | 4796.136445 | 1.500926 |
| gotip | 114.697760 | 4.320126 | 5207.753723 | 1.297117 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 126.013504 | 9.758145 | 5212.126053 | 14.126629 |
| 1.19 | 109.224388 | 11.979051 | 5231.855804 | 26.545817 |
| 1.20 | 114.818449 | 2.535769 | 6118.663371 | 25.836147 |
| gotip | 111.933681 | 4.516697 | 6599.400563 | 35.870957 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 127.966646 | 2.705986 | 6098.367847 | 88.520729 |
| 1.19 | 109.777310 | 3.214997 | 10422.541504 | 91.645650 |
| 1.20 | 117.275352 | 2.207800 | 10603.525180 | 50.512500 |
| gotip | 123.101136 | 5.003385 | 6974.979897 | 123.980035 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 124.014592 | 4.204012 | 5891.991786 | 66.049712 |
| 1.19 | 111.745570 | 2.355207 | 7515.164812 | 79.594757 |
| 1.20 | 117.034098 | 2.696334 | 8141.295200 | 56.862711 |
| gotip | 118.130291 | 2.739139 | 8494.297163 | 298.488519 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 141.658591 | 10.678932 | 5980.076135 | 5.000520 |
| 1.19 | 119.076553 | 4.220345 | 5919.602684 | 3.134895 |
| 1.20 | 130.617036 | 8.088665 | 5979.477091 | 3.232115 |
| gotip | 136.109345 | 6.597196 | 5962.165016 | 8.806748 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 126.917980 | 3.680739 | 3542.001782 | 2.108792 |
| 1.19 | 109.471724 | 4.241544 | 360.725099 | 0.516936 |
| 1.20 | 118.517460 | 4.557527 | 362.209038 | 0.478732 |
| gotip | 115.420057 | 3.736803 | 367.376181 | 1.543592 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 147.287305 | 8.144033 | 4819.314055 | 677.096046 |
| 1.19 | 127.294525 | 18.779056 | 5309.979914 | 541.065767 |
| 1.20 | 136.241760 | 8.901255 | 5268.461382 | 433.230527 |
| gotip | 132.455691 | 7.469145 | 5341.084163 | 188.401034 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 120.443761 | 3.538361 | 5219.417991 | 1.459058 |
| 1.19 | 108.139662 | 3.469564 | 2706.926505 | 5.353164 |
| 1.20 | 115.871126 | 5.463331 | 2706.436957 | 1.232997 |
| gotip | 114.398786 | 4.315437 | 2706.937249 | 0.901690 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 150.529486 | 10.841604 | 1120.038575 | 7.572405 |
| 1.19 | 136.330783 | 3.795383 | 1151.710497 | 7.850087 |
| 1.20 | 138.868759 | 3.915446 | 1107.244759 | 3.545821 |
| gotip | 136.649012 | 2.634283 | 1079.455967 | 12.174706 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 159.500881 | 5.490449 | 2342.660687 | 24.468909 |
| 1.19 | 137.353428 | 3.052046 | 2195.035195 | 25.304309 |
| 1.20 | 147.008896 | 5.737101 | 2166.671802 | 47.162883 |
| gotip | 149.782593 | 4.077021 | 2165.748814 | 28.003428 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 187.017438 | 39.784469 | 5925.590923 | 23.804921 |
| 1.19 | 167.396033 | 3.448470 | 3505.732822 | 12.748454 |
| 1.20 | 169.849084 | 4.304930 | 3528.941747 | 6.161912 |
| gotip | 164.979298 | 9.477654 | 2826.644040 | 2.493670 |

![switch_case](./switch_case__725e73000e.png)

