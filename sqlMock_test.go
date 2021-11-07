package go_UT_learn

import (
	/*
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	 */
	"gorm.io/gorm"
)
// sqlMock, 无需关注db连接，模拟实现了sql driver mock


type Goods struct {
	name string
}
type Repository struct {
	db *gorm.DB
}


func(p *Repository)ListAll()([]*Goods, error) {
	var goods []*Goods
	err := p.db.Find(&goods).Error
	return goods, err
}

/*
func TestSql(t *testing.T) {
	// 使用 sqlmock.New() 创建 *sql.DB 的模拟实例和模拟控制器
	db, mock, err := sqlmock.New()
	if nil != err {
		log.Fatalf("Init sqlmock failed, err %v", err)
	}
	// create gorm link with sqlmock
	dbLink, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn: db,
	}), &gorm.Config{})
	if nil != err {
		log.Fatalf("Init DB with sqlmock failed, err %v", err)
	}

	repo := Repository{
		db: dbLink,
	}

	// create mock record
	const sqlSelectAll = `SELECT * FROM "blogs"`
	mock.ExpectQuery(sqlSelectAll).WillReturnRows(sqlmock.NewRows(nil))

	goods, err := repo.ListAll()
	assert.Nil(t, err)
	assert.Nil(t, goods)

}

*/