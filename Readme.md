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

### curl

#### GetBoards
```bash
https://www.pinterest.com/parapeng
```
in
```
<script id="__PWS_INITIAL_PROPS__" type="application/json">
```
#### GetPins
```bash
https://www.pinterest.com/resource/BoardFeedResource/get/?source_url=/parapeng/wallpaper/&data={"options":{"board_id":"946107902908880006","board_url":"/parapeng/wallpaper/","page_size":250}}
```