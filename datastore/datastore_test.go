package datastore

import (
	"errors"
	"testing"
	"github.com/bapjiws/timezones_mc/datastore/mocks"
	"github.com/bapjiws/timezones_mc/revel_app/app/models"

	"github.com/stretchr/testify/assert"
)

// AddDocument(doc models.Document) error
func TestAddDocument(t *testing.T) {
	assert := assert.New(t)

	testDataStore := new(mocks.Datastore)

	testCases := []struct {
		document models.Document
		error    error
	}{
		{
			document: mocks.BadDocument,
			error:    errors.New("Bad document!"),
		},
		{
			document: mocks.GoodDocument,
			error:    nil,
		},
	}

	for _, testCase := range testCases {
		testDataStore.On("AddDocument", testCase.document).Return(testCase.error)
		err := testDataStore.AddDocument(testCase.document)
		assert.Equal(testCase.error, err)
	}

	// Assert that everything specified with On and Return was in fact called as expected.
	testDataStore.AssertExpectations(t)
}

// FindDocumentById(id string) (models.Document, error)
func TestFindDocumentById(t *testing.T) {
	assert := assert.New(t)

	testDataStore := new(mocks.Datastore)

	testCases := []struct {
		id       string
		document models.Document
		error    error
	}{
		{
			id:       mocks.BadId,
			document: nil,
			error:    errors.New("Bad id!"),
		},
		{
			id:       mocks.GoodId,
			document: mocks.GoodDocument,
			error:    nil,
		},
	}

	for _, testCase := range testCases {
		testDataStore.On("FindDocumentById", testCase.id).Return(testCase.document, testCase.error)
		document, err := testDataStore.FindDocumentById(testCase.id)
		assert.Equal(testCase.document, document)
		assert.Equal(testCase.error, err)
	}

	// Assert that everything specified with On and Return was in fact called as expected.
	testDataStore.AssertExpectations(t)
}

// SuggestDocuments(suggesterName string, text string, field string, payloadKey string) ([]models.Document, error)
func TestSuggestDocuments(t *testing.T) {
	assert := assert.New(t)

	testDataStore := new(mocks.Datastore)

	suggesterName := mocks.SuggesterName
	field := mocks.Field
	payloadKey := mocks.PayloadKey

	testCases := []struct {
		text        string
		suggestions []models.Document
		error       error
	}{
		{
			text:        mocks.BadText,
			suggestions: nil,
			error:       errors.New("Bad suggestion text!"),
		},
		{
			text:        mocks.GoodTextWithNoSuggestions,
			suggestions: mocks.EmptyDocumentList,
			error:       nil,
		},
		{
			text:        mocks.GoodTextWithSomeSuggestions,
			suggestions: mocks.GoodDocuments,
			error:       nil,
		},
	}

	for _, testCase := range testCases {
		testDataStore.
			On("SuggestDocuments", suggesterName, testCase.text, field, payloadKey).
			Return(testCase.suggestions, testCase.error)
		suggestions, err := testDataStore.SuggestDocuments(suggesterName, testCase.text, field, payloadKey)
		assert.Equal(testCase.suggestions, suggestions)
		assert.Equal(testCase.error, err)
	}

	// Assert that everything specified with On and Return was in fact called as expected.
	testDataStore.AssertExpectations(t)
}
