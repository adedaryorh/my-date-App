package helpers

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"sort"
	"text/template"

	"strconv"
	"strings"
	"time"

	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/google/uuid"
	"github.com/leekchan/accounting"
	gonanoid "github.com/matoous/go-nanoid"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	alphabet               = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	MODEL_PREFIX_SEPARATOR = "."
)

type TypeModelIdTag string
type Map map[string]interface{}

func ToSlug(value string) string {
	var slug string

	splitted := strings.Split(value, " ")
	for i := 0; i < len(splitted); i++ {
		slug += splitted[i]
		if i < len(splitted)-1 {
			slug += "-"
		}
	}
	return strings.ToLower(slug)
}

func ToTitleCase(s string) string {
	return strings.ToTitle(strings.Split(s, "@")[0])
}

func GenerateUniqueReferenceId(length int) string {
	alphanumeric := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	id, _ := gonanoid.Generate(alphanumeric, length)
	return id
}
func HashString512(s string) string {
	h := sha512.New()

	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
func StringToUUID(s string) uuid.UUID {
	u, _ := uuid.Parse(s)
	return u
}

func GenerateRandomNumber(length int) string {
	return generateRandom(length, 10)
}

func GenerateRandomByte(length int) string {
	return generateRandom(length, len(alphabet))
}

func GenerateRandomAccNumber() string {
	return generateRandom(10, 10)
}

func BytesToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeHex(input []byte) ([]byte, error) {
	db := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(db, input)
	if err != nil {
		return nil, err
	}
	return db, nil
}
func base64Encode(input []byte) []byte {
	eb := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
	base64.StdEncoding.Encode(eb, input)

	return eb
}

func HexToBase64(hex string) string {
	db, _ := decodeHex([]byte(hex))
	return base64.StdEncoding.EncodeToString(db)
}

func generateRandom(length, randomRange int) string {
	var random string
	for i := 0; i < length; i++ {
		random += fmt.Sprintf("%s", string(alphabet[rand.Intn(randomRange)]))
	}
	return random
}

func Base64StringToByte(s string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}
func StructToBase64(v interface{}) (string, error) {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	err := json.NewEncoder(encoder).Encode(v)
	if err != nil {
		return "", err
	}
	encoder.Close()
	return buf.String(), nil
}

func Base64ToStruct(v interface{}, enc string) error {
	return json.NewDecoder(base64.NewDecoder(base64.StdEncoding, strings.NewReader(enc))).Decode(v)
}

func Hash(s string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 14)
	fmt.Println(err)

	return string(bytes)
}

func HashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func CompareHashString(s, hash string) bool {
	return HashString(s) == hash
}

func CompareHash(hashed, s string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(s))
	return err == nil
}
func StringToLower(s string) string {
	return strings.ToLower(s)
}
func StringToTitle(s string) string {
	c := cases.Title(language.English, cases.NoLower)
	return c.String(s)
}
func StringToObjectID(s string) *uuid.UUID {
	var id uuid.UUID
	id, err := uuid.Parse(s)
	if err != nil {
		return nil
	}
	return &id
}

func GetDurationFromTimeString(s string) time.Duration {
	duration, _ := time.ParseDuration(s)
	return duration
}

func GenerateRandomUppercase(length int) string {
	s, _ := uuid.NewRandom()
	return strings.ToUpper(strings.Replace(s.String(), "-", "", -1))
}

func StringToBoolean(s string) bool {
	b, _ := strconv.ParseBool(s)
	return b
}

func GenerateSerialNumber(serialNumber, length int) string {
	numString := fmt.Sprintf("%d", serialNumber)
	num := length - len(numString)
	var result string
	for i := 0; i < num; i++ {
		result += "0"
	}
	result += numString
	return result
}

func StructToMap(v any) (map[string]interface{}, error) {
	var converted map[string]interface{}
	// convert to map
	byte, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(byte, &converted)
	return converted, nil
}

