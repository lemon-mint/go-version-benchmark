# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.2

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz

Cores: 1

Mhz: 2294.684000

CacheSize: 51200

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz

Cores: 1

Mhz: 2294.684000

CacheSize: 51200

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 585.342855 | 32.073794 | 4511.579198 | 186.895340 |
| 1.18 | 576.399541 | 31.024594 | 4126.350553 | 123.142621 |
| 1.19 | 625.545045 | 26.266802 | 3987.417874 | 122.592063 |
| gotip | 4910.366109 | 110.478412 | 4219.021469 | 168.181241 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 171.473534 | 8.652678 | 4451.890691 | 134.073727 |
| 1.18 | 166.124766 | 19.269201 | 4426.293915 | 145.309653 |
| 1.19 | 188.227273 | 8.928960 | 4475.597096 | 93.065832 |
| gotip | 3862.336928 | 110.973400 | 4703.425735 | 66.840252 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 168.472045 | 17.772251 | 7195.645318 | 60.229351 |
| 1.18 | 188.139603 | 9.549287 | 7198.527286 | 135.315205 |
| 1.19 | 187.033024 | 10.455716 | 7147.275089 | 79.429082 |
| gotip | 3762.110007 | 142.212917 | 7738.868844 | 210.018357 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 177.696968 | 8.864861 | 6143.342707 | 163.964255 |
| 1.18 | 188.656624 | 10.958008 | 5926.432747 | 296.552818 |
| 1.19 | 192.870220 | 10.664054 | 6026.850226 | 194.489670 |
| gotip | 3730.794246 | 76.927625 | 6115.944472 | 191.767628 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 158.149529 | 8.329332 | 7637.083557 | 200.422949 |
| 1.18 | 179.464581 | 10.854196 | 7350.875686 | 129.860070 |
| 1.19 | 177.903495 | 10.626238 | 9494.771727 | 193.013751 |
| gotip | 3591.168360 | 117.781555 | 10872.386948 | 297.052403 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 191.188553 | 15.356845 | 5969.888882 | 117.390232 |
| 1.18 | 204.957737 | 5.511032 | 5684.263369 | 189.762243 |
| 1.19 | 189.126419 | 12.983674 | 5676.856190 | 183.571285 |
| gotip | 4340.166979 | 132.198933 | 5704.405016 | 173.054344 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 177.856551 | 8.751838 | 4551.349605 | 92.716202 |
| 1.18 | 196.906241 | 10.435662 | 4364.827779 | 209.222036 |
| 1.19 | 168.380228 | 9.604451 | 423.916309 | 17.847384 |
| gotip | 3951.970836 | 125.715772 | 419.203296 | 16.465612 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 192.067270 | 133.431916 | 6149.936763 | 795.471273 |
| 1.18 | 205.110577 | 117.796770 | 5115.533506 | 849.986548 |
| 1.19 | 199.985927 | 81.492726 | 5291.614086 | 879.555700 |
| gotip | 5131.861623 | 96.163464 | 7021.963884 | 778.555604 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 169.832380 | 9.914332 | 3962.029943 | 110.439250 |
| 1.18 | 170.930059 | 8.149441 | 2792.228499 | 80.165189 |
| 1.19 | 189.255153 | 8.106871 | 2766.167922 | 74.754384 |
| gotip | 3794.266761 | 125.103079 | 2783.083381 | 94.840265 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 191.139574 | 42.140004 | 1287.631220 | 43.157449 |
| 1.18 | 204.456518 | 22.085874 | 1334.306934 | 68.425030 |
| 1.19 | 212.281040 | 13.863891 | 1436.448005 | 77.227822 |
| gotip | 5095.087612 | 131.583782 | 1421.730442 | 80.580540 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 222.365218 | 15.472276 | 2991.456985 | 75.103529 |
| 1.18 | 224.232180 | 14.838213 | 2930.335887 | 110.050398 |
| 1.19 | 217.833599 | 10.494619 | 2617.678305 | 110.331493 |
| gotip | 5320.106470 | 158.756170 | 2752.962403 | 80.998774 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 245.184835 | 24.680292 | 6695.640965 | 183.827623 |
| 1.18 | 235.104778 | 30.088261 | 6671.506037 | 254.396719 |
| 1.19 | 250.409517 | 16.064058 | 3927.259550 | 153.774377 |
| gotip | 6874.483243 | 131.442262 | 3606.786203 | 133.953980 |

![switch_case](./switch_case__725e73000e.png)

