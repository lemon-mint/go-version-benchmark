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
| 1.17 | 122.167062 | 4.747663 | 3761.604409 | 34.014400 |
| 1.18 | 128.289286 | 4.373011 | 5218.203557 | 3.513240 |
| 1.19beta1 | 267.642941 | 3.851626 | 5218.095638 | 9.823698 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 121.890845 | 3.415377 | 4793.133323 | 3.802604 |
| 1.18 | 125.378086 | 2.163420 | 4796.441763 | 4.704512 |
| 1.19beta1 | 265.569633 | 4.886159 | 4797.451684 | 4.821213 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 130.422298 | 10.676950 | 6110.050374 | 109.911705 |
| 1.18 | 131.450428 | 7.686704 | 5758.390935 | 70.135230 |
| 1.19beta1 | 272.253505 | 5.026212 | 7162.949183 | 164.117890 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 160.577315 | 4.600696 | 1097.421109 | 6.873958 |
| 1.18 | 164.478685 | 3.039244 | 1116.835419 | 4.113997 |
| 1.19beta1 | 300.463178 | 3.702514 | 1133.691478 | 7.679445 |

![regexp_Compile](./b52c0e0ed5be138613a41e4ac82fa786572d3635aa9d38700ddd7703cdee0d33.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 141.232910 | 1.998446 | 6081.527550 | 10.378376 |
| 1.18 | 150.752638 | 3.976077 | 5986.114310 | 11.552898 |
| 1.19beta1 | 284.015613 | 6.030166 | 5941.645024 | 10.801971 |

![Sort_Random](./7a0a58c9e3b5825d5c91544e7e01469f5aeb4b3af178a861bf75b9731df604c0.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 133.703283 | 8.590928 | 3667.257472 | 11.513228 |
| 1.18 | 137.096441 | 2.082064 | 3545.025308 | 2.971866 |
| 1.19beta1 | 270.557262 | 6.504691 | 394.453921 | 3.413652 |

![Sort_Reversed](./4f239a2e282214a7bf7c377fcf6bb4540d0934ce7ce00fadb75e8eeeb38d843d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 190.857931 | 3.740841 | 5941.315350 | 17.209932 |
| 1.18 | 196.031323 | 6.494343 | 5884.635573 | 32.305005 |
| 1.19beta1 | 326.736449 | 11.678704 | 3421.350344 | 4.713627 |

![switch_case](./725e73000e499ff7420aa0f5b1c7dfb379e3381a122e47e5d482cb597e03166a.png)

