package main

import (
	"net/http"

	"io"
	"io/ioutil"
	"os"
	"time"

	//"fmt"
	"html/template"
	"log"

	"github.com/julienschmidt/httprouter"
)

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("C:/Users/z50001851/Desktop/MyGO/src/GO语言实战流媒体视频网站/streamsever/upload.html")
	t.Execute(w, nil)
}

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//获取参数信息
	vid := p.ByName("vid-id")
	//得到完整链接，打开链接
	vl := VIDEO_DIR + vid

	video, err := os.Open(vl)

	if err != nil {
		log.Printf("Error when try to open files : %v", err)
		sendErrorResponse(w, "Interal server error", http.StatusInternalServerError)
	}
	//设置输出格式
	w.Header().Set("Content-Type", "video/mp4")
	//http发送内容
	http.ServeContent(w, r, "", time.Now(), video)

	defer video.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//先校验file的大小是否超过限制
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, "File is too big", http.StatusInternalServerError)
	}
	//从表单中拿到file文件
	file, _, err := r.FormFile("file")
	if err != nil {
		sendErrorResponse(w, "Internal", http.StatusInternalServerError)
	}
	//从file文件中读成2进制文件
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		sendErrorResponse(w, "Read file error", http.StatusInternalServerError)
	}
	//把文件写入上传
	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+fn, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, "write file error", http.StatusInternalServerError)
		return
	}
	//上传成功后返回正确码
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")
}
