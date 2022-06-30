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

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 193.429387 | 11.565849 | 6350.396681 | 505.525762 |
| 1.18 | 222.823425 | 53.098942 | 4750.369874 | 968.643903 |
| 1.19beta1 | 397.069106 | 26.128806 | 6275.319030 | 929.138353 |

![TimeAfterFunc](./b4a2fe2bf5600625b3bbe08e356e7f255f29db9268c853a512b4a253305d979a.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 164.236382 | 24.243432 | 4607.046571 | 82.486387 |
| 1.18 | 174.664654 | 12.484870 | 4680.139321 | 50.483724 |
| 1.19beta1 | 390.075599 | 10.843502 | 4664.365781 | 59.626451 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 181.215461 | 6.043694 | 4650.037378 | 116.872724 |
| 1.18 | 190.414309 | 7.412445 | 4611.699658 | 107.737603 |
| 1.19beta1 | 384.801297 | 23.723890 | 470.196307 | 5.288336 |

![Sort_Reversed](./4f239a2e282214a7bf7c377fcf6bb4540d0934ce7ce00fadb75e8eeeb38d843d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 268.173939 | 31.561952 | 7237.668740 | 76.674657 |
| 1.18 | 294.991929 | 29.573299 | 7216.838953 | 116.074623 |
| 1.19beta1 | 466.087400 | 18.428312 | 4109.168691 | 92.735040 |

![switch_case](./725e73000e499ff7420aa0f5b1c7dfb379e3381a122e47e5d482cb597e03166a.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 181.642336 | 5.587123 | 4108.760852 | 59.313029 |
| 1.18 | 182.451907 | 7.042177 | 2870.992412 | 33.171456 |
| 1.19beta1 | 388.510174 | 8.823202 | 2853.998408 | 78.142950 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 167.595813 | 8.691041 | 8472.069895 | 149.069126 |
| 1.18 | 180.315095 | 5.809604 | 7992.704400 | 190.441930 |
| 1.19beta1 | 377.514734 | 17.026575 | 9992.519097 | 123.961878 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 220.963326 | 27.166147 | 1477.180195 | 57.056362 |
| 1.18 | 219.119000 | 26.350617 | 1432.514978 | 22.251199 |
| 1.19beta1 | 401.618263 | 43.401874 | 1465.349835 | 33.214662 |

![regexp_Compile](./b52c0e0ed5be138613a41e4ac82fa786572d3635aa9d38700ddd7703cdee0d33.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 218.287618 | 6.219865 | 3031.832602 | 56.251746 |
| 1.18 | 225.273855 | 9.791390 | 3065.422478 | 70.350622 |
| 1.19beta1 | 414.530099 | 16.232722 | 2886.168784 | 57.526397 |

![regexp_FindAllString](./efbe67306d3132a2dcfa4c74e1ad1b2c51fd7423e2e5a5e3e4878c640f2a526d.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 197.369101 | 12.278352 | 5932.034092 | 90.520161 |
| 1.18 | 203.816145 | 78.046271 | 5926.978794 | 105.933034 |
| 1.19beta1 | 388.927796 | 95.010190 | 5991.298670 | 49.475127 |

![Sort_Random](./7a0a58c9e3b5825d5c91544e7e01469f5aeb4b3af178a861bf75b9731df604c0.png)

