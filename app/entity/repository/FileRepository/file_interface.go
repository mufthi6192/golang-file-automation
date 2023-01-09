package FileRepository

import "fileAutomation/app/entity/model"

type FileInterface interface {
	InsertFileInterface(*model.ProductModels) (string, error)
}
