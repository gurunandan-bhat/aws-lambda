package dtypes

type Category struct {
	PK                string
	SK                string
	IPCatID           uint32
	VName             string
	VURLName          string
	IParentID         uint32
	VShortDescription *string
	MImages           Images
	CTypeStatus       string
	LAttributes       []CategoryAttribute
	LChildren         []CategorySummary
}

type CategorySummary struct {
	IPCatID    uint32  `db:"iPCatID"`
	VName      string  `db:"vName"`
	VURLName   string  `db:"vUrlName"`
	IParentID  uint32  `db:"iParentID"`
	VShortDesc *string `db:"vShortDesc"`
	Images
	CStatus    string              `db:"cStatus"`
	Attributes []CategoryAttribute `db:"-"`
	Children   []CategorySummary   `db:"-"`
}

type Images struct {
	VSmallImage       *string `db:"vSmallImage" json:"vSmallImage"`
	VSmallImageAltTag *string `db:"vSmallImage_AltTag" json:"vSmallImage_AltTag"`
	VImage            *string `db:"vImage" json:"vImage"`
	VImageAltTag      *string `db:"vImage_AltTag" json:"vImage_AltTag"`
}

type CategoryAttribute struct {
	IAttribDatID uint32  `db:"iAttribDatID" json:"iAttribDatID" diff:"iAttribDatID"`
	IPCatID      uint32  `db:"iPCatID" json:"iPCatID" diff:"iPCatID"`
	IAttribID    uint32  `db:"iAttribID" json:"iAttribID" diff:"iAttribID"`
	VAttribName  *string `db:"vAttribName" json:"vAttribName" diff:"-"`
	VName        *string `db:"vName" json:"vName" diff:"vName"`
	IRank        int     `db:"iRank" json:"iRank" diff:"iRank"`
}

type ProductValue struct {
	PK                string
	SK                string
	IProdID           uint32
	IPCatID           uint32
	CCode             *string
	VName             string
	VURLName          string
	VCategoryName     string
	VCategoryURLName  string
	VShortDescription *string
	VDescription      *string
	MPrices           ProdPrice
	MImages           Images
	CTypeStatus       string
	VYTID             *string
	LAttributes       []ProductAttribute
	LSKUs             []SKU
}

type ProdPrice struct {
	FRetailPrice      float64 `db:"fRetailPrice" json:"fRetailPrice"`
	FRetailOPrice     float64 `db:"fRetailOPrice" json:"fRetailOPrice"`
	FShipping         float64 `db:"fShipping" json:"fShipping"`
	FPrice            float64 `db:"fPrice" json:"fPrice"`
	FOPrice           float64 `db:"fOPrice" json:"fOPrice"`
	FActualWeight     float64 `db:"fActualWeight" json:"fActualWeight"`
	FVolumetricWeight float64 `db:"fVolumetricWeight" json:"fVolumetricWeight"`
}

type ProductAttribute struct {
	IProdAttribID uint32  `db:"iProdAttribID" json:"iProdAttribID" diff:"iProdAttribID"`
	IProdID       uint32  `db:"iProdID" json:"iProdID" diff:"iProdID"`
	IAttribID     uint32  `db:"iAttribID" json:"iAttribID" diff:"iAttribID"`
	VAttribName   *string `db:"vAttribName" json:"vAttribName" diff:"-"`
	VValue        *string `db:"vValue" json:"vValue" diff:"vValue"`
	IPCID         uint32  `db:"iAttribPCID" json:"iAttribPCID" diff:"iPCID"`
	FRetailPrice  float64 `db:"fRetailPrice" json:"fRetailPrice" diff:"fRetailPrice"`
	FRetailOPrice float64 `db:"fRetailOPrice" json:"fRetailOPrice" diff:"fRetailOPrice"`
	FPrice        float64 `db:"fPrice" json:"fPrice" diff:"fPrice"`
	FOPrice       float64 `db:"fOPrice" json:"fOPrice" diff:"fOPrice"`
	CDefault      *string `db:"cDefault" json:"cDefault" diff:"cDefault"`
	CStock        *string `db:"cStock" json:"cStock" diff:"cStock"`
}

type SKUAttrib struct {
	VAttribName  string
	VAttribValue string
}
type SKU struct {
	Attributes    []SKUAttrib
	FRetailPrice  float64
	FRetailOPrice float64
	FPrice        float64
	FOPrice       float64
	CDefault      string
	CStock        string
}
