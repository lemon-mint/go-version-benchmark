# Benchmarks

## CPU Info

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

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 129.819614 | 3.338043 | 12073.237913 | 142.834718 |
| 1.18 | 137.941286 | 8.353118 | 11536.469635 | 150.324726 |
| 1.19beta1 | 282.542988 | 7.661180 | 14094.946709 | 122.967870 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## Sort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 134.858718 | 3.517649 | 3712.481442 | 3.339277 |
| 1.18 | 139.637633 | 3.943321 | 3573.066056 | 1.125061 |
| 1.19beta1 | 275.145682 | 12.922005 | 369.611612 | 0.432114 |

![Sort](./bec69036aa27e7fab7d44cad3909477b76631c39ba46fd7841ea71aae7e5a735.png)

## switch-case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 193.304819 | 3.702105 | 5201.271532 | 2.697496 |
| 1.18 | 197.155397 | 4.098484 | 5723.562771 | 2.285060 |
| 1.19beta1 | 333.342906 | 4.113717 | 3272.027380 | 0.604245 |

![switch-case](./1af1469d75e77ed39c58041d45b37b329137876f59fb4c03529ebb65c78b40aa.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 130.602540 | 4.543614 | 3174.555205 | 50.825925 |
| 1.18 | 135.096543 | 5.866101 | 4348.360743 | 1.770363 |
| 1.19beta1 | 273.128653 | 6.444817 | 4347.478340 | 0.880110 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 124.035034 | 3.497839 | 3902.296188 | 1.938311 |
| 1.18 | 127.531948 | 3.636813 | 4011.249149 | 1.047679 |
| 1.19beta1 | 275.046732 | 5.527460 | 4012.903111 | 1.428728 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

