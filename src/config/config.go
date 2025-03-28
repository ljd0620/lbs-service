package config

type Config struct {
    BaiduAPIKey   string `json:"baidu_api_key"`
    GaodeAPIKey   string `json:"gaode_api_key"`
    TencentAPIKey string `json:"tencent_api_key"`
    RedisAddress   string `json:"redis_address"`
    RedisPassword  string `json:"redis_password"`
    RedisDB       int    `json:"redis_db"`
}