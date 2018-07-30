# Model

* 與 database table 對應的 struct
* struct 轉成 table 以蛇形命名法（snake_case）命名

## gorm.Model

預設 struct

```go

type Model struct {
    ID        uint `gorm:"primary_key"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time
}


```

### 嵌入預設 Model

```go

type Account struct {
    gorm.Model
    Password stringsss
}


```

## 建立 Model

```go

type Student struct {       // table: table
    ID int                  // column: id
    Name string             // column: name
    Birthday time.Time      // column: birthday
    Gender byte             // column: gender
    CreateID string         // column: create_id 
    CreateTime time.Time    // column: create_time
}

```

### 自定義 table 名稱

#### 多元化名稱

修改 default table 名稱

```go

// default table name student
type Student struct {}

// 自定義 table 名稱
func (s Student) TableName() string {
  return "custom_student"
}

// 啟用多元 table 名稱
// true: 啟用預設 table 名稱，struct `Student` 預設為 `Student`，TableName() 方法將不會被啟用
db.SingularTable(true)

```

#### 異動預設 table 名稱

```go

// 修改預設 table 名稱
gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
    return "prefix_" + defaultTableName";
}

```

### 自定義 column 名稱

```go

type Student struct {       
    ID int                  `gorm:"column:student_id"`      // column: student_id
    Name string             `gorm:"column:student_name"`    // column: student_name
    Birthday time.Time      `gorm:"column:birthday"`        // column: birthday
    Gender byte             `gorm:"column:sex"`             // column: sex
    CreateID string         `gorm:"column:create_id"`       // column: create_id 
    CreateTime time.Time    `gorm:"column:create_date"`     // column: create_date
}

```

----

[struct](/Common/Note/Struct.md)