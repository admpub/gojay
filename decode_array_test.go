package gojay

import (
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testSliceInts []int

func (t *testSliceInts) UnmarshalJSONArray(dec *Decoder) error {
	i := 0
	if err := dec.AddInt(&i); err != nil {
		return err
	}
	*t = append(*t, i)
	return nil
}

func TestSliceInts(t *testing.T) {
	testCases := []struct {
		name           string
		json           string
		expectedResult testSliceInts
		err            bool
		errType        interface{}
	}{
		{
			name:           "basic-test",
			json:           "[1,2,3,43567788543,45777655,432,0]",
			expectedResult: testSliceInts{1, 2, 3, 43567788543, 45777655, 432, 0},
		},
		{
			name:           "basic-test",
			json:           "[1,2,3,43567788543,null,432,0]",
			expectedResult: testSliceInts{1, 2, 3, 43567788543, 0, 432, 0},
		},
		{
			name:           "empty",
			json:           "[]",
			expectedResult: testSliceInts{},
		},
		{
			name:           "floats",
			json:           "[1,2,3,43567788543,457.7765,432,0,0.45]",
			expectedResult: testSliceInts{1, 2, 3, 43567788543, 457, 432, 0, 0},
		},
		{
			name:           "invalid-type",
			json:           `[1,2,3,43567788543,457.7765,432,0,"test"]`,
			expectedResult: testSliceInts{1, 2, 3, 43567788543, 457, 432, 0, 0},
			err:            true,
			errType:        InvalidUnmarshalError(""),
		},
		{
			name:           "invalid-json",
			json:           `[1,2,3",43567788543,457.7765,432,0,"test"]`,
			expectedResult: testSliceInts{1, 2, 3, 43567788543, 457, 432, 0, 0},
			err:            true,
			errType:        InvalidJSONError(""),
		},
	}

	for _, testCase := range testCases {
		s := make(testSliceInts, 0)
		dec := BorrowDecoder(strings.NewReader(testCase.json))
		defer dec.Release()
		err := dec.Decode(&s)
		if testCase.err {
			assert.NotNil(t, err, "err should not be nil")
			if testCase.errType != nil {
				assert.IsType(t, testCase.errType, err, "err should be of the given type")
			}
			continue
		}
		log.Print(s)
		for k, v := range testCase.expectedResult {
			assert.Equal(t, v, s[k], "value at given index should be the same as expected results")
		}
	}
}

type testSliceStrings []string

func (t *testSliceStrings) UnmarshalJSONArray(dec *Decoder) error {
	str := ""
	if err := dec.AddString(&str); err != nil {
		return err
	}
	*t = append(*t, str)
	return nil
}

func TestSliceStrings(t *testing.T) {
	testCases := []struct {
		name           string
		json           string
		expectedResult testSliceStrings
		err            bool
		errType        interface{}
	}{
		{
			name:           "basic-test",
			json:           `["hello world", "hey" , "foo","bar"]`,
			expectedResult: testSliceStrings{"hello world", "hey", "foo", "bar"},
		},
		{
			name:           "basic-test",
			json:           `["hello world", "hey" , "foo","bar \\n escape"]`,
			expectedResult: testSliceStrings{"hello world", "hey", "foo", "bar \n escape"},
		},
		{
			name:           "basic-test",
			json:           `["hello world", "hey" , null,"bar \\n escape"]`,
			expectedResult: testSliceStrings{"hello world", "hey", "", "bar \n escape"},
		},
		{
			name:           "invalid-type",
			json:           `["foo",1,2,3,"test"]`,
			expectedResult: testSliceStrings{},
			err:            true,
			errType:        InvalidUnmarshalError(""),
		},
		{
			name:           "invalid-json",
			json:           `["hello world]`,
			expectedResult: testSliceStrings{},
			err:            true,
			errType:        InvalidJSONError(""),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			s := make(testSliceStrings, 0)
			dec := BorrowDecoder(strings.NewReader(testCase.json))
			defer dec.Release()
			err := dec.Decode(&s)
			if testCase.err {
				assert.NotNil(t, err, "err should not be nil")
				if testCase.errType != nil {
					assert.IsType(t, testCase.errType, err, "err should be of the given type")
				}
				return
			}
			assert.Nil(t, err, "err should be nil")
			for k, v := range testCase.expectedResult {
				assert.Equal(t, v, s[k], "value at given index should be the same as expected results")
			}
		})
	}
}

