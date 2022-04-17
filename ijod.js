"use strict";

const video   = document.getElementById("video");
const list    = document.getElementById("list");
const status  = document.getElementById("status");
const ytdl    = document.getElementById("ytdl");
const log     = document.getElementById("log");
const seen    = JSON.parse(window.localStorage.getItem("seen")) || {};

var self;

// function userfmt(name) { ... } is defined in userfmt.js

function write(msg) {
    let li = document.createElement("li");
    li.innerHTML = msg;
    li.title = new Date();
    log.appendChild(li);
    log.scrollTop = log.scrollHeight - 12;
}

// Apply a state
function load(state) {
    const diff = new Date() - new Date(state.timestamp);
    if (video.src != state.video) {
        status.innerText = state.video;
        video.src = state.video;
        video.load();
        video.currentTime = state.position +
            (state.playing ? diff : 0);

        let base = video.src.substr(video.src.lastIndexOf('/') + 1);
        write(userfmt(state.user) + " selected " + decodeURI(base));
    }
    if (Math.abs(state.position - video.currentTime) > 1) {
        video.currentTime = state.position;
    }
    if (state.playing) {
        video.play();
    } else {
        video.pause();
    }
}

// Log new users
var current;
function users(list) {
    list = new Set(list);
    if (!current) {
        list.forEach(user => {
            if (user != self) {
                write(userfmt(user) + " is here");
            }
        });
        current = list;
    }
    list.forEach(user => {
        if (!current.has(user)) {
            write(userfmt(user) + " joined");
        }
    });
    current.forEach(user => {
        if (!list.has(user)) {
            write(userfmt(user) + " left");
        }
    });
    current = list;
}

// File selection event handler generator
function select(file) {
    return (event) => {
        if (file.match(/\.vtt$/)) {
            while (video.firstChild) {
                video.firstChild.remove();
            }

            var track = document.createElement("track");
            track.kind = "subtitles";
            track.src = "/data/" + file;
            track.label = "Subtitles";
            video.appendChild(track);
            video.textTracks[0].mode = "showing";
        } else {
            video.src = "/data/" + file;
            video.currentTime = 0;
            video.load();

            status.innerText = file;

            let li = event.currentTarget;
            if (li) {
                li.classList.add("seen");
            }
        }
    }
}

// Recursive tree generation function
function display(tree, parent) {
    let ul = document.createElement("ul");

    for (let name in tree) {
        let li = document.createElement("li");
        let span = document.createElement("span");
        span.appendChild(document.createTextNode(name));
        li.appendChild(span);

        if ((typeof tree[name]) === 'string') {
            span.onclick = select(tree[name]);
            span.onauxclick = () => {
                let uri = "/data/" + tree[name];
                window.open(uri, '_blank').focus();
            }

            if (tree[name] in seen) {
                span.classList.add("seen");
                li.title = seen[tree[name]];
            }
        } else if (tree[name]) {
            span.onclick = _ => li.classList.toggle("opened");
            li.classList.add("dir");

            li.appendChild(display(tree[name], name));
        }

        ul.appendChild(li);
    }

    return ul;
}

// Websocket event handler
function recv(socket) {
    return (event) => {
        console.log(event);
        var msg = JSON.parse(event.data);

        switch (msg.type) {
        case "state":
            load(msg.data);
            break;

        case "users":
            users(msg.data);
            break;

        case "files":
            // Clear the previous file tree
            while (list.firstChild) {
                list.firstChild.remove();
            }

            // Build a new tree
            list.appendChild(display(msg.data, "/"));
            break;

        case "write":
            write(msg.data);
            break;

        case "self":
            self = msg.data;
            write("You are " + userfmt(self));
            break;

        case "disable-dl":
            ytdl.style.display = "none";

        case "ping":
            socket.send(JSON.stringify({
                "type": "pong",
                "data": msg.data
            }));
            break;

        default:
            console.log("Received unknown message: " + msg);
        }
    };
}

// Initialisation function
function connect() {
    let uri = location.href
        .replace(/^http/, "ws")
        .replace(/\/room/, "/socket");

    let socket = new WebSocket(uri);
    socket.onmessage = recv(socket);
    socket.onerror   = event => {
        console.error(event);
        write("Connection error");
        setTimeout(connect, 250);
    };
    socket.onclose   = event => {
        write("Connection closed");
        setTimeout(connect, 250);
    };

    let sync = (event) => {
        seen[video.src] = new Date();
        window.localStorage.setItem("seen", JSON.stringify(seen));

        socket.send(JSON.stringify({
            "type": "state",
            "data": {
                "timestamp":  new Date(),
                "position":   video.currentTime,
                "playing":    !video.paused,
                "video":      video.src,
                "user":       self
            }
        }));
    };
    video.onseeked  = sync;
    video.onpause   = sync;
    video.onplay    = sync;
    video.oncanplay = sync;
    ytdl.onkeyup    = (event) => {
        if (event.keyCode !== 13) {
            return;
        }

        event.preventDefault();
        try {
            socket.send(JSON.stringify({
                "type": "download",
                "data": ytdl.value
            }));
            ytdl.value = "";
        } catch (err) {
            console.error(err);
        }

    };

    status.onclick = (event) => {
        socket.send(JSON.stringify({"type": "refresh"}))
    };

    return socket;
}

// Local Variables:
// indent-tabs-mode: nil
// End:
