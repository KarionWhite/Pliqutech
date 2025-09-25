//useLogin

import { ref } from 'vue';
import { sPost } from '@/assets/js/api_com.mjs'
import {call_action, destroy_event} from '@/assets/js/queu.mjs'


export function useAuth() {


    const isLoggedIn = ref(false);
    const error = ref(null);

    const login_event = 'login';
    const login_error_event = 'login_error'

    const logout_event = 'logout';
    const logout_error_event = 'logout_error'

    function login(username, password) {
        error.value = null;

        // Die Erfolgs-Callback-Funktion
        const handleSuccess = (responseData) => {
            isLoggedIn.value = true;
            console.log("Login erfolgreich:", responseData);
            call_action(login_event,responseData);
            destroy_event(login_event) //Wir brauchen das jetzt nicht mehr.
            destroy_event(login_error_event);
        };

        // Die Fehler-Callback-Funktion
        const handleError = (err) => {
            error.value = err.message || "Unbekannter Fehler";
            console.error("Login fehlgeschlagen:", err);
            call_action(login_error_event,err); //Was auch immer hier geschehen soll.
            destroy_event(login_event) //Wir brauchen das jetzt nicht mehr.
            destroy_event(login_error_event);
        };

        sPost('login', { username, password }, handleSuccess, handleError,login_event,login_error_event);
    }

    /**
     * 
     * @param {String} username
     * @param {String} token der Token vom User für die auth
     */
    function logout(username, token){

        // Die Erfolgs-Callback-Funktion
        const handleSuccess = (responseData) => {
            isLoggedIn.value = false;
            console.log("Logout erfolgreich");
            call_action(logout_event, responseData);
            destroy_event(logout_event);
            destroy_event(logout_error_event);
        };

        // Die Fehler-Callback-Funktion
        const handleError = (err) => {
            error.value = err.message || "Unbekannter Fehler";
            console.error("Logout fehlgeschlagen", err);
            call_action(logout_error_event,err); //Was auch immer hier geschehen soll.
            destroy_event(logout_event);
            destroy_event(logout_error_event);
        };

        sPost('logout',{username, token}, handleSuccess, handleError)

    }

    return { isLoggedIn, error, login, logout,login_event,login_error_event,logout_event,logout_error_event };
}