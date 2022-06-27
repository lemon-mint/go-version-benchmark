# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.18.3

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz

Cores: 1

Mhz: 2294.686000

CacheSize: 51200

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz

Cores: 1

Mhz: 2294.686000

CacheSize: 51200

Microcode: 0xffffffff

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 164.401961 | 6.875235 | 5302.976543 | 144.204720 |
| 1.18 | 172.473406 | 9.305985 | 5367.325130 | 188.042482 |
| 1.19beta1 | 304.643447 | 9.895290 | 4720.956710 | 208.670681 |

![Sort_Random](./7a0a58c9e3b5825d5c91544e7e01469f5aeb4b3af178a861bf75b9731df604c0.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 157.459120 | 14.861404 | 3879.186996 | 202.606499 |
| 1.18 | 141.273684 | 12.048387 | 4102.306669 | 332.690858 |
| 1.19beta1 | 295.914716 | 24.498688 | 422.911722 | 29.093459 |

![Sort_Reversed](./4f239a2e282214a7bf7c377fcf6bb4540d0934ce7ce00fadb75e8eeeb38d843d.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 158.418804 | 9.488177 | 4027.168491 | 173.152608 |
| 1.18 | 148.097748 | 6.798323 | 2491.209759 | 49.807934 |
| 1.19beta1 | 311.726564 | 14.042033 | 2429.725280 | 65.362186 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 137.452573 | 4.297051 | 3915.613866 | 83.601273 |
| 1.18 | 141.124485 | 18.351688 | 3831.071543 | 82.621615 |
| 1.19beta1 | 302.168770 | 11.072523 | 4196.691814 | 274.606686 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 159.449307 | 11.968988 | 7918.005900 | 364.174046 |
| 1.18 | 153.803915 | 5.618598 | 6821.637343 | 119.597465 |
| 1.19beta1 | 345.588589 | 14.758747 | 8956.967161 | 281.280436 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 199.704995 | 27.628449 | 1314.195888 | 68.353307 |
| 1.18 | 199.937596 | 8.875915 | 1198.716873 | 66.219152 |
| 1.19beta1 | 329.389172 | 31.646462 | 1231.040954 | 53.350517 |

![regexp_Compile](./b52c0e0ed5be138613a41e4ac82fa786572d3635aa9d38700ddd7703cdee0d33.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 196.416742 | 8.621521 | 2813.602387 | 104.135036 |
| 1.18 | 217.175414 | 8.509159 | 2699.918253 | 89.212503 |
| 1.19beta1 | 347.201993 | 10.463990 | 2535.165117 | 87.542737 |

![regexp_FindAllString](./efbe67306d3132a2dcfa4c74e1ad1b2c51fd7423e2e5a5e3e4878c640f2a526d.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 166.501675 | 6.541883 | 6133.418363 | 651.121841 |
| 1.18 | 171.375614 | 60.056210 | 4670.801406 | 365.867466 |
| 1.19beta1 | 319.206841 | 12.728087 | 4918.073984 | 220.214286 |

![TimeAfterFunc](./b4a2fe2bf5600625b3bbe08e356e7f255f29db9268c853a512b4a253305d979a.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 218.225746 | 28.549362 | 5536.367814 | 130.843156 |
| 1.18 | 203.949653 | 10.559090 | 5616.359965 | 204.198010 |
| 1.19beta1 | 387.238809 | 32.712706 | 3431.527118 | 226.787550 |

![switch_case](./725e73000e499ff7420aa0f5b1c7dfb379e3381a122e47e5d482cb597e03166a.png)