type testSliceBools []bool

func (t *testSliceBools) UnmarshalJSONArray(dec *Decoder) error {
	b := false
	if err := dec.AddBool(&b); err != nil {
		return err
	}
	*t = append(*t, b)
	return nil
}

func TestSliceBools(t *testing.T) {
	testCases := []struct {
		name           string
		json           string
		expectedResult testSliceBools
		err            bool
		errType        interface{}
	}{
		{
			name:           "basic-test",
			json:           `[true, false, false, true, true, false]`,
			expectedResult: testSliceBools{true, false, false, true, true, false},
		},
		{
			name:           "basic-test2",
			json:           `[true, false, false, true, null,null,true,false]`,
			expectedResult: testSliceBools{true, false, false, true, false, false, true, false},
		},
		{
			name:           "invalid-type",
			json:           `["foo",1,2,3,"test"]`,
			expectedResult: testSliceBools{},
			err:            true,
			errType:        InvalidUnmarshalError(""),
		},
		{
			name:           "invalid-json",
			json:           `["hello world]`,
			expectedResult: testSliceBools{},
			err:            true,
			errType:        InvalidJSONError(""),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			s := make(testSliceBools, 0)
			dec := BorrowDecoder(strings.NewReader(testCase.json))
			defer dec.Release()
			err := dec.Decode(&s)
			if testCase.err {
				assert.NotNil(t, err, "err should not be nil")
				if testCase.errType != nil {
					assert.IsType(t, testCase.errType, err, "err should be of the given type")
				}
				return
			}
			log.Print(s, testCase.name)
			assert.Nil(t, err, "err should be nil")
			for k, v := range testCase.expectedResult {
				assert.Equal(t, v, s[k], "value at given index should be the same as expected results")
			}
		})
	}
}

type testSliceSlicesSlices []testSliceInts

func (t *testSliceSlicesSlices) UnmarshalJSONArray(dec *Decoder) error {
	sl := make(testSliceInts, 0)
	if err := dec.AddArray(&sl); err != nil {
		return err
	}
	*t = append(*t, sl)
	return nil
}

func TestSliceSlices(t *testing.T) {
	testCases := []struct {
		name           string
		json           string
		expectedResult testSliceSlicesSlices
		err            bool
		errType        interface{}
	}{
		{
			name:           "basic-test",
			json:           `[[1,2],[1,2],[1,2]]`,
			expectedResult: testSliceSlicesSlices{testSliceInts{1, 2}, testSliceInts{1, 2}, testSliceInts{1, 2}},
		},
		{
			name:           "basic-test",
			json:           `[[1,2],null,[1,2]]`,
			expectedResult: testSliceSlicesSlices{testSliceInts{1, 2}, testSliceInts{}, testSliceInts{1, 2}},
		},
		{
			name:           "invalid-type",
			json:           `["foo",1,2,3,"test"]`,
			expectedResult: testSliceSlicesSlices{},
			err:            true,
			errType:        InvalidUnmarshalError(""),
		},
		{
			name:           "invalid-json",
			json:           `["hello world]`,
			expectedResult: testSliceSlicesSlices{},
			err:            true,
			errType:        InvalidJSONError(""),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			s := make(testSliceSlicesSlices, 0)
			dec := BorrowDecoder(strings.NewReader(testCase.json))
			defer dec.Release()
			err := dec.Decode(&s)
			if testCase.err {
				assert.NotNil(t, err, "err should not be nil")
				if testCase.errType != nil {
					assert.IsType(t, testCase.errType, err, "err should be of the given type")
				}
				return
			}
			log.Print(s, testCase.name)
			assert.Nil(t, err, "err should be nil")
			for k, v := range testCase.expectedResult {
				assert.Equal(t, v, s[k], "value at given index should be the same as expected results")
			}
		})
	}
}

type testSliceObjects []*testObject

func (t *testSliceObjects) UnmarshalJSONArray(dec *Decoder) error {
	obj := &testObject{}
	*t = append(*t, obj)
	return dec.AddObject(obj)
}

