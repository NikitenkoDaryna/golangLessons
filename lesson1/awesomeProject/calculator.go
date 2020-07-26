package main

func sum(x, y int) int {
	return x + y
}
func divide(x,y float64) float64{
	return x/y
}
func Calculator(x int) func(int) int {
	return func(f int) int {
		return sum(x, f)
	}

}
func Division(n float64) func(float64) float64{
	return func(t float64) float64 {
		return divide(n,t)
	}
}
