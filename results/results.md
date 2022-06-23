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

Mhz: 2593.906000

CacheSize: 36608

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz

Cores: 1

Mhz: 2593.906000

CacheSize: 36608

Microcode: 0xffffffff

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 159.856584 | 2.879386 | 5019.795490 | 1.644943 |
| 1.18 | 164.250064 | 4.161797 | 5033.865998 | 1.771644 |
| 1.19beta1 | 317.106388 | 3.755534 | 5029.123192 | 1.698958 |

![Sort_Random](./7a0a58c9e3b5825d5c91544e7e01469f5aeb4b3af178a861bf75b9731df604c0.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 147.199486 | 2.999353 | 3719.427214 | 2.242194 |
| 1.18 | 152.491865 | 2.865855 | 3580.015490 | 1.792752 |
| 1.19beta1 | 310.926674 | 6.825579 | 370.052666 | 0.536161 |

![Sort_Reversed](./4f239a2e282214a7bf7c377fcf6bb4540d0934ce7ce00fadb75e8eeeb38d843d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 215.038207 | 4.195223 | 5214.994766 | 4.110874 |
| 1.18 | 219.888183 | 9.265658 | 5731.721851 | 2.693217 |
| 1.19beta1 | 374.961073 | 18.726131 | 3276.956567 | 1.513825 |

![switch_case](./725e73000e499ff7420aa0f5b1c7dfb379e3381a122e47e5d482cb597e03166a.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 134.626254 | 2.672496 | 3290.378230 | 38.277531 |
| 1.18 | 143.689879 | 3.937408 | 4353.873814 | 1.314177 |
| 1.19beta1 | 304.420236 | 7.049001 | 4351.520157 | 2.027315 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 133.716998 | 3.745606 | 3905.907496 | 1.294716 |
| 1.18 | 140.360861 | 2.681159 | 4017.653210 | 4.103295 |
| 1.19beta1 | 298.728782 | 5.534966 | 4017.340821 | 1.101720 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 139.448529 | 2.198415 | 6981.564871 | 126.778922 |
| 1.18 | 145.941812 | 7.788279 | 6726.343674 | 62.896768 |
| 1.19beta1 | 306.577937 | 3.680184 | 8385.738930 | 121.072950 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 177.605651 | 4.835891 | 1142.967441 | 10.735899 |
| 1.18 | 179.971174 | 2.521681 | 1195.028439 | 11.578506 |
| 1.19beta1 | 342.393686 | 5.856475 | 1218.638688 | 11.894412 |

![regexp_Compile](./b52c0e0ed5be138613a41e4ac82fa786572d3635aa9d38700ddd7703cdee0d33.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 181.642097 | 3.765831 | 2389.415276 | 8.474248 |
| 1.18 | 191.684672 | 2.976975 | 2403.301944 | 9.973820 |
| 1.19beta1 | 349.510466 | 3.202956 | 2280.877165 | 18.185910 |

![regexp_FindAllString](./efbe67306d3132a2dcfa4c74e1ad1b2c51fd7423e2e5a5e3e4878c640f2a526d.png)

