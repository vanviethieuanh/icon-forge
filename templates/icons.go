package {{.Env.PKG_NAME}} 

type Props struct {
	Size        int
	StrokeWidth int
}

func (p *Props) getProps(overrides Props) {
	if overrides.Size != 0 {
		p.Size = overrides.Size
	}
	if overrides.StrokeWidth != 0 {
		p.StrokeWidth = overrides.StrokeWidth
	}
}

func GetOrDefaults(overrides ...Props) Props {
	p := Props{Size: 24, StrokeWidth: 2}
	for _, o := range overrides {
		p.getProps(o)
	}
	return p
}
