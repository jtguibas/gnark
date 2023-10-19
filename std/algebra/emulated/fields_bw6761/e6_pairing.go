package fields_bw6761

func (e Ext6) nSquareCompressed(z *E6, n int) *E6 {
	for i := 0; i < n; i++ {
		z = e.CyclotomicSquareCompressed(z)
	}
	return z
}

// ExpX0Minus1 set z to z^{x₀-1} in E6 and return z
// x₀-1 = 91893752504881257682351033800651177983
func (e Ext6) ExpX0Minus1(z *E6) *E6 {
	z = e.Reduce(z)
	result := e.Copy(z)
	result = e.nSquareCompressed(result, 5)
	result = e.DecompressKarabina(result)
	result = e.Mul(result, z)
	z33 := e.Copy(result)
	result = e.nSquareCompressed(result, 7)
	result = e.DecompressKarabina(result)
	result = e.Mul(result, z33)
	result = e.nSquareCompressed(result, 4)
	result = e.DecompressKarabina(result)
	result = e.Mul(result, z)
	result = e.CyclotomicSquare(result)
	result = e.Mul(result, z)
	result = e.nSquareCompressed(result, 46)
	result = e.DecompressKarabina(result)

	return result
}

// ExpX0Minus1Square set z to z^{(x₀-1)²} in E6 and return z
// (x₀-1)² = 91893752504881257682351033800651177984
func (e Ext6) ExpX0Minus1Square(z *E6) *E6 {
	z = e.Reduce(z)
	result := e.Copy(z)
	result = e.CyclotomicSquare(result)
	t0 := e.Mul(z, result)
	t1 := e.CyclotomicSquare(t0)
	t0 = e.Mul(t0, t1)
	result = e.Mul(result, t0)
	t1 = e.Mul(t1, result)
	t0 = e.Mul(t0, t1)
	t2 := e.CyclotomicSquare(t0)
	t2 = e.Mul(t1, t2)
	t0 = e.Mul(t0, t2)
	t2 = e.nSquareCompressed(t2, 7)
	t2 = e.DecompressKarabina(t2)
	t1 = e.Mul(t1, t2)
	t1 = e.nSquareCompressed(t1, 11)
	t1 = e.DecompressKarabina(t1)
	t1 = e.Mul(t0, t1)
	t1 = e.nSquareCompressed(t1, 9)
	t1 = e.DecompressKarabina(t1)
	t0 = e.Mul(t0, t1)
	t0 = e.CyclotomicSquare(t0)
	result = e.Mul(result, t0)
	result = e.nSquareCompressed(result, 92)
	result = e.DecompressKarabina(result)

	return result

}

// ExpX0Plus1 set z to z^(x₀+1) in E6 and return z
// x₀+1 = 91893752504881257682351033800651177985
func (e Ext6) ExpX0Plus1(z *E6) *E6 {
	result := e.ExpX0Minus1(z)
	t := e.CyclotomicSquare(z)
	result = e.Mul(result, t)
	return result
}

// ExpX0Minus1Div3 set z to z^(x₀-1)/3 in E6 and return z
// (x₀-1)/3 = 3195374304363544576
func (e Ext6) ExptMinus1Div3(z *E6) *E6 {
	z = e.Reduce(z)
	result := e.Copy(z)
	result = e.CyclotomicSquare(result)
	result = e.Mul(result, z)
	t0 := e.Mul(result, z)
	t0 = e.CyclotomicSquare(t0)
	result = e.Mul(result, t0)
	t0 = result
	t0 = e.nSquareCompressed(t0, 7)
	t0 = e.DecompressKarabina(t0)
	result = e.Mul(result, t0)
	result = e.nSquareCompressed(result, 5)
	result = e.DecompressKarabina(result)
	result = e.Mul(result, z)
	result = e.nSquareCompressed(result, 46)
	result = e.DecompressKarabina(result)

	return result
}

// ExpC1 set z to z^C1 in E6 and return z
// ht, hy = 13, 9
// C1 = (ht+hy)/2 = 11
func (e Ext6) ExpC1(z *E6) *E6 {
	z = e.Reduce(z)
	result := e.CyclotomicSquare(z)
	result = e.Mul(result, z)
	t0 := e.Mul(z, result)
	t0 = e.CyclotomicSquare(t0)
	result = e.Mul(result, t0)

	return result
}

