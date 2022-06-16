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

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 205.344853 | 10.990402 | 5203.281793 | 3222.828305 |
| 1.18 | 195.561371 | 9.742071 | 5484.111714 | 320.697001 |
| 1.19beta1 | 334.179001 | 25.744481 | 3271.506151 | 188.991017 |

![switch_case](./725e73000e499ff7420aa0f5b1c7dfb379e3381a122e47e5d482cb597e03166a.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 120.512678 | 5.075915 | 3045.504969 | 109.737534 |
| 1.18 | 127.881827 | 2.418626 | 3836.350623 | 251.501415 |
| 1.19beta1 | 274.373325 | 31.661159 | 3856.089093 | 228.949772 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 124.900925 | 8.015991 | 3901.345850 | 183.601371 |
| 1.18 | 123.507203 | 12.188626 | 3540.467756 | 189.263095 |
| 1.19beta1 | 269.241334 | 44.011362 | 3539.643225 | 141.804610 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 127.947092 | 5.106229 | 6193.017255 | 100.836777 |
| 1.18 | 136.297852 | 14.226096 | 6012.866913 | 173.181320 |
| 1.19beta1 | 264.160446 | 11.828088 | 7490.761937 | 336.021821 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 156.182684 | 5.997625 | 1063.558895 | 38.888650 |
| 1.18 | 164.500082 | 5.055710 | 1130.888544 | 226.958244 |
| 1.19beta1 | 303.933720 | 7.202300 | 1127.972798 | 56.793324 |

![regexp_Compile](./b52c0e0ed5be138613a41e4ac82fa786572d3635aa9d38700ddd7703cdee0d33.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 148.672497 | 7.984954 | 5016.244124 | 2.648825 |
| 1.18 | 160.084283 | 3.934883 | 4449.982410 | 268.816501 |
| 1.19beta1 | 313.242282 | 20.643041 | 4435.481830 | 294.273930 |

![Sort_Random](./7a0a58c9e3b5825d5c91544e7e01469f5aeb4b3af178a861bf75b9731df604c0.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 143.056827 | 11.014990 | 3284.655134 | 173.507238 |
| 1.18 | 144.435837 | 12.169646 | 3159.590778 | 194.039211 |
| 1.19beta1 | 290.328238 | 17.144808 | 369.667692 | 18.527883 |

![Sort_Reversed](./4f239a2e282214a7bf7c377fcf6bb4540d0934ce7ce00fadb75e8eeeb38d843d.png)

