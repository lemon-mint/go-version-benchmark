# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.3

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
| 1.17 | 482.903313 | 14.085680 | 3870.428139 | 3.819260 |
| 1.18 | 487.259353 | 7.999256 | 3516.481751 | 70.819667 |
| 1.19 | 535.037819 | 3.947909 | 3541.262029 | 1.101685 |
| gotip | 4107.203481 | 52.036386 | 3557.748821 | 2.950195 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 136.693222 | 10.121651 | 3902.162182 | 4.828623 |
| 1.18 | 136.815909 | 5.154307 | 4012.442733 | 1.459773 |
| 1.19 | 146.714893 | 5.995307 | 3902.353640 | 1.049716 |
| gotip | 2896.814575 | 34.203341 | 3906.218764 | 2.909328 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 139.494596 | 2.814638 | 5702.791199 | 20.351358 |
| 1.18 | 147.256578 | 3.118216 | 5697.116566 | 46.609215 |
| 1.19 | 144.172976 | 4.900126 | 5567.710793 | 29.303265 |
| gotip | 2899.395852 | 26.411793 | 6218.404854 | 27.158141 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 151.138821 | 22.611673 | 8987.166314 | 36.470763 |
| 1.18 | 151.071263 | 6.234152 | 5229.259361 | 40.373803 |
| 1.19 | 146.983497 | 4.716853 | 8957.729446 | 61.359471 |
| gotip | 3010.240147 | 27.024471 | 9002.561640 | 17.801185 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 141.914033 | 24.959121 | 6614.059585 | 67.899783 |
| 1.18 | 141.528679 | 3.780508 | 5837.280150 | 55.993002 |
| 1.19 | 155.687514 | 6.082720 | 7474.258252 | 70.944181 |
| gotip | 2904.225374 | 17.047429 | 8778.237643 | 58.118031 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 157.348183 | 3.732237 | 5014.347548 | 3.142144 |
| 1.18 | 161.464915 | 4.981831 | 5028.374865 | 2.328766 |
| 1.19 | 158.183142 | 25.538388 | 5082.808067 | 4.756377 |
| gotip | 3375.976243 | 35.629276 | 5078.686398 | 1.804388 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 146.915305 | 4.076317 | 3715.262209 | 2.041918 |
| 1.18 | 146.912659 | 3.380969 | 3574.369308 | 0.751025 |
| 1.19 | 148.033207 | 2.508317 | 370.054136 | 0.423134 |
| gotip | 3212.973605 | 31.576480 | 369.834668 | 0.576953 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 158.663818 | 86.857451 | 4994.944033 | 393.954470 |
| 1.18 | 163.171477 | 103.920680 | 4561.881853 | 487.174779 |
| 1.19 | 160.918280 | 7.638818 | 4540.884814 | 239.122437 |
| gotip | 3896.059989 | 59.498834 | 5463.168642 | 190.849155 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 135.304824 | 5.078327 | 3167.156690 | 15.122995 |
| 1.18 | 138.812721 | 4.677940 | 4347.023691 | 1.085995 |
| 1.19 | 148.031589 | 5.521261 | 2191.598899 | 0.694735 |
| gotip | 2892.857487 | 25.388679 | 2191.438171 | 1.125791 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 173.915276 | 27.533756 | 1093.369136 | 16.650976 |
| 1.18 | 175.976496 | 15.451687 | 1155.927889 | 10.437157 |
| 1.19 | 175.388596 | 3.456571 | 1165.975686 | 10.859675 |
| gotip | 4046.999983 | 33.989149 | 1161.492383 | 6.742846 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 178.553037 | 6.420128 | 2359.823206 | 10.934278 |
| 1.18 | 181.016818 | 6.422954 | 2370.270297 | 14.562221 |
| 1.19 | 181.684123 | 7.277388 | 2250.986478 | 11.653346 |
| gotip | 4088.245655 | 35.282430 | 2215.521577 | 8.724933 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 215.612079 | 20.183776 | 5206.994913 | 3.351554 |
| 1.18 | 211.910799 | 21.491304 | 5724.000109 | 2.700502 |
| 1.19 | 211.405991 | 2.743772 | 3222.381780 | 2.099547 |
| gotip | 5345.999661 | 49.996457 | 3223.918094 | 1.698703 |

![switch_case](./switch_case__725e73000e.png)

