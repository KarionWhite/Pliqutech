<script setup>
  import { ref } from 'vue'; // Importiere ref für reaktive Variablen
  import { isLoggedIn, isMainMenuOpen, toggleMainMenu, handleLogin, handleLogout } from './assets/js/menu.mjs';
  import LoginForm from '@/components/Login.vue'; // Pfad zu deiner neuen LoginForm.vue

  // Zustandsvariable für die Sichtbarkeit des Login-Overlays
  const showLoginDialog = ref(false);

  // Anpassung der handleLogin Funktion
  const openLoginDialog = () => {
    showLoginDialog.value = true;
    // Optional: Wenn dein Hauptmenü offen ist, schließe es, wenn der Login-Dialog geöffnet wird.
    if (isMainMenuOpen.value) {
      toggleMainMenu();
    }
  };

  // Funktion zum Schließen des Login-Dialogs
  const closeLoginDialog = () => {
    showLoginDialog.value = false;
    // Optional: Nach dem Schließen des Dialogs könntest du das Hauptmenü wieder öffnen, wenn es vorher offen war.
  };

  // Funktion, die aufgerufen wird, wenn der Login in LoginForm erfolgreich war
  const onLoginSuccess = () => {
    // Hier kannst du zusätzliche Aktionen nach einem erfolgreichen Login ausführen,
    // z.B. eine Begrüßungsnachricht anzeigen oder zu einer anderen Seite navigieren.
    console.log("Login war erfolgreich! Header weiß Bescheid.");
    // handleLogin() wird wahrscheinlich in menu.mjs isLoggedin setzen.
  };

</script>

<template>
  <header class="banner-header">
    <div class="logo-container">
      <img src="/favicon.ico" alt="Pliqutech Logo" class="logo" />
    </div>

    <nav :class="{ 'navigation-menu': true, 'is-open': isMainMenuOpen }">
      <ul class="main-menu">
        <li class="menu-item"><a href="#">Startseite</a></li>
        <li class="menu-item"><a href="#">Über uns</a></li>
        <li class="menu-item"><a href="#">Dienste</a></li>
      </ul>
    </nav>

    <div class="header-controls">
      <div class="auth-section">
        <button v-if="!isLoggedIn" @click="openLoginDialog">Login</button>
        <button v-else @click="handleLogout">Logout</button>
      </div>

      <button
        type="button"
        @click="toggleMainMenu"
        :class="{ 'menu-toggle-button': true, 'is-active': isMainMenuOpen }"
        aria-label="Menü umschalten"
      >
        <span class="bar"></span>
        <span class="bar"></span>
        <span class="bar"></span>
      </button>
    </div>
  </header>

  <main>
    <LoginForm
      v-if="showLoginDialog"
      @close="closeLoginDialog"
      @loginSuccess="onLoginSuccess"
    />
  </main>
</template>

<style scoped>
/* Dein bestehendes CSS für den Header und das Menü */
.banner-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.5rem 1rem;
  background-color: black;
  border-bottom: 1px solid #333;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 1000;
  box-sizing: border-box;
}

.logo-container {
  z-index: 1001;
}

.logo {
  height: 40px;
}

.header-controls {
  display: flex;
  align-items: center;
  gap: 1rem;
  z-index: 1001;
}

.auth-section button {
  background-color: #4A90E2;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 5px;
  cursor: pointer;
}

.menu-toggle-button {
  width: 40px;
  height: 40px;
  background: transparent;
  border: none;
  cursor: pointer;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  padding: 8px;
  box-sizing: border-box;
}

.bar {
  display: block;
  width: 100%;
  height: 3px;
  background-color: white;
  border-radius: 3px;
  transition: all 0.3s ease-in-out;
}

.menu-toggle-button.is-active .bar {
  background-color: white;
}

.menu-toggle-button.is-active .bar:nth-child(1) {
  transform: translateY(8px) rotate(45deg);
}
.menu-toggle-button.is-active .bar:nth-child(2) {
  opacity: 0;
}
.menu-toggle-button.is-active .bar:nth-child(3) {
  transform: translateY(-8px) rotate(-45deg);
}

.navigation-menu {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100vh;
  background-color: rgba(10, 10, 10, 0.9);
  backdrop-filter: blur(8px);
  
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  z-index: 1000;

  transform: translateX(100%);
  visibility: hidden;
  transition: transform 0.5s cubic-bezier(0.77, 0, 0.175, 1), visibility 0s 1s;
}

.navigation-menu.is-open {
  transform: translateX(0);
  visibility: visible;
  transition: transform 0.35s ease-in-out;
}

.main-menu {
  list-style: none;
  padding: 0;
  margin: 0;
  text-align: center;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1.5rem;
}

.menu-item {
  opacity: 0;
  transform: translateX(20px);
  transition: opacity 0.3s ease, transform 0.3s ease;
}

.navigation-menu.is-open .menu-item {
  opacity: 1;
  transform: translateX(0);
}

.navigation-menu.is-open .menu-item:nth-child(1) { transition-delay: 0.2s; }
.navigation-menu.is-open .menu-item:nth-child(2) { transition-delay: 0.3s; }
.navigation-menu.is-open .menu-item:nth-child(3) { transition-delay: 0.4s; }

.main-menu a {
  display: inline-block;
  font-size: 2rem;
  font-weight: bold;
  color: white;
  text-decoration: none;
  padding: 0.5rem 1rem;
  margin: 0.5rem 0;
  transition: color 0.2s ease, transform 0.2s ease;
}

.main-menu a:hover {
  color: #4A90E2;
  transform: scale(1.1);
}

@media (min-width: 768px) {
  .menu-toggle-button {
    display: none;
  }

  .navigation-menu {
    position: static;
    height: auto;
    width: auto;
    background-color: transparent;
    backdrop-filter: none;
    flex-direction: row !important;
    justify-content: flex-start;
    
    visibility: visible;
    transform: none;
    z-index: auto;
    transition: none;
    
    flex-grow: 1;
    margin-left: 2rem;
  }

  .menu-item {
    opacity: 1;
    transform: none;
  }
  
  .main-menu {
    gap: 1.5rem;
    flex-direction: row;
  }

  .main-menu a {
    font-size: 1rem;
    font-weight: normal;
  }
  
  .main-menu a:hover {
    transform: none;
  }
}

main {
  padding-top: 70px;
}
</style>