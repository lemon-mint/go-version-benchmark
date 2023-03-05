# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.6

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz

Cores: 1

Mhz: 2593.905000

CacheSize: 36608

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz

Cores: 1

Mhz: 2593.905000

CacheSize: 36608

Microcode: 0xffffffff

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 483.913020 | 9.019863 | 3518.201807 | 8.912027 |
| 1.19 | 500.223063 | 8.849569 | 3542.845548 | 1.765307 |
| 1.20 | 426.678137 | 4.102993 | 3555.231459 | 1.052180 |
| gotip | 465.332036 | 3.141677 | 3554.155886 | 0.966723 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 118.467891 | 8.101434 | 4012.430191 | 1.366543 |
| 1.19 | 111.694874 | 9.352774 | 3903.292465 | 1.779203 |
| 1.20 | 117.695354 | 3.927670 | 3909.461773 | 4.997579 |
| gotip | 121.561126 | 3.088550 | 4246.630181 | 2.139774 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 129.633273 | 4.331700 | 5741.432110 | 125.023347 |
| 1.19 | 114.802136 | 12.315609 | 5556.150696 | 33.693035 |
| 1.20 | 119.026042 | 3.527976 | 6244.316282 | 20.102467 |
| gotip | 121.806825 | 4.794333 | 6394.609583 | 19.029882 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 133.505579 | 3.858013 | 5388.159892 | 52.336807 |
| 1.19 | 113.874212 | 3.151696 | 9077.491035 | 34.052079 |
| 1.20 | 122.691914 | 2.567937 | 9103.683264 | 53.758502 |
| gotip | 123.368694 | 4.737808 | 5465.954491 | 64.180543 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 124.691786 | 2.050977 | 5995.591605 | 78.541964 |
| 1.19 | 116.217448 | 2.929759 | 7636.384727 | 116.312438 |
| 1.20 | 122.088738 | 3.674153 | 8794.877557 | 44.813403 |
| gotip | 125.591269 | 3.041037 | 8916.501056 | 107.487460 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 142.645519 | 10.890260 | 5028.795471 | 1.718708 |
| 1.19 | 120.538131 | 3.522537 | 5084.603992 | 2.345389 |
| 1.20 | 134.316530 | 3.762150 | 5079.766997 | 1.813998 |
| gotip | 131.677474 | 3.501184 | 5064.384018 | 2.399189 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 132.455064 | 5.024169 | 3574.000611 | 2.744895 |
| 1.19 | 112.719863 | 4.047020 | 369.677739 | 0.483123 |
| 1.20 | 121.071729 | 3.513502 | 370.432261 | 0.685213 |
| gotip | 124.964552 | 3.625264 | 374.764344 | 0.414533 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 150.230979 | 4.066274 | 4463.976754 | 493.259815 |
| 1.19 | 125.037533 | 3.907125 | 5056.150212 | 464.239478 |
| 1.20 | 136.406372 | 3.968772 | 5459.077651 | 270.888075 |
| gotip | 140.472111 | 6.149593 | 5653.433596 | 354.980107 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 123.571977 | 4.231264 | 4348.503458 | 1.458518 |
| 1.19 | 112.935305 | 2.671132 | 2192.403955 | 0.938453 |
| 1.20 | 120.642645 | 3.769394 | 2191.606028 | 1.268772 |
| gotip | 123.421809 | 3.728191 | 4348.766856 | 1.670135 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 156.818308 | 4.130668 | 1155.252757 | 13.691115 |
| 1.19 | 135.929734 | 8.911152 | 1175.734143 | 6.389139 |
| 1.20 | 141.648422 | 6.812845 | 1170.808481 | 7.687605 |
| gotip | 145.091067 | 3.976296 | 1157.002363 | 8.177479 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 165.423185 | 5.554483 | 2376.043978 | 10.969028 |
| 1.19 | 143.109685 | 3.512990 | 2241.585084 | 3.790785 |
| 1.20 | 147.835944 | 3.965590 | 2222.752992 | 18.013580 |
| gotip | 152.956800 | 2.509602 | 2229.079687 | 10.998431 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 199.428601 | 21.240409 | 5725.319940 | 1.359280 |
| 1.19 | 167.981103 | 9.638111 | 3224.500326 | 1.241702 |
| 1.20 | 173.629526 | 5.144558 | 3224.711159 | 2.209193 |
| gotip | 175.577123 | 4.965019 | 2759.530612 | 1.005257 |

![switch_case](./switch_case__725e73000e.png)

