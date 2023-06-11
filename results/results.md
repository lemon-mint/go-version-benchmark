# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.19.9

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

## CGO_CALL_C_FUNC

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 488.210930 | 3.554116 | 3519.381602 | 6.640979 |
| 1.19 | 509.317849 | 9.797422 | 3546.771587 | 28.836632 |
| 1.20 | 439.363925 | 9.365078 | 3555.707373 | 4.613928 |
| gotip | 487.835538 | 8.637629 | 3535.470341 | 4.202427 |

![CGO_CALL_C_FUNC](./CGO_CALL_C_FUNC__1eb049ef6b.png)

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 123.725075 | 8.476014 | 4016.722359 | 7.068879 |
| 1.19 | 117.015893 | 3.398804 | 3903.624918 | 1.434452 |
| 1.20 | 118.855852 | 6.009185 | 3921.522014 | 11.804543 |
| gotip | 121.790402 | 5.580930 | 4251.658170 | 3.896435 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 127.295296 | 3.160689 | 5731.341979 | 17.228736 |
| 1.19 | 119.348490 | 21.012300 | 5597.392793 | 18.099327 |
| 1.20 | 125.445659 | 3.185196 | 6228.974404 | 27.109322 |
| gotip | 126.535330 | 4.382636 | 6571.626657 | 11.040613 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 140.382074 | 5.315436 | 5679.145803 | 193.086712 |
| 1.19 | 117.521803 | 3.370642 | 9265.169579 | 63.481669 |
| 1.20 | 122.347044 | 3.694330 | 9155.429886 | 116.261511 |
| gotip | 123.069270 | 2.582134 | 5760.666731 | 53.204809 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 129.078507 | 3.192467 | 6229.301492 | 230.429681 |
| 1.19 | 110.815787 | 9.939187 | 7970.199064 | 129.708382 |
| 1.20 | 120.821299 | 3.484339 | 9276.338362 | 247.711763 |
| gotip | 122.162167 | 5.145805 | 9519.002088 | 241.718610 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 148.422892 | 6.350757 | 5031.931026 | 3.340502 |
| 1.19 | 119.414284 | 3.566224 | 5090.863812 | 7.849346 |
| 1.20 | 130.992959 | 4.816383 | 5084.332382 | 1.300222 |
| gotip | 136.479092 | 4.254570 | 5136.934127 | 2.683345 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 134.179786 | 2.887859 | 3576.937641 | 1.932986 |
| 1.19 | 116.722541 | 16.942968 | 370.357878 | 0.527751 |
| 1.20 | 120.878182 | 4.385449 | 369.780200 | 0.662338 |
| gotip | 124.600212 | 4.102258 | 375.414277 | 0.294324 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 149.255198 | 4.148196 | 4088.518199 | 671.727430 |
| 1.19 | 136.403027 | 12.901094 | 5010.645653 | 433.668603 |
| 1.20 | 134.224801 | 5.022666 | 5554.854187 | 626.681910 |
| gotip | 137.608616 | 5.124577 | 5383.281882 | 244.518206 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 128.631546 | 2.930589 | 4350.582712 | 2.054344 |
| 1.19 | 113.781935 | 7.118853 | 2193.754679 | 1.173077 |
| 1.20 | 119.938822 | 3.804133 | 2193.047985 | 1.105325 |
| gotip | 123.138822 | 4.274960 | 4351.003515 | 0.991547 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 154.826007 | 3.327560 | 1187.249468 | 7.114094 |
| 1.19 | 135.006860 | 6.556740 | 1214.505664 | 20.196573 |
| 1.20 | 145.822860 | 4.307684 | 1221.046744 | 18.717279 |
| gotip | 144.549562 | 3.851327 | 1174.982517 | 7.403617 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 165.381253 | 12.721399 | 2410.378190 | 7.263587 |
| 1.19 | 140.267155 | 5.896786 | 2308.377499 | 10.508802 |
| 1.20 | 151.419741 | 5.016166 | 2262.421059 | 7.653405 |
| gotip | 147.082544 | 5.277924 | 2208.695650 | 12.722318 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.18 | 199.050296 | 5.336185 | 5843.845713 | 3.464603 |
| 1.19 | 163.860916 | 6.076408 | 3221.072293 | 1.555660 |
| 1.20 | 175.100204 | 5.070120 | 3221.212378 | 1.100656 |
| gotip | 166.451888 | 4.811887 | 2776.276893 | 1.931340 |

![switch_case](./switch_case__725e73000e.png)

