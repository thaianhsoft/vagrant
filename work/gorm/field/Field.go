package field

type IField interface {
	GetSqlType() string
	DefaultSize(size int) IField
	AI() IField
	Unique() IField
	Nullable() IField
}