package repository

import (
	"database/sql"
	"errors"
	"github.com/sculptorvoid/langi_backend/pkg/entity"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func TestTodoItemPostgres_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewWordPsql(db)

	type args struct {
		dictId int
		item   entity.Word
	}
	type mockBehavior func(args args, id int)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			input: args{
				dictId: 1,
				item: entity.Word{
					Word:        "test word",
					Translation: "test tr",
				},
			},
			want: 2,
			mock: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO words").
					WithArgs(args.item.Word, args.item.Translation).WillReturnRows(rows)

				mock.ExpectExec("INSERT INTO dictionaries_words").WithArgs(args.dictId, id).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
		},
		{
			name: "Empty Fields",
			input: args{
				dictId: 1,
				item: entity.Word{
					Word:        "",
					Translation: "tr",
				},
			},
			mock: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id).RowError(0, errors.New("insert error"))
				mock.ExpectQuery("INSERT INTO words").
					WithArgs(args.item.Word, args.item.Translation).WillReturnRows(rows)

				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name: "Failed 2nd Insert",
			input: args{
				dictId: 1,
				item: entity.Word{
					Word:        "word",
					Translation: "translation",
				},
			},
			mock: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO words").
					WithArgs(args.item.Word, args.item.Translation).WillReturnRows(rows)

				mock.ExpectExec("INSERT INTO dictionaries_words").WithArgs(args.dictId, id).
					WillReturnError(errors.New("insert error"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input, tt.want)

			got, err := r.Create(tt.input.dictId, tt.input.item)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestTodoItemPostgres_GetAll(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewWordPsql(db)

	type args struct {
		dictId int
		userId int
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		want    []entity.Word
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "word", "translation"}).
					AddRow(1, "cat", "кот").
					AddRow(2, "dog", "пес").
					AddRow(3, "phone", "телефон")

				mock.ExpectQuery("SELECT (.+) FROM words word INNER JOIN dictionaries_words dictsWords ON (.+) INNER JOIN users_dictionaries usersDicts ON (.+) WHERE (.+)").
					WithArgs(1, 1).WillReturnRows(rows)
			},
			input: args{
				dictId: 1,
				userId: 1,
			},
			want: []entity.Word{
				{1, "cat", "кот"},
				{2, "dog", "пес"},
				{3, "phone", "телефон"},
			},
		},
		{
			name: "No Records",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "word", "translation"})

				mock.ExpectQuery("SELECT (.+) FROM words word INNER JOIN dictionaries_words dictsWords ON (.+) INNER JOIN users_dictionaries usersDicts ON (.+) WHERE (.+)").
					WithArgs(1, 1).WillReturnRows(rows)
			},
			input: args{
				dictId: 1,
				userId: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetAll(tt.input.userId, tt.input.dictId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestTodoItemPostgres_GetById(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewWordPsql(db)

	type args struct {
		itemId int
		userId int
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		want    entity.Word
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "word", "translation"}).
					AddRow(1, "w1", "t1")

				mock.ExpectQuery("SELECT (.+) FROM words word INNER JOIN dictionaries_words dictsWords ON (.+) INNER JOIN users_dictionaries usersDicts ON (.+) WHERE (.+)").
					WithArgs(1, 1).WillReturnRows(rows)
			},
			input: args{
				itemId: 1,
				userId: 1,
			},
			want: entity.Word{Id: 1, Word: "w1", Translation: "t1"},
		},
		{
			name: "Not Found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "word", "translation", "done"})

				mock.ExpectQuery("SELECT (.+) FROM words word INNER JOIN dictionaries_words dictsWords ON (.+) INNER JOIN users_dictionaries usersDicts ON (.+) WHERE (.+)").
					WithArgs(404, 1).WillReturnRows(rows)
			},
			input: args{
				itemId: 404,
				userId: 1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetById(tt.input.userId, tt.input.itemId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestTodoItemPostgres_Delete(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewWordPsql(db)

	type args struct {
		itemId int
		userId int
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				mock.ExpectExec("DELETE FROM words word USING dictionaries_words dictsWords, users_dictionaries usersDicts WHERE (.+)").
					WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				itemId: 1,
				userId: 1,
			},
		},
		{
			name: "Not Found",
			mock: func() {
				mock.ExpectExec("DELETE FROM words word USING dictionaries_words dictsWords, users_dictionaries usersDicts WHERE (.+)").
					WithArgs(1, 404).WillReturnError(sql.ErrNoRows)
			},
			input: args{
				itemId: 404,
				userId: 1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := r.Delete(tt.input.userId, tt.input.itemId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestTodoItemPostgres_Update(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewWordPsql(db)

	type args struct {
		itemId int
		userId int
		input  entity.UpdateWordInput
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		wantErr bool
	}{
		{
			name: "OK_AllFields",
			mock: func() {
				mock.ExpectExec("UPDATE words word SET (.+) FROM dictionaries_words dictsWords, users_dictionaries usersDicts WHERE (.+)").
					WithArgs("new w", "new t", 1, 1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				itemId: 1,
				userId: 1,
				input: entity.UpdateWordInput{
					Word:        stringPointer("new w"),
					Translation: stringPointer("new t"),
				},
			},
		},
		{
			name: "OK_WithoutDone",
			mock: func() {
				mock.ExpectExec("UPDATE words word SET (.+) FROM dictionaries_words dictsWords, users_dictionaries usersDicts WHERE (.+)").
					WithArgs("new w", "new t", 1, 1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				itemId: 1,
				userId: 1,
				input: entity.UpdateWordInput{
					Word:        stringPointer("new w"),
					Translation: stringPointer("new t"),
				},
			},
		},
		{
			name: "OK_WithoutDoneAndDescription",
			mock: func() {
				mock.ExpectExec("UPDATE words word SET (.+) FROM dictionaries_words dictsWords, users_dictionaries usersDicts WHERE (.+)").
					WithArgs("new word", 1, 1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				itemId: 1,
				userId: 1,
				input: entity.UpdateWordInput{
					Word: stringPointer("new word"),
				},
			},
		},
		{
			name: "OK_NoInputFields",
			mock: func() {
				mock.ExpectExec("UPDATE words word SET FROM dictionaries_words dictsWords, users_dictionaries usersDicts WHERE (.+)").
					WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				itemId: 1,
				userId: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := r.Update(tt.input.userId, tt.input.itemId, tt.input.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
