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

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 153.266627 | 12.718227 | 1019.355075 | 11.867544 |
| 1.18 | 157.687047 | 13.285672 | 1079.884464 | 47.367734 |
| 1.19beta1 | 279.948305 | 3.918142 | 1070.675326 | 19.851103 |

![regexp_Compile](./b52c0e0ed5be138613a41e4ac82fa786572d3635aa9d38700ddd7703cdee0d33.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 156.826780 | 2.894035 | 2145.179525 | 68.859506 |
| 1.18 | 164.412001 | 3.296776 | 2346.556510 | 54.400402 |
| 1.19beta1 | 301.756884 | 2.573493 | 2134.364444 | 49.199279 |

![regexp_FindAllString](./efbe67306d3132a2dcfa4c74e1ad1b2c51fd7423e2e5a5e3e4878c640f2a526d.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 125.129065 | 3.832293 | 3225.162422 | 260.329289 |
| 1.18 | 141.819033 | 5.583947 | 3574.380010 | 243.916958 |
| 1.19beta1 | 262.567602 | 5.810563 | 323.944697 | 23.183471 |

![Sort_Reversed](./4f239a2e282214a7bf7c377fcf6bb4540d0934ce7ce00fadb75e8eeeb38d843d.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 140.112460 | 3.772011 | 4591.759638 | 315.746141 |
| 1.18 | 147.427353 | 37.567225 | 4367.216525 | 547.854592 |
| 1.19beta1 | 292.735545 | 13.472084 | 4778.400067 | 365.610857 |

![TimeAfterFunc](./b4a2fe2bf5600625b3bbe08e356e7f255f29db9268c853a512b4a253305d979a.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 112.113641 | 2.862982 | 2893.081217 | 60.714787 |
| 1.18 | 125.622406 | 4.601157 | 4346.339353 | 316.606875 |
| 1.19beta1 | 274.042231 | 6.268323 | 4188.644123 | 299.102375 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 119.219696 | 4.335019 | 3901.182771 | 6.687492 |
| 1.18 | 130.901529 | 10.283159 | 4012.645976 | 174.510319 |
| 1.19beta1 | 255.305412 | 8.518654 | 3502.753147 | 254.914056 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 123.784697 | 5.254134 | 5667.352922 | 52.010245 |
| 1.18 | 131.110722 | 4.022907 | 5747.479433 | 49.846374 |
| 1.19beta1 | 256.348345 | 15.419947 | 6751.474233 | 176.909263 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 126.853386 | 2.606617 | 4690.969654 | 359.516505 |
| 1.18 | 141.876850 | 6.076170 | 4389.871989 | 316.140854 |
| 1.19beta1 | 287.241449 | 5.000814 | 5021.414415 | 318.348416 |

![Sort_Random](./7a0a58c9e3b5825d5c91544e7e01469f5aeb4b3af178a861bf75b9731df604c0.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 197.248077 | 17.815259 | 4475.082592 | 371.010634 |
| 1.18 | 188.272915 | 19.485433 | 4957.332965 | 365.960023 |
| 1.19beta1 | 313.396945 | 7.432006 | 3271.414908 | 218.381103 |

![switch_case](./725e73000e499ff7420aa0f5b1c7dfb379e3381a122e47e5d482cb597e03166a.png)

