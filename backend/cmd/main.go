package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "os"
    _ "github.com/lib/pq"  // PostgreSQLドライバ
    "github.com/rs/cors"   // CORSミドルウェア
)

var db *sql.DB

type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
}

func main() {
    // 環境変数からDB情報を取得
    user := os.Getenv("POSTGRES_USER")
    password := os.Getenv("POSTGRES_PASSWORD")
    dbname := os.Getenv("POSTGRES_DB")

    // デバッグ用に環境変数をログ出力
    log.Printf("DB User: %s, DB Password: %s, DB Name: %s\n", user, password, dbname)

    // 接続文字列を構築
    connStr := "postgres://" + user + ":" + password + "@db:5432/" + dbname + "?sslmode=disable"

    var err error
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/signup", SignupHandler)
    http.HandleFunc("/signin", SigninHandler)
    // CORS設定を追加
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"}, // フロントエンドのオリジンを許可
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders:   []string{"Content-Type"},
        AllowCredentials: true,
    })

    handler := c.Handler(http.DefaultServeMux)

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", handler))
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)

    // パスワードのハッシュ化や入力バリデーションを行う
    // データベースにユーザーを追加
    _, err := db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
    if err != nil {
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("User created")
}

func SigninHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)

    // データベースからユーザーを検索して認証
    var dbUser User
    err := db.QueryRow("SELECT id, username, password FROM users WHERE username=$1", user.Username).Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password)
    if err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    // パスワードチェック（省略）
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Signed in")
}
