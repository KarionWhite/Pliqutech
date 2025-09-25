import { ref } from 'vue'

// Zustand für Login (true = eingeloggt, false = ausgeloggt)
const isLoggedIn = ref(false) // Du müsstest hier deine echte Login-Logik integrieren

// Zustand für das Hauptmenü (true = offen, false = geschlossen)
const isMainMenuOpen = ref(false)

// Funktion, um das Hauptmenü umzuschalten
function toggleMainMenu() {
  isMainMenuOpen.value = !isMainMenuOpen.value
}

// Platzhalter-Funktionen für Login/Logout
function handleLogin() {
  // Hier deine Login-Logik
  isLoggedIn.value = true
  console.log('Login geklickt')
}

function handleLogout() {
  // Hier deine Logout-Logik
  isLoggedIn.value = false
  isMainMenuOpen.value = false // Menü ggf. schließen
  console.log('Logout geklickt')
}

export {isLoggedIn, isMainMenuOpen, toggleMainMenu, handleLogin, handleLogout }