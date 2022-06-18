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

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 196.623435 | 4.571330 | 5204.027833 | 4.179953 |
| 1.18 | 202.210091 | 3.814019 | 5725.577301 | 1.911929 |
| 1.19beta1 | 340.338765 | 13.328175 | 3272.254526 | 2.868894 |

![switch_case](./725e73000e499ff7420aa0f5b1c7dfb379e3381a122e47e5d482cb597e03166a.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 126.997335 | 4.102121 | 3198.527338 | 40.552758 |
| 1.18 | 135.411457 | 5.704832 | 4348.557617 | 1.305753 |
| 1.19beta1 | 278.514831 | 12.882475 | 4349.113792 | 1.233321 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 129.100238 | 4.077621 | 3902.053206 | 2.587060 |
| 1.18 | 131.747248 | 4.905085 | 4012.550203 | 1.671272 |
| 1.19beta1 | 273.585365 | 5.987111 | 4012.819294 | 3.094866 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 134.705604 | 6.646446 | 6136.335468 | 94.234304 |
| 1.18 | 137.467299 | 6.958743 | 6037.539917 | 23.328990 |
| 1.19beta1 | 281.860902 | 3.801578 | 7352.691544 | 47.704197 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 164.827026 | 6.127984 | 1095.054782 | 4.052619 |
| 1.18 | 169.766589 | 3.323650 | 1149.730138 | 10.291894 |
| 1.19beta1 | 310.387855 | 3.906127 | 1158.332505 | 13.894106 |

![regexp_Compile](./b52c0e0ed5be138613a41e4ac82fa786572d3635aa9d38700ddd7703cdee0d33.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 170.691238 | 3.967830 | 2355.561722 | 10.482861 |
| 1.18 | 174.205736 | 3.664326 | 2362.963196 | 9.060574 |
| 1.19beta1 | 315.224063 | 3.866413 | 2257.535094 | 17.622396 |

![regexp_FindAllString](./efbe67306d3132a2dcfa4c74e1ad1b2c51fd7423e2e5a5e3e4878c640f2a526d.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 149.389718 | 2.877315 | 5012.106342 | 1.355906 |
| 1.18 | 154.444433 | 8.565898 | 5029.471292 | 2.772067 |
| 1.19beta1 | 292.927557 | 3.457657 | 5022.027684 | 1.961053 |

![Sort_Random](./7a0a58c9e3b5825d5c91544e7e01469f5aeb4b3af178a861bf75b9731df604c0.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 136.204664 | 2.333586 | 3711.913048 | 2.531920 |
| 1.18 | 144.276491 | 4.170468 | 3575.079828 | 0.979682 |
| 1.19beta1 | 285.033295 | 5.393480 | 370.099220 | 0.410228 |

![Sort_Reversed](./4f239a2e282214a7bf7c377fcf6bb4540d0934ce7ce00fadb75e8eeeb38d843d.png)

