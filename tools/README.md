# data conversion tool
## convert excel to json file

```shell
go run main.go --dirPath data --promptCol E --completionCol F --startRow 2
```

```shell
export OPENAI_API_KEY="sk-xxxx"

openai tools fine_tunes.prepare_data -f json.txt
openai api fine_tunes.create -t json_prepared.jsonl
openai api fine_tunes.list
openai api fine_tunes.follow -i <id>
openai api completions.create -m curie:ft-personal-2023-03-26-07-52-31 -p <YOUR_PROMPT>
```