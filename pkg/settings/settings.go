package settings

type Settings struct {
	ID      string `bson:"_id,omitempty" json:"-"`
	AutoBuy bool   `bson:"autoBuyCards" json:"autoBuyCards"`
}
