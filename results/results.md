# Benchmarks

## Environment

NumCPU: 2
Arch: amd64
OS: linux
Version: go1.18.3
Itercount: 10
### CPU 0

Model: Intel(R) Xeon(R) CPU E5-2673 v3 @ 2.40GHz
Cores: 1
Mhz: 2394.453000
CacheSize: 30720
Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) CPU E5-2673 v3 @ 2.40GHz
Cores: 1
Mhz: 2394.453000
CacheSize: 30720
Microcode: 0xffffffff

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 145.015261 | 8.319905 | 13182.647651 | 258.712197 |
| 1.18 | 149.589567 | 8.198342 | 12416.121677 | 129.698702 |
| 1.19beta1 | 329.000157 | 7.860174 | 15721.141179 | 266.229265 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## Sort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 149.956386 | 6.230666 | 4136.879067 | 30.533581 |
| 1.18 | 150.586974 | 7.222723 | 4074.882202 | 33.112807 |
| 1.19beta1 | 309.738509 | 7.209895 | 405.068995 | 8.693974 |

![Sort](./bec69036aa27e7fab7d44cad3909477b76631c39ba46fd7841ea71aae7e5a735.png)

## switch-case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 219.851857 | 6.342803 | 6219.657838 | 38.178715 |
| 1.18 | 219.320198 | 8.523360 | 6251.103044 | 56.207995 |
| 1.19beta1 | 376.014311 | 11.598205 | 3607.498052 | 18.949124 |

![switch-case](./1af1469d75e77ed39c58041d45b37b329137876f59fb4c03529ebb65c78b40aa.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 134.285194 | 6.633934 | 3366.108740 | 52.414169 |
| 1.18 | 141.484410 | 5.776822 | 2579.568101 | 30.240660 |
| 1.19beta1 | 302.627759 | 14.370665 | 2585.801922 | 33.682198 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 141.151370 | 3.720658 | 4288.976950 | 24.493554 |
| 1.18 | 141.619769 | 6.497153 | 4294.679445 | 31.558407 |
| 1.19beta1 | 303.854956 | 4.878252 | 4298.000307 | 22.988802 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

