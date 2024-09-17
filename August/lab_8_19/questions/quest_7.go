package questions

type Rect struct {
    X      int
    Y      int
    Area   float64
    Perimeter float64
}

func CalculateArea(r *Rect) float64 {
    return float64(r.X * r.Y)
}

func CalculatePerimeter(r *Rect) float64 {
    return 2 * float64(r.X + r.Y)
}

func (r *Rect) CalculateAreaMethod() {
    r.Area = CalculateArea(r)
}

func (r *Rect) CalculatePerimeterMethod() {
    r.Perimeter = CalculatePerimeter(r)
}