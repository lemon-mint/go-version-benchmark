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

Mhz: 2095.080000

CacheSize: 36608

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHz

Cores: 1

Mhz: 2095.080000

CacheSize: 36608

Microcode: 0xffffffff

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 230.752133 | 2.070787 | 6181.927846 | 17.214066 |
| 1.18 | 238.044464 | 16.756473 | 6822.000491 | 22.539068 |
| 1.19beta1 | 392.783679 | 4.889262 | 3880.054579 | 10.011625 |

![switch_case](./725e73000e499ff7420aa0f5b1c7dfb379e3381a122e47e5d482cb597e03166a.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 147.340793 | 3.032557 | 3798.274833 | 59.174420 |
| 1.18 | 158.492229 | 4.970820 | 5162.338434 | 15.393926 |
| 1.19beta1 | 326.772550 | 6.816592 | 5161.861528 | 20.659080 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 147.234807 | 4.788912 | 4616.174619 | 14.884704 |
| 1.18 | 153.022639 | 2.026895 | 4756.232902 | 11.458198 |
| 1.19beta1 | 326.895706 | 15.557667 | 4751.040788 | 21.697282 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 153.774015 | 10.097110 | 7019.789983 | 125.331697 |
| 1.18 | 163.428477 | 3.563165 | 6880.676706 | 68.125705 |
| 1.19beta1 | 337.293434 | 5.076519 | 8548.545223 | 78.903080 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 191.835223 | 6.795793 | 1307.959483 | 20.210136 |
| 1.18 | 203.183963 | 5.664558 | 1377.262461 | 3.979135 |
| 1.19beta1 | 365.756101 | 3.780530 | 1394.407977 | 14.120936 |

![regexp_Compile](./b52c0e0ed5be138613a41e4ac82fa786572d3635aa9d38700ddd7703cdee0d33.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 199.315781 | 4.038674 | 2850.031414 | 18.608887 |
| 1.18 | 208.690753 | 6.033475 | 3012.495085 | 23.640851 |
| 1.19beta1 | 375.471494 | 6.806498 | 2807.231858 | 20.874864 |

![regexp_FindAllString](./efbe67306d3132a2dcfa4c74e1ad1b2c51fd7423e2e5a5e3e4878c640f2a526d.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 176.063801 | 3.184925 | 6005.617514 | 12.153743 |
| 1.18 | 181.563561 | 3.524699 | 6027.495721 | 6.475807 |
| 1.19beta1 | 343.196575 | 4.168459 | 6031.858361 | 12.391960 |

![Sort_Random](./7a0a58c9e3b5825d5c91544e7e01469f5aeb4b3af178a861bf75b9731df604c0.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 161.846271 | 5.402080 | 4431.856864 | 13.543581 |
| 1.18 | 167.565555 | 4.008864 | 4249.277990 | 14.176975 |
| 1.19beta1 | 328.907945 | 9.638674 | 441.489701 | 3.178855 |

![Sort_Reversed](./4f239a2e282214a7bf7c377fcf6bb4540d0934ce7ce00fadb75e8eeeb38d843d.png)

