# Benchmarks

## CPU Info

NumCPU: 2
Arch: amd64
OS: linux
Version: go1.18.3

### CPU 0

Model: Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz
Cores: 1
Mhz: 2294.685000
CacheSize: 51200
Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz
Cores: 1
Mhz: 2294.685000
CacheSize: 51200
Microcode: 0xffffffff

## switch-case

| Version | Build Time (ms) | Run Time (ms) |
| ------ | ------ | ------ |
| 1.16 | 215.655319 | 5456.643753 |
| 1.17 | 225.387460 | 6897.375397 |
| 1.18 | 275.286495 | 7016.742861 |
| 1.19beta1 | 396.389189 | 3592.663293 |

![switch-case](./1af1469d75e77ed39c58041d45b37b329137876f59fb4c03529ebb65c78b40aa.png)

## alloc_1.5k

| Version | Build Time (ms) | Run Time (ms) |
| ------ | ------ | ------ |
| 1.16 | 179.131571 | 3349.043069 |
| 1.17 | 136.694153 | 3379.319380 |
| 1.18 | 153.571280 | 2667.943246 |
| 1.19beta1 | 340.287556 | 2519.011813 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Run Time (ms) |
| ------ | ------ | ------ |
| 1.16 | 163.916598 | 5162.110484 |
| 1.17 | 130.113346 | 3852.944455 |
| 1.18 | 150.178895 | 4348.912294 |
| 1.19beta1 | 349.038208 | 3955.290530 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

## MergeSort

| Version | Build Time (ms) | Run Time (ms) |
| ------ | ------ | ------ |
| 1.16 | 146.740746 | 13313.492852 |
| 1.17 | 177.235566 | 14488.376907 |
| 1.18 | 147.256432 | 12930.224878 |
| 1.19beta1 | 293.037745 | 17146.495567 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## Sort

| Version | Build Time (ms) | Run Time (ms) |
| ------ | ------ | ------ |
| 1.16 | 170.361016 | 5164.624163 |
| 1.17 | 149.457768 | 3694.298368 |
| 1.18 | 143.356908 | 3623.404334 |
| 1.19beta1 | 300.130360 | 350.855496 |

![Sort](./bec69036aa27e7fab7d44cad3909477b76631c39ba46fd7841ea71aae7e5a735.png)

