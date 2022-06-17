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
Mhz: 2394.458000
CacheSize: 30720
Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) CPU E5-2673 v3 @ 2.40GHz
Cores: 1
Mhz: 2394.458000
CacheSize: 30720
Microcode: 0xffffffff

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 245.171811 | 11.699348 | 6867.743138 | 122.518549 |
| 1.18 | 242.279220 | 12.522992 | 6693.125340 | 86.060557 |
| 1.19beta1 | 404.572906 | 32.839189 | 3876.413620 | 25.961967 |

![switch_case](./725e73000e499ff7420aa0f5b1c7dfb379e3381a122e47e5d482cb597e03166a.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 154.695673 | 10.924099 | 3686.707317 | 34.979425 |
| 1.18 | 165.584548 | 6.709491 | 2673.011686 | 87.515756 |
| 1.19beta1 | 332.652374 | 13.647011 | 2730.103418 | 136.910460 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 154.321147 | 6.015481 | 4794.516156 | 22.005376 |
| 1.18 | 158.245968 | 7.273246 | 4767.225050 | 45.126625 |
| 1.19beta1 | 342.487996 | 10.192496 | 4747.898387 | 32.699539 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 165.017169 | 8.799689 | 7373.928251 | 68.637240 |
| 1.18 | 167.046509 | 10.909593 | 7134.061489 | 123.116101 |
| 1.19beta1 | 351.458969 | 8.206448 | 8808.071891 | 135.094964 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 195.359782 | 6.647899 | 1312.334673 | 12.613213 |
| 1.18 | 204.087414 | 6.928354 | 1372.223231 | 14.732664 |
| 1.19beta1 | 374.486043 | 10.032287 | 1393.639848 | 17.850572 |

![regexp_Compile](./b52c0e0ed5be138613a41e4ac82fa786572d3635aa9d38700ddd7703cdee0d33.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 178.333080 | 4.367473 | 5872.687106 | 41.332737 |
| 1.18 | 188.280809 | 14.314752 | 5840.125417 | 59.507398 |
| 1.19beta1 | 377.219681 | 12.778984 | 5806.259708 | 126.514307 |

![Sort_Random](./7a0a58c9e3b5825d5c91544e7e01469f5aeb4b3af178a861bf75b9731df604c0.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 169.307197 | 12.505989 | 4493.683936 | 88.608198 |
| 1.18 | 176.143945 | 6.224475 | 4454.730698 | 92.608783 |
| 1.19beta1 | 339.461149 | 10.362614 | 435.318125 | 9.436758 |

![Sort_Reversed](./4f239a2e282214a7bf7c377fcf6bb4540d0934ce7ce00fadb75e8eeeb38d843d.png)

