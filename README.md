# Gosend

A simple form receiver

## Run

Simply clone the repository and run the app

```bash
#!/bin/bash
git clone github.com/TemaSaur/gosend && cd gosend
go run .
```

## Usage

Simply use the hosted address of your local gosend instance in your forms. You can use either POST or GET forms. The `source` parameter is used for distinguishing between different form sources.

```html
<form method="POST" action="http://127.0.0.1:8000/form/your_source_name">
  <!-- any form fields -->
  <input name="name" />
  <input type="email" name="email" />
  <textarea name="content"></textarea>
  <button type="submit">Submit</button>
</form>
```

Valid forms will result in a success page.

After that use the `/get` endpoint to get all submitted forms.

```
$ curl 127.0.0.1:8000/get
SOURCE: source1
CONTENT:
[name]
Bob
[message]
Lorem ipsum dolor sit amet consectetur adipiscing elit quisque faucibus ex
sapien vitae pellentesque sem placerat

--------------------------------------------------------------------------------

SOURCE: source2
CONTENT:
[email]
alice@gmail.com
[message]
Lorem ipsum dolor sit amet consectetur adipiscing elit quisque faucibus ex
sapien vitae pellentesque sem placerat
```
