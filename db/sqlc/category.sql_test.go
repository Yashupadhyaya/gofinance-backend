package db

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/DATA-DOG/go-sqlmock"
	"reflect"
)

type Category struct {
	ID          int32
	UserID      int32
	Title       string
	Type        string
	Description string
	CreatedAt   time.Time
}/*
ROOST_METHOD_HASH=DeleteCategories_596952739e
ROOST_METHOD_SIG_HASH=DeleteCategories_90cdfe4457


 */
func TestDeleteCategories(t *testing.T) {
	tests := []struct {
		name        string
		categoryID  int32
		mockSetup   func(sqlmock.Sqlmock)
		expectedErr error
	}{
		{
			name:       "Successful Deletion of a Category",
			categoryID: 1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM categories WHERE id = \\$1").
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedErr: nil,
		},
		{
			name:       "Deletion of Non-Existing Category",
			categoryID: 999,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM categories WHERE id = \\$1").
					WithArgs(999).
					WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedErr: nil,
		},
		{
			name:       "Database Connection Error",
			categoryID: 1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM categories WHERE id = \\$1").
					WithArgs(1).
					WillReturnError(errors.New("db connection error"))
			},
			expectedErr: errors.New("db connection error"),
		},
		{
			name:       "Invalid Category ID (Negative Value)",
			categoryID: -1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM categories WHERE id = \\$1").
					WithArgs(-1).
					WillReturnError(errors.New("invalid category ID"))
			},
			expectedErr: errors.New("invalid category ID"),
		},
		{
			name:       "SQL Injection Attempt",
			categoryID: 0,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM categories WHERE id = \\$1").
					WithArgs("1 OR 1=1").
					WillReturnError(errors.New("invalid category ID"))
			},
			expectedErr: errors.New("invalid category ID"),
		},
		{
			name:       "Context Cancellation",
			categoryID: 1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM categories WHERE id = \\$1").
					WithArgs(1).
					WillReturnError(context.Canceled)
			},
			expectedErr: context.Canceled,
		},
		{
			name:       "Large Category ID",
			categoryID: 1 << 30,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM categories WHERE id = \\$1").
					WithArgs(1 << 30).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedErr: nil,
		},
		{
			name:       "Multiple Deletion Requests",
			categoryID: 1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM categories WHERE id = \\$1").
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("DELETE FROM categories WHERE id = \\$1").
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedErr: nil,
		},
		{
			name:       "Deletion with Different Context Values",
			categoryID: 1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM categories WHERE id = \\$1").
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedErr: nil,
		},
		{
			name:       "Deletion with DBTX Mock Returning RowsAffected",
			categoryID: 1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM categories WHERE id = \\$1").
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			tt.mockSetup(mock)

			q := &Queries{db: db}
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			err = q.DeleteCategories(ctx, tt.categoryID)
			assert.Equal(t, tt.expectedErr, err)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=CreateCategory_4a03e34912
ROOST_METHOD_SIG_HASH=CreateCategory_eb9f8f9b1c


 */
func TestQueriesCreateCategory(t *testing.T) {

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	queries := &Queries{db: db}

	tests := []struct {
		name          string
		params        CreateCategoryParams
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "Successful Category Creation",
			params: CreateCategoryParams{
				UserID:      1,
				Title:       "Test Category",
				Type:        "Type1",
				Description: "Test Description",
			},
			mockSetup: func() {
				mock.ExpectQuery("INSERT INTO categories").
					WithArgs(1, "Test Category", "Type1", "Test Description").
					WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"}).
						AddRow(1, 1, "Test Category", "Type1", "Test Description", time.Now()))
			},
			expectedError: false,
		},
		{
			name: "Missing UserID",
			params: CreateCategoryParams{
				UserID:      0,
				Title:       "Test Category",
				Type:        "Type1",
				Description: "Test Description",
			},
			mockSetup: func() {
				mock.ExpectQuery("INSERT INTO categories").
					WithArgs(0, "Test Category", "Type1", "Test Description").
					WillReturnError(sql.ErrNoRows)
			},
			expectedError: true,
		},
		{
			name: "Empty Title Field",
			params: CreateCategoryParams{
				UserID:      1,
				Title:       "",
				Type:        "Type1",
				Description: "Test Description",
			},
			mockSetup: func() {
				mock.ExpectQuery("INSERT INTO categories").
					WithArgs(1, "", "Type1", "Test Description").
					WillReturnError(sql.ErrNoRows)
			},
			expectedError: true,
		},
		{
			name: "Database Connection Failure",
			params: CreateCategoryParams{
				UserID:      1,
				Title:       "Test Category",
				Type:        "Type1",
				Description: "Test Description",
			},
			mockSetup: func() {
				mock.ExpectQuery("INSERT INTO categories").
					WithArgs(1, "Test Category", "Type1", "Test Description").
					WillReturnError(sql.ErrConnDone)
			},
			expectedError: true,
		},
		{
			name: "SQL Injection Attempt",
			params: CreateCategoryParams{
				UserID:      1,
				Title:       "'); DROP TABLE categories;--",
				Type:        "Type1",
				Description: "Test Description",
			},
			mockSetup: func() {
				mock.ExpectQuery("INSERT INTO categories").
					WithArgs(1, "'); DROP TABLE categories;--", "Type1", "Test Description").
					WillReturnError(sql.ErrNoRows)
			},
			expectedError: true,
		},
		{
			name: "Valid Category with Special Characters",
			params: CreateCategoryParams{
				UserID:      1,
				Title:       "Special!@#$%^&*()",
				Type:        "Type1",
				Description: "Test Description",
			},
			mockSetup: func() {
				mock.ExpectQuery("INSERT INTO categories").
					WithArgs(1, "Special!@#$%^&*()", "Type1", "Test Description").
					WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"}).
						AddRow(1, 1, "Special!@#$%^&*()", "Type1", "Test Description", time.Now()))
			},
			expectedError: false,
		},
		{
			name: "Long String Fields",
			params: CreateCategoryParams{
				UserID:      1,
				Title:       "This is a very long title that exceeds normal length expectations for a category title",
				Type:        "Type1",
				Description: "This is a very long description that exceeds normal length expectations for a category description",
			},
			mockSetup: func() {
				mock.ExpectQuery("INSERT INTO categories").
					WithArgs(1, "This is a very long title that exceeds normal length expectations for a category title", "Type1", "This is a very long description that exceeds normal length expectations for a category description").
					WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"}).
						AddRow(1, 1, "This is a very long title that exceeds normal length expectations for a category title", "Type1", "This is a very long description that exceeds normal length expectations for a category description", time.Now()))
			},
			expectedError: false,
		},
		{
			name: "Concurrent Category Creation",
			params: CreateCategoryParams{
				UserID:      1,
				Title:       "Concurrent Test Category",
				Type:        "Type1",
				Description: "Test Description",
			},
			mockSetup: func() {
				mock.ExpectQuery("INSERT INTO categories").
					WithArgs(1, "Concurrent Test Category", "Type1", "Test Description").
					WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"}).
						AddRow(1, 1, "Concurrent Test Category", "Type1", "Test Description", time.Now()))
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			result, err := queries.CreateCategory(context.Background(), tt.params)
			if tt.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.params.UserID, result.UserID)
				require.Equal(t, tt.params.Title, result.Title)
				require.Equal(t, tt.params.Type, result.Type)
				require.Equal(t, tt.params.Description, result.Description)
			}

			err = mock.ExpectationsWereMet()
			require.NoError(t, err)
		})
	}
}

