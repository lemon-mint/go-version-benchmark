# Benchmarks

## Environment

NumCPU: 2
Arch: amd64
OS: linux
Version: go1.18.3
Itercount: 10
### CPU 0

Model: Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz
Cores: 1
Mhz: 2294.686000
CacheSize: 51200
Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz
Cores: 1
Mhz: 2294.686000
CacheSize: 51200
Microcode: 0xffffffff

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 171.306866 | 6.558069 | 8286.251431 | 165.599226 |
| 1.18 | 171.922831 | 8.499789 | 7870.144636 | 89.507511 |
| 1.19beta1 | 367.707614 | 11.647955 | 9878.945376 | 68.562277 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 215.005777 | 10.446894 | 1400.407634 | 15.281294 |
| 1.18 | 214.900352 | 4.507706 | 1450.422738 | 8.613564 |
| 1.19beta1 | 409.088260 | 11.854537 | 1475.642725 | 14.199305 |

![regexp_Compile](./b52c0e0ed5be138613a41e4ac82fa786572d3635aa9d38700ddd7703cdee0d33.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 188.496110 | 3.389482 | 6065.161976 | 65.743042 |
| 1.18 | 189.183566 | 7.871678 | 6139.215395 | 85.626781 |
| 1.19beta1 | 363.263202 | 6.619181 | 6145.781752 | 51.714516 |

![Sort_Random](./7a0a58c9e3b5825d5c91544e7e01469f5aeb4b3af178a861bf75b9731df604c0.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 172.009862 | 7.224168 | 4722.875034 | 86.019836 |
| 1.18 | 180.018834 | 3.703357 | 4708.456491 | 29.395918 |
| 1.19beta1 | 354.916050 | 10.764193 | 470.967434 | 13.387209 |

![Sort_Reversed](./4f239a2e282214a7bf7c377fcf6bb4540d0934ce7ce00fadb75e8eeeb38d843d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 255.050864 | 7.270621 | 7294.410199 | 132.629927 |
| 1.18 | 264.536034 | 10.919195 | 7322.264546 | 44.593419 |
| 1.19beta1 | 441.873533 | 12.466586 | 4251.904177 | 45.225862 |

![switch_case](./725e73000e499ff7420aa0f5b1c7dfb379e3381a122e47e5d482cb597e03166a.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 158.603174 | 4.954289 | 3975.950599 | 68.293532 |
| 1.18 | 171.336821 | 7.056185 | 2920.254424 | 40.428458 |
| 1.19beta1 | 368.710922 | 12.159776 | 2902.763192 | 66.484796 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 158.433870 | 5.442830 | 4720.891535 | 49.113995 |
| 1.18 | 170.290263 | 5.453912 | 4746.046774 | 101.952423 |
| 1.19beta1 | 365.941289 | 8.929114 | 4732.817980 | 48.151531 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

