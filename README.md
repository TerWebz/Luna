Luna
====

A SNES, N64, and GameCube game downloader made in [Go](http://golang.org)

Introduction
------------

Welcome to Luna! Luna is a downloader that downloads Super Nintendo games, Nintendo 64 games and \*Nintendo GameCube games (\* = select titles only)

To install Luna, you must have **Go** installed, which you can find [here](https://golang.org/dl/).

Once you have Go installed, you have 2 options to get the executable up and running.

1.  You may compile the source code, reccomended option for all platforms
2.  You may download the executable from the **Releases** section

If you choose to do 1, run `go build /path/to/luna/main.go` in your command prompt or terminal. It may take a little bit, so go ahead and continue reading this while you are at it.

Running the Program
-------------------

Once you have Luna installed, you'd wanna of course, run it.
**Note:** The game has to be typed in a specific format, on both platforms.

### Windows
===

Run the program like so, `./luna.exe`. It would go ahead and prompt you to enter a gamename.

```
Microsoft Windows (Version 10.0.10240)
(c) 2015 Microsoft Corporation. All Rights Reserved.

C:\Users\user> luna.exe
Usage: ./luna.exe <gamename>
```

### MacOS/Linux
===

Run the program like so, `./luna`. It would go ahead and prompt you to enter a gamename.

```
javier@Javier:~$ ./luna
Usage: ./luna <gamename>
```

The gamename would have to be written again, in a specific format. Here is an example of what that would look like.

| Correct Ways        	| Incorrect Ways        	| 
|---------------------	|-----------------------	|
| supermarioworld.zip 	| SuPeRMaRiOWoRlD.ZiP   	|
| supermarioworld       | Super Mario World.zip 	|                    	


We hope to find a way to make it so you could be able to do the third option. :)
