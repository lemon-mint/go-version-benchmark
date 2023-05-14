# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.9

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz

Cores: 1

Mhz: 2593.907000

CacheSize: 36608

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz

Cores: 1

Mhz: 2593.907000

CacheSize: 36608

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 471.345437 | 7.660799 | 3518.965395 | 2.363569 |
| 1.19 | 502.176263 | 8.834683 | 3543.629885 | 3.284206 |
| 1.20 | 427.058323 | 5.225394 | 3555.713178 | 0.712156 |
| gotip | 468.481902 | 5.869387 | 3580.387466 | 8.100333 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 117.282673 | 2.633387 | 4013.487126 | 1.654344 |
| 1.19 | 115.491588 | 6.855483 | 3903.630194 | 1.224673 |
| 1.20 | 120.392101 | 3.516919 | 3925.108199 | 7.366462 |
| gotip | 118.278325 | 4.564841 | 4246.167879 | 9.119874 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 130.988090 | 3.238105 | 5709.887193 | 30.315510 |
| 1.19 | 116.470968 | 15.100674 | 5556.967118 | 20.801076 |
| 1.20 | 119.724591 | 3.069777 | 6207.468348 | 22.350957 |
| gotip | 116.693377 | 4.309920 | 6574.332224 | 17.179649 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 136.820812 | 3.434751 | 5438.050591 | 40.282751 |
| 1.19 | 120.186475 | 8.222848 | 9100.587865 | 48.455347 |
| 1.20 | 123.118562 | 4.332774 | 9224.962923 | 23.210627 |
| gotip | 120.993776 | 3.767791 | 5619.426073 | 64.447024 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 126.009743 | 4.578091 | 5958.250941 | 33.093344 |
| 1.19 | 118.642847 | 5.521274 | 7678.942204 | 68.031544 |
| 1.20 | 120.385447 | 6.102607 | 8795.917211 | 85.762928 |
| gotip | 120.425554 | 3.077916 | 9085.160992 | 122.542104 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 146.540088 | 5.764889 | 5030.968789 | 2.158608 |
| 1.19 | 124.682251 | 14.044864 | 5085.851572 | 6.057691 |
| 1.20 | 131.007636 | 3.296076 | 5080.937380 | 0.924293 |
| gotip | 134.045448 | 3.935393 | 5203.431735 | 1.343189 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 134.279969 | 3.377626 | 3576.463172 | 2.138447 |
| 1.19 | 117.828889 | 6.077625 | 369.990019 | 0.596535 |
| 1.20 | 123.140305 | 3.396227 | 369.944455 | 0.419433 |
| gotip | 119.760827 | 4.079885 | 374.598956 | 0.452594 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 145.537089 | 4.184278 | 5034.184186 | 611.001522 |
| 1.19 | 129.680808 | 29.492374 | 4783.722111 | 322.574311 |
| 1.20 | 135.559190 | 3.285137 | 5349.101140 | 214.953034 |
| gotip | 133.863621 | 2.857110 | 5365.895265 | 311.002039 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 127.161392 | 4.416032 | 4349.005485 | 1.378123 |
| 1.19 | 115.724274 | 6.866506 | 2193.002350 | 1.143298 |
| 1.20 | 119.766308 | 2.811720 | 2193.831477 | 1.689443 |
| gotip | 117.526226 | 4.059777 | 4347.982246 | 1.025575 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 159.238883 | 17.837250 | 1154.857099 | 8.138323 |
| 1.19 | 142.277698 | 6.191683 | 1177.607401 | 11.627626 |
| 1.20 | 145.959707 | 6.583501 | 1167.398839 | 13.199773 |
| gotip | 140.531320 | 4.398618 | 1150.174620 | 8.293364 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 162.950241 | 22.780210 | 2367.837036 | 15.269804 |
| 1.19 | 145.508045 | 4.958299 | 2252.273638 | 11.772839 |
| 1.20 | 150.894228 | 2.970945 | 2219.019486 | 12.076826 |
| gotip | 147.694639 | 4.317344 | 2134.760002 | 10.493231 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 192.605651 | 5.997397 | 5724.889728 | 1.144358 |
| 1.19 | 171.041259 | 8.362049 | 3223.733369 | 1.357850 |
| 1.20 | 175.297068 | 5.599990 | 3225.311613 | 1.871825 |
| gotip | 172.258747 | 5.698485 | 2765.055973 | 2.572262 |

![switch_case](./switch_case__725e73000e.png)