func TestSliceObjects(t *testing.T) {
	testCases := []struct {
		name           string
		json           string
		expectedResult testSliceObjects
		err            bool
		errType        interface{}
	}{
		{
			name: "basic-test",
			json: `[{"testStr":"foo bar","testInt":123},{"testStr":"foo bar","testInt":123}]`,
			expectedResult: testSliceObjects{
				&testObject{
					testStr: "foo bar",
					testInt: 123,
				},
				&testObject{
					testStr: "foo bar",
					testInt: 123,
				},
			},
		},
		{
			name: "basic-test",
			json: `[{"testStr":"foo bar","testInt":123},null,{"testStr":"foo bar","testInt":123}]`,
			expectedResult: testSliceObjects{
				&testObject{
					testStr: "foo bar",
					testInt: 123,
				},
				&testObject{},
				&testObject{
					testStr: "foo bar",
					testInt: 123,
				},
			},
		},
		{
			name:           "invalid-type",
			json:           `["foo",1,2,3,"test"]`,
			expectedResult: testSliceObjects{},
			err:            true,
			errType:        InvalidUnmarshalError(""),
		},
		{
			name:           "invalid-json",
			json:           `["hello world]`,
			expectedResult: testSliceObjects{},
			err:            true,
			errType:        InvalidJSONError(""),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			s := make(testSliceObjects, 0)
			dec := BorrowDecoder(strings.NewReader(testCase.json))
			defer dec.Release()
			err := dec.Decode(&s)
			if testCase.err {
				assert.NotNil(t, err, "err should not be nil")
				if testCase.errType != nil {
					assert.IsType(t, testCase.errType, err, "err should be of the given type")
				}
				return
			}
			assert.Nil(t, err, "err should be nil")
			for k, v := range testCase.expectedResult {
				assert.Equal(t, *v, *s[k], "value at given index should be the same as expected results")
			}
		})
	}
}

type testChannelArray chan *TestObj

func (c *testChannelArray) UnmarshalJSONArray(dec *Decoder) error {
	obj := &TestObj{}
	if err := dec.AddObject(obj); err != nil {
		return err
	}
	*c <- obj
	return nil
}

func TestDecoderSliceNull(t *testing.T) {
	json := []byte(`null`)
	v := &testSliceStrings{}
	err := Unmarshal(json, v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, len(*v), 0, "v must be of len 0")
}

func TestDecodeSliceInvalidType(t *testing.T) {
	result := testSliceObjects{}
	err := UnmarshalJSONArray([]byte(`{}`), &result)
	assert.NotNil(t, err, "err should not be nil")
	assert.IsType(t, InvalidUnmarshalError(""), err, "err should be of type InvalidUnmarshalError")
	assert.Equal(t, "Cannot unmarshall to array, wrong char '{' found at pos 0", err.Error(), "err should not be nil")
}

func TestDecoderChannelOfObjectsBasic(t *testing.T) {
	json := []byte(`[
		{
			"test": 245,
			"test2": -246,
			"test3": "string"
		},
		{
			"test": 247,
			"test2": 248,
			"test3": "string"
		},
		{
			"test": 777,
			"test2": 456,
			"test3": "string"
		}
	]`)
	testChan := testChannelArray(make(chan *TestObj, 3))
	err := UnmarshalJSONArray(json, &testChan)
	assert.Nil(t, err, "Err must be nil")
	ct := 0
	l := len(testChan)
	for _ = range testChan {
		ct++
		if ct == l {
			break
		}
	}
	assert.Equal(t, ct, 3)
}

func TestDecoderSliceInvalidJSON(t *testing.T) {
	json := []byte(`hello`)
	testArr := testSliceInts{}
	err := UnmarshalJSONArray(json, &testArr)
	assert.NotNil(t, err, "Err must not be nil as JSON is invalid")
	assert.IsType(t, InvalidJSONError(""), err, "err message must be 'Invalid JSON'")
}

func TestDecoderSliceDecoderAPI(t *testing.T) {
	json := `["string","string1"]`
	testArr := testSliceStrings{}
	dec := NewDecoder(strings.NewReader(json))
	err := dec.DecodeArray(&testArr)
	assert.Nil(t, err, "Err must be nil")
	assert.Len(t, testArr, 2, "testArr should be of len 2")
	assert.Equal(t, "string", testArr[0], "testArr[0] should be 'string'")
	assert.Equal(t, "string1", testArr[1], "testArr[1] should be 'string1'")
}

