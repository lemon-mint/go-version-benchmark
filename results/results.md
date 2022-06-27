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

Mhz: 2593.905000

CacheSize: 36608

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz

Cores: 1

Mhz: 2593.905000

CacheSize: 36608

Microcode: 0xffffffff

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 126.529563 | 13.170865 | 3901.960692 | 2.816690 |
| 1.18 | 129.518864 | 8.153027 | 4011.483261 | 0.757621 |
| 1.19beta1 | 268.866206 | 3.051894 | 4012.755202 | 1.528512 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 130.963846 | 3.380683 | 6036.934183 | 65.506917 |
| 1.18 | 132.081953 | 4.641129 | 5789.848738 | 51.257883 |
| 1.19beta1 | 274.359228 | 5.828963 | 7232.844352 | 172.928702 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 143.293616 | 3.429293 | 5011.758409 | 1.479815 |
| 1.18 | 149.309429 | 2.096492 | 5029.276272 | 2.414096 |
| 1.19beta1 | 282.519606 | 5.381495 | 5021.518431 | 2.438107 |

![Sort_Random](./7a0a58c9e3b5825d5c91544e7e01469f5aeb4b3af178a861bf75b9731df604c0.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 135.453751 | 8.776019 | 3709.373725 | 1.418797 |
| 1.18 | 140.460813 | 4.985484 | 3574.611305 | 0.941532 |
| 1.19beta1 | 272.645413 | 3.942732 | 369.395253 | 0.535128 |

![Sort_Reversed](./4f239a2e282214a7bf7c377fcf6bb4540d0934ce7ce00fadb75e8eeeb38d843d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 191.785112 | 13.117212 | 5204.364118 | 4.012546 |
| 1.18 | 194.573658 | 15.638838 | 5722.876160 | 0.943115 |
| 1.19beta1 | 333.674874 | 5.838246 | 3271.470219 | 0.882386 |

![switch_case](./725e73000e499ff7420aa0f5b1c7dfb379e3381a122e47e5d482cb597e03166a.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 145.518462 | 1.519631 | 4613.728126 | 291.603078 |
| 1.18 | 155.507072 | 8.578211 | 4652.227184 | 546.571071 |
| 1.19beta1 | 292.063703 | 8.403088 | 5097.789550 | 259.575902 |

![TimeAfterFunc](./b4a2fe2bf5600625b3bbe08e356e7f255f29db9268c853a512b4a253305d979a.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 123.401739 | 4.817416 | 3197.577110 | 13.276968 |
| 1.18 | 132.091275 | 3.565301 | 4347.696447 | 0.844774 |
| 1.19beta1 | 271.547246 | 5.078909 | 4348.716179 | 1.313869 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 162.003522 | 12.093269 | 1093.402380 | 2.781971 |
| 1.18 | 166.802326 | 10.689449 | 1144.960736 | 7.415803 |
| 1.19beta1 | 298.805180 | 5.747742 | 1153.251593 | 8.415772 |

![regexp_Compile](./b52c0e0ed5be138613a41e4ac82fa786572d3635aa9d38700ddd7703cdee0d33.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 166.016232 | 3.148530 | 2348.533715 | 16.519539 |
| 1.18 | 168.711083 | 3.041064 | 2363.513897 | 8.344223 |
| 1.19beta1 | 305.481395 | 2.267068 | 2246.757957 | 11.573495 |

![regexp_FindAllString](./efbe67306d3132a2dcfa4c74e1ad1b2c51fd7423e2e5a5e3e4878c640f2a526d.png)