// ExpC2 set z to z^C2 in E6 and return z
// ht, hy = 13, 9
// C2 = (ht**2+3*hy**2)/4 = 103
func (e Ext6) ExpC2(z *E6) *E6 {
	z = e.Reduce(z)

	result := e.CyclotomicSquare(z)
	result = e.Mul(result, z)
	t0 := result
	t0 = e.nSquareCompressed(t0, 4)
	t0 = e.DecompressKarabina(t0)
	result = e.Mul(result, t0)
	result = e.CyclotomicSquare(result)
	result = e.Mul(result, z)

	return result
}

// MulBy014 multiplies z by an E6 sparse element of the form
//
//	E6{
//		B0: E3{A0: c0, A1: c1, A2: 0},
//		B1: E3{A0: 0,  A1: 1,  A2: 0},
//	}
func (e *Ext6) MulBy014(z *E6, c0, c1 *baseEl) *E6 {

	a := e.MulBy01(&z.B0, c0, c1)

	var b E3
	// Mul by E3{0, 1, 0}
	b.A0 = *mulFpByNonResidue(e.fp, &z.B1.A2)
	b.A2 = z.B1.A1
	b.A1 = z.B1.A0

	one := e.fp.One()
	d := e.fp.Add(c1, one)

	zC1 := e.Ext3.Add(&z.B1, &z.B0)
	zC1 = e.Ext3.MulBy01(zC1, c0, d)
	zC1 = e.Ext3.Sub(zC1, a)
	zC1 = e.Ext3.Sub(zC1, &b)
	zC0 := e.Ext3.MulByNonResidue(&b)
	zC0 = e.Ext3.Add(zC0, a)

	return &E6{
		B0: *zC0,
		B1: *zC1,
	}
}

//	multiplies two E6 sparse element of the form:
//
//	E6{
//		B0: E3{A0: c0, A1: c1, A2: 0},
//		B1: E3{A0: 0,  A1: 1,  A2: 0},
//	}
//
// and
//
//	E6{
//		B0: E3{A0: d0, A1: d1, A2: 0},
//		B1: E3{A0: 0,  A1: 1,  A2: 0},
//	}
func (e Ext6) Mul014By014(d0, d1, c0, c1 *baseEl) [5]*baseEl {
	one := e.fp.One()
	x0 := e.fp.Mul(c0, d0)
	x1 := e.fp.Mul(c1, d1)
	tmp := e.fp.Add(c0, one)
	x04 := e.fp.Add(d0, one)
	x04 = e.fp.Mul(x04, tmp)
	x04 = e.fp.Sub(x04, x0)
	x04 = e.fp.Sub(x04, one)
	tmp = e.fp.Add(c0, c1)
	x01 := e.fp.Add(d0, d1)
	x01 = e.fp.Mul(x01, tmp)
	x01 = e.fp.Sub(x01, x0)
	x01 = e.fp.Sub(x01, x1)
	tmp = e.fp.Add(c1, one)
	x14 := e.fp.Add(d1, one)
	x14 = e.fp.Mul(x14, tmp)
	x14 = e.fp.Sub(x14, x1)
	x14 = e.fp.Sub(x14, one)

	zC0B0 := e.fp.Add(one, one)
	zC0B0 = e.fp.Add(zC0B0, zC0B0)
	zC0B0 = e.fp.Neg(zC0B0)

	zC0B0 = e.fp.Add(zC0B0, x0)

	return [5]*baseEl{zC0B0, x01, x1, x04, x14}
}

// MulBy01245 multiplies z by an E6 sparse element of the form
//
//	E6{
//		C0: E3{B0: c0, B1: c1, B2: c2},
//		C1: E3{B0: 0, B1: c4, B2: c5},
//	}
func (e *Ext6) MulBy0645(z *E6, x [5]*baseEl) *E6 {
	c0 := &E3{A0: *x[0], A1: *x[1], A2: *x[2]}
	c1 := &E3{A0: *e.fp.Zero(), A1: *x[3], A2: *x[4]}
	a := e.Ext3.Add(&z.B0, &z.B1)
	b := e.Ext3.Add(c0, c1)
	a = e.Ext3.Mul(a, b)
	b = e.Ext3.Mul(&z.B0, c0)
	c := e.Ext3.MulBy12(&z.B1, x[3], x[4])
	z1 := e.Ext3.Sub(a, b)
	z1 = e.Ext3.Sub(z1, c)
	z0 := e.Ext3.MulByNonResidue(c)
	z0 = e.Ext3.Add(z0, b)
	return &E6{
		B0: *z0,
		B1: *z1,
	}
}
