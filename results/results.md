# Benchmarks

## CPU Info

NumCPU: 2
Arch: amd64
OS: linux
Version: go1.18.3

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz
Cores: 1
Mhz: 2593.905000
CacheSize: 36608
Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz
Cores: 1
Mhz: 2593.905000
CacheSize: 36608
Microcode: 0xffffffff

## alloc_1.5k

| Version | Build Time (ms) | Run Time (ms) |
| ------ | ------ | ------ |
| 1.16 | 169.571933 | 3208.291802 |
| 1.17 | 121.555133 | 3176.370006 |
| 1.18 | 133.534566 | 4348.987022 |
| 1.19beta1 | 272.736350 | 4347.378953 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Run Time (ms) |
| ------ | ------ | ------ |
| 1.16 | 127.897972 | 5074.316737 |
| 1.17 | 118.388523 | 3902.961941 |
| 1.18 | 124.285140 | 4011.928501 |
| 1.19beta1 | 266.755086 | 4012.255694 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

## MergeSort

| Version | Build Time (ms) | Run Time (ms) |
| ------ | ------ | ------ |
| 1.16 | 131.946667 | 10639.375948 |
| 1.17 | 126.730380 | 11928.512172 |
| 1.18 | 135.953905 | 11461.306385 |
| 1.19beta1 | 287.039902 | 14131.717284 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## Sort

| Version | Build Time (ms) | Run Time (ms) |
| ------ | ------ | ------ |
| 1.16 | 142.988739 | 4703.223330 |
| 1.17 | 134.888248 | 3711.237824 |
| 1.18 | 135.284196 | 3574.597874 |
| 1.19beta1 | 272.409405 | 369.477110 |

![Sort](./bec69036aa27e7fab7d44cad3909477b76631c39ba46fd7841ea71aae7e5a735.png)

## switch-case

| Version | Build Time (ms) | Run Time (ms) |
| ------ | ------ | ------ |
| 1.16 | 204.026042 | 5287.559205 |
| 1.17 | 191.710399 | 5200.803164 |
| 1.18 | 192.677426 | 5722.456068 |
| 1.19beta1 | 323.804618 | 3271.220086 |

![switch-case](./1af1469d75e77ed39c58041d45b37b329137876f59fb4c03529ebb65c78b40aa.png)

