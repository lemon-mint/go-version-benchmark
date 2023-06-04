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
| 1.18 | 474.557984 | 10.702221 | 3517.364170 | 3.087365 |
| 1.19 | 500.919518 | 5.585612 | 3541.438780 | 1.970555 |
| 1.20 | 426.397294 | 7.285657 | 3554.286128 | 1.453099 |
| gotip | 464.575249 | 5.873222 | 3538.566287 | 43.744024 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 122.335025 | 2.865475 | 4013.472088 | 3.875770 |
| 1.19 | 112.258752 | 12.287510 | 3902.861025 | 2.032439 |
| 1.20 | 114.015896 | 3.018231 | 3912.203391 | 7.962482 |
| gotip | 118.518337 | 2.932535 | 4251.556409 | 1.824029 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 129.985433 | 6.145030 | 5710.895316 | 16.811936 |
| 1.19 | 116.794199 | 3.563245 | 5561.464536 | 37.519540 |
| 1.20 | 122.913298 | 3.081897 | 6221.540871 | 20.649533 |
| gotip | 117.338820 | 0.982731 | 6542.374891 | 20.999727 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 134.909008 | 3.024966 | 5301.773866 | 66.733254 |
| 1.19 | 118.085679 | 6.013570 | 9000.738502 | 45.896814 |
| 1.20 | 120.482998 | 2.571847 | 9063.768649 | 47.245238 |
| gotip | 119.348962 | 2.674335 | 5508.623729 | 34.923992 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 122.337477 | 4.537028 | 5843.325582 | 55.481569 |
| 1.19 | 116.377245 | 3.676618 | 7516.246582 | 42.648218 |
| 1.20 | 119.710715 | 6.166241 | 8710.292838 | 41.990794 |
| gotip | 119.045045 | 5.329743 | 8950.654528 | 66.844832 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 160.473021 | 3.729260 | 5038.713395 | 4.454602 |
| 1.19 | 143.758593 | 22.886555 | 5092.046231 | 6.825769 |
| 1.20 | 140.118632 | 4.916328 | 5080.820655 | 11.126608 |
| gotip | 134.927529 | 4.903282 | 5094.864608 | 1.831789 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 133.155641 | 2.666522 | 3574.921514 | 2.135916 |
| 1.19 | 117.906324 | 3.033752 | 370.890537 | 0.832604 |
| 1.20 | 123.015792 | 4.517363 | 370.241403 | 1.005482 |
| gotip | 121.082151 | 2.459455 | 374.399301 | 0.503608 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 150.345557 | 2.707327 | 4379.582932 | 612.835694 |
| 1.19 | 129.779835 | 19.832507 | 4876.397273 | 462.725915 |
| 1.20 | 135.465214 | 3.996338 | 5459.921903 | 330.479954 |
| gotip | 131.454567 | 6.705344 | 5557.196491 | 207.560562 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 125.618935 | 2.512644 | 4348.296086 | 4.992372 |
| 1.19 | 111.513035 | 5.379244 | 2191.839070 | 0.687266 |
| 1.20 | 117.871327 | 4.328064 | 2191.121677 | 0.837521 |
| gotip | 116.963775 | 3.450447 | 2191.345200 | 0.755475 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 157.289107 | 13.258118 | 1151.281635 | 13.846023 |
| 1.19 | 141.074510 | 6.728394 | 1160.818739 | 9.928100 |
| 1.20 | 144.536425 | 6.320040 | 1163.490952 | 14.457938 |
| gotip | 143.834809 | 3.560089 | 1121.492751 | 14.215008 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 163.571901 | 9.221296 | 2364.029545 | 10.696755 |
| 1.19 | 145.595493 | 2.978831 | 2254.247854 | 8.210143 |
| 1.20 | 150.771398 | 7.090434 | 2240.201214 | 11.067463 |
| gotip | 155.283030 | 4.356709 | 2189.233097 | 13.679777 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 191.402168 | 11.351701 | 5723.835203 | 1.859888 |
| 1.19 | 166.317336 | 8.938365 | 3223.404897 | 4.662375 |
| 1.20 | 171.408042 | 2.973654 | 3224.916997 | 1.155627 |
| gotip | 173.276072 | 5.651203 | 2679.972762 | 1.994318 |

![switch_case](./switch_case__725e73000e.png)

