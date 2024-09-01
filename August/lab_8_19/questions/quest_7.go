package questions

type Rect struct {
    x      int
    y      int
    area   float64
    perimeter float64
}

func CalculateArea(r Rect) float64 {
    return float64(r.x * r.y)
}

func CalculatePerimeter(r Rect) float64 {
    return 2 * float64(r.x + r.y)
}

func (r *Rect) CalculateAreaMethod() {
    r.area = CalculateArea(*r)
}

func (r *Rect) CalculatePerimeterMethod() {
    r.perimeter = CalculatePerimeter(*r)
}