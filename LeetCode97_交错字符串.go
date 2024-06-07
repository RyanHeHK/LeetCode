package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	ls1 := len(s1)
	ls2 := len(s2)
	ls3 := len(s3)
	if ls1+ls2 != ls3 {
		return false
	}
	f := make([][]bool, ls1)
	for i := 0; i <= ls1; i++ {
		f[i] = make([]bool, ls2+1)
	}
	for {

	}
	return true
}
