package api

import (
	mockdb "9bany/simple_bank/db/mock"
	db "9bany/simple_bank/db/sqlc"
	"9bany/simple_bank/util"
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	user, password := randomUser(t)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username":  user.Username,
				"email":     user.Email,
				"full_name": user.FullName,
				"password":  password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).Times(1).Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name: "Invalid PassWord",
			body: gin.H{
				"username":  user.Username,
				"email":     user.Email,
				"full_name": user.FullName,
				"password":  "123",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Invalid Email",
			body: gin.H{
				"username":  user.Username,
				"email":     "",
				"full_name": user.FullName,
				"password":  password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "StatusInternalServerError",
			body: gin.H{
				"username":  "ban",
				"email":     user.Email,
				"full_name": user.FullName,
				"password":  password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).Times(1).Return(db.Users{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "StatusForbidden",
			body: gin.H{
				"username":  user.Username,
				"email":     user.Email,
				"full_name": user.FullName,
				"password":  password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).Times(1).Return(db.Users{}, &pq.Error{
					Code: "23505",
				})
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)

			// build stubs
			tc.buildStubs(store)
			// start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			// marshal body data to json
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/users"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)
			server.router.ServeHTTP(recorder, request)
			// check response
			tc.checkResponse(t, recorder)
		})
	}

}

func randomUser(t *testing.T) (user db.Users, password string) {

	password = util.RandomString(8)
	hashedPassword, err := util.HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	user = db.Users{
		Username:     util.RandomOwnerName(),
		HashPassword: hashedPassword,
		FullName:     util.RandomOwnerName(),
		Email:        util.RandomEmail(),
	}
	return
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.Users) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var goUser db.Users
	err = json.Unmarshal(data, &goUser)
	require.NoError(t, err)
	require.Equal(t, goUser.FullName, user.FullName)
	require.Equal(t, goUser.Username, user.Username)
	require.Equal(t, goUser.Email, user.Email)
}