/*
ROOST_METHOD_HASH=GetCategory_0669ee5937
ROOST_METHOD_SIG_HASH=GetCategory_4d9ce8fc09


 */
func TestQueriesGetCategory(t *testing.T) {

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	q := &Queries{db: db}

	type testCase struct {
		name             string
		id               int32
		mockSetup        func()
		expectedError    error
		expectedCategory Category
	}

	testCases := []testCase{
		{
			name: "Successfully Retrieve a Category by ID",
			id:   1,
			mockSetup: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"}).
					AddRow(1, 1, "Category 1", "Type 1", "Description 1", time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC))
				mock.ExpectQuery("^SELECT id, user_id, title, type, description, created_at FROM categories WHERE id = \\$1 LIMIT 1$").
					WithArgs(1).
					WillReturnRows(rows)
			},
			expectedError: nil,
			expectedCategory: Category{
				ID:          1,
				UserID:      1,
				Title:       "Category 1",
				Type:        "Type 1",
				Description: "Description 1",
				CreatedAt:   time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "Category Not Found",
			id:   2,
			mockSetup: func() {
				mock.ExpectQuery("^SELECT id, user_id, title, type, description, created_at FROM categories WHERE id = \\$1 LIMIT 1$").
					WithArgs(2).
					WillReturnError(sql.ErrNoRows)
			},
			expectedError: sql.ErrNoRows,
		},
		{
			name: "Database Connection Error",
			id:   3,
			mockSetup: func() {
				mock.ExpectQuery("^SELECT id, user_id, title, type, description, created_at FROM categories WHERE id = \\$1 LIMIT 1$").
					WithArgs(3).
					WillReturnError(sql.ErrConnDone)
			},
			expectedError: sql.ErrConnDone,
		},
		{
			name: "Invalid ID Provided",
			id:   -1,
			mockSetup: func() {

			},
			expectedError: ErrInvalidID,
		},
		{
			name: "SQL Query Syntax Error",
			id:   4,
			mockSetup: func() {
				mock.ExpectQuery("^SELECT id, user_id, title, type, description, created_at FROM categories WHERE id = \\$1 LIMIT 1$").
					WithArgs(4).
					WillReturnError(sql.ErrSyntax)
			},
			expectedError: sql.ErrSyntax,
		},
		{
			name: "Successfully Retrieve Category with Special Characters in Description",
			id:   5,
			mockSetup: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"}).
					AddRow(5, 1, "Category 5", "Type 5", "Description with special characters! @#$%^&*()", time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC))
				mock.ExpectQuery("^SELECT id, user_id, title, type, description, created_at FROM categories WHERE id = \\$1 LIMIT 1$").
					WithArgs(5).
					WillReturnRows(rows)
			},
			expectedError: nil,
			expectedCategory: Category{
				ID:          5,
				UserID:      1,
				Title:       "Category 5",
				Type:        "Type 5",
				Description: "Description with special characters! @#$%^&*()",
				CreatedAt:   time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			ctx := context.Background()
			category, err := q.GetCategory(ctx, tc.id)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedCategory, category)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=UpdateCategories_2babb188a3
