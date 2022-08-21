package http_dto

type DatasetCreateReq struct {
	Name         string `json:"name" binding:"required"`
	Mid          int    `json:"mid" binding:"required"`
	UserId       int    `json:"userId" binding:"required"`
	CenterId     int    `json:"centerId" binding:"required"`
	Desc         string `json:"desc" binding:"required"`
	NonPublic    bool   `json:"nonPublic" binding:"required"`
	VolumeId     int    `json:"volumeId"`
	ProviderType int    `json:"providerType" binding:"required"`
	SubPath      int    `json:"subPath"`
}

type DatasetCreateResp struct {
	Name         string `json:"name"`
	Mid          int    `json:"mid"`
	NonPublic    bool   `json:"nonPublic"`
	VolumeId     int    `json:"volumeId,omitempty"`
	ProviderType int    `json:"providerType"`
	SubPath      int    `json:"subPath,omitempty"`
	Sid          int    `json:"sid" binding:"required"`
}
