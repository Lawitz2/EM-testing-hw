package hw6_test

import (
	"errors"
	"github.com/lawitz2/em-testing-hw/hw6"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestDB(t *testing.T) {
	// создаем мок базы данных
	ctrl := gomock.NewController(t)
	db := hw6.NewMockDB(ctrl)

	user1 := hw6.User{"Alice", 25}
	user2 := hw6.User{"Bob", 15}

	// говорим моку базы что именно мы ожидаем получить
	db.EXPECT().AddToDb(user1).Return(nil).Times(1)

	expectedErr := errors.New("err adding to database")
	db.EXPECT().AddToDb(user2).Return(expectedErr).Times(1)

	// делаем вызовы к моку бд так, будто обращаемся к реальной бд
	err := db.AddToDb(user1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = db.AddToDb(user2)
	if err == nil {
		t.Errorf("Expected error: %v", expectedErr)
	}
}
