package global

type GetScale struct {
	MinS int
	MaxS int
}

var GetScaleReq = &GetScale{
	MinS: MinScale,
	MaxS: MaxScale,
}
