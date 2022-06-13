# Benchmarks

## Environment

NumCPU: 8
Arch: amd64
OS: windows
Version: go1.18.3
Itercount: 10
### CPU 0

Model: Intel(R) Core(TM) i7-4790K CPU @ 4.00GHz
Cores: 8
Mhz: 4001.000000
CacheSize: 0
Microcode: 

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 1199.963550 | 74.015515 | 3210.825850 | 55.136208 |
| 1.18 | 1194.551300 | 40.685614 | 3197.762150 | 48.393861 |
| 1.19beta1 | 1850.572350 | 18.395682 | 3203.531350 | 36.587768 |

![Fibonacci](./016be0f0bc3aacaadb309d0adc2b1024980e3775065236c79ab0d186380b4f83.png)

