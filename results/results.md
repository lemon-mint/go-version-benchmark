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

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 149.025218 | 3.210943 | 5012.483867 | 1.298902 |
| 1.18 | 154.323166 | 8.764213 | 5029.099093 | 3.301391 |
| 1.19beta1 | 288.869303 | 5.517072 | 5021.811563 | 2.412746 |

![Sort_Random](./7a0a58c9e3b5825d5c91544e7e01469f5aeb4b3af178a861bf75b9731df604c0.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 138.183438 | 7.862011 | 3711.304982 | 2.453920 |
| 1.18 | 141.482681 | 3.249706 | 3574.189271 | 1.569204 |
| 1.19beta1 | 284.170873 | 12.813685 | 369.956335 | 0.613477 |

![Sort_Reversed](./4f239a2e282214a7bf7c377fcf6bb4540d0934ce7ce00fadb75e8eeeb38d843d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 198.066586 | 3.832596 | 5200.838316 | 1.834378 |
| 1.18 | 199.528608 | 4.658550 | 5723.408046 | 0.883593 |
| 1.19beta1 | 338.165619 | 8.553585 | 3272.545383 | 1.095524 |

![switch_case](./725e73000e499ff7420aa0f5b1c7dfb379e3381a122e47e5d482cb597e03166a.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 125.265533 | 2.774457 | 3171.687461 | 43.013833 |
| 1.18 | 135.988695 | 3.995366 | 4347.271296 | 1.217283 |
| 1.19beta1 | 277.527017 | 7.061961 | 4348.782205 | 2.876309 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 126.184054 | 4.470175 | 3901.732560 | 1.210425 |
| 1.18 | 131.121369 | 2.983239 | 4011.253661 | 2.043020 |
| 1.19beta1 | 269.651226 | 4.531856 | 4011.881888 | 1.138846 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 133.667636 | 4.188043 | 5992.875213 | 54.233712 |
| 1.18 | 139.117504 | 8.492498 | 5751.629361 | 46.293777 |
| 1.19beta1 | 282.125858 | 4.556720 | 7217.773674 | 73.666748 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 161.602192 | 2.384027 | 1091.747516 | 2.678558 |
| 1.18 | 165.938223 | 2.805524 | 1152.136745 | 10.088476 |
| 1.19beta1 | 301.994587 | 2.661146 | 1150.907970 | 6.279374 |

![regexp_Compile](./b52c0e0ed5be138613a41e4ac82fa786572d3635aa9d38700ddd7703cdee0d33.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 166.575056 | 2.721672 | 2354.173319 | 6.979085 |
| 1.18 | 171.270818 | 4.372614 | 2361.459080 | 16.041534 |
| 1.19beta1 | 308.976095 | 8.825943 | 2247.334241 | 10.826702 |

![regexp_FindAllString](./efbe67306d3132a2dcfa4c74e1ad1b2c51fd7423e2e5a5e3e4878c640f2a526d.png)

