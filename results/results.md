# Benchmarks

## CPU Info

NumCPU: 2
Arch: amd64
OS: linux
Version: go1.18.3

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz
Cores: 1
Mhz: 2593.904000
CacheSize: 36608
Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz
Cores: 1
Mhz: 2593.904000
CacheSize: 36608
Microcode: 0xffffffff

## Sort

| Version | Build Time (ms) | Run Time (ms) |
| ------ | ------ | ------ |
| 1.16 | 149.300396 | 4575.250194 |
| 1.17 | 135.574461 | 3605.172704 |
| 1.18 | 146.058098 | 3471.601359 |
| 1.19beta1 | 302.077587 | 337.396262 |

![Sort](./bec69036aa27e7fab7d44cad3909477b76631c39ba46fd7841ea71aae7e5a735.png)

## switch-case

| Version | Build Time (ms) | Run Time (ms) |
| ------ | ------ | ------ |
| 1.16 | 218.218116 | 4982.341845 |
| 1.17 | 211.514962 | 5179.724460 |
| 1.18 | 208.589555 | 5386.396403 |
| 1.19beta1 | 352.092150 | 3087.759795 |

![switch-case](./1af1469d75e77ed39c58041d45b37b329137876f59fb4c03529ebb65c78b40aa.png)

## alloc_1.5k

| Version | Build Time (ms) | Run Time (ms) |
| ------ | ------ | ------ |
| 1.16 | 197.030668 | 3159.196022 |
| 1.17 | 129.434073 | 3209.027380 |
| 1.18 | 134.660054 | 4092.084307 |
| 1.19beta1 | 285.820829 | 4093.018595 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Run Time (ms) |
| ------ | ------ | ------ |
| 1.16 | 132.239690 | 5090.766628 |
| 1.17 | 131.663178 | 3675.411511 |
| 1.18 | 136.580488 | 3804.581846 |
| 1.19beta1 | 291.051092 | 3912.827452 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

## MergeSort

| Version | Build Time (ms) | Run Time (ms) |
| ------ | ------ | ------ |
| 1.16 | 134.587049 | 11154.556639 |
| 1.17 | 138.103078 | 12498.919962 |
| 1.18 | 141.037240 | 12305.395636 |
| 1.19beta1 | 293.195168 | 14988.345242 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