func SortMapKeys(entity map[string]interface{}) map[string]interface{} {
	var keys []string
	for k := range entity {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	sorted := make(map[string]interface{})
	for _, k := range keys {
		sorted[k] = entity[k]
	}

	return sorted
}

func HmacHashHex(secret, message string, encoding func() hash.Hash) string {
	h := hmac.New(encoding, []byte(secret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func HmacHashBase64(secret, message string, encoding func() hash.Hash) string {
	h := hmac.New(encoding, []byte(secret))
	h.Write([]byte(message))
	hexHash := hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(hexHash))
}

func FormatPhoneNumber(phone string) string {
	s := phone[len(phone)-10:]
	return fmt.Sprintf("0%s", s)
}

func ParseTemplate(fileName string, content map[string]interface{}) (*string, error) {
	tml, err := template.ParseFiles(fileName)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = tml.Execute(buf, content)
	if err != nil {
		return nil, err
	}

	text := buf.String()

	return &text, nil
}

func FormatMoney(currency string, amount int64) string {
	currencySymbol := GetCurrencySymbol(currency)
	ac := accounting.Accounting{Symbol: currencySymbol, Precision: 2}

	return ac.FormatMoney(float64(amount) / 100)
}

func GetCurrencySymbol(currency string) string {
	if currency == "NGN" {
		return "₦"
	}
	if currency == "USD" {
		return "$"
	}

	return currency
}

func FormatTransactionType(txnType string) string {
	var s string
	splitted := strings.Split(txnType, "-")

	for _, v := range splitted {
		s += fmt.Sprintf("%s ", strings.Title(v))
	}
	return s
}

func GetEntityComputedHash(entity any) (map[string]interface{}, string) {
	entityMap, err := StructToMap(entity)

	if err != nil {
		return nil, err.Error()
	}

	if entityMap["id"] != nil {
		delete(entityMap, "id")
	}

	if entityMap["Operation"] != nil {
		delete(entityMap, "Operation")
	}
	if entityMap["BalanceId"] != nil {
		delete(entityMap, "BalanceId")
	}
	if entityMap["txn_time"] != nil {
		delete(entityMap, "txn_time")
	}
	if entityMap["created_at"] != nil {
		delete(entityMap, "created_at")
	}
	if entityMap["updated_at"] != nil {
		delete(entityMap, "updated_at")
	}
	sortedEntity := SortMapKeys(entityMap)

	b, _ := json.Marshal(sortedEntity)

	hash := HashString(string(b))
	return entityMap, hash
}

type RestUriCredentials struct {
	BaseUrl string
	Id      string
	Secret  string
}

func ExtractURICredentials(uri string, protocol ...string) *RestUriCredentials {
	var url []string

	if len(protocol) > 0 {
		url = strings.Split(uri, protocol[0])
	} else {
		url = strings.Split(uri, "rest://")
	}

	url = strings.Split(url[1], "@")

	if len(url) == 1 {
		return &RestUriCredentials{
			BaseUrl: url[0],
		}
	}

	credentials := strings.Split((url[0]), ":")
	baseUrl := url[1]

	return &RestUriCredentials{
		BaseUrl: baseUrl,
		Id:      credentials[0],
		Secret:  credentials[1],
	}
}

func Getenv(variable string, defaultValue ...string) string {
	env := os.Getenv(variable)
	if env == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return ""
	}
	return env
}

func PointerString(s string) *string {
	return &s
}

func PointerTime(t time.Time) *time.Time {
	return &t
}

func PointerInt(i int) *int {
	return &i
}
func Int64ToPointer(i int64) *int64 {
	return &i
}
func PointerToInt64(i *int64) int64 {
	return *i
}

func PointerUint64(i uint64) *uint64 {
	return &i
}

func Base58Encode(input []byte) string {
	encode := base58.Encode(input)
	return encode
}

func Base58Decode(input string) []byte {
	decode := base58.Decode(input)
	return decode
}

func Base58EncodeString(input string) string {
	return Base58Encode([]byte(input))
}

func Base58DecodeString(input string) string {
	return string(Base58Decode(input))
}

func Base58EncodeFromHex(input string) (*string, error) {
	data, err := hex.DecodeString(input)
	if err != nil {
		return nil, err
	}

	encoded := Base58Encode(data)
	return &encoded, nil
}

func Base58DecodeToHex(input string) (*string, error) {
	data := Base58Decode(input)
	decoded := hex.EncodeToString(data)
	return &decoded, nil
}

func TagHexId(tag TypeModelIdTag, id string) string {
	encoded, _ := Base58EncodeFromHex(id)
	return string(tag) + "." + *encoded
}

func TagId(tag TypeModelIdTag, id string) string {
	return string(tag) + "." + id
}

func TagIdToString(tagId string) (*string, error) {

	split := strings.Split(tagId, MODEL_PREFIX_SEPARATOR)
	var id string
	if len(split) == 2 {
		id = split[1]
	} else if len(split) == 3 {
		id = split[2]
	} else {
		return nil, errors.New("invalid tag id")
	}

	hex, err := Base58DecodeToHex(id)
	if err != nil {
		return nil, err
	}

	return hex, nil
}
func GenerateCSV(header []string, records [][]string) ([]byte, error) {
	allRecords := [][]string{
		header,
	}
	allRecords = append(allRecords, records...)
	csvBuffer := new(bytes.Buffer)
	writer := csv.NewWriter(csvBuffer)
	writer.WriteAll(allRecords)

	return csvBuffer.Bytes(), nil
}

func MultipartHeaderToBytes(fileHeader *multipart.FileHeader) ([]byte, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}

func GetFileType(filename, separator string) string {
	fileTypeSlice := strings.Split(filename, separator)
	return fileTypeSlice[len(fileTypeSlice)-1]
}

func GetTransactionId(clientCode string) string {
	current_time := time.Now()
	year := fmt.Sprintf("%d", current_time.Year())
	year = string(year[len(year)-2:])

	random := generateRandom(12, 10)

	s := fmt.Sprintf("%s%s%02d%02d%02d%02d%02d%s", clientCode, year, current_time.Month(), current_time.Day(), current_time.Hour(), current_time.Minute(), current_time.Second(), random)
	return s
}

func StringToDate(s, layout string) time.Time {
	t, _ := time.Parse(layout, s)
	return t
}

func DateToString(date time.Time, layout string) string {

	return date.Format(layout)
}

func FormatTime(sentTime time.Time, timeZone string) string {
	location, _ := time.LoadLocation(timeZone)
	locationTime := sentTime.In(location)
	return fmt.Sprintf("%v", locationTime.Format("2 Jan 2006, 3:04 pm"))
}
func CalcOrigBinaryLength(datas string) int {

	l := len(datas)

	// count how many trailing '=' there are (if any)
	eq := 0
	if l >= 2 {
		if datas[l-1] == '=' {
			eq++
		}
		if datas[l-2] == '=' {
			eq++
		}

		l -= eq
	}

	// basically:
	//
	// eq == 0 :    bits-wasted = 0
	// eq == 1 :    bits-wasted = 2
	// eq == 2 :    bits-wasted = 4

	// each base64 character = 6 bits

	// so orig length ==  (l*6 - eq*2) / 8

	return (l*3 - eq) / 4
}

func JSONToStruct(body string, v any) error {
	return json.Unmarshal([]byte(body), &v)
}

func ToHigherDecimal(num int64) float64 { // kobo to naira
	return float64(num / 100)
}

func ToLowerDecimal(num float64) int64 { // naira to kobo
	return int64(num * 100)
}

func RemoveDuplicates[T comparable](slice []T) []T {
	var final []T
	sliceMap := make(map[T]struct{})

	for _, val := range slice {
		if _, ok := sliceMap[val]; !ok {
			final = append(final, val)
			sliceMap[val] = struct{}{}
		}
	}
	return final
}
func Reverse[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
