//api_com.mjs

const url = "https://pliqutech.de/api/"; 
const wsUrl = "wss://pliqutech.de/ws";
const listeners = {};
let socket = null;
let isOpen = false;
let queue = [];

function connect() {
    socket = new WebSocket(wsUrl);

    socket.onopen = () => {
        isOpen = true;
        queue.forEach(msg => socket.send(msg));
        queue = [];
    };

    socket.onmessage = (event) => {
        let msg;
        try {
            msg = JSON.parse(event.data);
        } catch (e) {
            console.error("Ungültige Nachricht:", event.data);
            return;
        }

        const key = msg.cmd || msg.type || msg.channel;
        if (listeners[key]) {
            listeners[key](msg.data, msg); // payload + roher msg
        } else {
            console.warn("Keine Reaktion auf:", key, msg);
        }
    };

    socket.onclose = () => {
        isOpen = false;
        console.warn("WebSocket-Verbindung geschlossen");
    };

    socket.onerror = (err) => {
        console.error("WebSocket-Fehler:", err);
    };
}

function pipe(cmd, payload = {}, reaction = () => {}) {
    if (!socket) connect();

    listeners[cmd] = reaction;

    const message = JSON.stringify({ cmd, ...payload });
    if (isOpen) {
        socket.send(message);
    } else {
        queue.push(message);
    }
}

const emptyFunction = (data)=>{console.log(data)}
const emptyErrFunction = (data)=>{console.log(data)}

/**
 * Sendet eine GET-Anfrage.
 * @param {string} path Der API-Pfad, z.B. 'posts/1'
 * @param {function} action Die Funktion, die bei Erfolg aufgerufen wird.
 * @param {function} erraction Die Funktion, die bei Misserfolt aufgerufen wird.
 */
function sGet(path, action = emptyFunction, erraction = emptyErrFunction) {
    fetch(url + path)
        .then(response => {
            if (!response.ok) {
                throw new Error("HTTP-Fehler " + response.status);
            }
            return response.json();
        })
        .then(action)
        .catch(erraction);
}

/**
 * Sendet eine POST-Anfrage mit JSON-Payload.
 * @param {string} path Der API-Pfad, z.B. 'posts/1'
 * @param {function} action Die Funktion, die bei Erfolg aufgerufen wird.
 * @param {function} erraction Die Funktion, die bei Misserfolt aufgerufen wird.
 */
function sPost(path, payload, action = emptyFunction, erraction = emptyErrFunction) {
    fetch(url + path, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(payload)
    })
    .then(response => {
        if (!response.ok) {
            // Wir versuchen, die JSON-Fehlermeldung vom Server zu lesen
            return response.json().then(err => { throw err; });
        }
        return response.json();
    })
    .then(action)
    .catch(erraction);
}

export {sGet, sPost, pipe}