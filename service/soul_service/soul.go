package soul_service

import "go-eth/models"

/**
获取消息
*/
func GetSouls(maps map[string]interface{}, offset int, limit int) ([]models.Soul, error) {
	souls, err := models.GetSouls(maps, offset, limit)
	if err != nil {
		return nil, err
	}
	return souls, nil
}
