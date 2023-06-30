package utils

type IData struct {
	ID int
}

func IsExist(data *IData) bool {
	return data != nil && data.ID > 0
}
