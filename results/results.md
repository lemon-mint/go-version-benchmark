# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.10

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.436000

CacheSize: 49152

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.436000

CacheSize: 49152

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 464.432769 | 23.851778 | 3616.838810 | 2.985981 |
| 1.19 | 485.647717 | 7.086635 | 3613.210882 | 64.833616 |
| 1.20 | 415.202224 | 19.264644 | 3633.899291 | 2.449774 |
| gotip | 454.764178 | 3.872729 | 3555.435674 | 2.026205 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 114.144206 | 2.877760 | 4797.031547 | 1.128633 |
| 1.19 | 108.913235 | 4.580570 | 4794.773391 | 0.722760 |
| 1.20 | 114.532692 | 1.956327 | 4793.585284 | 0.800142 |
| gotip | 111.781283 | 3.083664 | 5214.410288 | 1.012230 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 123.248717 | 3.104146 | 5196.556804 | 356.979354 |
| 1.19 | 110.839764 | 2.722645 | 5199.844370 | 24.041070 |
| 1.20 | 114.753562 | 4.395771 | 6127.954664 | 54.790056 |
| gotip | 115.248317 | 5.759080 | 6604.318959 | 63.102024 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 127.481669 | 3.908482 | 6261.936796 | 66.948832 |
| 1.19 | 110.685382 | 11.934961 | 10545.069106 | 39.901873 |
| 1.20 | 116.971602 | 3.685811 | 10458.458322 | 53.419098 |
| gotip | 111.744612 | 3.978794 | 6315.290741 | 92.872808 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 123.991406 | 4.096576 | 5886.125584 | 36.052843 |
| 1.19 | 114.971654 | 3.866005 | 7436.555083 | 37.439119 |
| 1.20 | 115.981751 | 3.278732 | 8054.127657 | 39.237293 |
| gotip | 115.168376 | 2.121931 | 8286.066050 | 63.208762 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 139.042410 | 19.559005 | 5984.892165 | 7.172943 |
| 1.19 | 119.470779 | 3.934869 | 5918.844595 | 2.426417 |
| 1.20 | 124.863510 | 3.525350 | 5968.802460 | 2.230808 |
| gotip | 128.046755 | 2.344378 | 6028.219467 | 3.511439 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 126.078000 | 2.611662 | 3542.574779 | 1.775769 |
| 1.19 | 109.749317 | 16.307921 | 361.115769 | 0.555353 |
| 1.20 | 118.170360 | 2.041185 | 360.460116 | 1.079490 |
| gotip | 115.309860 | 3.438206 | 371.158886 | 0.455141 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 140.137061 | 15.571677 | 5100.215656 | 758.335468 |
| 1.19 | 120.632394 | 22.973969 | 4983.192761 | 428.920393 |
| 1.20 | 128.377811 | 6.210243 | 5195.064093 | 351.110444 |
| gotip | 130.437903 | 4.617658 | 5100.810095 | 327.946571 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 116.765647 | 2.463171 | 5218.463428 | 2.004262 |
| 1.19 | 108.492211 | 2.292676 | 2705.879075 | 0.477344 |
| 1.20 | 112.477325 | 4.205813 | 2705.997423 | 0.607733 |
| gotip | 113.425011 | 5.024148 | 5218.661609 | 1.176448 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 151.248692 | 5.201877 | 1116.419540 | 4.723251 |
| 1.19 | 134.276610 | 6.136057 | 1145.423411 | 10.815662 |
| 1.20 | 137.201969 | 4.295264 | 1104.212906 | 18.042715 |
| gotip | 135.003584 | 1.689302 | 1081.419602 | 5.649576 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 158.867800 | 13.276893 | 2343.277409 | 22.139230 |
| 1.19 | 137.491242 | 3.462083 | 2196.092577 | 20.505997 |
| 1.20 | 144.172017 | 2.776612 | 2178.188248 | 34.286811 |
| gotip | 142.388660 | 3.901949 | 2114.871784 | 20.275131 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 183.056573 | 4.222243 | 5929.918220 | 25.008926 |
| 1.19 | 162.048958 | 7.048953 | 3535.843200 | 2.497108 |
| 1.20 | 168.147316 | 6.162117 | 3520.115268 | 4.785665 |
| gotip | 161.052837 | 3.768189 | 3183.282576 | 1.885648 |

![switch_case](./switch_case__725e73000e.png)

