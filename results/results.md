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

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 160.629415 | 3.417633 | 1088.939486 | 5.549823 |
| 1.18 | 164.609256 | 2.920637 | 1147.710765 | 7.363137 |
| 1.19beta1 | 303.318383 | 3.286932 | 1148.614680 | 9.174759 |

![regexp_Compile](./b52c0e0ed5be138613a41e4ac82fa786572d3635aa9d38700ddd7703cdee0d33.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 162.372051 | 3.131094 | 2347.095251 | 10.137859 |
| 1.18 | 174.058844 | 6.127154 | 2354.719818 | 12.268600 |
| 1.19beta1 | 306.633414 | 3.711527 | 2247.969920 | 23.848700 |

![regexp_FindAllString](./efbe67306d3132a2dcfa4c74e1ad1b2c51fd7423e2e5a5e3e4878c640f2a526d.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 142.892417 | 2.072038 | 5011.963295 | 1.548888 |
| 1.18 | 150.868558 | 5.166093 | 5027.606360 | 1.498324 |
| 1.19beta1 | 287.659879 | 5.621640 | 5022.304674 | 2.953625 |

![Sort_Random](./7a0a58c9e3b5825d5c91544e7e01469f5aeb4b3af178a861bf75b9731df604c0.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 134.182127 | 3.451312 | 3713.579802 | 2.023669 |
| 1.18 | 139.960785 | 4.766059 | 3574.076833 | 1.301122 |
| 1.19beta1 | 273.401669 | 3.325615 | 369.616710 | 0.466358 |

![Sort_Reversed](./4f239a2e282214a7bf7c377fcf6bb4540d0934ce7ce00fadb75e8eeeb38d843d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 192.753836 | 2.966797 | 5204.019662 | 3.202257 |
| 1.18 | 196.831595 | 5.153122 | 5724.074865 | 1.316052 |
| 1.19beta1 | 331.547858 | 10.512352 | 3272.110776 | 1.722527 |

![switch_case](./725e73000e499ff7420aa0f5b1c7dfb379e3381a122e47e5d482cb597e03166a.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 122.681633 | 4.030618 | 3168.683360 | 19.068930 |
| 1.18 | 132.894480 | 4.108752 | 4347.957378 | 2.306032 |
| 1.19beta1 | 268.762550 | 5.285587 | 4347.530819 | 10.466845 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 124.334955 | 4.405839 | 3901.603579 | 1.910702 |
| 1.18 | 129.167569 | 3.685242 | 4013.426862 | 2.901063 |
| 1.19beta1 | 272.298996 | 7.238927 | 4012.525556 | 1.328931 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 131.059760 | 8.578956 | 5928.816656 | 90.051923 |
| 1.18 | 137.092714 | 5.601085 | 5686.954776 | 29.267949 |
| 1.19beta1 | 277.014577 | 2.159649 | 7066.934645 | 49.044659 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

