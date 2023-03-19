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

Mhz: 2793.594000

CacheSize: 49152

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.594000

CacheSize: 49152

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 456.982836 | 3.588947 | 3617.348888 | 173.368546 |
| 1.19 | 479.826575 | 3.382710 | 3615.310354 | 2.281334 |
| 1.20 | 409.399911 | 4.108599 | 3632.495380 | 94.937353 |
| gotip | 445.696261 | 3.343488 | 3575.267666 | 14.030816 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 112.424438 | 8.488749 | 4797.627100 | 3.246375 |
| 1.19 | 106.133701 | 5.960945 | 4792.443401 | 0.897912 |
| 1.20 | 110.232969 | 3.538162 | 4792.427034 | 0.772373 |
| gotip | 113.418065 | 4.396946 | 5205.864878 | 5.010495 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 122.367880 | 2.846953 | 5189.398461 | 14.739991 |
| 1.19 | 108.877634 | 2.957857 | 5192.345515 | 18.217624 |
| 1.20 | 113.400319 | 3.634816 | 6093.884507 | 72.483667 |
| gotip | 116.400515 | 3.546867 | 6348.517558 | 24.767123 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 126.512504 | 3.960538 | 6103.261503 | 95.553640 |
| 1.19 | 111.131266 | 16.199154 | 10342.996833 | 52.536526 |
| 1.20 | 115.662123 | 4.253175 | 10373.661887 | 64.703244 |
| gotip | 116.122677 | 4.520610 | 6167.324474 | 118.667503 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 117.187666 | 4.565810 | 5640.628464 | 32.279783 |
| 1.19 | 111.140394 | 2.498368 | 7277.152351 | 35.172677 |
| 1.20 | 116.532381 | 2.901219 | 7966.310736 | 73.480778 |
| gotip | 115.900479 | 4.976843 | 7860.787935 | 31.485321 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 137.208928 | 3.850618 | 5981.351803 | 5.422742 |
| 1.19 | 118.654179 | 5.617011 | 5926.862529 | 2.695151 |
| 1.20 | 122.678754 | 3.764273 | 6012.765533 | 13.582422 |
| gotip | 125.973015 | 6.751323 | 6161.723218 | 3.868690 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 127.654115 | 3.630739 | 3539.009959 | 1.885067 |
| 1.19 | 109.651381 | 4.570835 | 360.988958 | 1.765236 |
| 1.20 | 116.208193 | 2.267940 | 361.771691 | 0.440000 |
| gotip | 117.659988 | 2.914836 | 363.441024 | 5.383755 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 143.849933 | 6.145431 | 4210.154242 | 834.799309 |
| 1.19 | 120.673644 | 18.315619 | 5092.907849 | 375.533415 |
| 1.20 | 127.865841 | 5.095440 | 5265.969128 | 258.333824 |
| gotip | 129.701798 | 4.411689 | 5070.350861 | 273.462393 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 116.394360 | 15.444555 | 5218.227093 | 1.414339 |
| 1.19 | 106.105016 | 3.624859 | 2706.182201 | 0.871087 |
| 1.20 | 109.860128 | 3.263636 | 2705.571098 | 0.393345 |
| gotip | 114.007156 | 3.337043 | 2706.449677 | 1.490058 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 152.217029 | 4.736004 | 1119.725659 | 10.569035 |
| 1.19 | 133.289408 | 4.986218 | 1148.871294 | 9.400491 |
| 1.20 | 138.817284 | 2.570237 | 1098.527148 | 8.569941 |
| gotip | 139.620043 | 1.818894 | 1084.759304 | 4.815422 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 155.764637 | 3.428914 | 2327.681003 | 26.060877 |
| 1.19 | 134.003788 | 2.738376 | 2198.342152 | 7.778179 |
| 1.20 | 140.263107 | 3.506333 | 2164.706256 | 24.175167 |
| gotip | 145.139561 | 2.833426 | 2134.093015 | 16.592056 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 186.629966 | 19.757828 | 5944.076801 | 28.086003 |
| 1.19 | 156.821412 | 9.790979 | 3517.950243 | 4.323071 |
| 1.20 | 166.780453 | 3.282105 | 3516.292427 | 5.856756 |
| gotip | 163.295582 | 3.500529 | 2827.025623 | 1.975330 |

![switch_case](./switch_case__725e73000e.png)

