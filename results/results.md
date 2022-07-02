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

Mhz: 2593.907000

CacheSize: 36608

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz

Cores: 1

Mhz: 2593.907000

CacheSize: 36608

Microcode: 0xffffffff

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 142.690403 | 2.089697 | 4720.349678 | 246.351134 |
| 1.18 | 157.610028 | 13.176282 | 4765.803809 | 486.842349 |
| 1.19beta1 | 291.367874 | 20.304494 | 4904.046419 | 391.945495 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 133.765127 | 2.971553 | 6147.519602 | 104.192230 |
| 1.18 | 142.203131 | 3.721546 | 6050.137280 | 41.129776 |
| 1.19beta1 | 286.743335 | 4.330032 | 7593.299778 | 86.897196 |

![MergeSort](./MergeSort__619024e898.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 171.622746 | 4.074518 | 2372.486056 | 10.359176 |
| 1.18 | 177.109617 | 3.188711 | 2379.349317 | 13.778001 |
| 1.19beta1 | 321.182991 | 5.397452 | 2259.509575 | 9.801176 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 149.416295 | 4.060500 | 5012.791584 | 1.149166 |
| 1.18 | 155.585667 | 4.408142 | 5030.925087 | 1.925187 |
| 1.19beta1 | 304.667278 | 5.921368 | 5023.588708 | 2.412395 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 137.617427 | 3.097616 | 3715.412055 | 2.348508 |
| 1.18 | 146.289327 | 16.968380 | 3575.811291 | 4.404916 |
| 1.19beta1 | 287.164740 | 2.879353 | 369.743709 | 0.345714 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 125.392887 | 2.726568 | 3196.326248 | 56.033740 |
| 1.18 | 139.167413 | 7.578207 | 4347.558912 | 1.320503 |
| 1.19beta1 | 286.328866 | 4.524891 | 4347.861203 | 1.575064 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 131.219131 | 4.872041 | 3902.639237 | 0.834604 |
| 1.18 | 134.851322 | 4.667338 | 4013.840009 | 7.297932 |
| 1.19beta1 | 282.515520 | 3.773813 | 4013.172282 | 2.354937 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 167.287585 | 19.129696 | 1105.283802 | 3.810773 |
| 1.18 | 174.136632 | 14.051893 | 1153.866165 | 6.700718 |
| 1.19beta1 | 313.708433 | 3.518152 | 1154.761333 | 4.172652 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 204.153418 | 16.201103 | 5203.756999 | 2.757813 |
| 1.18 | 206.040559 | 20.542967 | 5724.438445 | 2.266548 |
| 1.19beta1 | 358.017814 | 4.975278 | 3273.048084 | 1.816070 |

![switch_case](./switch_case__725e73000e.png)

