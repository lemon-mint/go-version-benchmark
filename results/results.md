# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.18.3

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

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 146.800268 | 2.161788 | 4913.203768 | 545.387115 |
| 1.18 | 157.856622 | 71.597170 | 4607.960577 | 398.787858 |
| 1.19beta1 | 288.725663 | 8.958865 | 5133.946073 | 348.972071 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 126.756162 | 11.207484 | 3901.759424 | 1.442481 |
| 1.18 | 128.632135 | 4.272147 | 4013.639673 | 3.611185 |
| 1.19beta1 | 271.038715 | 4.117397 | 4011.567961 | 15.459233 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 135.047832 | 2.553887 | 3715.263544 | 2.448485 |
| 1.18 | 139.239710 | 10.326546 | 3575.258080 | 1.293835 |
| 1.19beta1 | 274.830580 | 6.504168 | 369.585051 | 0.573944 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 197.437035 | 12.357713 | 5201.140632 | 2.508699 |
| 1.18 | 200.566771 | 8.378045 | 5723.368131 | 1.402876 |
| 1.19beta1 | 330.974082 | 8.303880 | 3272.800675 | 1.244038 |

![switch_case](./switch_case__725e73000e.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 127.690708 | 3.667346 | 3174.018969 | 39.412687 |
| 1.18 | 134.858877 | 1.677155 | 4348.140901 | 2.180060 |
| 1.19beta1 | 270.917807 | 4.881543 | 4347.322423 | 1.760121 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 128.446076 | 4.452151 | 5842.407361 | 46.982998 |
| 1.18 | 134.744392 | 2.299700 | 5604.275989 | 46.505621 |
| 1.19beta1 | 275.108702 | 3.486610 | 6955.565737 | 46.269748 |

![MergeSort](./MergeSort__619024e898.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 160.553492 | 12.376319 | 1087.868645 | 2.229155 |
| 1.18 | 165.159197 | 4.279996 | 1142.260955 | 8.128395 |
| 1.19beta1 | 298.972248 | 4.325553 | 1150.298458 | 6.932044 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 163.495369 | 1.643247 | 2345.282248 | 7.727896 |
| 1.18 | 173.430812 | 4.961032 | 2350.875943 | 7.048934 |
| 1.19beta1 | 302.479617 | 3.177314 | 2247.183250 | 11.352427 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 144.264508 | 3.601754 | 5011.505618 | 2.497874 |
| 1.18 | 151.280075 | 4.621358 | 5028.281542 | 3.120270 |
| 1.19beta1 | 283.921497 | 4.416669 | 5022.126820 | 1.621632 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

