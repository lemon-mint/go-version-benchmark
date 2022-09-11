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

Mhz: 2095.078000

CacheSize: 36608

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHz

Cores: 1

Mhz: 2095.078000

CacheSize: 36608

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 586.603460 | 21.814519 | 4634.803455 | 16.486654 |
| 1.18 | 640.386082 | 17.764158 | 4247.983217 | 46.679249 |
| 1.19 | 670.470750 | 28.781688 | 4255.306326 | 59.932649 |
| gotip | 683.216591 | 11.777314 | 4245.448685 | 7.660966 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 153.661242 | 10.144181 | 4680.428303 | 16.939083 |
| 1.18 | 157.628488 | 3.722190 | 4816.847620 | 34.404726 |
| 1.19 | 183.551830 | 7.569546 | 4696.669857 | 91.985625 |
| gotip | 183.433380 | 5.821461 | 4686.038904 | 37.471250 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 163.559067 | 6.851849 | 6805.071635 | 137.735235 |
| 1.18 | 186.107448 | 9.693389 | 7069.272801 | 126.360629 |
| 1.19 | 189.329776 | 9.483365 | 6852.164395 | 126.057229 |
| gotip | 196.406463 | 6.578442 | 7717.340974 | 71.742753 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 177.760519 | 7.008681 | 10766.042493 | 77.275340 |
| 1.18 | 200.466463 | 10.217963 | 6359.758658 | 237.870946 |
| 1.19 | 184.025169 | 10.313102 | 10646.023175 | 47.279571 |
| gotip | 185.042540 | 9.763386 | 10773.527736 | 187.672004 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 163.065592 | 9.375175 | 7532.031856 | 228.680135 |
| 1.18 | 169.500824 | 9.116501 | 7053.390327 | 130.254568 |
| 1.19 | 186.317887 | 6.378093 | 9343.225159 | 242.772678 |
| gotip | 187.159074 | 5.232932 | 10785.822796 | 276.472824 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 192.423673 | 6.230560 | 6040.737327 | 104.891458 |
| 1.18 | 195.311822 | 8.735893 | 6032.172695 | 17.425813 |
| 1.19 | 191.709059 | 6.172908 | 6112.112327 | 16.943984 |
| gotip | 214.534970 | 10.552720 | 6112.916515 | 14.264298 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 170.032465 | 7.284101 | 4458.313796 | 102.283996 |
| 1.18 | 188.860990 | 9.621404 | 4344.737832 | 106.573508 |
| 1.19 | 182.262779 | 4.876211 | 465.691275 | 12.174094 |
| gotip | 195.007306 | 6.297064 | 473.148464 | 5.751523 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 199.651131 | 8.870448 | 7413.014376 | 520.588523 |
| 1.18 | 206.999738 | 207.755752 | 4639.181322 | 690.314688 |
| 1.19 | 213.536708 | 26.820398 | 5537.609581 | 1151.846002 |
| gotip | 207.343559 | 8.984101 | 6410.903691 | 899.027219 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 160.116224 | 8.852997 | 3919.239109 | 51.764434 |
| 1.18 | 170.286013 | 6.648092 | 5231.801161 | 65.368729 |
| 1.19 | 184.389617 | 7.203214 | 2768.295112 | 66.703493 |
| gotip | 189.400154 | 15.141361 | 2629.723104 | 63.078764 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 201.045827 | 32.801966 | 1404.530920 | 19.475377 |
| 1.18 | 210.754742 | 59.761960 | 1464.360310 | 27.166649 |
| 1.19 | 213.901616 | 79.427095 | 1515.283238 | 49.110414 |
| gotip | 233.629224 | 16.639819 | 1462.602860 | 48.235375 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 211.802684 | 6.153020 | 3010.360921 | 46.648207 |
| 1.18 | 219.813603 | 8.030128 | 2998.228394 | 26.022614 |
| 1.19 | 227.469026 | 17.217492 | 2884.225215 | 39.564336 |
| gotip | 236.420965 | 9.135338 | 2831.424498 | 27.764751 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 257.078973 | 24.895067 | 6254.592808 | 123.461508 |
| 1.18 | 261.280386 | 23.991406 | 7085.307720 | 222.822727 |
| 1.19 | 248.815485 | 29.150063 | 3880.899071 | 11.289370 |
| gotip | 285.921062 | 15.099846 | 3886.503549 | 22.271389 |

![switch_case](./switch_case__725e73000e.png)

