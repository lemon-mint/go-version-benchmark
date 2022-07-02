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

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 157.685239 | 12.159274 | 1073.622914 | 30.094233 |
| 1.18 | 171.222941 | 10.661107 | 1068.999192 | 47.918840 |
| 1.19beta1 | 284.264481 | 6.190756 | 1094.572334 | 31.119427 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 184.433933 | 19.648590 | 4586.629369 | 8.726257 |
| 1.18 | 195.948617 | 15.956745 | 5045.122380 | 337.211192 |
| 1.19beta1 | 324.225731 | 17.687850 | 2886.540610 | 155.324352 |

![switch_case](./switch_case__725e73000e.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 159.254671 | 6.138786 | 2145.326950 | 71.230904 |
| 1.18 | 170.640170 | 3.452257 | 2127.887838 | 29.924489 |
| 1.19beta1 | 311.146710 | 4.264449 | 2185.901767 | 48.112095 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 146.495710 | 3.445449 | 4437.500670 | 288.191240 |
| 1.18 | 153.108250 | 3.618312 | 4446.879859 | 291.073265 |
| 1.19beta1 | 272.152416 | 5.807848 | 4437.793615 | 175.114711 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 128.965312 | 3.415491 | 3280.732516 | 176.780995 |
| 1.18 | 130.784697 | 10.141244 | 3152.299866 | 10.925121 |
| 1.19beta1 | 257.151643 | 5.125740 | 325.809398 | 0.881236 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 146.320156 | 2.686933 | 4856.234210 | 479.039168 |
| 1.18 | 151.997422 | 118.341087 | 4308.172508 | 550.964742 |
| 1.19beta1 | 290.514781 | 25.937871 | 5174.301436 | 264.628955 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 128.249318 | 4.085351 | 3159.141840 | 150.193013 |
| 1.18 | 138.476986 | 5.283299 | 3919.231607 | 224.690586 |
| 1.19beta1 | 264.665916 | 6.636673 | 3838.887581 | 1.423564 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 121.589463 | 10.306477 | 3443.110492 | 140.720625 |
| 1.18 | 129.175790 | 3.826242 | 3543.239396 | 145.782666 |
| 1.19beta1 | 262.555654 | 8.501618 | 3536.598761 | 142.179020 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 126.790264 | 3.824262 | 5971.371189 | 89.751276 |
| 1.18 | 137.222800 | 3.011121 | 6099.486314 | 100.361356 |
| 1.19beta1 | 282.004786 | 7.037221 | 7664.124298 | 151.501784 |

![MergeSort](./MergeSort__619024e898.png)

