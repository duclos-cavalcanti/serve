#!/bin/bash


checkDependency() {
    cmd=$1
    if ! command -v ${cmd} &>/dev/null; then
        echo "The following program is not found: ${cmd}"
        exit 1
    fi
}

checkDependency ffmpeg
# checkDependency xwininfo

# geometry=$(xwininfo | grep geometry | awk '{print $2}')

ffmpeg -f x11grab -video_size 1920x1080 -framerate 50 -i :0.0+1920 -vf format=yuv420p ./.assets/example.mp4
# ffmpeg -ss 30 -t 3 -i ./.assets/example.mp4 -vf "fps=10,scale=320:-1:flags=lanczos,split[s0][s1];[s0]palettegen[p];[s1][p]paletteuse" -loop 0 ./.assets/example.gif

