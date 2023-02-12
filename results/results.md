# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.5

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

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 465.826790 | 17.065428 | 3516.429704 | 2.151925 |
| 1.19 | 483.485018 | 5.925623 | 3542.682636 | 1.275290 |
| 1.20 | 419.935015 | 4.364099 | 3554.245139 | 4.277758 |
| gotip | 458.871976 | 5.856472 | 3584.167324 | 35.672858 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 119.304758 | 2.913458 | 4012.817221 | 4.221428 |
| 1.19 | 109.054440 | 7.002654 | 3903.337263 | 3.837364 |
| 1.20 | 115.164129 | 2.462829 | 3903.573787 | 6.271371 |
| gotip | 117.279325 | 4.194242 | 4246.147244 | 4.902626 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 127.704418 | 4.033644 | 5696.060965 | 26.752577 |
| 1.19 | 110.905032 | 12.821819 | 5565.911356 | 32.587439 |
| 1.20 | 118.916582 | 2.754815 | 6199.821458 | 21.517470 |
| gotip | 118.293658 | 5.753015 | 6453.094478 | 26.513987 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 130.568228 | 4.153086 | 5199.527566 | 61.594546 |
| 1.19 | 115.981217 | 8.576003 | 8773.425186 | 53.193913 |
| 1.20 | 118.531594 | 2.350751 | 8822.129393 | 75.874893 |
| gotip | 118.244045 | 4.141705 | 5341.844541 | 78.752607 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 124.400723 | 2.490303 | 5746.597004 | 34.589789 |
| 1.19 | 113.028709 | 11.265132 | 7371.598381 | 32.123479 |
| 1.20 | 116.577325 | 4.502241 | 8484.764093 | 73.936482 |
| gotip | 122.639279 | 3.957023 | 8923.721644 | 76.356701 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 140.703807 | 2.332661 | 5027.865430 | 1.929124 |
| 1.19 | 121.878448 | 4.222663 | 5086.232315 | 5.238663 |
| 1.20 | 128.697834 | 3.027080 | 5080.145786 | 1.874767 |
| gotip | 133.167846 | 3.580307 | 5276.613658 | 8.788934 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 131.991613 | 3.669967 | 3575.566961 | 5.677268 |
| 1.19 | 113.953862 | 14.057643 | 370.373178 | 0.425501 |
| 1.20 | 115.889982 | 1.821017 | 370.386574 | 0.821544 |
| gotip | 121.446357 | 2.351099 | 374.849024 | 0.612919 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 144.402344 | 31.291992 | 5067.958596 | 828.893983 |
| 1.19 | 122.661921 | 33.820451 | 4755.268351 | 262.851629 |
| 1.20 | 129.105073 | 4.415169 | 5271.275607 | 171.347475 |
| gotip | 132.252703 | 3.118540 | 5338.906655 | 216.590489 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 122.872251 | 5.978433 | 4347.960644 | 1.407515 |
| 1.19 | 110.363005 | 1.812802 | 2190.927725 | 0.679675 |
| 1.20 | 114.407137 | 2.149250 | 2191.345998 | 0.651863 |
| gotip | 117.192950 | 2.625105 | 4347.023528 | 3.728593 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 159.669347 | 20.593646 | 1151.000554 | 13.375617 |
| 1.19 | 135.376213 | 4.933735 | 1168.042534 | 8.668427 |
| 1.20 | 137.742446 | 2.526817 | 1163.118011 | 7.498661 |
| gotip | 142.805432 | 3.855360 | 1138.498122 | 9.352444 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 162.068389 | 9.311352 | 2366.517560 | 19.415568 |
| 1.19 | 139.205511 | 4.298676 | 2249.989957 | 7.775709 |
| 1.20 | 141.467638 | 4.353212 | 2230.922154 | 14.451466 |
| gotip | 149.240923 | 4.033554 | 2160.587688 | 16.074283 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 190.752961 | 14.501503 | 5723.622549 | 1.422428 |
| 1.19 | 164.647354 | 9.577143 | 3222.958514 | 5.066328 |
| 1.20 | 167.479857 | 2.770202 | 3224.347271 | 1.618589 |
| gotip | 168.852039 | 5.356045 | 2696.957917 | 4.881979 |

![switch_case](./switch_case__725e73000e.png)

