# Subtitles Renamer

This is a small program to rename your subtitles as the same name with your video file, written in Golang.

You are required to place the subtitle files that has the extensions `.srt` or `.ass` together with your video files of extensions `.mp4` or `.mkv` or `.avi` in the same folder.
Also make sure your video file and subtitles contains the following pattern `S01EXX` or `EpXX`.
<br/><br/>
Simply clone the repository and run the following:

    USAGE: go run . -dir="<path containing video files and subtitles>"

This program is tested on Windows and Linux, assuming macOS works as well as it works on Linux. 
The process will only proceed if you have exactly the same number of subtitles as per video files.