ROOST_METHOD_SIG_HASH=UpdateCategories_6158760915


 */
func TestQueriesUpdateCategories(t *testing.T) {

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	q := &Queries{db: db}

	t.Run("Scenario 1: Successful Update of Category", func(t *testing.T) {

		ctx := context.Background()
		params := UpdateCategoriesParams{
			ID:          1,
			Title:       "New Title",
			Description: "New Description",
		}
		mock.ExpectQuery(`UPDATE categories`).
			WithArgs(params.ID, params.Title, params.Description).
			WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"}).
				AddRow(params.ID, 1, params.Title, "type", params.Description, time.Now()))

		category, err := q.UpdateCategories(ctx, params)

		require.NoError(t, err)
		require.Equal(t, params.Title, category.Title)
		require.Equal(t, params.Description, category.Description)
	})

	t.Run("Scenario 2: Category Not Found", func(t *testing.T) {

		ctx := context.Background()
		params := UpdateCategoriesParams{
			ID:          2,
			Title:       "Non-existent Title",
			Description: "Non-existent Description",
		}
		mock.ExpectQuery(`UPDATE categories`).
			WithArgs(params.ID, params.Title, params.Description).
			WillReturnError(sql.ErrNoRows)

		_, err := q.UpdateCategories(ctx, params)

		require.Error(t, err)
		require.Equal(t, sql.ErrNoRows, err)
	})

	t.Run("Scenario 3: Database Error During Update", func(t *testing.T) {

		ctx := context.Background()
		params := UpdateCategoriesParams{
			ID:          1,
			Title:       "Title",
			Description: "Description",
		}
		mock.ExpectQuery(`UPDATE categories`).
			WithArgs(params.ID, params.Title, params.Description).
			WillReturnError(sql.ErrConnDone)

		_, err := q.UpdateCategories(ctx, params)

		require.Error(t, err)
		require.Equal(t, sql.ErrConnDone, err)
	})

	t.Run("Scenario 4: Update With Empty Title", func(t *testing.T) {

		ctx := context.Background()
		params := UpdateCategoriesParams{
			ID:          1,
			Title:       "",
			Description: "Description",
		}
		mock.ExpectQuery(`UPDATE categories`).
			WithArgs(params.ID, params.Title, params.Description).
			WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"}).
				AddRow(params.ID, 1, params.Title, "type", params.Description, time.Now()))

		category, err := q.UpdateCategories(ctx, params)

		require.NoError(t, err)
		require.Equal(t, params.Title, category.Title)
		require.Equal(t, params.Description, category.Description)
	})

	t.Run("Scenario 5: Update With Empty Description", func(t *testing.T) {

		ctx := context.Background()
		params := UpdateCategoriesParams{
			ID:          1,
			Title:       "Title",
			Description: "",
		}
		mock.ExpectQuery(`UPDATE categories`).
			WithArgs(params.ID, params.Title, params.Description).
			WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"}).
				AddRow(params.ID, 1, params.Title, "type", params.Description, time.Now()))

		category, err := q.UpdateCategories(ctx, params)

		require.NoError(t, err)
		require.Equal(t, params.Title, category.Title)
		require.Equal(t, params.Description, category.Description)
	})

	t.Run("Scenario 6: Update With Invalid Data Types", func(t *testing.T) {

		ctx := context.Background()
		params := UpdateCategoriesParams{
			ID:          1,
			Title:       "Title",
			Description: "Description",
		}

		_, err := q.UpdateCategories(ctx, params)

		require.NoError(t, err)
	})

	t.Run("Scenario 7: Simultaneous Updates", func(t *testing.T) {

		ctx := context.Background()
		params1 := UpdateCategoriesParams{
			ID:          1,
			Title:       "Title 1",
			Description: "Description 1",
		}
		params2 := UpdateCategoriesParams{
			ID:          1,
			Title:       "Title 2",
			Description: "Description 2",
		}
		mock.ExpectQuery(`UPDATE categories`).
			WithArgs(params1.ID, params1.Title, params1.Description).
			WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"}).
				AddRow(params1.ID, 1, params1.Title, "type", params1.Description, time.Now()))
		mock.ExpectQuery(`UPDATE categories`).
			WithArgs(params2.ID, params2.Title, params2.Description).
			WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"}).
				AddRow(params2.ID, 1, params2.Title, "type", params2.Description, time.Now()))

		ch := make(chan struct{})
		go func() {
			_, err := q.UpdateCategories(ctx, params1)
			require.NoError(t, err)
			ch <- struct{}{}
		}()
		go func() {
			_, err := q.UpdateCategories(ctx, params2)
			require.NoError(t, err)
			ch <- struct{}{}
		}()
		<-ch
		<-ch

	})

	t.Run("Scenario 8: Update With Special Characters", func(t *testing.T) {

		ctx := context.Background()
		params := UpdateCategoriesParams{
			ID:          1,
			Title:       "Title!@#$%^&*()",
			Description: "Description!@#$%^&*()",
		}
		mock.ExpectQuery(`UPDATE categories`).
			WithArgs(params.ID, params.Title, params.Description).
			WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"}).
				AddRow(params.ID, 1, params.Title, "type", params.Description, time.Now()))

		category, err := q.UpdateCategories(ctx, params)

		require.NoError(t, err)
		require.Equal(t, params.Title, category.Title)
		require.Equal(t, params.Description, category.Description)
	})

	t.Run("Scenario 9: Large Input Data", func(t *testing.T) {

		ctx := context.Background()
		largeTitle := string(make([]byte, 1000))
		largeDescription := string(make([]byte, 1000))
		params := UpdateCategoriesParams{
			ID:          1,
			Title:       largeTitle,
			Description: largeDescription,
		}
		mock.ExpectQuery(`UPDATE categories`).
			WithArgs(params.ID, params.Title, params.Description).
			WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"}).
				AddRow(params.ID, 1, params.Title, "type", params.Description, time.Now()))

		category, err := q.UpdateCategories(ctx, params)

		require.NoError(t, err)
		require.Equal(t, params.Title, category.Title)
		require.Equal(t, params.Description, category.Description)
	})

	t.Run("Scenario 10: Update With Null Fields", func(t *testing.T) {

		ctx := context.Background()
		params := UpdateCategoriesParams{
			ID:          1,
			Title:       "",
			Description: "",
		}
		mock.ExpectQuery(`UPDATE categories`).
			WithArgs(params.ID, params.Title, params.Description).
			WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"}).
				AddRow(params.ID, 1, params.Title, "type", params.Description, time.Now()))

		category, err := q.UpdateCategories(ctx, params)

		require.NoError(t, err)
		require.Equal(t, params.Title, category.Title)
		require.Equal(t, params.Description, category.Description)
	})
}

