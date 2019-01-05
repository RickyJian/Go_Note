# FK&AssociationFK

Foreign Key：定義關聯，被動，一定存在，預設為被嵌入的 model 型別再加上它的 FK

Association ForeignKey：幫助關聯將質一併寫入雙 model，主動，預設為嵌入的 model FK 的值會當作 FK 的值

## 關聯：一對一

### Belongs To

```go

// 搜尋語法
db.Model(&role).Related(&account)
// SQL：SELECT * FROM account where account_id = XXX

```

#### 預設

```go

type Account struct {
    gorm.Model
    Name    string
    Gender  byte
}

// `Role` Belongs to `Accoount`，AccountID is the FK 
type Role struct {
    gorm.Model
    Name    string
    Account     Account
    AccountID   string
}

```

#### 指定 FK

```go

type Account struct {
    gorm.Model
    Name    string
    Gender  byte
}

// `Role` Belongs to `Accoount`，AccountID is the FK 
type Role struct {
    gorm.Model
    Name        string
    Account     Account `gorm:"foreignkey:AccountRefer"` // ，AccountRefer 當做 FK
    AccountRefer string
}

```

#### 指定 AssociationFK

```go

type Account struct {
    gorm.Model
    Name    string
    Gender  byte
    Refer   string
}

// `Role` Belongs to `Accoount`，AccountID is the FK 
type Role struct {
    gorm.Model
    Name        string
    Account     Account `gorm:"association_foreignkey:Refer"` // Refer 當做 AssociationFK
    AccountID   string
}

```

#### 指定 FK&AssociationFK

```go

type Account struct {
    gorm.Model
    Name    string
    Gender  byte
    Refer   string
}

// `Role` Belongs to `Accoount`，AccountID is the FK 
type Role struct {
    gorm.Model
    Name        string
    Account     Account `gorm:"foreignkey:AccountID;association_foreignkey:Refer"` //  AccountID 當做 FK，Refer 當做 AssociationFK
    AccountID   string
}

```

### Has One

與 `belongsto` 不同的是此 model 擁有另一個 model


```go

// 搜尋語法
db.Model(&role).Related(&account)
// SQL：SELECT * FROM account where account_id = XXX


```

#### 預設

```go

// `Account` has one `Role`
type Account struct {
    gorm.Model
    Name    string
    Gender  byte
    Role    Role
    RoleID  uint 
}

type Role struct {
    gorm.Model
    Name    string
}

```

#### 指定 FK

```go

// `Account` has one `Role`
type Account struct {
    gorm.Model
    Name    string
    Gender  byte
    Role    Role    `gorm:"foreignkey:RoleRefer"`
    RoleRefer uint 
}

type Role struct {
    gorm.Model
    Name    string
}

```

#### 指定 AssociationFK

```go

// `Account` has one `Role`
type Account struct {
    gorm.Model
    Name    string
    Gender  byte
    Role    Role    `gorm:"association_foreignkey:Refer"`
    RoleID  uint
}

type Role struct {
    gorm.Model
    Refer   uint
    Name    string
}

```

#### 指定 FK&AssociationFK

```go

// `Account` has one `Role`
type Account struct {
    gorm.Model
    Name    string
    Gender  byte
    Role    Role    `gorm:"foreignkey:RoleID;association_foreignkey:Refer"`
    RoleID  uint
}

type Role struct {
    gorm.Model
    Refer   uint
    Name    string
}

```

## 關聯：一對多

### Has Many

```go

db.Model(&user).Related(&email)
// SELECT * FROM email where user_id = 111

```

#### 預設

```go

// One
type User struct {
    gorm.Model
    Emails  []Email
}

// Many
type Email struct {
    gorm.Model
    Email   string
    UserID  uint
}

```

#### 指定 FK

```go

// One
type User struct {
    gorm.Model
    Emails  []Email `gorm:"foreignkey:UserRefer"`
}

// Many
type Email struct {
    gorm.Model
    Email   string
    UserRefer  uint 
}

```

#### 指定 Association FK


```go

// One
type User struct {
    gorm.Model
    Refer   uint
    Emails  []Email `gorm:"association_foreignkey:Refer"`
}

// Many
type Email struct {
    gorm.Model
    Email   string
    UserID  uint 
}

```


#### 指定 FK&Association FK

```go

// One
type User struct {
    gorm.Model
    Refer   uint
    Emails  []Email `gorm:"foreignkey:UserID;association_foreignkey:Refer"`
}

// Many
type Email struct {
    gorm.Model
    Email   string
    UserID  uint 
}

```

## 關聯：多對多

### Many To Many

#### 單方綁定

```go

type User struct {
    gorm.Model
    Languages         []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
    gorm.Model
    Name string
}

```

#### 雙方綁定

```go

type User struct {
    gorm.Model
    Languages         []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
    gorm.Model
    Name string
    Users             []User     `gorm:"many2many:user_languages;"`
}

```

#### 指定 FK&Association FK

```go

type Role struct {
    RoleID string `gorm:"primary_key"`
    Name string
}

type Account struct {
    AccountID string `gorm:"primary_key"`
    Roles []Role `gorm:"many2many:AccountRole;association_foreignkey:RoleID;foreignkey:AccountID"`
}


```

#### Jointable ForeignKey

倘若你想更改 join table's foreign key ， 須要使用 `association_jointable_foreignkey` , `jointable_foreignkey`

```go

type Role struct {
    RoleID string `gorm:"primary_key"`
    Name string
}

type Account struct {
    AccountID string `gorm:"primary_key"`
    Roles []Role `gorm:"many2many:AcountRole;foreignkey:AccountID;association_foreignkey:RoleID;association_jointable_foreignkey:role_id;jointable_foreignkey:account_id;"`
}


```

#### Self-Referencing

```go

type Account struct {
    AccountID string `gorm:"primary_key"`
    Friends []*Account `gorm:"many2many:friendships;association_jointable_foreignkey:friend_id"`
}

```

## 多形

Has One 、 Has Many 支援

```go

type Cat struct {
    Id    int
    Name  string
    Toy   Toy `gorm:"polymorphic:Owner;"`
}

type Dog struct {
    Id   int
    Name string
    Toy  Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
    Id        int
    Name      string
    OwnerId   int
    OwnerType string
}

```