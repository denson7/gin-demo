package utils

import (
	"bytes"
	"errors"
	"github.com/valyala/fasthttp"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func PostRawData(url string,data []byte) ([]byte,error){
	request := fasthttp.AcquireRequest()
	request.SetRequestURI(url)
	request.Header.SetMethod("POST")
	request.SetBody(data)
	response := fasthttp.AcquireResponse()
	err := requestFellowRedirect(request,response)
	if err != nil{
		return nil,err
	}
	return response.Body(),err
}

func PostFields(url string,fields map[string]string) ([]byte,error){
	//新建一个缓冲，用于存放文件内容
	bodyBufer := &bytes.Buffer{}
	//创建一个multipart文件写入器，方便按照http规定格式写入内容
	bodyWriter := multipart.NewWriter(bodyBufer)
	err := writeFieldsData(bodyWriter,fields)
	if err != nil{
		bodyWriter.Close()
		return nil,err
	}
	//关闭bodyWriter停止写入数据
	bodyWriter.Close()
	contentType := bodyWriter.FormDataContentType()
	//构建request，发送请求
	req := fasthttp.AcquireRequest()
	req.Header.SetContentType(contentType)
	//直接将构建好的数据放入post的body中
	req.SetBody(bodyBufer.Bytes())
	req.Header.SetMethod("POST")
	req.SetRequestURI(url)
	resp := fasthttp.AcquireResponse()
	err = requestFellowRedirect(req,resp)
	if err != nil{
		return nil,err
	}
	return resp.Body(),nil
}
//支持多文件上传
func PostFiles(url string,files map[string]string) ([]byte,error){
	//新建一个缓冲，用于存放文件内容
	bodyBufer := &bytes.Buffer{}
	//创建一个multipart文件写入器，方便按照http规定格式写入内容
	bodyWriter := multipart.NewWriter(bodyBufer)

	err := writeFilesData(bodyWriter,files)
	if err != nil{
		bodyWriter.Close()
		return nil,err
	}
	//关闭bodyWriter停止写入数据
	bodyWriter.Close()
	contentType := bodyWriter.FormDataContentType()
	//构建request，发送请求
	request := fasthttp.AcquireRequest()
	request.Header.SetContentType(contentType)
	//直接将构建好的数据放入post的body中
	request.SetBody(bodyBufer.Bytes())
	request.Header.SetMethod("POST")
	request.SetRequestURI(url)
	response := fasthttp.AcquireResponse()
	err = requestFellowRedirect(request,response)
	if err != nil{
		return nil,err
	}
	return response.Body(),nil
}

//支持多文件上传
func PostFilesAndFields(url string,files map[string]string,fields map[string]string) ([]byte,error){
	//新建一个缓冲，用于存放文件内容
	bodyBufer := &bytes.Buffer{}
	//创建一个multipart文件写入器，方便按照http规定格式写入内容
	bodyWriter := multipart.NewWriter(bodyBufer)

	err := writeFilesData(bodyWriter,files)
	if err != nil{
		bodyWriter.Close()
		return nil,err
	}
	err = writeFieldsData(bodyWriter,fields)
	if err != nil{
		bodyWriter.Close()
		return nil,err
	}
	//关闭bodyWriter停止写入数据
	bodyWriter.Close()
	contentType := bodyWriter.FormDataContentType()
	//构建request，发送请求
	request := fasthttp.AcquireRequest()
	request.Header.SetContentType(contentType)
	//直接将构建好的数据放入post的body中
	request.SetBody(bodyBufer.Bytes())
	request.Header.SetMethod("POST")
	request.SetRequestURI(url)
	response := fasthttp.AcquireResponse()
	err = requestFellowRedirect(request,response)
	if err != nil{
		return nil,err
	}
	return response.Body(),nil
}

func Get(url string) ([]byte,error){
	client := fasthttp.Client{}
	_,resp,err := client.Get(nil,url)
	return resp,err
}

//支持重定向文件下载
func DownLoadFileByGet(url,savePath string) (string,error){
	//fasthttp.Do不支持自动重定向，此处自己实现
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	resp := fasthttp.AcquireResponse()
	err := requestFellowRedirect(req,resp)
	if err != nil{
		return "",err
	}
	return saveFile(savePath,resp)
}
//用于需要上传form表单下载文件的情况
func DownloadFileByPostFields(url,savePath string,fields map[string]string) (string,error){
	//新建一个缓冲，用于存放文件内容
	bodyBufer := &bytes.Buffer{}
	//创建一个multipart文件写入器，方便按照http规定格式写入内容
	bodyWriter := multipart.NewWriter(bodyBufer)
	err := writeFieldsData(bodyWriter,fields)
	if err != nil{
		bodyWriter.Close()
		return "",err
	}
	//关闭bodyWriter停止写入数据
	bodyWriter.Close()
	contentType := bodyWriter.FormDataContentType()
	//构建request，发送请求
	req := fasthttp.AcquireRequest()
	req.Header.SetContentType(contentType)
	//直接将构建好的数据放入post的body中
	req.SetBody(bodyBufer.Bytes())
	req.Header.SetMethod("POST")
	req.SetRequestURI(url)
	resp := fasthttp.AcquireResponse()
	err = requestFellowRedirect(req,resp)
	if err != nil{
		return "",err
	}
	return saveFile(savePath,resp)
}
//用于需要上传json串等数据下载文件的情况
func DownloadFileByPostRawData(url,savePath string,rawData []byte)(string,error){
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.SetBody(rawData)
	resp := fasthttp.AcquireResponse()
	err := requestFellowRedirect(req,resp)
	if err != nil{
		return "",err
	}
	return saveFile(savePath,resp)
}

func writeFieldsData(bodyWriter *multipart.Writer,fields map[string]string) error{
	for k,v := range fields{
		fieldWriter,err := bodyWriter.CreateFormField(k)
		if err != nil{
			return err
		}
		fieldWriter.Write([]byte(v))
	}
	return nil
}

func writeFilesData(bodyWriter *multipart.Writer,files map[string]string) error{
	for key,val := range files{
		//从bodyWriter生成fileWriter,并将文件内容写入fileWriter,多个文件可进行多次
		fileWriter,err := bodyWriter.CreateFormFile(key,path.Base(val))
		if err != nil{
			return err
		}
		file,err := os.Open(val)
		if err != nil{
			return err
		}
		//不要忘记关闭打开的文件
		_,err = io.Copy(fileWriter,file)
		if err != nil{
			return err
		}
		file.Close()
	}
	return nil
}
//发出请求，如果返回重定向则继续请求重定向地址，直到返回非重定向结果
func requestFellowRedirect(req *fasthttp.Request,resp *fasthttp.Response)error{
	for{
		err := fasthttp.Do(req,resp)
		if err != nil{
			return err
		}
		statusCode := resp.StatusCode()
		if statusCode != fasthttp.StatusMovedPermanently &&
			statusCode !=  fasthttp.StatusFound &&
			statusCode !=  fasthttp.StatusSeeOther &&
			statusCode !=  fasthttp.StatusTemporaryRedirect &&
			statusCode !=  fasthttp.StatusPermanentRedirect {
			break
		}
		location := ""
		//获取重定向链接地址
		resp.Header.VisitAll(func (key,val []byte){
			if (string(key) == "Location"){
				location = string(val)
			}
		})

		if location == ""{
			return errors.New("redirect response without location")
		}
		resp.Reset()
		req.SetRequestURI(location)
	}
	return nil
}
//根据http请求的response保存文件
func saveFile(savePath string,resp *fasthttp.Response)(string,error){
	//先返回文件下载地址
	fileName := ""
	//从header头的Content-disposition获得文件名
	resp.Header.VisitAll(func (key,val []byte){
		if (string(key) == "Content-Disposition"){
			infoStr := string(val)
			stringArr := strings.Split(infoStr,";")
			for _,str := range stringArr{
				if  strings.Index(str,"filename") == 1{
					strArr := strings.Split(str,"=")
					fileName = strings.Trim(strArr[1],"\"")
					break
				}
			}
		}
	})
	if fileName == ""{
		strArr := strings.Split(string(resp.Header.ContentType()),"/")
		fileName = GetRandomStr(30)+"."+strArr[1]
	}
	//有时候上传图片的名称中包含路径名
	fileName =filepath.Base(fileName)
	if savePath == ""{
		savePath = os.TempDir()
	}
	fullFile := savePath+"/"+fileName
	file,err := os.Create(fullFile)
	defer file.Close()
	if err != nil{
		return "",err
	}
	_,err = file.Write(resp.Body())
	if err != nil{
		return "",nil
	}
	return fullFile,nil
}