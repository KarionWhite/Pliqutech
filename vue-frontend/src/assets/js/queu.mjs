/**
 * Das Herzstück unseres Event-Systems.
 * Dieses Objekt speichert alle Event-Namen als Schlüssel
 * und ein Array von Listener-Funktionen als deren Wert.
 * Da es auf der obersten Ebene des Moduls lebt, agiert es als "Singleton",
 * d.h. die gesamte Anwendung teilt sich diese eine Instanz.
 */
const qeventO = {};

/**
 * Registriert eine Funktion (einen "Listener") für ein bestimmtes Event.
 * @param {string} qevent - Der Name des Events, auf das gelauscht werden soll (z.B. 'login:success').
 * @param {Function} fun - Die Funktion, die ausgeführt werden soll, wenn das Event gefeuert wird.
 */
function que_action(qevent, fun) {
    // Falls für dieses Event noch kein Array existiert, erstellen wir ein leeres.
    if (!qeventO[qevent]) {
        qeventO[qevent] = [];
    }

    // Wir fügen die Funktion nur hinzu, wenn sie "truthy" ist.
    // Das verhindert, dass 'null' oder 'undefined' als Listener registriert werden.
    if (fun) {
        qeventO[qevent].push(fun);
    }
}

/**
 * Feuert ein Event und führt alle zugehörigen Listener-Funktionen aus.
 * @param {string} qevent - Der Name des Events, das ausgelöst werden soll.
 * @param {*} [data] - Optionale Daten, die an jeden Listener übergeben werden.
 */
function call_action(qevent, data) {
    // Holt alle Listener für das Event. Das '|| []' ist eine Absicherung,
    // damit nichts schiefgeht, falls ein Event ohne Listener aufgerufen wird.
    const events = qeventO[qevent] || [];

    // Wir gehen durch alle registrierten Listener und rufen sie mit den Daten auf.
    for (const myevent of events) {
        myevent(data);
    }
}

/**
 * Entfernt eine spezifische Listener-Funktion von einem Event.
 * Wichtig für "Cleanup", um Memory-Leaks zu vermeiden, wenn Komponenten
 * nicht mehr existieren, deren Listener aber noch registriert sind.
 * @param {string} qevent - Der Name des Events.
 * @param {Function} fun - Genau die Funktion, die entfernt werden soll.
 */
function remove_action(qevent, fun) {
    const events = qeventO[qevent];

    // Wenn es keine Listener für dieses Event gibt, gibt es nichts zu tun.
    if (!events) return;

    // Wir filtern das Array und behalten nur die Funktionen,
    // die nicht mit der zu entfernenden Funktion identisch sind.
    qeventO[qevent] = events.filter(handler => handler !== fun);
}

/**
 * Zerstört ein komplettes Event und entfernt alle zugehörigen Listener.
 * @param {string} qevent - Der Name des Events, das komplett entfernt werden soll.
 */
function destroy_event(qevent) {
    if (qeventO[qevent]) {
        delete qeventO[qevent];
        return true;
    }
    return false;
}

// Wir exportieren die fertigen Funktionen, damit sie in der ganzen App
// importiert und für die Kommunikation zwischen Modulen genutzt werden können. 🚀
export {
    que_action,
    call_action,
    remove_action,
    destroy_event // Die neue Funktion hier hinzufügen
};