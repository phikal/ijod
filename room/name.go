package room

import (
	"math/rand"
	"time"
)

var (
	chars  []byte
	markov = map[byte]map[byte]float64{
		'a': {
			'a': 59.0 / 11380, 'i': 2093.0 / 11380, 'k': 843.0 / 11380,
			'o': 84.0 / 11380, 's': 3760.0 / 11380, 'u': 1216.0 / 11380,
			'v': 978.0 / 11380, 'w': 652.0 / 11380, 'x': 228.0 / 11380,
			'y': 1124.0 / 11380, 'z': 343.0 / 11380,
		}, 'i': {
			'a': 3155.0 / 19061, 'i': 21.0 / 19061, 'k': 301.0 / 19061,
			'o': 5298.0 / 19061, 's': 6518.0 / 19061, 'u': 306.0 / 19061,
			'v': 1722.0 / 19061, 'w': 36.0 / 19061, 'x': 172.0 / 19061,
			'y': 16.0 / 19061, 'z': 1516.0 / 19061,
		}, 'k': {
			'a': 394.0 / 3110, 'i': 1367.0 / 3110, 'k': 41.0 / 3110,
			'o': 176.0 / 3110, 's': 739.0 / 3110, 'u': 104.0 / 3110,
			'w': 71.0 / 3110, 'y': 218.0 / 3110,
		}, 'o': {
			'a': 954.0 / 17140, 'i': 877.0 / 17140, 'k': 564.0 / 17140,
			'o': 2225.0 / 17140, 's': 2274.0 / 17140, 'u': 3611.0 / 17140,
			'v': 1181.0 / 17140, 'w': 1564.0 / 17140, 'x': 271.0 / 17140,
			'y': 370.0 / 17140, 'z': 139.0 / 17140,
		}, 's': {
			'a': 1883.0 / 15254, 'i': 3686.0 / 15254, 'k': 589.0 / 15254,
			'o': 1727.0 / 15254, 's': 4646.0 / 15254, 'u': 1756.0 / 15254,
			'v': 23.0 / 15254, 'w': 463.0 / 15254, 'y': 479.0 / 15254,
			'z': 2.0 / 15254,
		}, 'u': {
			'a': 1067.0 / 6350, 'i': 1059.0 / 6350, 'k': 114.0 / 6350,
			'o': 212.0 / 6350, 's': 3539.0 / 6350, 'u': 15.0 / 6350,
			'v': 83.0 / 6350, 'w': 13.0 / 6350, 'x': 85.0 / 6350,
			'y': 44.0 / 6350, 'z': 119.0 / 6350,
		}, 'v': {
			'a': 1211.0 / 3707, 'i': 1726.0 / 3707, 'o': 585.0 / 3707,
			's': 18.0 / 3707, 'u': 95.0 / 3707, 'v': 21.0 / 3707,
			'y': 51.0 / 3707,
		}, 'w': {
			'a': 1521.0 / 3817, 'i': 1052.0 / 3817, 'k': 47.0 / 3817,
			'o': 834.0 / 3817, 's': 309.0 / 3817, 'u': 17.0 / 3817,
			'w': 8.0 / 3817, 'y': 23.0 / 3817, 'z': 6.0 / 3817,
		}, 'x': {
			'a': 123.0 / 661, 'i': 331.0 / 661, 'o': 76.0 / 661,
			's': 3.0 / 661, 'u': 65.0 / 661, 'v': 1.0 / 661,
			'w': 9.0 / 661, 'x': 2.0 / 661, 'y': 49.0 / 661,
			'z': 2.0 / 661,
		}, 'y': {
			'a': 405.0 / 1849, 'i': 372.0 / 1849, 'k': 18.0 / 1849,
			'o': 221.0 / 1849, 's': 615.0 / 1849, 'u': 60.0 / 1849,
			'v': 2.0 / 1849, 'w': 102.0 / 1849, 'x': 23.0 / 1849,
			'y': 5.0 / 1849, 'z': 26.0 / 1849,
		}, 'z': {
			'a': 457.0 / 232, 'i': 528.0 / 232, 'k': 4.0 / 232,
			'o': 176.0 / 232, 's': 4.0 / 232, 'u': 33.0 / 232,
			'v': 9.0 / 232, 'w': 7.0 / 232, 'y': 54.0 / 232,
		},
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())

	chars = make([]byte, 0, len(markov))
	for c := range markov {
		chars = append(chars, c)
	}
}

func randName() string {
	for {
		var (
			n    = 5
			c    = chars[rand.Intn(len(chars))]
			name = make([]byte, 1, n)
			p0   float64
		)

		name[0] = c
		for i := 1; i < n; i++ {
			p := rand.Float64()

			for C, P := range markov[c] {
				if P+p0 > p {
					c = C
					break
				}

				p0 += P
			}

			name = append(name, c)
		}

		if _, ok := rooms[string(name)]; !ok {
			return string(name)
		}
	}
}