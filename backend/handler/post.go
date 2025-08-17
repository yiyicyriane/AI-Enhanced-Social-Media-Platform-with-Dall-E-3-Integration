package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"socialai/model"
	"socialai/service"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/pborman/uuid"
)


var (
   mediaTypes = map[string]string{
       ".jpeg": "image",
       ".jpg":  "image",
       ".gif":  "image",
       ".png":  "image",
       ".mov":  "video",
       ".mp4":  "video",
       ".avi":  "video",
       ".flv":  "video",
       ".wmv":  "video",
   }
)


func uploadHandler(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Received one upload request")


   // 1. process Request
   // 1.1 form-data text -> Post
   // 1.2 geerate id
   // 1.3 form-data file -> file
   // 1.4 file -> type


   token := r.Context().Value("user") // user:token <- jwt middleware
   claims := token.(*jwt.Token).Claims
   username := claims.(jwt.MapClaims)["username"]


   p := model.Post{
       Id:      uuid.New(),//1.1 generate id
       User:    username.(string), 
       Message: r.FormValue("message"),////1.2 form-data text --> post
   }
   file, header, err := r.FormFile("media_file")////1.3 form-data file --> file
   if err != nil {
       http.Error(w, "Media file is not available", http.StatusBadRequest)
       fmt.Printf("Media file is not available %v\n", err)
       return
   }
   // .jpg => image
   // .mp4 => video
   suffix := filepath.Ext(header.Filename)//1.4 get file type， for exemple, .jpg --> image, .mp4 --> video
   if t, ok := mediaTypes[suffix]; ok {
       p.Type = t
   } else {
       p.Type = "unknown"
   }


   // 2. call service to handle business logic
   err = service.SavePost(&p, file)
   if err != nil {
       http.Error(w, "Failed to save post to backend", http.StatusInternalServerError)
       fmt.Printf("Failed to save post to backend %v\n", err)
       return
   }


   fmt.Println("Post is saved successfully.")


   // 3. Response


}


func searchHandler(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Received one request for search")
   w.Header().Set("Content-Type", "application/json")


   // 1. process Request
   // URL -> string
   user := r.URL.Query().Get("user")
   keywords := r.URL.Query().Get("keywords")
   fmt.Println("user:" + user)
   fmt.Println("keywords:" + keywords)


   // 2. call service to handle business logic
   var posts []model.Post
   var err error
   if user != "" {
       posts, err = service.SearchPostsByUser(user)
   } else {
       posts, err = service.SearchPostsByKeywords(keywords)
   }
   if err != nil {
       http.Error(w, "Failed to read post from backend", http.StatusInternalServerError)
       fmt.Printf("Failed to read post from backend %v.\n", err)
       return
   }


   // 3. Response
   // model.POST => JSON string
   js, err := json.Marshal(posts)
   if err != nil {
       http.Error(w, "Failed to parse posts into JSON format", http.StatusInternalServerError)
       fmt.Printf("Failed to parse posts into JSON format %v.\n", err)
       return
   }
   w.Write(js)


}
