# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.439000

CacheSize: 49152

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.439000

CacheSize: 49152

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 465.709489 | 12.774847 | 4020.280268 | 3.133260 |
| 1.18 | 474.745497 | 12.058906 | 3615.785075 | 2.158153 |
| 1.19 | 518.756358 | 8.415238 | 3611.757164 | 1.767905 |
| gotip | 524.785571 | 5.072768 | 3596.514755 | 0.823414 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 127.408072 | 9.417513 | 4791.067581 | 1.263055 |
| 1.18 | 126.901606 | 1.758190 | 4795.383804 | 0.697202 |
| 1.19 | 138.845496 | 1.955729 | 4791.712080 | 1.908821 |
| gotip | 139.651492 | 7.557099 | 4795.368094 | 1.695412 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 126.882606 | 3.525596 | 5226.019561 | 52.849425 |
| 1.18 | 134.594476 | 1.716025 | 5179.449337 | 25.536152 |
| 1.19 | 135.333731 | 3.285892 | 5180.121389 | 21.261824 |
| gotip | 141.823252 | 9.832992 | 6574.925439 | 17.796004 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 131.859762 | 4.607911 | 9667.542083 | 35.670834 |
| 1.18 | 138.804642 | 4.324983 | 5569.322079 | 32.460208 |
| 1.19 | 138.669231 | 1.316268 | 9695.095336 | 59.841432 |
| gotip | 142.852232 | 4.342619 | 5709.191835 | 83.744338 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 125.261107 | 2.707370 | 5818.335060 | 45.925471 |
| 1.18 | 130.413322 | 3.689137 | 5539.576396 | 24.676363 |
| 1.19 | 140.775946 | 4.511790 | 7063.757244 | 37.975768 |
| gotip | 145.357566 | 2.415774 | 7890.928587 | 90.514649 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 141.096606 | 2.868904 | 6096.359871 | 2.940653 |
| 1.18 | 144.728904 | 4.503948 | 5974.376448 | 5.843090 |
| 1.19 | 147.740613 | 5.122167 | 5923.348254 | 2.521407 |
| gotip | 150.682988 | 3.276488 | 5929.239768 | 11.080939 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 128.998804 | 2.183057 | 3703.617164 | 5.315416 |
| 1.18 | 135.989076 | 2.682893 | 3540.211658 | 1.268481 |
| 1.19 | 140.635631 | 4.440676 | 360.440883 | 0.548821 |
| gotip | 143.945042 | 3.375807 | 400.630432 | 2.324601 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 146.397145 | 6.132679 | 4517.405305 | 594.802873 |
| 1.18 | 155.260344 | 58.729915 | 5134.919503 | 913.826375 |
| 1.19 | 154.379221 | 80.026260 | 4998.936729 | 361.436552 |
| gotip | 165.040599 | 48.567760 | 5223.440493 | 180.330235 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 123.113344 | 37.905350 | 3670.221429 | 13.402132 |
| 1.18 | 126.853403 | 4.278446 | 5218.596639 | 1.176734 |
| 1.19 | 139.876496 | 3.286065 | 2706.428273 | 1.620295 |
| gotip | 141.441345 | 2.952810 | 5220.188105 | 1.472734 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 155.707256 | 9.029327 | 1060.943345 | 1.374096 |
| 1.18 | 161.526233 | 8.475300 | 1103.383164 | 3.090566 |
| 1.19 | 167.678454 | 7.943021 | 1132.560332 | 8.206702 |
| gotip | 168.865014 | 8.161341 | 1078.034202 | 4.701638 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 158.963670 | 1.301944 | 2345.567643 | 61.153227 |
| 1.18 | 166.595348 | 4.591108 | 2314.409976 | 11.613263 |
| 1.19 | 166.264955 | 3.522696 | 2206.487320 | 22.706416 |
| gotip | 173.206985 | 3.543364 | 2174.829425 | 59.517778 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 188.756907 | 11.027128 | 5935.170910 | 13.638380 |
| 1.18 | 193.177217 | 9.233312 | 5937.458174 | 30.395462 |
| 1.19 | 197.172123 | 11.470435 | 3557.776794 | 4.943717 |
| gotip | 202.822329 | 9.334895 | 3415.208693 | 2.518322 |

![switch_case](./switch_case__725e73000e.png)

