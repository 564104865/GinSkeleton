package upload_file

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/md5_encrypt"
	"goskeleton/app/utils/response"
	"goskeleton/app/utils/snow_flake"
	"net/http"
	"path"
)

func Upload(context *gin.Context, save_path string) bool {

	//  1.获取上传的文件名(参数验证器已经验证完成了第一步错误，这里简化)
	file, _ := context.FormFile(variable.UploadFileField) //  file 是一个文件结构体（文件对象）

	//  保存文件，原始文件名进行全局唯一编码加密、md5 加密，保证在后台存储不重复
	var saveErr error
	if unique_id, err := snow_flake.CreateSnowFlakeFactory().GetId(); err == nil {
		saveFileName := fmt.Sprintf("%d%s", unique_id, file.Filename)
		saveFileName = md5_encrypt.MD5([]byte(saveFileName)) + path.Ext(saveFileName)
		if saveErr = context.SaveUploadedFile(file, save_path+saveFileName); saveErr == nil {
			//  上传成功,返回资源的存储路径，这里请根据实际返回绝对路径或者相对路径
			succ := gin.H{
				"path": save_path + saveFileName,
			}
			response.ReturnJson(context, http.StatusCreated, consts.CurdStatusOkCode, consts.CurdStatusOkMsg, succ)
			return true
		}
	} else {
		saveErr = errors.New(my_errors.ErrorsSnowflakeGetIdFail)
	}

	response.ReturnJson(context, http.StatusBadRequest, consts.FilesUploadFailCode, consts.FilesUploadFailMsg+", 文件保存失败!", saveErr.Error())
	return false

}