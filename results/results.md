# Benchmarks

## Environment

NumCPU: 2
Arch: amd64
OS: linux
Version: go1.18.3
Itercount: 10
### CPU 0

Model: Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHz
Cores: 1
Mhz: 2095.195000
CacheSize: 36608
Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHz
Cores: 1
Mhz: 2095.195000
CacheSize: 36608
Microcode: 0xffffffff

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 141.553154 | 2.769149 | 3776.955514 | 33.687471 |
| 1.18 | 152.840068 | 5.221140 | 5220.650002 | 2.403427 |
| 1.19beta1 | 314.559811 | 7.641251 | 5216.332502 | 2.306407 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 138.904246 | 5.177669 | 4678.649234 | 3.590854 |
| 1.18 | 148.628546 | 5.143435 | 4812.085052 | 5.351677 |
| 1.19beta1 | 316.033606 | 9.468483 | 4811.930056 | 2.472726 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 145.757404 | 9.651345 | 6970.884937 | 129.270235 |
| 1.18 | 151.714944 | 2.555399 | 6667.464937 | 28.846891 |
| 1.19beta1 | 321.881269 | 4.814185 | 8323.964600 | 95.085254 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## regexp (Compile)

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 188.109441 | 3.063800 | 1305.700509 | 5.014184 |
| 1.18 | 188.980199 | 5.253778 | 1370.634546 | 7.948662 |
| 1.19beta1 | 351.592278 | 4.281332 | 1381.363477 | 8.937739 |

![regexp (Compile)](./ff7ab1cc39d9f8604be0c37e0bdeb4c7cf02cf9cc420d7410e411ce3835d9a42.png)

## Sort (Random)

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 167.231211 | 3.537029 | 6014.870011 | 2.895976 |
| 1.18 | 172.445561 | 11.281116 | 6026.083722 | 4.239135 |
| 1.19beta1 | 329.176706 | 8.925725 | 6033.752869 | 3.046905 |

![Sort (Random)](./3c1d173f078fb01a5525f5ae6f8bce14fbf8318a5b39da592c1847510b1fa8f1.png)

## Sort (Reversed)

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 189.578656 | 22.425505 | 4451.800763 | 3.859409 |
| 1.18 | 159.831327 | 2.763941 | 4290.532682 | 3.137230 |
| 1.19beta1 | 321.109706 | 8.160241 | 443.689702 | 1.273763 |

![Sort (Reversed)](./a0f4a74015cf54e2efe82a24aa116a48b4f983e1fc0126e0d3e84f2560003ee7.png)

## switch-case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 225.642793 | 3.890284 | 6235.729124 | 4.850952 |
| 1.18 | 226.949578 | 16.247222 | 6865.329798 | 1.361651 |
| 1.19beta1 | 383.623663 | 5.411964 | 3925.325838 | 3.256768 |

![switch-case](./1af1469d75e77ed39c58041d45b37b329137876f59fb4c03529ebb65c78b40aa.png)

