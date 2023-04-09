package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomString(t *testing.T) Tsarka{
	arg := CreateStringParams{
		String: "zhaisss",
		MaxSubstring: "zhais",
	}

	result, err := testQueries.CreateString(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, arg.String, result.String)

	require.NotZero(t, result.ID)

	return result
}

func TestCreateString(t *testing.T) {
	createRandomString(t)
}

func TestGetString(t *testing.T) {
	arg := CreateStringParams{
		String: "saaars",
		MaxSubstring: "ars",
	}
	result2, err := testQueries.CreateString(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result2)

	require.Equal(t, arg.String, result2.String)
	require.Equal(t, arg.MaxSubstring, result2.MaxSubstring)
}

func TestUpdateString(t *testing.T) {
	result1 := createRandomString(t)

	arg := UpdateStringParams{
		MaxSubstring: result1.MaxSubstring,
		ID:           result1.ID,
	}

	result2, err := testQueries.UpdateString(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result2)

	require.Equal(t, result1.ID, result2.ID)
	require.Equal(t, result1.String, result2.String)
	require.Equal(t, result1.MaxSubstring, result2.MaxSubstring)
}

func TestDeleteString(t *testing.T) {
	arg := CreateStringParams{
		String:       "test",
		MaxSubstring: "est",
	}
	createdString, err := testQueries.CreateString(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, createdString)

	_, err = testQueries.DeleteString(context.Background(), createdString.ID)
	require.NoError(t, err)

	var count int
	row := testQueries.db.QueryRowContext(context.Background(), "SELECT COUNT(*) FROM tsarka WHERE id = $1", createdString.ID)
	err = row.Scan(&count)
	require.NoError(t, err)
	require.Zero(t, count)
}

func TestListString(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomString(t)
	}

	arg := ListStringsParams{
		Limit:  3,
		Offset: 3,
	}

	results, err := testQueries.ListStrings(context.Background(), arg)
	require.NoError(t, err)

	for _, result := range results {
		require.NotEmpty(t, result)
	}
}