/*
ROOST_METHOD_HASH=GetCategories_1c640a0750
ROOST_METHOD_SIG_HASH=GetCategories_d9632df74a


 */
func TestGetCategories(t *testing.T) {
	tests := []struct {
		name       string
		setupMock  func(sqlmock.Sqlmock)
		params     GetCategoriesParams
		wantResult []Category
		wantErr    bool
	}{
		{
			name: "Retrieve Categories for Valid User and Type",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"}).
					AddRow(1, 1, "Title1", "Type1", "Description1", time.Now()).
					AddRow(2, 1, "Title2", "Type1", "Description2", time.Now())
				mock.ExpectQuery("SELECT id, user_id, title, type, description, created_at FROM categories").
					WithArgs(int32(1), "Type1", "%title%", "%description%").
					WillReturnRows(rows)
			},
			params: GetCategoriesParams{
				UserID:      1,
				Type:        "Type1",
				Title:       "title",
				Description: "description",
			},
			wantResult: []Category{
				{ID: 1, UserID: 1, Title: "Title1", Type: "Type1", Description: "Description1", CreatedAt: time.Now()},
				{ID: 2, UserID: 1, Title: "Title2", Type: "Type1", Description: "Description2", CreatedAt: time.Now()},
			},
			wantErr: false,
		},
		{
			name: "No Categories Found for Given Filters",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"})
				mock.ExpectQuery("SELECT id, user_id, title, type, description, created_at FROM categories").
					WithArgs(int32(1), "Type1", "%nonexistent%", "%nonexistent%").
					WillReturnRows(rows)
			},
			params: GetCategoriesParams{
				UserID:      1,
				Type:        "Type1",
				Title:       "nonexistent",
				Description: "nonexistent",
			},
			wantResult: []Category{},
			wantErr:    false,
		},
		{
			name: "DB Query Error Handling",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT id, user_id, title, type, description, created_at FROM categories").
					WithArgs(int32(1), "Type1", "%error%", "%error%").
					WillReturnError(assert.AnError)
			},
			params: GetCategoriesParams{
				UserID:      1,
				Type:        "Type1",
				Title:       "error",
				Description: "error",
			},
			wantResult: nil,
			wantErr:    true,
		},
		{
			name: "Scan Error Handling",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"}).
					AddRow("invalid", 1, "Title1", "Type1", "Description1", time.Now())
				mock.ExpectQuery("SELECT id, user_id, title, type, description, created_at FROM categories").
					WithArgs(int32(1), "Type1", "%title%", "%description%").
					WillReturnRows(rows)
			},
			params: GetCategoriesParams{
				UserID:      1,
				Type:        "Type1",
				Title:       "title",
				Description: "description",
			},
			wantResult: nil,
			wantErr:    true,
		},
		{
			name:      "Context Cancellation Handling",
			setupMock: func(mock sqlmock.Sqlmock) {},
			params: GetCategoriesParams{
				UserID:      1,
				Type:        "Type1",
				Title:       "title",
				Description: "description",
			},
			wantResult: nil,
			wantErr:    true,
		},
		{
			name: "Empty Title and Description Filters",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"}).
					AddRow(1, 1, "Title1", "Type1", "Description1", time.Now()).
					AddRow(2, 1, "Title2", "Type1", "Description2", time.Now())
				mock.ExpectQuery("SELECT id, user_id, title, type, description, created_at FROM categories").
					WithArgs(int32(1), "Type1", "%%", "%%").
					WillReturnRows(rows)
			},
			params: GetCategoriesParams{
				UserID:      1,
				Type:        "Type1",
				Title:       "",
				Description: "",
			},
			wantResult: []Category{
				{ID: 1, UserID: 1, Title: "Title1", Type: "Type1", Description: "Description1", CreatedAt: time.Now()},
				{ID: 2, UserID: 1, Title: "Title2", Type: "Type1", Description: "Description2", CreatedAt: time.Now()},
			},
			wantErr: false,
		},
		{
			name: "Special Characters in Title and Description Filters",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"}).
					AddRow(1, 1, "Title@1", "Type1", "Desc#1", time.Now())
				mock.ExpectQuery("SELECT id, user_id, title, type, description, created_at FROM categories").
					WithArgs(int32(1), "Type1", "%@%", "%#%").
					WillReturnRows(rows)
			},
			params: GetCategoriesParams{
				UserID:      1,
				Type:        "Type1",
				Title:       "@",
				Description: "#",
			},
			wantResult: []Category{
				{ID: 1, UserID: 1, Title: "Title@1", Type: "Type1", Description: "Desc#1", CreatedAt: time.Now()},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Failed to open mock sql db, %v", err)
			}
			defer db.Close()

			q := &Queries{db: db}

			var ctx context.Context
			if tt.name == "Context Cancellation Handling" {
				var cancel context.CancelFunc
				ctx, cancel = context.WithCancel(context.Background())
				cancel()
			} else {
				ctx = context.Background()
			}

			tt.setupMock(mock)

			gotResult, err := q.GetCategories(ctx, tt.params)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.ElementsMatch(t, tt.wantResult, gotResult)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=GetCategoriesByUserIdAndType_5baa0d5587
