# Get lesson (**for regenerate lessons**)

```sh
cd tool/getlessons
curl https://go.dev/tour/lesson/ -o lessons.json
rm -rf ../../lessons/
go run .

# It will make `../../lessons/`
```
