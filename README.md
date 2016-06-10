# ya-go-config
Yet Another Go Config library, simple and tiny, with JSON format

Example of config file: config.json
```
{
    "answer": 53
}
```
Define default values and read values from config.json:
```
defaults := map[string]interface{}{
		"question": "Meaning of life",
		"answer": 42,
}
config := New(defaults)
config.Load("./config.json")
fmt.Printf("%s is %d", config.GetStr("question"), config.GetInt("answer"))
```
It would print:
```
Meaning of life is 53
```
Also see config_test.go