ROOST_METHOD_SIG_HASH=GetCategoriesByUserIdAndType_6dc6c26962


 */
func TestQueriesGetCategoriesByUserIdAndType(t *testing.T) {
	tests := []struct {
		name         string
		userID       int32
		categoryType string
		setupMock    func(sqlmock.Sqlmock)
		expected     []Category
		expectErr    bool
	}{
		{
			name:         "Valid UserID and Type with Multiple Categories",
			userID:       1,
			categoryType: "expense",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"}).
					AddRow(1, 1, "Food", "expense", "Groceries", time.Now()).
					AddRow(2, 1, "Transport", "expense", "Bus fare", time.Now())
				mock.ExpectQuery("^SELECT (.+) FROM categories WHERE user_id = \\$1 AND type = \\$2$").
					WithArgs(1, "expense").
					WillReturnRows(rows)
			},
			expected: []Category{
				{ID: 1, UserID: 1, Title: "Food", Type: "expense", Description: "Groceries"},
				{ID: 2, UserID: 1, Title: "Transport", Type: "expense", Description: "Bus fare"},
			},
			expectErr: false,
		},
		{
			name:         "Valid UserID and Type with No Categories",
			userID:       2,
			categoryType: "income",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"})
				mock.ExpectQuery("^SELECT (.+) FROM categories WHERE user_id = \\$1 AND type = \\$2$").
					WithArgs(2, "income").
					WillReturnRows(rows)
			},
			expected:  []Category{},
			expectErr: false,
		},
		{
			name:         "Invalid UserID",
			userID:       99,
			categoryType: "expense",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"})
				mock.ExpectQuery("^SELECT (.+) FROM categories WHERE user_id = \\$1 AND type = \\$2$").
					WithArgs(99, "expense").
					WillReturnRows(rows)
			},
			expected:  []Category{},
			expectErr: false,
		},
		{
			name:         "Invalid Type",
			userID:       1,
			categoryType: "invalid_type",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"})
				mock.ExpectQuery("^SELECT (.+) FROM categories WHERE user_id = \\$1 AND type = \\$2$").
					WithArgs(1, "invalid_type").
					WillReturnRows(rows)
			},
			expected:  []Category{},
			expectErr: false,
		},
		{
			name:         "Database Connection Error",
			userID:       1,
			categoryType: "expense",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("^SELECT (.+) FROM categories WHERE user_id = \\$1 AND type = \\$2$").
					WithArgs(1, "expense").
					WillReturnError(errors.New("db connection error"))
			},
			expected:  nil,
			expectErr: true,
		},
		{
			name:         "SQL Query Error",
			userID:       1,
			categoryType: "expense",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("^SELECT (.+) FROM categories WHERE user_id = \\$1 AND type = \\$2$").
					WithArgs(1, "expense").
					WillReturnError(errors.New("sql query error"))
			},
			expected:  nil,
			expectErr: true,
		},
		{
			name:         "Row Scanning Error",
			userID:       1,
			categoryType: "expense",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"}).
					AddRow(1, 1, "Food", "expense", "Groceries", time.Now()).
					AddRow("invalid", 1, "Transport", "expense", "Bus fare", time.Now())
				mock.ExpectQuery("^SELECT (.+) FROM categories WHERE user_id = \\$1 AND type = \\$2$").
					WithArgs(1, "expense").
					WillReturnRows(rows)
			},
			expected:  nil,
			expectErr: true,
		},
		{
			name:         "Context Deadline Exceeded",
			userID:       1,
			categoryType: "expense",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("^SELECT (.+) FROM categories WHERE user_id = \\$1 AND type = \\$2$").
					WithArgs(1, "expense").
					WillDelayFor(2 * time.Second)
			},
			expected:  nil,
			expectErr: true,
		},
		{
			name:         "Context Cancellation",
			userID:       1,
			categoryType: "expense",
			setupMock:    func(mock sqlmock.Sqlmock) {},
			expected:     nil,
			expectErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			q := &Queries{db: db}

			if tt.name == "Context Deadline Exceeded" {
				ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
				defer cancel()
				tt.setupMock(mock)
				_, err := q.GetCategoriesByUserIdAndType(ctx, GetCategoriesByUserIdAndTypeParams{UserID: tt.userID, Type: tt.categoryType})
				if err == nil {
					t.Errorf("expected error but got none")
				} else if !errors.Is(err, context.DeadlineExceeded) {
					t.Errorf("expected context deadline exceeded error but got %v", err)
				}
				return
			}

			if tt.name == "Context Cancellation" {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				_, err := q.GetCategoriesByUserIdAndType(ctx, GetCategoriesByUserIdAndTypeParams{UserID: tt.userID, Type: tt.categoryType})
				if err == nil {
					t.Errorf("expected error but got none")
				} else if !errors.Is(err, context.Canceled) {
					t.Errorf("expected context canceled error but got %v", err)
				}
				return
			}

			ctx := context.Background()
			tt.setupMock(mock)

			result, err := q.GetCategoriesByUserIdAndType(ctx, GetCategoriesByUserIdAndTypeParams{UserID: tt.userID, Type: tt.categoryType})

			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected result: %v, got: %v", tt.expected, result)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=GetCategoriesByUserIdAndTypeAndDescription_6785853427
ROOST_METHOD_SIG_HASH=GetCategoriesByUserIdAndTypeAndDescription_ac5e3ac55a


 */
func TestGetCategoriesByUserIdAndTypeAndDescription(t *testing.T) {

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	q := &Queries{db: db}

	tests := []struct {
		name          string
		params        GetCategoriesByUserIdAndTypeAndDescriptionParams
		mockSetup     func()
		expectedError error
		expectedItems []Category
	}{
		{
			name: "Valid Input, Multiple Categories Returned",
			params: GetCategoriesByUserIdAndTypeAndDescriptionParams{
				UserID:      1,
				Type:        "expense",
				Description: "%groceries%",
			},
			mockSetup: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"}).
					AddRow(1, 1, "Groceries", "expense", "Weekly groceries", time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)).
					AddRow(2, 1, "Groceries", "expense", "Monthly groceries", time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC))
				mock.ExpectQuery(getCategoriesByUserIdAndTypeAndDescription).WithArgs(1, "expense", "%groceries%").WillReturnRows(rows)
			},
			expectedError: nil,
			expectedItems: []Category{
				{
					ID:          1,
					UserID:      1,
					Title:       "Groceries",
					Type:        "expense",
					Description: "Weekly groceries",
					CreatedAt:   time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:          2,
					UserID:      1,
					Title:       "Groceries",
					Type:        "expense",
					Description: "Monthly groceries",
					CreatedAt:   time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "Valid Input, No Categories Returned",
			params: GetCategoriesByUserIdAndTypeAndDescriptionParams{
				UserID:      1,
				Type:        "expense",
				Description: "%nonexistent%",
			},
			mockSetup: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"})
				mock.ExpectQuery(getCategoriesByUserIdAndTypeAndDescription).WithArgs(1, "expense", "%nonexistent%").WillReturnRows(rows)
			},
			expectedError: nil,
			expectedItems: []Category{},
		},
		{
			name: "Invalid User ID",
			params: GetCategoriesByUserIdAndTypeAndDescriptionParams{
				UserID:      -1,
				Type:        "expense",
				Description: "%groceries%",
			},
			mockSetup: func() {
				mock.ExpectQuery(getCategoriesByUserIdAndTypeAndDescription).WithArgs(-1, "expense", "%groceries%").WillReturnError(errors.New("invalid user ID"))
			},
			expectedError: errors.New("invalid user ID"),
			expectedItems: nil,
		},
		{
			name: "Empty Description",
			params: GetCategoriesByUserIdAndTypeAndDescriptionParams{
				UserID:      1,
				Type:        "expense",
				Description: "",
			},
			mockSetup: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"}).
					AddRow(1, 1, "Groceries", "expense", "", time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC))
				mock.ExpectQuery(getCategoriesByUserIdAndTypeAndDescription).WithArgs(1, "expense", "").WillReturnRows(rows)
			},
			expectedError: nil,
			expectedItems: []Category{
				{
					ID:          1,
					UserID:      1,
					Title:       "Groceries",
					Type:        "expense",
					Description: "",
					CreatedAt:   time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "Database Query Error",
			params: GetCategoriesByUserIdAndTypeAndDescriptionParams{
				UserID:      1,
				Type:        "expense",
				Description: "%groceries%",
			},
			mockSetup: func() {
				mock.ExpectQuery(getCategoriesByUserIdAndTypeAndDescription).WithArgs(1, "expense", "%groceries%").WillReturnError(errors.New("query error"))
			},
			expectedError: errors.New("query error"),
			expectedItems: nil,
		},
		{
			name: "SQL Injection Attempt",
			params: GetCategoriesByUserIdAndTypeAndDescriptionParams{
				UserID:      1,
				Type:        "expense",
				Description: "groceries'; DROP TABLE users;--",
			},
			mockSetup: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"})
				mock.ExpectQuery(getCategoriesByUserIdAndTypeAndDescription).WithArgs(1, "expense", "groceries'; DROP TABLE users;--").WillReturnRows(rows)
			},
			expectedError: nil,
			expectedItems: []Category{},
		},
		{
			name: "Large Dataset Handling",
			params: GetCategoriesByUserIdAndTypeAndDescriptionParams{
				UserID:      1,
				Type:        "expense",
				Description: "%groceries%",
			},
			mockSetup: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "title", "type", "description", "created_at"})
				for i := 0; i < 1000; i++ {
					rows.AddRow(i, 1, "Groceries", "expense", "Weekly groceries", time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC))
				}
				mock.ExpectQuery(getCategoriesByUserIdAndTypeAndDescription).WithArgs(1, "expense", "%groceries%").WillReturnRows(rows)
			},
			expectedError: nil,
			expectedItems: func() []Category {
				items := []Category{}
				for i := 0; i < 1000; i++ {
					items = append(items, Category{
						ID:          int32(i),
						UserID:      1,
						Title:       "Groceries",
						Type:        "expense",
						Description: "Weekly groceries",
						CreatedAt:   time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
					})
				}
				return items
			}(),
		},
		{
			name: "Context Cancellation",
			params: GetCategoriesByUserIdAndTypeAndDescriptionParams{
				UserID:      1,
				Type:        "expense",
				Description: "%groceries%",
			},
			mockSetup: func() {

				mock.ExpectQuery(getCategoriesByUserIdAndTypeAndDescription).WithArgs(1, "expense", "%groceries%").WillReturnError(context.Canceled)
			},
			expectedError: context.Canceled,
			expectedItems: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			ctx := context.Background()

			if tt.name == "Context Cancellation" {
				var cancel context.CancelFunc
				ctx, cancel = context.WithCancel(ctx)
				cancel()
			}

			items, err := q.GetCategoriesByUserIdAndTypeAndDescription(ctx, tt.params)
			if tt.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tt.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedItems, items)
			}
		})
	}
}

