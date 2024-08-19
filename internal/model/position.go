package model

type PositionCreateUpdateBase struct {
	PicUrl    string 
	GoodsName string 
	Link      string 
	Sort      int    
	GoodsId   int   
}

type PositionInfoCreateInput struct {
	PositionCreateUpdateBase
}

type PositionInfoCreateOutput struct {
	Id int
}
