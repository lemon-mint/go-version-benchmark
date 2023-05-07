# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.8

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
| 1.18 | 481.752770 | 23.983093 | 3524.489645 | 18.149018 |
| 1.19 | 492.335029 | 14.010851 | 3544.964452 | 1.575227 |
| 1.20 | 420.410086 | 3.159773 | 3561.287295 | 59.873414 |
| gotip | 456.192651 | 4.320807 | 3547.602324 | 4.283655 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 121.040887 | 4.933266 | 4014.358511 | 33.887597 |
| 1.19 | 108.559922 | 3.353670 | 3909.816648 | 18.533988 |
| 1.20 | 114.213681 | 4.582441 | 3918.201420 | 57.987960 |
| gotip | 112.735183 | 2.720752 | 4245.314964 | 23.185869 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 127.180442 | 2.719339 | 5755.003579 | 44.258846 |
| 1.19 | 113.261958 | 5.836274 | 5567.117584 | 42.226494 |
| 1.20 | 115.900944 | 5.166270 | 6217.844038 | 17.755489 |
| gotip | 115.116739 | 6.139284 | 6595.732988 | 29.661013 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 131.059199 | 3.034586 | 5186.731692 | 40.410436 |
| 1.19 | 114.788025 | 3.750340 | 8826.361206 | 37.363929 |
| 1.20 | 118.535640 | 6.069651 | 8901.340030 | 49.243455 |
| gotip | 116.879556 | 4.054097 | 5373.695989 | 38.134010 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 122.355041 | 7.922367 | 5868.043692 | 73.181583 |
| 1.19 | 116.637133 | 6.757894 | 7525.837342 | 42.415847 |
| 1.20 | 118.393831 | 2.523456 | 8632.725913 | 53.613186 |
| gotip | 117.374131 | 2.053700 | 8849.071478 | 57.572989 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 145.588302 | 7.613979 | 5029.940728 | 24.168121 |
| 1.19 | 123.694278 | 2.893367 | 5086.559444 | 29.413164 |
| 1.20 | 130.301661 | 3.642373 | 5082.805447 | 10.697444 |
| gotip | 133.273359 | 4.620915 | 5206.702434 | 2.393839 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 134.879065 | 4.747379 | 3577.701031 | 6.540286 |
| 1.19 | 117.869227 | 5.950709 | 370.116128 | 5.531846 |
| 1.20 | 121.248437 | 5.495511 | 370.109923 | 0.538811 |
| gotip | 119.237153 | 3.139437 | 374.901000 | 9.088462 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 151.269803 | 8.666539 | 4619.808047 | 679.143708 |
| 1.19 | 125.930115 | 4.330713 | 4741.144824 | 386.133389 |
| 1.20 | 131.970923 | 3.507484 | 5296.117705 | 236.420157 |
| gotip | 131.750466 | 3.660887 | 5346.226550 | 224.937263 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 121.300034 | 3.870764 | 4349.618108 | 57.958631 |
| 1.19 | 111.810316 | 6.772655 | 2194.589257 | 32.278504 |
| 1.20 | 115.823354 | 3.388987 | 2191.606292 | 27.157114 |
| gotip | 114.953635 | 4.649394 | 4349.171121 | 26.140899 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 155.406675 | 17.280031 | 1160.280514 | 18.456770 |
| 1.19 | 136.844355 | 19.456017 | 1183.282862 | 16.568415 |
| 1.20 | 140.542273 | 5.245101 | 1183.556197 | 16.880722 |
| gotip | 138.294620 | 2.884724 | 1167.942207 | 19.197503 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 162.166127 | 5.359869 | 2379.298590 | 20.212867 |
| 1.19 | 141.569423 | 5.130572 | 2283.125563 | 32.203633 |
| 1.20 | 144.040580 | 5.276536 | 2250.641636 | 23.427380 |
| gotip | 147.989446 | 5.416537 | 2152.466228 | 20.312579 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 193.491926 | 5.742106 | 5727.803367 | 86.850035 |
| 1.19 | 166.441519 | 12.888690 | 3223.968444 | 44.372653 |
| 1.20 | 172.477598 | 6.209494 | 3224.847550 | 3.112751 |
| gotip | 165.367747 | 3.924735 | 2758.095623 | 2.233843 |

![switch_case](./switch_case__725e73000e.png)

