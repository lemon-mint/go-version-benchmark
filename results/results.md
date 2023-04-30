# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.8

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHz

Cores: 1

Mhz: 2095.170000

CacheSize: 36608

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHz

Cores: 1

Mhz: 2095.170000

CacheSize: 36608

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 563.359302 | 22.897348 | 4306.829949 | 69.428715 |
| 1.19 | 590.282379 | 9.110408 | 4309.490170 | 81.198978 |
| 1.20 | 495.673639 | 4.865698 | 4268.996819 | 78.621737 |
| gotip | 544.076337 | 9.187033 | 4278.033982 | 73.754750 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 139.585048 | 26.695227 | 4913.511792 | 100.956572 |
| 1.19 | 131.714192 | 3.558382 | 4757.597857 | 83.420096 |
| 1.20 | 135.561168 | 5.918979 | 4720.859577 | 102.767110 |
| gotip | 131.442800 | 3.150110 | 5257.491686 | 87.539304 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 150.607635 | 17.074758 | 6881.689437 | 51.215797 |
| 1.19 | 131.061919 | 7.018067 | 6705.554216 | 57.778470 |
| 1.20 | 138.675821 | 4.431881 | 7524.502895 | 52.584254 |
| gotip | 133.580531 | 3.840519 | 7980.225770 | 77.875936 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 158.640335 | 20.865609 | 6040.932528 | 55.891438 |
| 1.19 | 136.442930 | 4.689897 | 10415.786112 | 119.109974 |
| 1.20 | 140.307451 | 4.481335 | 10444.184195 | 54.762124 |
| gotip | 138.220874 | 6.474226 | 6275.761866 | 48.804331 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 142.467638 | 4.162995 | 6906.371435 | 44.869062 |
| 1.19 | 134.680886 | 3.467316 | 8854.916404 | 107.753014 |
| 1.20 | 139.513505 | 5.698359 | 10176.407135 | 176.335662 |
| gotip | 135.777155 | 3.009497 | 10683.890878 | 130.893181 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 163.195872 | 4.804944 | 6092.167983 | 101.020696 |
| 1.19 | 139.822711 | 3.788682 | 6105.625100 | 76.813864 |
| 1.20 | 149.578547 | 4.967133 | 6096.477222 | 92.722177 |
| gotip | 149.735777 | 5.584309 | 6373.891337 | 112.194291 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 148.924504 | 13.096216 | 4299.929098 | 82.051685 |
| 1.19 | 133.044111 | 5.734975 | 450.644819 | 6.655708 |
| 1.20 | 138.645358 | 8.743943 | 446.936314 | 7.623311 |
| gotip | 134.353944 | 3.335010 | 464.389577 | 11.849071 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 171.259167 | 5.740921 | 4284.657752 | 1050.240964 |
| 1.19 | 150.704105 | 34.612449 | 4657.721585 | 864.472021 |
| 1.20 | 154.092671 | 10.217632 | 5864.295056 | 1019.554984 |
| gotip | 149.713840 | 5.596646 | 5692.889235 | 784.888294 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 139.500370 | 4.343168 | 5225.833903 | 102.716654 |
| 1.19 | 135.590528 | 5.581013 | 2668.291038 | 40.770936 |
| 1.20 | 135.106915 | 3.039165 | 2669.036998 | 47.305081 |
| gotip | 130.972365 | 2.769622 | 5239.464818 | 131.695244 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 191.596144 | 15.624752 | 1414.231780 | 24.445320 |
| 1.19 | 163.115089 | 7.632089 | 1430.341975 | 18.875038 |
| 1.20 | 163.295220 | 5.188267 | 1416.005359 | 20.565713 |
| gotip | 160.536743 | 3.983550 | 1406.766474 | 25.835700 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 193.251927 | 6.219711 | 2970.131739 | 22.874455 |
| 1.19 | 164.925477 | 17.621360 | 2834.539663 | 29.372505 |
| 1.20 | 173.420294 | 4.923201 | 2791.558723 | 33.023733 |
| gotip | 167.418015 | 5.733763 | 2669.160420 | 18.392309 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 221.220306 | 4.666801 | 6875.077851 | 151.112797 |
| 1.19 | 192.799856 | 21.701851 | 3968.758203 | 66.266209 |
| 1.20 | 193.652193 | 7.695613 | 3879.535504 | 114.623271 |
| gotip | 195.123985 | 7.619469 | 3230.586623 | 30.470221 |

![switch_case](./switch_case__725e73000e.png)

