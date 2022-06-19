# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.18.3

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHz

Cores: 1

Mhz: 2095.077000

CacheSize: 36608

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHz

Cores: 1

Mhz: 2095.077000

CacheSize: 36608

Microcode: 0xffffffff

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 214.930273 | 7.650648 | 5532.039404 | 88.566098 |
| 1.18 | 215.991814 | 15.399553 | 6063.905959 | 103.889618 |
| 1.19beta1 | 351.072474 | 11.230201 | 3431.304386 | 39.855681 |

![switch_case](./725e73000e499ff7420aa0f5b1c7dfb379e3381a122e47e5d482cb597e03166a.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 135.487318 | 5.811948 | 3434.489822 | 65.689317 |
| 1.18 | 136.799861 | 3.352601 | 4597.795388 | 92.782417 |
| 1.19beta1 | 289.625345 | 5.142307 | 4559.109132 | 43.594765 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 134.072103 | 4.282798 | 4080.541065 | 69.297356 |
| 1.18 | 140.686585 | 4.472295 | 4314.453971 | 85.785948 |
| 1.19beta1 | 299.925988 | 7.523397 | 4293.445478 | 56.171730 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 138.317396 | 7.684890 | 6541.546170 | 85.438974 |
| 1.18 | 141.352985 | 14.172188 | 6247.537560 | 55.897918 |
| 1.19beta1 | 295.601008 | 13.035737 | 7729.066671 | 109.377939 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 178.655173 | 8.565709 | 1205.936194 | 32.710144 |
| 1.18 | 177.313864 | 3.521938 | 1247.267188 | 17.311022 |
| 1.19beta1 | 334.489609 | 13.638612 | 1252.963778 | 24.926324 |

![regexp_Compile](./b52c0e0ed5be138613a41e4ac82fa786572d3635aa9d38700ddd7703cdee0d33.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 172.288088 | 3.219036 | 2567.849968 | 43.122536 |
| 1.18 | 178.303269 | 6.330506 | 2578.226063 | 36.994078 |
| 1.19beta1 | 320.517325 | 6.473568 | 2504.998221 | 54.755922 |

![regexp_FindAllString](./efbe67306d3132a2dcfa4c74e1ad1b2c51fd7423e2e5a5e3e4878c640f2a526d.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 155.008731 | 5.338080 | 5266.712716 | 65.238435 |
| 1.18 | 167.963551 | 6.215518 | 5288.282726 | 55.855852 |
| 1.19beta1 | 314.059783 | 8.102389 | 5298.144484 | 48.940076 |

![Sort_Random](./7a0a58c9e3b5825d5c91544e7e01469f5aeb4b3af178a861bf75b9731df604c0.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 142.246580 | 5.579108 | 4011.663699 | 88.012114 |
| 1.18 | 152.419392 | 4.551765 | 3798.964179 | 84.575841 |
| 1.19beta1 | 293.222960 | 7.410441 | 399.671421 | 9.507835 |

![Sort_Reversed](./4f239a2e282214a7bf7c377fcf6bb4540d0934ce7ce00fadb75e8eeeb38d843d.png)

