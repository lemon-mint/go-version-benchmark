# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.18.3

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.437000

CacheSize: 49152

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz

Cores: 1

Mhz: 2793.437000

CacheSize: 49152

Microcode: 0xffffffff

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 119.468657 | 3.292659 | 3680.340532 | 7.361474 |
| 1.18 | 128.955104 | 3.826981 | 5217.990592 | 1.737438 |
| 1.19beta1 | 267.701332 | 6.465261 | 5218.195765 | 1.521640 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 197.212006 | 17.152008 | 5928.379080 | 56.083470 |
| 1.18 | 200.247718 | 11.322260 | 5924.320804 | 37.735279 |
| 1.19beta1 | 325.892505 | 6.798516 | 3409.928176 | 3.766959 |

![switch_case](./725e73000e499ff7420aa0f5b1c7dfb379e3381a122e47e5d482cb597e03166a.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 130.958870 | 7.797277 | 3641.061104 | 6.620289 |
| 1.18 | 137.919291 | 4.354678 | 3544.905703 | 2.436691 |
| 1.19beta1 | 273.565300 | 8.959163 | 396.180380 | 1.590769 |

![Sort_Reversed](./4f239a2e282214a7bf7c377fcf6bb4540d0934ce7ce00fadb75e8eeeb38d843d.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 145.204665 | 91.250390 | 4424.195011 | 394.063761 |
| 1.18 | 156.568452 | 56.258457 | 4168.713246 | 859.271038 |
| 1.19beta1 | 282.109492 | 183.478234 | 4680.331162 | 467.492692 |

![TimeAfterFunc](./b4a2fe2bf5600625b3bbe08e356e7f255f29db9268c853a512b4a253305d979a.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 120.773983 | 7.017212 | 4791.791662 | 0.808917 |
| 1.18 | 125.903908 | 2.838863 | 4797.107824 | 3.928213 |
| 1.19beta1 | 266.354807 | 3.355034 | 4796.474601 | 2.038892 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 125.386381 | 3.806004 | 5867.578281 | 56.636625 |
| 1.18 | 132.199497 | 2.612130 | 5576.421338 | 22.898205 |
| 1.19beta1 | 270.145084 | 1.953961 | 6906.989722 | 46.182406 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 157.971141 | 14.127831 | 1066.121445 | 7.297492 |
| 1.18 | 162.749789 | 7.278464 | 1105.681924 | 2.885006 |
| 1.19beta1 | 295.723614 | 4.252033 | 1125.678441 | 9.515073 |

![regexp_Compile](./b52c0e0ed5be138613a41e4ac82fa786572d3635aa9d38700ddd7703cdee0d33.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 160.144229 | 2.120788 | 2328.569081 | 8.848627 |
| 1.18 | 166.853922 | 9.817507 | 2337.928146 | 29.924845 |
| 1.19beta1 | 299.279203 | 2.885920 | 2189.742053 | 22.303069 |

![regexp_FindAllString](./efbe67306d3132a2dcfa4c74e1ad1b2c51fd7423e2e5a5e3e4878c640f2a526d.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 142.177819 | 4.604170 | 6092.489337 | 13.043427 |
| 1.18 | 146.333724 | 2.277007 | 5990.216076 | 20.062046 |
| 1.19beta1 | 286.178279 | 7.828890 | 5931.710806 | 14.111824 |

![Sort_Random](./7a0a58c9e3b5825d5c91544e7e01469f5aeb4b3af178a861bf75b9731df604c0.png)

