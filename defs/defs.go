package defs

type GroupCommon struct {
	Name string `form:"name"`
}

type GroupAdd struct {
	GroupCommon
}
