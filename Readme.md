# download pinterest origin image

## How to use
download from [releases](https://github.com/pzx521521/pinterest-download/releases/):

```bash
pinterest  <username> <boradname> 
```
example:
```bash
pinterest parapeng wallpaper 
```
### if you want to download all board of user
```bash
  pinterest <username>
```
### use proxy
```bash
pinterest -p=http://localhost:7897 parapeng wallpaper 
```
### use outputdir
```bash
pinterest -o=./download parapeng wallpaper 
```

### use poolsize
```bash
pinterest -ps=5 parapeng wallpaper 
```