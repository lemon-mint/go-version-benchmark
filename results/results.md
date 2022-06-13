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
Mhz: 2095.077000
CacheSize: 36608
Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHz
Cores: 1
Mhz: 2095.077000
CacheSize: 36608
Microcode: 0xffffffff

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 155.589489 | 12.575483 | 7263.325654 | 59.605718 |
| 1.18 | 155.513370 | 4.571362 | 7331.540828 | 254.663665 |
| 1.19beta1 | 337.092711 | 9.059022 | 8534.964539 | 259.095717 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 197.210805 | 5.751530 | 1325.068336 | 8.058927 |
| 1.18 | 199.353373 | 3.694101 | 1409.741303 | 19.398943 |
| 1.19beta1 | 366.627098 | 5.900294 | 1421.643469 | 19.157693 |

![regexp_Compile](./b52c0e0ed5be138613a41e4ac82fa786572d3635aa9d38700ddd7703cdee0d33.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 173.851281 | 3.538170 | 6012.436835 | 8.234604 |
| 1.18 | 177.507880 | 4.974412 | 6023.640624 | 21.940467 |
| 1.19beta1 | 349.420269 | 6.815520 | 6044.810352 | 16.661184 |

![Sort_Random](./7a0a58c9e3b5825d5c91544e7e01469f5aeb4b3af178a861bf75b9731df604c0.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 156.471450 | 4.235613 | 4452.298125 | 4.781552 |
| 1.18 | 168.692375 | 3.989658 | 4295.380210 | 26.873414 |
| 1.19beta1 | 339.001866 | 14.551888 | 448.128335 | 6.004809 |

![Sort_Reversed](./4f239a2e282214a7bf7c377fcf6bb4540d0934ce7ce00fadb75e8eeeb38d843d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 239.172572 | 4.366567 | 6173.007585 | 59.893923 |
| 1.18 | 234.782187 | 4.711334 | 6800.995921 | 12.402392 |
| 1.19beta1 | 394.335913 | 12.293737 | 3891.233287 | 12.068320 |

![switch_case](./725e73000e499ff7420aa0f5b1c7dfb379e3381a122e47e5d482cb597e03166a.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 145.787188 | 9.275174 | 3909.973217 | 57.040892 |
| 1.18 | 156.383405 | 4.680529 | 5223.338309 | 13.239820 |
| 1.19beta1 | 337.784028 | 9.389872 | 5227.493473 | 7.703626 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 151.221813 | 6.759417 | 4683.503904 | 5.425329 |
| 1.18 | 153.490139 | 3.828121 | 4814.330661 | 8.756184 |
| 1.19beta1 | 328.765925 | 7.874955 | 4811.968298 | 13.173320 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

