package repositories

// type BlockSQL struct {
// 	DB *gorm.DB
// }

// func (b *BlockSQL) List() ([]models.Block, error) {
// 	var blocks []models.Block
// 	if err := b.DB.Find(&blocks).Error; err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
// 	return blocks, nil
// }

// func (b *BlockSQL) Create(block models.Block) (models.Block, error) {
// 	if err := b.DB.Create(&block).Error; err != nil {
// 		log.Println(err)
// 		return models.Block{}, err
// 	}
// 	return block, nil
// }
