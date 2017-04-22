package recocido

import (
	"image/color"
	"container/list"
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"github.com/gonum/plot/vg/draw"
)

// GraficaCosto crea una gráfica con el costo de las
// soluciones aceptadas.
func GraficaCosto(lst *list.List, l int) {
	pts := make(plotter.XYs, lst.Len() / l)
	i := 0
	j := 0
	for f := lst.Front(); f != nil; f = f.Next() {
		if (i % l == 0) && j < (lst.Len() / l) {
			pts[j].X = float64(i)
			pts[j].Y = f.Value.(float64)
			j++
		}
		i++
	}		

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Gráfica de Costos de Soluciones Aceptadas"
	p.Y.Label.Text = "Costo"
	p.Add(plotter.NewGrid())
	s, err := plotter.NewScatter(pts)
	if err != nil {
		panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}

	p.Add(s)

	if err := p.Save(18*vg.Inch, 10*vg.Inch, "costos.svg"); err != nil {
		panic(err)
	}
}

// GraficaMejoresSols crea una gráfica con el costo de las
// soluciones por número de semilla utilizado y muestra si
// son factibles o no.
func GraficaMejoresSols(b []bool, f []float64, nFac int, nNFac int, nom string) {
	ptsFac := make(plotter.XYs, nFac)
	ptsNFac := make(plotter.XYs, nNFac)

	j := 0
	k := 0
	for i := 0; i < len(b); i++ {
		if b[i] {
			ptsFac[j].X = float64(i)
			ptsFac[j].Y = f[i]
			j++
		} else {
			ptsNFac[k].X = float64(i)
			ptsNFac[k].Y = f[i]
			k++
		}
	}		

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	
	p.Title.Text = "Gráfica de mejores soluciones"
	p.Y.Label.Text = "Costo"
	p.X.Label.Text = "Semilla"
	
	p.Add(plotter.NewGrid())
	s1, err := plotter.NewScatter(ptsFac)
	s2, err := plotter.NewScatter(ptsNFac)
	if err != nil {
		panic(err)
	}
	s1.GlyphStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	s2.GlyphStyle.Color = color.RGBA{A: 255}
	s1.Shape = draw.CircleGlyph{}
	s2.Shape = draw.CircleGlyph{}
	
	p.Add(s1, s2)
	p.Legend.Add("Factibles", s1)
	p.Legend.Add("No factibles", s2)

	if err := p.Save(15*vg.Inch, 10*vg.Inch, nom); err != nil {
		panic(err)
	}
}
