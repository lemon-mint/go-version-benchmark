# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.18.3

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) CPU E5-2673 v3 @ 2.40GHz

Cores: 1

Mhz: 2394.456000

CacheSize: 30720

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) CPU E5-2673 v3 @ 2.40GHz

Cores: 1

Mhz: 2394.456000

CacheSize: 30720

Microcode: 0xffffffff

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 177.565928 | 5.109513 | 6373.677969 | 724.146649 |
| 1.18 | 179.839877 | 113.871065 | 4037.539048 | 338.298110 |
| 1.19beta1 | 347.524490 | 530.854595 | 4431.577541 | 290.074222 |
| gotip | 208.822380 | 8.433650 | 5011.958180 | 568.573615 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 195.116674 | 15.321537 | 1288.257616 | 16.508609 |
| 1.18 | 198.265456 | 15.269023 | 1359.414908 | 21.927456 |
| 1.19beta1 | 380.226955 | 26.789990 | 1394.547948 | 28.765417 |
| gotip | 221.136031 | 7.354457 | 1372.890373 | 26.266185 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 205.345619 | 7.880162 | 2940.163183 | 30.558241 |
| 1.18 | 202.659605 | 4.778487 | 2874.584630 | 20.686277 |
| 1.19beta1 | 362.254850 | 5.689631 | 2696.912666 | 47.719686 |
| gotip | 223.365536 | 15.091144 | 2659.843987 | 18.243743 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 227.558427 | 17.646585 | 6606.683672 | 80.314834 |
| 1.18 | 236.505848 | 31.063222 | 6601.741342 | 58.379684 |
| 1.19beta1 | 392.415818 | 10.506576 | 3822.184877 | 62.459770 |
| gotip | 263.515824 | 4.465633 | 3800.275785 | 23.498212 |

![switch_case](./switch_case__725e73000e.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 144.985497 | 3.881611 | 3701.260965 | 36.710944 |
| 1.18 | 164.742061 | 5.835767 | 2601.220112 | 41.139408 |
| 1.19beta1 | 337.291853 | 11.620972 | 2574.430173 | 46.738297 |
| gotip | 183.014536 | 4.534882 | 2524.237839 | 17.156697 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 145.751781 | 3.199424 | 4593.521016 | 72.885487 |
| 1.18 | 156.213294 | 5.828621 | 4591.339423 | 32.743321 |
| 1.19beta1 | 331.424074 | 10.580317 | 4659.473721 | 71.017954 |
| gotip | 181.506849 | 6.585229 | 4613.997144 | 43.262631 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 154.945359 | 4.986349 | 7263.878784 | 114.934669 |
| 1.18 | 157.048914 | 5.271421 | 6844.710339 | 92.537177 |
| 1.19beta1 | 340.732093 | 9.647568 | 8679.522302 | 89.804753 |
| gotip | 200.000543 | 14.723096 | 8621.415708 | 138.229645 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 166.778668 | 6.194060 | 5593.404691 | 72.197866 |
| 1.18 | 176.953307 | 10.087248 | 5607.552326 | 36.341084 |
| 1.19beta1 | 337.476465 | 6.838689 | 5727.976046 | 59.142035 |
| gotip | 199.301126 | 4.960109 | 5728.303980 | 78.824647 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 151.165475 | 2.931600 | 4330.676245 | 69.111058 |
| 1.18 | 164.179807 | 5.100553 | 4305.534497 | 26.631203 |
| 1.19beta1 | 329.697098 | 21.877010 | 423.193615 | 8.191201 |
| gotip | 182.571306 | 5.889571 | 422.750258 | 7.176877 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

