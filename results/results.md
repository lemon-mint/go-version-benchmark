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

Mhz: 2095.193000

CacheSize: 36608

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHz

Cores: 1

Mhz: 2095.193000

CacheSize: 36608

Microcode: 0xffffffff

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 159.334144 | 3.295402 | 5604.806330 | 67.927685 |
| 1.18 | 167.334158 | 6.918433 | 5880.936182 | 73.126057 |
| 1.19beta1 | 323.220057 | 6.174135 | 5891.368596 | 54.308169 |

![Sort_Random](./7a0a58c9e3b5825d5c91544e7e01469f5aeb4b3af178a861bf75b9731df604c0.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 151.531618 | 3.981720 | 4353.321656 | 24.782048 |
| 1.18 | 156.734364 | 3.330083 | 4079.970618 | 80.933421 |
| 1.19beta1 | 313.921110 | 5.961834 | 428.050592 | 3.822119 |

![Sort_Reversed](./4f239a2e282214a7bf7c377fcf6bb4540d0934ce7ce00fadb75e8eeeb38d843d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 221.394247 | 3.873602 | 6017.181666 | 80.726489 |
| 1.18 | 216.028160 | 6.287730 | 6577.701467 | 40.248242 |
| 1.19beta1 | 374.909482 | 9.569602 | 3797.649105 | 29.675157 |

![switch_case](./725e73000e499ff7420aa0f5b1c7dfb379e3381a122e47e5d482cb597e03166a.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 137.110430 | 3.748180 | 3596.909800 | 61.318721 |
| 1.18 | 140.204904 | 5.282953 | 4936.904372 | 14.950530 |
| 1.19beta1 | 305.747835 | 25.160304 | 4816.989047 | 98.434811 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 137.264362 | 2.895173 | 4324.602203 | 58.748159 |
| 1.18 | 140.861393 | 3.270338 | 4201.213659 | 112.326383 |
| 1.19beta1 | 285.247708 | 12.845209 | 4359.454802 | 78.749239 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 141.326449 | 12.441441 | 6598.696120 | 57.025950 |
| 1.18 | 145.503479 | 4.408270 | 6238.521997 | 61.713953 |
| 1.19beta1 | 306.959354 | 7.151570 | 7868.570145 | 168.689490 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 175.651960 | 3.869257 | 1210.786884 | 16.989775 |
| 1.18 | 175.118274 | 4.005609 | 1279.714242 | 13.676865 |
| 1.19beta1 | 337.645611 | 4.724078 | 1296.881313 | 12.876005 |

![regexp_Compile](./b52c0e0ed5be138613a41e4ac82fa786572d3635aa9d38700ddd7703cdee0d33.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 179.530763 | 5.753436 | 2635.154716 | 43.935716 |
| 1.18 | 186.280135 | 3.095965 | 2810.938003 | 35.830473 |
| 1.19beta1 | 337.279707 | 2.952155 | 2594.944828 | 32.891000 |

![regexp_FindAllString](./efbe67306d3132a2dcfa4c74e1ad1b2c51fd7423e2e5a5e3e4878c640f2a526d.png)

