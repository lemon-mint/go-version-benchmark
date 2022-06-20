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

Mhz: 2294.685000

CacheSize: 51200

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz

Cores: 1

Mhz: 2294.685000

CacheSize: 51200

Microcode: 0xffffffff

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 171.131656 | 7.068421 | 8477.688641 | 171.937483 |
| 1.18 | 178.709366 | 9.275966 | 7971.012485 | 235.864209 |
| 1.19beta1 | 381.978203 | 11.719940 | 9562.393876 | 331.139825 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 201.041247 | 8.155629 | 1376.186942 | 32.662350 |
| 1.18 | 216.546042 | 12.887089 | 1458.312973 | 42.833444 |
| 1.19beta1 | 395.523355 | 20.172939 | 1456.166049 | 22.558288 |

![regexp_Compile](./b52c0e0ed5be138613a41e4ac82fa786572d3635aa9d38700ddd7703cdee0d33.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 206.840796 | 8.228089 | 3063.483200 | 61.242947 |
| 1.18 | 212.702290 | 6.892809 | 2977.386597 | 45.941937 |
| 1.19beta1 | 404.623897 | 17.643960 | 2863.235786 | 63.121364 |

![regexp_FindAllString](./efbe67306d3132a2dcfa4c74e1ad1b2c51fd7423e2e5a5e3e4878c640f2a526d.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 178.190364 | 9.331282 | 5918.835619 | 92.887026 |
| 1.18 | 186.350250 | 9.244291 | 5938.774143 | 100.841252 |
| 1.19beta1 | 364.678720 | 13.092621 | 5878.028112 | 84.128425 |

![Sort_Random](./7a0a58c9e3b5825d5c91544e7e01469f5aeb4b3af178a861bf75b9731df604c0.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 170.895097 | 58.182342 | 4669.333945 | 96.281138 |
| 1.18 | 171.115641 | 5.677562 | 4525.809245 | 75.592078 |
| 1.19beta1 | 344.358186 | 12.178077 | 460.113797 | 10.745378 |

![Sort_Reversed](./4f239a2e282214a7bf7c377fcf6bb4540d0934ce7ce00fadb75e8eeeb38d843d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 253.833996 | 10.843233 | 7037.791209 | 126.834831 |
| 1.18 | 247.299353 | 25.853058 | 6711.511701 | 135.003922 |
| 1.19beta1 | 421.921354 | 17.845021 | 3888.920363 | 91.907059 |

![switch_case](./725e73000e499ff7420aa0f5b1c7dfb379e3381a122e47e5d482cb597e03166a.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 171.875489 | 7.280011 | 4054.528438 | 101.360174 |
| 1.18 | 164.159227 | 8.406548 | 2836.302115 | 49.003538 |
| 1.19beta1 | 379.936451 | 13.720530 | 2788.758090 | 40.359832 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 165.622832 | 3.802393 | 4709.313128 | 61.989498 |
| 1.18 | 172.547009 | 12.554827 | 4692.765828 | 49.991253 |
| 1.19beta1 | 381.981329 | 14.680415 | 4686.027122 | 59.381769 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

