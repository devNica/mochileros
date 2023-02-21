package serializers

import (
	"bytes"
	"io"
	"mime/multipart"
	"strings"

	"github.com/devNica/mochileros/commons"
	"github.com/devNica/mochileros/dto/request"
)

func SerializeUserAsset(forms map[string][]*multipart.FileHeader) ([]request.FileRequestModel, error) {

	var images []request.FileRequestModel

	for _, files := range forms {

		for _, file := range files {

			index := strings.Index(file.Filename, ".")
			filename := file.Filename[0:index]
			assetsType := commons.GetAssetDataDictionary()
			assetTypeId := commons.GetAssetTypeId(filename, assetsType)

			f, err := file.Open()

			if err != nil {
				return []request.FileRequestModel{}, err
			}

			defer f.Close()

			buf := bytes.NewBuffer(nil)
			io.Copy(buf, f)

			images = append(images, request.FileRequestModel{
				Buffer:      buf.Bytes(),
				Filetype:    file.Header.Get("Content-Type"),
				Filesize:    int(file.Size),
				AssetTypeId: assetTypeId,
			})
		}

	}

	return images, nil
}