func TestDecoderSliceDecoderAPIError(t *testing.T) {
	testArr := testSliceInts{}
	dec := NewDecoder(strings.NewReader(`hello`))
	err := dec.DecodeArray(&testArr)
	assert.NotNil(t, err, "Err must not be nil as JSON is invalid")
	assert.IsType(t, InvalidJSONError(""), err, "err message must be 'Invalid JSON'")
}

func TestUnmarshalJSONArrays(t *testing.T) {
	testCases := []struct {
		name         string
		v            UnmarshalerJSONArray
		d            []byte
		expectations func(err error, v interface{}, t *testing.T)
	}{
		{
			v:    new(testDecodeSlice),
			d:    []byte(`[{"test":"test"}]`),
			name: "test decode slice",
			expectations: func(err error, v interface{}, t *testing.T) {
				vtPtr := v.(*testDecodeSlice)
				vt := *vtPtr
				assert.Nil(t, err, "err must be nil")
				assert.Len(t, vt, 1, "len of vt must be 1")
				assert.Equal(t, "test", vt[0].test, "vt[0].test must be equal to 'test'")
			},
		},
		{
			v:    new(testDecodeSlice),
			d:    []byte(`[{"test":"test"},{"test":"test2"}]`),
			name: "test decode slice",
			expectations: func(err error, v interface{}, t *testing.T) {
				vtPtr := v.(*testDecodeSlice)
				vt := *vtPtr
				assert.Nil(t, err, "err must be nil")
				assert.Len(t, vt, 2, "len of vt must be 2")
				assert.Equal(t, "test", vt[0].test, "vt[0].test must be equal to 'test'")
				assert.Equal(t, "test2", vt[1].test, "vt[1].test must be equal to 'test2'")
			},
		},
		{
			v:    new(testDecodeSlice),
			d:    []byte(`invalid json`),
			name: "test decode object null",
			expectations: func(err error, v interface{}, t *testing.T) {
				assert.NotNil(t, err, "err must not be nil")
				assert.IsType(t, InvalidJSONError(""), err, "err must be of type InvalidJSONError")
			},
		},
	}
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(*testing.T) {
			err := UnmarshalJSONArray(testCase.d, testCase.v)
			testCase.expectations(err, testCase.v, t)
		})
	}
}

func TestSkipArray(t *testing.T) {
	testCases := []struct {
		json         string
		expectations func(*testing.T, int, error)
	}{
		{
			json: `"testbasic"]`,
			expectations: func(t *testing.T, i int, err error) {
				assert.Equal(t, len(`"testbasic"]`), i)
				assert.Nil(t, err)
			},
		},
		{
			json: `"test \\\\\" escape"]`,
			expectations: func(t *testing.T, i int, err error) {
				assert.Equal(t, len(`"test \\\\\" escape"]`), i)
				assert.Nil(t, err)
			},
		},
		{
			json: `"test \\\\\\"]`,
			expectations: func(t *testing.T, i int, err error) {
				assert.Equal(t, len(`"test \\\\\\"]`), i)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range testCases {
		dec := NewDecoder(strings.NewReader(test.json))
		i, err := dec.skipArray()
		test.expectations(t, i, err)
	}
}

func TestDecodeArrayEmpty(t *testing.T) {
	v := new(testDecodeSlice)
	dec := NewDecoder(strings.NewReader(""))
	err := dec.Decode(v)
	assert.NotNil(t, err, "err should not be nil")
	assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
}

func TestDecodeArraySkipError(t *testing.T) {
	v := new(testDecodeSlice)
	dec := NewDecoder(strings.NewReader("34fef"))
	err := dec.Decode(v)
	assert.NotNil(t, err, "err should not be nil")
	assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
}

func TestDecodeArrayNullError(t *testing.T) {
	v := new(testDecodeSlice)
	dec := NewDecoder(strings.NewReader("nall"))
	err := dec.Decode(v)
	assert.NotNil(t, err, "err should not be nil")
	assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
}
