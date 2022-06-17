package qo

/**
  @author: chenxi@cpgroup.cn
  @date:2022/5/19
  @description:
**/

type DownloadImgQO struct {
	FileName string
}

type UpdateImgQO struct {
	ImgLink string `binding:"required"`
}

type UpdateQO struct {
	Link string `binding:"required"`
}

type DeleteQO struct {
	Link string `binding:"required"`
}
