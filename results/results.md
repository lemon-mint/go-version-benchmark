# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHz

Cores: 1

Mhz: 2095.172000

CacheSize: 36608

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHz

Cores: 1

Mhz: 2095.172000

CacheSize: 36608

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 572.636772 | 72.577523 | 4301.281353 | 72.412251 |
| 1.18 | 580.356997 | 20.776607 | 3942.390705 | 48.822702 |
| 1.19 | 684.709398 | 34.032057 | 4011.099370 | 60.054583 |
| gotip | 653.531556 | 18.649729 | 3976.379259 | 58.292719 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 159.189718 | 21.352646 | 4392.156291 | 44.317652 |
| 1.18 | 162.802498 | 7.649728 | 4561.103472 | 69.355910 |
| 1.19 | 174.556270 | 8.020587 | 4368.025812 | 55.401504 |
| gotip | 174.554308 | 4.540038 | 4463.732496 | 61.713914 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 164.100064 | 5.505108 | 6486.492559 | 96.828574 |
| 1.18 | 168.897808 | 3.433640 | 6537.792027 | 74.596993 |
| 1.19 | 169.786243 | 5.241629 | 6364.134064 | 89.580884 |
| gotip | 174.677727 | 5.466791 | 7566.090304 | 88.896430 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 165.874975 | 7.050540 | 10307.033207 | 113.840130 |
| 1.18 | 178.464371 | 4.867679 | 6027.318912 | 64.481298 |
| 1.19 | 178.213121 | 4.386278 | 10254.611891 | 113.158622 |
| gotip | 181.910483 | 8.002129 | 6181.321035 | 73.716546 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 164.384632 | 5.860216 | 6928.094101 | 127.835204 |
| 1.18 | 165.306469 | 4.760819 | 6595.373497 | 80.073958 |
| 1.19 | 180.855874 | 4.731355 | 8589.302808 | 113.744594 |
| gotip | 181.880637 | 6.018876 | 10133.203621 | 74.152658 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 181.642184 | 5.052734 | 5667.768782 | 77.122301 |
| 1.18 | 187.065139 | 6.098200 | 5683.682343 | 100.335925 |
| 1.19 | 181.843338 | 5.188574 | 5646.944487 | 90.029912 |
| gotip | 192.060137 | 4.623331 | 5578.809345 | 66.672138 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 170.805471 | 6.683479 | 4130.990561 | 46.612885 |
| 1.18 | 186.980842 | 8.924762 | 3997.738347 | 75.486343 |
| 1.19 | 177.170122 | 5.833842 | 416.031476 | 8.446637 |
| gotip | 182.506844 | 4.741034 | 414.259748 | 15.466314 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 179.439568 | 23.516211 | 6815.865510 | 496.838770 |
| 1.18 | 188.161997 | 139.072072 | 5066.456581 | 672.700310 |
| 1.19 | 195.678978 | 82.605553 | 5125.043330 | 652.528569 |
| gotip | 191.532715 | 5.635311 | 6806.320056 | 766.906437 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 158.867964 | 141.443001 | 3638.513122 | 51.491506 |
| 1.18 | 166.627896 | 6.027147 | 4877.735435 | 78.111704 |
| 1.19 | 180.577269 | 10.750572 | 2488.693328 | 28.419866 |
| gotip | 176.885586 | 7.274323 | 4810.736800 | 47.572157 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 216.755629 | 20.025966 | 1292.725218 | 22.385688 |
| 1.18 | 210.727480 | 37.599489 | 1320.321196 | 22.891390 |
| 1.19 | 201.538316 | 21.026192 | 1354.060181 | 17.580135 |
| gotip | 214.278119 | 14.458360 | 1265.932800 | 10.890728 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 208.354562 | 7.125544 | 2748.562013 | 38.447358 |
| 1.18 | 211.682227 | 9.561596 | 2780.621370 | 31.187233 |
| 1.19 | 221.226087 | 8.036646 | 2672.426311 | 42.308735 |
| gotip | 223.186100 | 7.672759 | 2631.823013 | 24.127854 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 244.689819 | 25.798150 | 5750.398256 | 46.802726 |
| 1.18 | 241.482613 | 27.269519 | 6366.108126 | 67.600357 |
| 1.19 | 247.612917 | 48.317431 | 3567.493228 | 35.565965 |
| gotip | 266.827813 | 11.328090 | 3621.975495 | 69.774867 |

![switch_case](./switch_case__725e73000e.png)

