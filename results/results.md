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

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 134.227213 | 36.241942 | 3174.301825 | 18.232287 |
| 1.18 | 135.347770 | 6.104113 | 4348.324100 | 3.327691 |
| 1.19beta1 | 275.784523 | 6.466728 | 4347.303243 | 0.975965 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 131.343913 | 3.673537 | 3901.127035 | 1.239971 |
| 1.18 | 133.928591 | 2.261685 | 4013.298158 | 3.018516 |
| 1.19beta1 | 276.562147 | 2.723162 | 4012.658598 | 2.160671 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 136.902362 | 5.403329 | 6088.137938 | 99.261905 |
| 1.18 | 140.368933 | 8.368330 | 5903.053312 | 76.149588 |
| 1.19beta1 | 286.952811 | 5.990615 | 7347.839065 | 91.378188 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 165.423900 | 10.049688 | 1095.413310 | 3.959795 |
| 1.18 | 169.233220 | 5.465317 | 1154.388584 | 5.027723 |
| 1.19beta1 | 311.677400 | 3.794247 | 1160.086310 | 6.551528 |

![regexp_Compile](./b52c0e0ed5be138613a41e4ac82fa786572d3635aa9d38700ddd7703cdee0d33.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 172.297013 | 3.524127 | 2359.916250 | 9.425444 |
| 1.18 | 176.715973 | 3.295948 | 2365.007296 | 12.629223 |
| 1.19beta1 | 311.604977 | 4.764565 | 2242.967424 | 5.773166 |

![regexp_FindAllString](./efbe67306d3132a2dcfa4c74e1ad1b2c51fd7423e2e5a5e3e4878c640f2a526d.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 147.248857 | 3.477525 | 5012.297084 | 2.183742 |
| 1.18 | 156.471246 | 5.848280 | 5030.704846 | 2.283094 |
| 1.19beta1 | 290.824400 | 8.825430 | 5023.028924 | 1.501552 |

![Sort_Random](./7a0a58c9e3b5825d5c91544e7e01469f5aeb4b3af178a861bf75b9731df604c0.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 139.520026 | 4.770276 | 3714.625017 | 1.625947 |
| 1.18 | 143.700053 | 1.352668 | 3574.692983 | 1.282230 |
| 1.19beta1 | 281.209982 | 3.209932 | 369.559393 | 0.414870 |

![Sort_Reversed](./4f239a2e282214a7bf7c377fcf6bb4540d0934ce7ce00fadb75e8eeeb38d843d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 199.797980 | 6.400877 | 5205.576972 | 2.891593 |
| 1.18 | 200.248437 | 4.230596 | 5723.687454 | 2.533072 |
| 1.19beta1 | 337.917994 | 15.032096 | 3272.039380 | 1.796956 |

![switch_case](./725e73000e499ff7420aa0f5b1c7dfb379e3381a122e47e5d482cb597e03166a.png)

