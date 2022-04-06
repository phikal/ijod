"use strict";

const video   = document.getElementById("video");
const list    = document.getElementById("list");
const status  = document.getElementById("status");
const log     = document.getElementById("log");
const seen    = JSON.parse(window.localStorage.getItem("seen")) || {};

var self;

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
        video.src = state.video;
        video.load();
        video.currentTime = state.position +
            (state.playing ? diff : 0);

        seen[state.video] = new Date();
        window.localStorage.setItem("seen", JSON.stringify(seen));
        let base = video.src.substr(video.src.lastIndexOf('/') + 1);
        write("<q>" + state.user + "</q> selected " + base);
    }
    if (Math.abs(state.position - video.currentTime) > 0.2) {
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
                write("<q>" + user + "</q> is here");
            }
        });
        current = list;
    }
    list.forEach(user => {
        if (!current.has(user)) {
            write("<q>" + user + "</q> joined");
        }
    });
    current.forEach(user => {
        if (!list.has(user)) {
            write("<q>" + user + "</q> left");
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
                li.title = seen[file] = new Date();
                window.localStorage.setItem("seen", JSON.stringify(seen));
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
function recv(event) {
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
        write("You are <q>" + self + "</q>");
        break;

    default:
        console.log("Received unknown message: " + msg);
    }
}

// Initialisation function
function connect() {
    let uri = location.href
        .replace(/^http/, "ws")
        .replace(/\/room/, "/socket");

    let socket = new WebSocket(uri);
    socket.onmessage = recv;
    socket.onclose   = socket.onerror = event => {
        console.error(event);
        setTimeout(connect, 250);
    };

    let sync = (event) => {
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

    status.onclick = (event) => {
        socket.send(JSON.stringify({"type": "refresh"}))
    };

    return socket;
}

// Local Variables:
// indent-tabs-mode: nil
// End: