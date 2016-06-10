# ya-go-config
Yet Another Go Config library, simple and tiny, with JSON format

How to use it, see config_test.go


config.json contents:
```
{
    "answer": 53
}
```
Example of usage:
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
