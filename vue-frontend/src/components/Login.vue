<script setup>
  import { onMounted, ref } from 'vue'; // ref für reaktive Variablen, falls du Fehlermeldungen im Template anzeigen willst
  import { useAuth } from '@/assets/js/useLogin.mjs';
  import {que_action} from '@/assets/js/queu.mjs';

  // Deklariere useAuth direkt im Setup-Bereich, um reaktive Bindungen zu ermöglichen
  const {isLoggedIn, login, logout, login_event, login_error_event, logout_event, logout_error_event} = useAuth();

  // Beispiel für reaktive Variablen für Fehlermeldungen und Ladezustand
  const errorMessage = ref('');
  const isLoading = ref(false);
  const showLoginOverlay = ref(true); // Um das Template zu schließen/öffnen

  onMounted(() => {
    const loginForm = document.getElementById("login_form");

    if (loginForm) {
      loginForm.addEventListener("submit",(event)=>{
        event.preventDefault();

        // Fehlermeldung zurücksetzen und Ladezustand setzen
        errorMessage.value = '';
        isLoading.value = true;

        const element_username = document.getElementById("user");
        const element_password = document.getElementById("password");

        const user = element_username ? element_username.value : '';
        const pass = element_password ? element_password.value : '';

        // Definieren der Callback-Funktionen
        const login_func = (data) => {
          console.log("Login erfolgreich:", data);
          showLoginOverlay.value = false; // Template schließen
          isLoading.value = false; // Ladebalken beenden
        };

        const login_error_func = (data) => {
          console.error("Login Fehler:", data);
          errorMessage.value = data.message || "Anmeldung fehlgeschlagen. Bitte überprüfen Sie Ihre Eingaben."; // Zeige Fehler an
          isLoading.value = false; // Ladebalken beenden
        };

        // Deine Validierungslogik
        let validationFailed = false; // Eine Flag, um zu verfolgen, ob Validierung fehlschlägt

        if(user === ''){ // Verwende === für strengen Vergleich
          errorMessage.value = 'Bitte geben Sie Ihren Benutzernamen ein.'; // Setze spezifische Fehlermeldung
          validationFailed = true;
          // Du könntest hier auch ein visuelles Feedback direkt am Feld geben
        }
        if(pass === ''){ // Verwende === für strengen Vergleich
          if (errorMessage.value) { // Wenn schon eine Fehlermeldung da ist, ergänze oder überschreibe
            errorMessage.value += '\nBitte geben Sie Ihr Passwort ein.'; // Beispiel: Neue Zeile für mehrere Fehler
          } else {
            errorMessage.value = 'Bitte geben Sie Ihr Passwort ein.';
          }
          validationFailed = true;
          // Du könntest hier auch ein visuelles Feedback direkt am Feld geben
        }

        if (validationFailed) {
            isLoading.value = false; // Ladebalken beenden, da kein Login-Versuch
            return; // Beende die Funktion hier
        }

        // Wenn Validierung erfolgreich, registriere Callbacks und starte Login
        que_action(login_event, login_func);
        que_action(login_error_event, login_error_func);
        login(user, pass);
      });
    } else {
      console.error("Das Formular mit der ID 'login_form' konnte nicht gefunden werden.");
    }
  });
</script>

<template>
  <div class="login-overlay" v-if="showLoginOverlay">
    <div class="login-box">
      <h2 class="login-title">Anmelden</h2>
      <form @submit.prevent id="login_form">
        <div class="input-group">
          <label for="user">Username</label>
          <input type="text" id="user" name="user" required>
        </div>
        <div class="input-group">
          <label for="password">Passwort</label>
          <input type="password" id="password" name="password" required>
        </div>
        <p v-if="errorMessage" style="color: red; margin-top: 10px;">{{ errorMessage }}</p>
        <button type="submit" class="submit-btn" :disabled="isLoading">
            {{ isLoading ? 'Lädt...' : 'Login' }}
        </button>
      </form>
      <div class="login-footer">
        <router-link to="/passwort-vergessen">Passwort vergessen?</router-link> | <router-link to="/register">Registrieren</router-link>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* Dein CSS bleibt gleich, wurde aber der Vollständigkeit halber hier belassen */
.login-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.7);
  backdrop-filter: blur(5px);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 2000; /* Muss über dem Header liegen */
}

.login-box {
  background-color: #1e1e1e;
  padding: 2rem 3rem;
  border-radius: 10px;
  border: 1px solid #444;
  width: 100%;
  max-width: 400px;
  text-align: center;
  color: white;
}

.login-title {
  margin-bottom: 1.5rem;
}

.input-group {
  text-align: left;
  margin-bottom: 1rem;
}

.input-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-size: 0.9rem;
}

.input-group input {
  width: 100%;
  padding: 0.75rem;
  background-color: #333;
  border: 1px solid #555;
  border-radius: 5px;
  color: white;
  box-sizing: border-box; /* Wichtig für korrekte Breite */
}

.submit-btn {
  width: 100%;
  padding: 0.75rem;
  background-color: #4A90E2;
  color: white;
  border: none;
  border-radius: 5px;
  font-size: 1rem;
  font-weight: bold;
  cursor: pointer;
  margin-top: 1rem;
}

.submit-btn:disabled {
  background-color: #6c757d; /* Grau, wenn disabled */
  cursor: not-allowed;
}


.login-footer {
  margin-top: 1.5rem;
  font-size: 0.8rem;
}

.login-footer a {
  color: #4A90E2;
  text-decoration: none;
}
</style>