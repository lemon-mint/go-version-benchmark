# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.18.3

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.437000

CacheSize: 49152

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.437000

CacheSize: 49152

Microcode: 0xffffffff

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 117.848806 | 6.736945 | 3656.947491 | 9.175939 |
| 1.18 | 125.031518 | 3.262294 | 5219.233415 | 1.009268 |
| 1.19beta1 | 262.709006 | 2.289444 | 5219.855661 | 0.915739 |
| gotip | 145.904663 | 2.899820 | 2704.482780 | 0.409017 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 157.458015 | 8.442088 | 1060.544959 | 3.548667 |
| 1.18 | 158.920365 | 8.674039 | 1101.899078 | 10.849152 |
| 1.19beta1 | 293.145978 | 3.202766 | 1120.531293 | 13.702598 |
| gotip | 176.992313 | 3.698333 | 1126.632083 | 5.904077 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 187.125615 | 8.195869 | 5962.757761 | 54.480292 |
| 1.18 | 191.388369 | 10.030079 | 5896.134868 | 32.333231 |
| 1.19beta1 | 323.001039 | 6.581159 | 3417.742771 | 3.566998 |
| gotip | 216.114983 | 15.186141 | 3561.526268 | 4.444079 |

![switch_case](./switch_case__725e73000e.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 142.083007 | 2.066305 | 4742.507298 | 497.450201 |
| 1.18 | 147.661406 | 62.471831 | 5305.572759 | 1056.004992 |
| 1.19beta1 | 279.743035 | 198.246662 | 5112.649028 | 337.192558 |
| gotip | 163.141719 | 16.356020 | 4777.595331 | 566.975333 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 117.128660 | 1.670249 | 4790.548136 | 1.022476 |
| 1.18 | 123.867822 | 5.604552 | 4794.913028 | 0.505745 |
| 1.19beta1 | 262.044048 | 10.757957 | 4794.668032 | 1.162421 |
| gotip | 144.792839 | 3.315145 | 4790.429826 | 1.037933 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 123.500253 | 3.555972 | 5751.302258 | 40.141291 |
| 1.18 | 126.558300 | 2.133927 | 5455.899043 | 44.763267 |
| 1.19beta1 | 268.372593 | 4.981671 | 6760.188507 | 28.209478 |
| gotip | 152.376826 | 9.683327 | 7037.661831 | 37.205456 |

![MergeSort](./MergeSort__619024e898.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 158.266110 | 1.584226 | 2319.775131 | 26.177344 |
| 1.18 | 163.054167 | 2.465304 | 2335.054874 | 32.308189 |
| 1.19beta1 | 296.484695 | 3.819637 | 2179.319251 | 40.408226 |
| gotip | 180.590721 | 5.297800 | 2174.553966 | 17.797999 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 137.799227 | 1.695273 | 6069.005188 | 5.823577 |
| 1.18 | 144.683963 | 2.614682 | 5976.448828 | 5.453566 |
| 1.19beta1 | 276.203074 | 2.187505 | 5930.762195 | 4.999192 |
| gotip | 158.863209 | 2.318420 | 5928.071882 | 6.790533 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 125.696622 | 2.760281 | 3643.207594 | 17.913843 |
| 1.18 | 133.449213 | 1.458070 | 3541.095008 | 2.207148 |
| 1.19beta1 | 266.945517 | 3.858750 | 395.361826 | 1.632712 |
| gotip | 147.868860 | 1.879794 | 360.494976 | 0.267605 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

