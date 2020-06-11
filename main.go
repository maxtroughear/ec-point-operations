package main

import (
	"log"
	"math/big"
)

func main() {
	curve := ECCurve{
		a: big.NewInt(2),
		b: big.NewInt(3),
		p: big.NewInt(29),
	}

	basePoint := ECPoint{
		curve: &curve,
		x:     big.NewInt(8),
		y:     big.NewInt(3),
	}
	basePoint.Log(1)

	twoP := basePoint.Double()
	twoP.Log(2)

	fourP := twoP.Double()
	fourP.Log(4)

	eightP := fourP.Double()
	eightP.Log(8)

	sixteenP := eightP.Double()
	sixteenP.Log(16)

	seventeenP := sixteenP.Add(basePoint)
	seventeenP.Log(17)

}

// func DoubleandAdd(P ECPoint) float64 {
// 	T := P
// 	for i := t - 1; i >= 0; i-- {
// 		T =
// 		if di == 1 {
// 			T += math.Mod(P, n)
// 		}
// 	}
// 	return T
// }

// ECCurve struct
type ECCurve struct {
	a *big.Int
	b *big.Int
	p *big.Int
}

// ECPoint struct
type ECPoint struct {
	curve *ECCurve
	x     *big.Int
	y     *big.Int
}

// Log prints the values of a point to console
func (p ECPoint) Log(i int) {
	log.Printf("%dP (%s,%s)", i, p.x.String(), p.y.String())
}

// Add where points p and q are not the same
func (p ECPoint) Add(q ECPoint) ECPoint {
	h :=
		new(big.Int).Mul(
			new(big.Int).Sub(
				q.y,
				p.y,
			),
			modInverse(
				new(big.Int).Sub(
					q.x,
					p.x,
				),
				p.curve.p,
			),
		)
	// h :=
	// 	new(big.Int).Div(
	// 		new(big.Int).Sub(
	// 			q.y,
	// 			p.y,
	// 		),
	// 		new(big.Int).Sub(
	// 			q.x,
	// 			p.x,
	// 		),
	// 	)

	return p.internalAdd(q, h)
}

// Double where points p and q are at the same coordinates
func (p ECPoint) Double() ECPoint {
	q := p
	h :=
		new(big.Int).Mul(
			new(big.Int).Add(
				new(big.Int).Mul(
					big.NewInt(3),
					new(big.Int).Exp(
						p.x,
						big.NewInt(2),
						nil,
					),
				),
				p.curve.a,
			),
			modInverse(
				new(big.Int).Mul(
					p.y,
					big.NewInt(2)),
				p.curve.p,
			),
		)

	//h := (3*math.Pow(p.x, 2) + p.curve.a) /
	//	modInverse(new(big.Int).Mul(p.x, big.NewInt(2)), p.curve.p)
	return p.internalAdd(q, h)
}

func (p ECPoint) internalAdd(q ECPoint, h *big.Int) ECPoint {
	r := p
	r.x =
		new(big.Int).Mod(
			new(big.Int).Sub(
				new(big.Int).Sub(
					new(big.Int).Exp(
						h,
						big.NewInt(2),
						nil,
					),
					p.x,
				),
				q.x,
			),
			p.curve.p,
		)
	r.y =
		new(big.Int).Mod(
			new(big.Int).Sub(
				new(big.Int).Mul(
					h,
					new(big.Int).Sub(
						p.x,
						r.x,
					),
				),
				p.y,
			),
			p.curve.p,
		)

	//r.x = math.Pow(h, 2) - p.x - q.x
	//r.y = h*(p.x-r.x) - p.y
	return r
}

func modInverse(x *big.Int, p *big.Int) *big.Int {
	i := new(big.Int).ModInverse(x, p)
	if i == nil {
		return x
	}
	return i
}
