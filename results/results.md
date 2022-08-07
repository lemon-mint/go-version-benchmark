# Benchmarks

## Environment

NumCPU: 2

Arch: amd64

OS: linux

Version: go1.18.4

Itercount: 10

### CPU 0

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz

Cores: 1

Mhz: 2593.903000

CacheSize: 36608

Microcode: 0xffffffff

### CPU 1

Model: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz

Cores: 1

Mhz: 2593.903000

CacheSize: 36608

Microcode: 0xffffffff

## Fibonacci

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 129.896364 | 3.413744 | 3905.597340 | 2.758058 |
| 1.18 | 138.022179 | 7.085267 | 4016.225997 | 1.793128 |
| 1.19 | 149.239264 | 2.467985 | 3904.862264 | 1.569433 |
| gotip | 148.964405 | 3.577994 | 4014.592265 | 1.351490 |

![Fibonacci](./Fibonacci__016be0f0bc.png)

## Garbage_Collection

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 137.721201 | 9.064523 | 5730.006012 | 36.729996 |
| 1.18 | 146.438056 | 2.176185 | 5722.672539 | 22.928405 |
| 1.19 | 152.374907 | 3.664564 | 5549.149408 | 33.897421 |
| gotip | 151.738994 | 3.276652 | 5570.001846 | 23.077542 |

![Garbage_Collection](./Garbage_Collection__f27466590e.png)

## Goroutine_Creation

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 145.484424 | 6.345210 | 9093.240107 | 39.244847 |
| 1.18 | 146.872431 | 5.724687 | 5422.495221 | 39.099367 |
| 1.19 | 150.787325 | 8.416683 | 9045.902793 | 27.534904 |
| gotip | 154.171563 | 8.742510 | 5590.308162 | 12.892470 |

![Goroutine_Creation](./Goroutine_Creation__c0773f341a.png)

## MergeSort

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 136.194559 | 3.655458 | 6381.369946 | 57.987416 |
| 1.18 | 143.688971 | 3.947558 | 6176.921706 | 58.146913 |
| 1.19 | 156.600437 | 2.959674 | 8032.984750 | 34.057155 |
| gotip | 160.071922 | 8.757119 | 8242.669539 | 111.860858 |

![MergeSort](./MergeSort__619024e898.png)

## Sort_Random

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 150.229426 | 1.639456 | 5015.146389 | 2.466841 |
| 1.18 | 162.667291 | 5.025463 | 5030.662326 | 2.687365 |
| 1.19 | 161.322119 | 5.353193 | 5084.946557 | 4.817064 |
| gotip | 164.382284 | 2.651650 | 5008.540104 | 1.865462 |

![Sort_Random](./Sort_Random__7a0a58c9e3.png)

## Sort_Reversed

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 140.879878 | 3.987394 | 3715.980610 | 2.454265 |
| 1.18 | 149.884027 | 2.380478 | 3578.924747 | 1.168541 |
| 1.19 | 156.215152 | 3.029633 | 370.255210 | 0.436626 |
| gotip | 155.203193 | 4.906466 | 370.236708 | 0.401406 |

![Sort_Reversed](./Sort_Reversed__4f239a2e28.png)

## TimeAfterFunc

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 158.795431 | 4.966730 | 5015.107648 | 744.992337 |
| 1.18 | 165.553756 | 51.719177 | 4431.167123 | 481.312483 |
| 1.19 | 165.794243 | 51.738411 | 4667.656990 | 390.135928 |
| gotip | 168.091312 | 10.236736 | 4720.810634 | 296.596226 |

![TimeAfterFunc](./TimeAfterFunc__b4a2fe2bf5.png)

## alloc_1.5k

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 131.346306 | 3.307822 | 3215.484999 | 23.535998 |
| 1.18 | 138.514925 | 2.213314 | 4352.326663 | 1.728080 |
| 1.19 | 151.117739 | 3.292933 | 2194.734143 | 1.012787 |
| gotip | 151.175753 | 3.157143 | 4352.144904 | 1.141544 |

![alloc_1.5k](./alloc_1.5k__78691b2f49.png)

## regexp_Compile

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 173.924603 | 2.701692 | 1136.289399 | 6.473490 |
| 1.18 | 178.128699 | 14.605817 | 1188.476084 | 12.928164 |
| 1.19 | 184.916507 | 5.958700 | 1206.514721 | 8.046620 |
| gotip | 183.835348 | 5.793756 | 1202.623926 | 13.053806 |

![regexp_Compile](./regexp_Compile__b52c0e0ed5.png)

## regexp_FindAllString

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 176.072674 | 4.203954 | 2372.846706 | 16.063711 |
| 1.18 | 188.504993 | 4.643442 | 2386.197176 | 10.889640 |
| 1.19 | 188.375439 | 2.998255 | 2256.969834 | 8.330314 |
| gotip | 181.756273 | 2.297521 | 2262.190213 | 7.813984 |

![regexp_FindAllString](./regexp_FindAllString__efbe67306d.png)

## switch_case

| Version | Build Time (ms) | Standard Deviation | Run Time (ms) | Standard Deviation |
| ------ | ------ | ------ | ------ | ------ |
| 1.17 | 204.935147 | 3.184413 | 5213.198905 | 2.395526 |
| 1.18 | 214.817947 | 15.563146 | 5728.552736 | 2.500772 |
| 1.19 | 214.622041 | 9.346428 | 3230.999929 | 1.476607 |
| gotip | 219.267508 | 4.040476 | 3276.541242 | 0.628707 |

![switch_case](./switch_case__725e73000e.png)

