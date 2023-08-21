
# Jagaad Technical Interview Assignment

I'm using golang version `1.21.0`

Library that I used:
* [Cobra](https://github.com/spf13/cobra) for easier making CLI app
* [Testify](https://github.com/stretchr/testify) for easier making unit test
* [Mockery](https://github.com/vektra/mockery) for easier making mock for unit test

# Config
By default, I already set array of `base_url` to be called, the application will loop through the array to call the API, feel free to add some / remove if you want.

    {
        "base_url": [
            "https://run.mocky.io/v3/03d2a7bd-f12f-4275-9e9a-84e41f9c2aae",
            "https://run.mocky.io/v3/87931203-8086-43ef-ba16-4c8903d8fa88"
        ]
    }

# How to use
On terminal:

    go run main.go 

It will show us the list of available command

# Useful command
There's some useful command such as:

    go run main.go fetch_and_save_user

It will call the API from `base_url` inside `config.json` file, after it called the API and get the data, it will create a new csv file called `result.csv` and save the data inside it.

If you want to search the users inside csv data:

    go run main.go read_and_search 

It will read the csv data and will filter the data by cli flag `--tags`, not putting the cli flag `--tags` will return all of the data from csv. You can put multiple tags by appending the tags using `,`, for example if you want to put tags Apple & Orange, you can put on the terminal like this

    go run main.go read_and_search --tags apple,orange


# Time spent

The time I spent to develop the core logic for this assignment is around 90 mins, I finished the core logic on this [commit message](https://github.com/papannn/jagaat-technical-task/commit/d6a5491c7e8e69040768a9a09784311d115036c8), after that I slow down my pace and start to refactor my code, doing bug fixing and writing unit test so the reviewer can easily review the code than before. 