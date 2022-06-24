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

Mhz: 2294.687000

CacheSize: 51200

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz

Cores: 1

Mhz: 2294.687000

CacheSize: 51200

Microcode: 0xffffffff

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 155.704613 | 3.334066 | 7915.195097 | 154.014875 |
| 1.18 | 164.322389 | 8.152969 | 7463.938787 | 63.891374 |
| 1.19beta1 | 349.827591 | 8.695661 | 9399.234528 | 139.235812 |

![MergeSort](./619024e898d5dcaadcf23d3b2f3a22d86c871a7b76284aafd1eb289200c2e49a.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 193.617132 | 6.353361 | 1312.385395 | 18.244011 |
| 1.18 | 205.521970 | 9.428361 | 1355.921805 | 37.985387 |
| 1.19beta1 | 367.463117 | 5.350560 | 1411.012256 | 32.318098 |

![regexp_Compile](./b52c0e0ed5be138613a41e4ac82fa786572d3635aa9d38700ddd7703cdee0d33.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 205.887842 | 9.979878 | 2866.731844 | 32.429598 |
| 1.18 | 206.647219 | 5.448453 | 2887.782912 | 58.592036 |
| 1.19beta1 | 371.024083 | 5.930790 | 2736.347781 | 33.643238 |

![regexp_FindAllString](./efbe67306d3132a2dcfa4c74e1ad1b2c51fd7423e2e5a5e3e4878c640f2a526d.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 177.677316 | 8.256310 | 5687.740601 | 88.856091 |
| 1.18 | 179.424108 | 9.118760 | 5692.474125 | 80.046579 |
| 1.19beta1 | 358.951812 | 8.432154 | 5814.562885 | 70.152430 |

![Sort_Random](./7a0a58c9e3b5825d5c91544e7e01469f5aeb4b3af178a861bf75b9731df604c0.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 168.222336 | 10.245871 | 4514.152853 | 62.150439 |
| 1.18 | 172.180616 | 6.393684 | 4391.094287 | 62.701789 |
| 1.19beta1 | 344.910190 | 13.572672 | 443.583194 | 14.039211 |

![Sort_Reversed](./4f239a2e282214a7bf7c377fcf6bb4540d0934ce7ce00fadb75e8eeeb38d843d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 239.890109 | 12.834129 | 6821.579416 | 81.458227 |
| 1.18 | 244.064733 | 18.115065 | 6861.083290 | 53.505628 |
| 1.19beta1 | 431.736005 | 39.865797 | 4054.302904 | 45.794767 |

![switch_case](./725e73000e499ff7420aa0f5b1c7dfb379e3381a122e47e5d482cb597e03166a.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 151.970861 | 3.379235 | 3900.644575 | 88.243038 |
| 1.18 | 169.636127 | 13.826481 | 2680.066795 | 23.453476 |
| 1.19beta1 | 334.317524 | 8.757255 | 2650.994877 | 63.213884 |

![alloc_1.5k](./78691b2f49e91d20e4fc03ba30be4e2828c5acd9ddd58fbf8d3e5b21bed97b8d.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 152.840731 | 5.713992 | 4459.519083 | 113.756183 |
| 1.18 | 155.804309 | 7.734328 | 4459.064770 | 35.823564 |
| 1.19beta1 | 333.539649 | 8.157274 | 4474.440094 | 63.481616 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

