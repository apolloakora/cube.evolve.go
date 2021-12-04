# How To Evolve Cube State
___

A 2x2x2 Rubiks cube can be represented by the following state string:
```
0   1   2   3   4   5   6  
0xyz0xyz0xyz0xyz0xyz0xyz0xyz
```
**The state is relative to a corner (our core)**
- with colour (red, green, white)
- in position (front, left, top)

**Through the cube's evolution, this corner never moves, which means**
- x rotations spin the back layer
- y rotations spin the right layer
- z rotations spin the bottom layer

**In a solved state all the other pieces are now relative to our core**

| # | cell name | x pos | y pos | z pos | x face | y face | z face | pos |
|---|-----------|-------|-------|-------|--------|--------|--------|-----|
|0|x adjacent|back|left|top|orange|green|white|xa|
|1|y adjacent|front|right|top|red|blue|white|ya|
|2|z adjacent|front|left|bottom|red|green|yellow|za|
|3|x diagonal|front|right|bottom|red|blue|yellow|xd|
|4|y diagonal|back|left|bottom|orange|green|yellow|yd|
|5|z diagonal|back|right|top|orange|blue|white|zd|
|6|polar opposite|back|right|bottom|orange|blue|yellow|po|


## Orientation State

___

At any point in time, the orientation state of the cells will be **one** of 6 possibles values\
which are xyz, xzy, yxz, yzx, zxy and zyx.

###**What do these orientations mean and how to make them?**

Hold the cube so that the **front, left, top** piece has the order **red, green, white** respectively.\
The three letters x, y or z will always be in the order of the x face, y face and then z face.

**For the first letter**\
If the colour on the x face is 
- orange/red - the letter will be x
- blue/green - the letter will be y
- white/yellow - the letter will be z

**For the second letter**\
If the colour on the y face is
- orange/red - the letter will be x
- blue/green - the letter will be y
- white/yellow - the letter will be z

**For the third letter**\
If the colour on the z face is
- orange/red - the letter will be x
- blue/green - the letter will be y
- white/yellow - the letter will be z

The three letters will represent where the faces on your piece are currently, compared to where they should be on a solved cube

## How Rotation Changes Orientation
___





